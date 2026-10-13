// Copyright (c) 2026, Oracle and/or its affiliates. All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package transfer

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/objectstorage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	uploadStreamOriginalBody = "v1, original object"
	uploadStreamOriginalETag = "destination-v1-etag"
	uploadStreamNewBody      = "v2, overwriting via UploadManager.UploadStream"
	uploadStreamNewETag      = "destination-v2-etag"
	uploadStreamUpdatedBody  = "intervening writer's object"
	uploadStreamUpdatedETag  = "intervening-writer-etag"
)

// An unconditional overwrite must continue to upload and commit successfully
func TestUploadManager_UploadStreamUnconditionalOverwrite(t *testing.T) {
	body, etag, err := uploadStreamWithTestServer(t, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, uploadStreamNewBody, body)
	assert.Equal(t, uploadStreamNewETag, etag)
}

// The destination object's ETag must condition the final overwrite, not a new part
func TestUploadManager_UploadStreamIfMatchUnchanged(t *testing.T) {
	body, etag, err := uploadStreamWithTestServer(t, common.String(uploadStreamOriginalETag), nil)
	require.NoError(t, err)
	assert.Equal(t, uploadStreamNewBody, body)
	assert.Equal(t, uploadStreamNewETag, etag)
}

// A competing write must reject commit and preserve the intervening object
func TestUploadManager_UploadStreamIfMatchConcurrentUpdate(t *testing.T) {
	beforeCommit := func() (string, string) {
		return uploadStreamUpdatedBody, uploadStreamUpdatedETag
	}
	body, etag, err := uploadStreamWithTestServer(t, common.String(uploadStreamOriginalETag), beforeCommit)
	require.Error(t, err)
	serviceErr, ok := common.IsServiceErrorRichInfo(err)
	require.True(t, ok, "expected a service error, got %v", err)
	assert.Equal(t, http.StatusPreconditionFailed, serviceErr.GetHTTPStatusCode())
	assert.Equal(t, "IfMatchFailed", serviceErr.GetCode())
	assert.Equal(t, "CommitMultipartUpload", serviceErr.GetOperationName())
	assert.Equal(t, uploadStreamUpdatedBody, body)
	assert.Equal(t, uploadStreamUpdatedETag, etag)
}

func uploadStreamWithTestServer(t *testing.T, ifMatch *string, beforeCommit func() (string, string)) (string, string, error) {
	t.Helper()
	const (
		objectName = "if-match-repro-object"
		createPath = "/n/namespace/b/bucket/u"
		uploadPath = createPath + "/" + objectName
		uploadID   = "new-upload-id"
		partETag   = "part-1-etag"
	)
	expectedIfMatch := ""
	if ifMatch != nil {
		expectedIfMatch = *ifMatch
	}

	// Start with an existing destination object
	objectBody, objectETag := uploadStreamOriginalBody, uploadStreamOriginalETag
	var partBody string
	var requests []string
	var mu sync.Mutex
	failMatch := func(w http.ResponseWriter) {
		w.WriteHeader(http.StatusPreconditionFailed)
		_, _ = io.WriteString(w, `{"code":"IfMatchFailed","message":"The current entity tag did not match the if-match condition"}`)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		w.Header().Set("content-type", "application/json")
		w.Header().Set("opc-request-id", "if-match-test-request-id")
		operation := r.Method + " " + r.URL.Path
		requests = append(requests, operation)
		switch operation {
		case "POST " + createPath:
			assert.Equal(t, expectedIfMatch, r.Header.Get("if-match"))
			_, _ = io.WriteString(w, `{"uploadId":"new-upload-id"}`)
		case "PUT " + uploadPath:
			assert.Equal(t, uploadID, r.URL.Query().Get("uploadId"))
			assert.Equal(t, "1", r.URL.Query().Get("uploadPartNum"))
			assert.Empty(t, r.Header.Get("if-match"), "the destination ETag must not condition a new part")
			// A new part has no ETag, so any IfMatch fails before its body is read
			if r.Header.Get("if-match") != "" {
				failMatch(w)
				return
			}
			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)
			partBody = string(body)
			w.Header().Set("etag", partETag)
		case "POST " + uploadPath:
			assert.Equal(t, uploadID, r.URL.Query().Get("uploadId"))
			assert.Equal(t, expectedIfMatch, r.Header.Get("if-match"), "commit must preserve the destination condition")
			var details objectstorage.CommitMultipartUploadDetails
			assert.NoError(t, json.NewDecoder(r.Body).Decode(&details))
			assert.Equal(t, []objectstorage.CommitMultipartUploadPartDetails{{
				PartNum: common.Int(1), Etag: common.String(partETag),
			}}, details.PartsToCommit)
			if beforeCommit != nil {
				objectBody, objectETag = beforeCommit()
			}
			if match := r.Header.Get("if-match"); match != "" && match != objectETag {
				failMatch(w)
				return
			}
			objectBody, objectETag = partBody, uploadStreamNewETag
			w.Header().Set("etag", objectETag)
		default:
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.String())
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()
	client := &objectstorage.ObjectStorageClient{
		BaseClient: common.BaseClient{
			HTTPClient: server.Client(),
			Signer:     fakeSigner{},
			Host:       server.URL,
			UserAgent:  "upload-manager-if-match-test",
		},
	}
	noRetryPolicy := common.NoRetryPolicy()
	response, err := NewUploadManager().UploadStream(context.Background(), UploadStreamRequest{
		UploadRequest: UploadRequest{
			NamespaceName:       common.String("namespace"),
			BucketName:          common.String("bucket"),
			ObjectName:          common.String(objectName),
			ObjectStorageClient: client,
			IfMatch:             ifMatch,
			NumberOfGoroutines:  common.Int(1),
			RequestMetadata:     common.RequestMetadata{RetryPolicy: &noRetryPolicy},
		},
		StreamReader: strings.NewReader(uploadStreamNewBody),
	})

	mu.Lock()
	defer mu.Unlock()
	assert.Equal(t, MultipartUpload, response.Type)
	assert.Equal(t, []string{"POST " + createPath, "PUT " + uploadPath, "POST " + uploadPath}, requests)
	return objectBody, objectETag, err
}

// IfNoneMatch describes the destination and must allow uploading a part again
func TestMultipartUpload_IfNoneMatch(t *testing.T) {
	const uploadPath = "/n/namespace/b/bucket/u/object"
	var requests []string
	var mu sync.Mutex
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		w.Header().Set("content-type", "application/json")
		assert.Empty(t, r.Header.Get("if-match"))
		operation := r.Method + " " + r.URL.Path
		requests = append(requests, operation)
		switch operation {
		case "POST /n/namespace/b/bucket/u":
			assert.Equal(t, "*", r.Header.Get("if-none-match"))
			_, _ = io.WriteString(w, `{"uploadId":"new-upload-id"}`)
		case "PUT " + uploadPath:
			assert.Equal(t, "new-upload-id", r.URL.Query().Get("uploadId"))
			assert.Equal(t, "1", r.URL.Query().Get("uploadPartNum"))
			assert.Empty(t, r.Header.Get("if-none-match"), "the destination condition must not apply to parts")
			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)
			assert.Equal(t, "part body", string(body))
			w.Header().Set("etag", "part-etag")
		case "POST " + uploadPath:
			assert.Equal(t, "new-upload-id", r.URL.Query().Get("uploadId"))
			assert.Equal(t, "*", r.Header.Get("if-none-match"), "commit must preserve the destination condition")
			w.Header().Set("etag", "object-etag")
		default:
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.String())
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()
	client := &objectstorage.ObjectStorageClient{
		BaseClient: common.BaseClient{
			HTTPClient: server.Client(),
			Signer:     fakeSigner{},
			Host:       server.URL,
			UserAgent:  "multipart-uploader-test",
		},
	}
	noRetryPolicy := common.NoRetryPolicy()
	request := UploadRequest{
		NamespaceName:       common.String("namespace"),
		BucketName:          common.String("bucket"),
		ObjectName:          common.String("object"),
		ObjectStorageClient: client,
		IfNoneMatch:         common.String("*"),
		RequestMetadata:     common.RequestMetadata{RetryPolicy: &noRetryPolicy},
	}
	uploader := multipartUpload{}
	ctx := context.Background()
	uploadID, err := uploader.createMultipartUpload(ctx, request)
	require.NoError(t, err)
	part := uploadPart{partNum: 1, partBody: []byte("part body"), size: int64(len("part body"))}
	for i := 0; i < 2; i++ {
		response, err := uploader.uploadPart(ctx, request, part, uploadID)
		require.NoError(t, err)
		part.etag = response.ETag
	}
	_, err = uploader.commit(ctx, request, map[int]uploadPart{1: part}, uploadID)
	require.NoError(t, err)

	mu.Lock()
	defer mu.Unlock()
	assert.Equal(t, []string{"POST /n/namespace/b/bucket/u", "PUT " + uploadPath, "PUT " + uploadPath, "POST " + uploadPath}, requests)
}
