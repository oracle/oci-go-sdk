// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// APIs for managing buckets and objects.
//

package integtest

import (
	"github.com/oci-go-sdk/common"
	"github.com/oci-go-sdk/objectstorage"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testRegionForObjectStorage = common.REGION_PHX
)

func TestObjectStorageClient_AbortMultipartUpload(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.AbortMultipartUploadRequest{}
	err := c.AbortMultipartUpload(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_CommitMultipartUpload(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.CommitMultipartUploadRequest{}
	err := c.CommitMultipartUpload(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_CreateBucket(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.CreateBucketRequest{}
	r, err := c.CreateBucket(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_CreateMultipartUpload(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.CreateMultipartUploadRequest{}
	r, err := c.CreateMultipartUpload(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_CreatePreauthenticatedRequest(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.CreatePreauthenticatedRequestRequest{}
	r, err := c.CreatePreauthenticatedRequest(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_DeleteBucket(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.DeleteBucketRequest{}
	err := c.DeleteBucket(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_DeleteObject(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.DeleteObjectRequest{}
	err := c.DeleteObject(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_DeletePreauthenticatedRequest(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.DeletePreauthenticatedRequestRequest{}
	err := c.DeletePreauthenticatedRequest(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_GetBucket(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.GetBucketRequest{}
	r, err := c.GetBucket(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_GetNamespace(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.GetNamespaceRequest{}
	r, err := c.GetNamespace(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_GetObject(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.GetObjectRequest{}
	r, err := c.GetObject(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_GetPreauthenticatedRequest(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.GetPreauthenticatedRequestRequest{}
	r, err := c.GetPreauthenticatedRequest(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_HeadBucket(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.HeadBucketRequest{}
	err := c.HeadBucket(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_HeadObject(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.HeadObjectRequest{}
	err := c.HeadObject(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_ListBuckets(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.ListBucketsRequest{}
	r, err := c.ListBuckets(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_ListMultipartUploadParts(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.ListMultipartUploadPartsRequest{}
	r, err := c.ListMultipartUploadParts(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_ListMultipartUploads(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.ListMultipartUploadsRequest{}
	r, err := c.ListMultipartUploads(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_ListObjects(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.ListObjectsRequest{}
	r, err := c.ListObjects(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_ListPreauthenticatedRequests(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.ListPreauthenticatedRequestsRequest{}
	r, err := c.ListPreauthenticatedRequests(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_PutObject(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.PutObjectRequest{}
	err := c.PutObject(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_UpdateBucket(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.UpdateBucketRequest{}
	r, err := c.UpdateBucket(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_UploadPart(t *testing.T) {
	t.Skip("Not implemented")
	c := objectstorage.NewObjectStorageClientForRegion(testRegionForObjectStorage)
	request := objectstorage.UploadPartRequest{}
	err := c.UploadPart(context.Background(), request)
	assert.NoError(t, err)
	return
}
