// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// APIs for managing buckets and objects.
//

package objectstorage

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//ObjectStorageClientData a client for ObjectStorage
type ObjectStorageClientData struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

func buildSigner(configProvider common.ConfigurationProvider) common.HTTPRequestSigner {
	objStorageHeaders := []string{"date", "(request-target)", "host"}
	defaultBodyHeaders := []string{"content-length", "content-type", "x-content-sha256"}
	shouldHashBody := func(r *http.Request) bool {
		return r.Method == http.MethodPost
	}
	signer := common.RequestSignerWithBodyHashingPredicate(configProvider, objStorageHeaders, defaultBodyHeaders, shouldHashBody)
	return signer
}

// NewObjectStorageClientWithConfigurationProvider Creates a new default ObjectStorage client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewObjectStorageClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (ObjectStorageClient, error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return nil, err
	}
	baseClient.Signer = buildSigner(configProvider)

	client := ObjectStorageClientData{BaseClient: baseClient}
	err = client.setConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	return &client, err
}

// GetBaseClient get the BaseClient object of this client
func (client *ObjectStorageClientData) GetBaseClient() *common.BaseClient {
	return &client.BaseClient
}

// SetRegion overrides the region of this client.
func (client *ObjectStorageClientData) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "objectstorage", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ObjectStorageClientData) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.config = &configProvider
	client.SetRegion(region)
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *ObjectStorageClientData) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AbortMultipartUpload Aborts an in-progress multipart upload and deletes all parts that have been uploaded.
func (client ObjectStorageClientData) AbortMultipartUpload(ctx context.Context, request AbortMultipartUploadRequest) (response AbortMultipartUploadResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// CommitMultipartUpload Commits a multipart upload, which involves checking part numbers and ETags of the parts, to create an aggregate object.
func (client ObjectStorageClientData) CommitMultipartUpload(ctx context.Context, request CommitMultipartUploadRequest) (response CommitMultipartUploadResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// CreateBucket Creates a bucket in the given namespace with a bucket name and optional user-defined metadata.
// To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies ({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClientData) CreateBucket(ctx context.Context, request CreateBucketRequest) (response CreateBucketResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/n/{namespaceName}/b/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// CreateMultipartUpload Starts a new multipart upload to a specific object in the given bucket in the given namespace.
func (client ObjectStorageClientData) CreateMultipartUpload(ctx context.Context, request CreateMultipartUploadRequest) (response CreateMultipartUploadResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/u", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// CreatePreauthenticatedRequest Create a pre-authenticated request specific to the bucket
func (client ObjectStorageClientData) CreatePreauthenticatedRequest(ctx context.Context, request CreatePreauthenticatedRequestRequest) (response CreatePreauthenticatedRequestResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/p/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// DeleteBucket Deletes a bucket if it is already empty. If the bucket is not empty, use DeleteObject first.
func (client ObjectStorageClientData) DeleteBucket(ctx context.Context, request DeleteBucketRequest) (response DeleteBucketResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// DeleteObject Deletes an object.
func (client ObjectStorageClientData) DeleteObject(ctx context.Context, request DeleteObjectRequest) (response DeleteObjectResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// DeletePreauthenticatedRequest Deletes the bucket level pre-authenticateted request
func (client ObjectStorageClientData) DeletePreauthenticatedRequest(ctx context.Context, request DeletePreauthenticatedRequestRequest) (response DeletePreauthenticatedRequestResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/p/{parId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// GetBucket Gets the current representation of the given bucket in the given namespace.
func (client ObjectStorageClientData) GetBucket(ctx context.Context, request GetBucketRequest) (response GetBucketResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// GetNamespace Gets the name of the namespace for the user making the request. An account name must be unique, must start with a
// letter, and can have up to 15 lowercase letters and numbers. You cannot use spaces or special characters.
func (client ObjectStorageClientData) GetNamespace(ctx context.Context, request GetNamespaceRequest) (response GetNamespaceResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// GetObject Gets the metadata and body of an object.
func (client ObjectStorageClientData) GetObject(ctx context.Context, request GetObjectRequest) (response GetObjectResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// GetPreauthenticatedRequest Get the bucket level pre-authenticateted request
func (client ObjectStorageClientData) GetPreauthenticatedRequest(ctx context.Context, request GetPreauthenticatedRequestRequest) (response GetPreauthenticatedRequestResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/p/{parId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// HeadBucket Efficiently checks if a bucket exists and gets the current ETag for the bucket.
func (client ObjectStorageClientData) HeadBucket(ctx context.Context, request HeadBucketRequest) (response HeadBucketResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodHead, "/n/{namespaceName}/b/{bucketName}/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// HeadObject Gets the user-defined metadata and entity tag for an object.
func (client ObjectStorageClientData) HeadObject(ctx context.Context, request HeadObjectRequest) (response HeadObjectResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodHead, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// ListBuckets Gets a list of all `BucketSummary`s in a compartment. A `BucketSummary` contains only summary fields for the bucket
// and does not contain fields like the user-defined metadata.
// To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies ({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClientData) ListBuckets(ctx context.Context, request ListBucketsRequest) (response ListBucketsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// ListMultipartUploadParts Lists the parts of an in-progress multipart upload.
func (client ObjectStorageClientData) ListMultipartUploadParts(ctx context.Context, request ListMultipartUploadPartsRequest) (response ListMultipartUploadPartsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// ListMultipartUploads Lists all in-progress multipart uploads for the given bucket in the given namespace.
func (client ObjectStorageClientData) ListMultipartUploads(ctx context.Context, request ListMultipartUploadsRequest) (response ListMultipartUploadsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/u", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// ListObjects Lists the objects in a bucket.
// To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies ({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClientData) ListObjects(ctx context.Context, request ListObjectsRequest) (response ListObjectsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/o", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// ListPreauthenticatedRequests List pre-authenticated requests for the bucket
func (client ObjectStorageClientData) ListPreauthenticatedRequests(ctx context.Context, request ListPreauthenticatedRequestsRequest) (response ListPreauthenticatedRequestsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/p/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// PutObject Creates a new object or overwrites an existing one.
// To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies ({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClientData) PutObject(ctx context.Context, request PutObjectRequest) (response PutObjectResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/n/{namespaceName}/b/{bucketName}/o/{objectName}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// UpdateBucket Performs a partial or full update of a bucket's user-defined metadata.
func (client ObjectStorageClientData) UpdateBucket(ctx context.Context, request UpdateBucketRequest) (response UpdateBucketResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// UploadPart Uploads a single part of a multipart upload.
func (client ObjectStorageClientData) UploadPart(ctx context.Context, request UploadPartRequest) (response UploadPartResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/n/{namespaceName}/b/{bucketName}/u/{objectName}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// ObjectStorageClient is client interface for ObjectStorage
type ObjectStorageClient interface {
	GetBaseClient() *common.BaseClient
	ConfigurationProvider() *common.ConfigurationProvider
	AbortMultipartUpload(ctx context.Context, request AbortMultipartUploadRequest) (response AbortMultipartUploadResponse, err error)
	CommitMultipartUpload(ctx context.Context, request CommitMultipartUploadRequest) (response CommitMultipartUploadResponse, err error)
	CreateBucket(ctx context.Context, request CreateBucketRequest) (response CreateBucketResponse, err error)
	CreateMultipartUpload(ctx context.Context, request CreateMultipartUploadRequest) (response CreateMultipartUploadResponse, err error)
	CreatePreauthenticatedRequest(ctx context.Context, request CreatePreauthenticatedRequestRequest) (response CreatePreauthenticatedRequestResponse, err error)
	DeleteBucket(ctx context.Context, request DeleteBucketRequest) (response DeleteBucketResponse, err error)
	DeleteObject(ctx context.Context, request DeleteObjectRequest) (response DeleteObjectResponse, err error)
	DeletePreauthenticatedRequest(ctx context.Context, request DeletePreauthenticatedRequestRequest) (response DeletePreauthenticatedRequestResponse, err error)
	GetBucket(ctx context.Context, request GetBucketRequest) (response GetBucketResponse, err error)
	GetNamespace(ctx context.Context, request GetNamespaceRequest) (response GetNamespaceResponse, err error)
	GetObject(ctx context.Context, request GetObjectRequest) (response GetObjectResponse, err error)
	GetPreauthenticatedRequest(ctx context.Context, request GetPreauthenticatedRequestRequest) (response GetPreauthenticatedRequestResponse, err error)
	HeadBucket(ctx context.Context, request HeadBucketRequest) (response HeadBucketResponse, err error)
	HeadObject(ctx context.Context, request HeadObjectRequest) (response HeadObjectResponse, err error)
	ListBuckets(ctx context.Context, request ListBucketsRequest) (response ListBucketsResponse, err error)
	ListMultipartUploadParts(ctx context.Context, request ListMultipartUploadPartsRequest) (response ListMultipartUploadPartsResponse, err error)
	ListMultipartUploads(ctx context.Context, request ListMultipartUploadsRequest) (response ListMultipartUploadsResponse, err error)
	ListObjects(ctx context.Context, request ListObjectsRequest) (response ListObjectsResponse, err error)
	ListPreauthenticatedRequests(ctx context.Context, request ListPreauthenticatedRequestsRequest) (response ListPreauthenticatedRequestsResponse, err error)
	PutObject(ctx context.Context, request PutObjectRequest) (response PutObjectResponse, err error)
	UpdateBucket(ctx context.Context, request UpdateBucketRequest) (response UpdateBucketResponse, err error)
	UploadPart(ctx context.Context, request UploadPartRequest) (response UploadPartResponse, err error)
}
