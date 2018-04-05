// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Common set of Object and Archive Storage APIs for managing buckets and objects.
//

package objectstorage

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//ObjectStorageClient a client for ObjectStorage
type ObjectStorageClient struct {
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
func NewObjectStorageClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ObjectStorageClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}
	baseClient.Signer = buildSigner(configProvider)

	client = ObjectStorageClient{BaseClient: baseClient}
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ObjectStorageClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "objectstorage", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ObjectStorageClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ObjectStorageClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AbortMultipartUpload Aborts an in-progress multipart upload and deletes all parts that have been uploaded.
func (client ObjectStorageClient) AbortMultipartUpload(ctx context.Context, request AbortMultipartUploadRequest) (response AbortMultipartUploadResponse, err error) {
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
func (client ObjectStorageClient) CommitMultipartUpload(ctx context.Context, request CommitMultipartUploadRequest) (response CommitMultipartUploadResponse, err error) {
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
func (client ObjectStorageClient) CreateBucket(ctx context.Context, request CreateBucketRequest) (response CreateBucketResponse, err error) {
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
func (client ObjectStorageClient) CreateMultipartUpload(ctx context.Context, request CreateMultipartUploadRequest) (response CreateMultipartUploadResponse, err error) {
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

// CreatePreauthenticatedRequest Creates a pre-authenticated request specific to the bucket.
func (client ObjectStorageClient) CreatePreauthenticatedRequest(ctx context.Context, request CreatePreauthenticatedRequestRequest) (response CreatePreauthenticatedRequestResponse, err error) {
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
func (client ObjectStorageClient) DeleteBucket(ctx context.Context, request DeleteBucketRequest) (response DeleteBucketResponse, err error) {
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
func (client ObjectStorageClient) DeleteObject(ctx context.Context, request DeleteObjectRequest) (response DeleteObjectResponse, err error) {
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

// DeletePreauthenticatedRequest Deletes the pre-authenticated request for the bucket.
func (client ObjectStorageClient) DeletePreauthenticatedRequest(ctx context.Context, request DeletePreauthenticatedRequestRequest) (response DeletePreauthenticatedRequestResponse, err error) {
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
func (client ObjectStorageClient) GetBucket(ctx context.Context, request GetBucketRequest) (response GetBucketResponse, err error) {
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

// GetNamespace Namespaces are unique. Namespaces are either the tenancy name or a random string automatically generated during
// account creation. You cannot edit a namespace.
func (client ObjectStorageClient) GetNamespace(ctx context.Context, request GetNamespaceRequest) (response GetNamespaceResponse, err error) {
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

// GetNamespaceMetadata Get the metadata for the namespace, which contains defaultS3CompartmentId and defaultSwiftCompartmentId.
// Any user with the NAMESPACE_READ permission will be able to see the current metadata. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write
// policies to give users access, see Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClient) GetNamespaceMetadata(ctx context.Context, request GetNamespaceMetadataRequest) (response GetNamespaceMetadataResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/n/{namespaceName}", request)
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
func (client ObjectStorageClient) GetObject(ctx context.Context, request GetObjectRequest) (response GetObjectResponse, err error) {
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

// GetPreauthenticatedRequest Gets the pre-authenticated request for the bucket.
func (client ObjectStorageClient) GetPreauthenticatedRequest(ctx context.Context, request GetPreauthenticatedRequestRequest) (response GetPreauthenticatedRequestResponse, err error) {
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

// HeadBucket Efficiently checks to see if a bucket exists and gets the current ETag for the bucket.
func (client ObjectStorageClient) HeadBucket(ctx context.Context, request HeadBucketRequest) (response HeadBucketResponse, err error) {
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
func (client ObjectStorageClient) HeadObject(ctx context.Context, request HeadObjectRequest) (response HeadObjectResponse, err error) {
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
// Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClient) ListBuckets(ctx context.Context, request ListBucketsRequest) (response ListBucketsResponse, err error) {
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
func (client ObjectStorageClient) ListMultipartUploadParts(ctx context.Context, request ListMultipartUploadPartsRequest) (response ListMultipartUploadPartsResponse, err error) {
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
func (client ObjectStorageClient) ListMultipartUploads(ctx context.Context, request ListMultipartUploadsRequest) (response ListMultipartUploadsResponse, err error) {
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
// Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClient) ListObjects(ctx context.Context, request ListObjectsRequest) (response ListObjectsResponse, err error) {
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

// ListPreauthenticatedRequests Lists pre-authenticated requests for the bucket.
func (client ObjectStorageClient) ListPreauthenticatedRequests(ctx context.Context, request ListPreauthenticatedRequestsRequest) (response ListPreauthenticatedRequestsResponse, err error) {
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
func (client ObjectStorageClient) PutObject(ctx context.Context, request PutObjectRequest) (response PutObjectResponse, err error) {
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

// RenameObject Rename an object from source key to target key in the given namespace.
func (client ObjectStorageClient) RenameObject(ctx context.Context, request RenameObjectRequest) (response RenameObjectResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/renameObject", request)
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

// RestoreObjects Restore one or more objects specified by objectName parameter.
// By default object will be restored for 24 hours.Duration can be configured using hours parameter.
func (client ObjectStorageClient) RestoreObjects(ctx context.Context, request RestoreObjectsRequest) (response RestoreObjectsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/actions/restoreObjects", request)
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
func (client ObjectStorageClient) UpdateBucket(ctx context.Context, request UpdateBucketRequest) (response UpdateBucketResponse, err error) {
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

// UpdateNamespaceMetadata Change the default Swift/S3 compartmentId of user's namespace into the user-defined compartmentId. Upon doing
// this, all subsequent bucket creations will use the new default compartment, but no previously created
// buckets will be modified. A user must have the NAMESPACE_UPDATE permission to make changes to the default
// compartments for S3 and Swift.
func (client ObjectStorageClient) UpdateNamespaceMetadata(ctx context.Context, request UpdateNamespaceMetadataRequest) (response UpdateNamespaceMetadataResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/n/{namespaceName}", request)
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
func (client ObjectStorageClient) UploadPart(ctx context.Context, request UploadPartRequest) (response UploadPartResponse, err error) {
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
