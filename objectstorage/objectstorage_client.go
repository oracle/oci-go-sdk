// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
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
	ociResponse, e := client.Retry(ctx, request, client.abortMultipartUpload, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(AbortMultipartUploadResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) abortMultipartUpload(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/u/{objectName}")
	if err != nil {
		return nil, err
	}

	var response AbortMultipartUploadResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CommitMultipartUpload Commits a multipart upload, which involves checking part numbers and ETags of the parts, to create an aggregate object.
func (client ObjectStorageClient) CommitMultipartUpload(ctx context.Context, request CommitMultipartUploadRequest) (response CommitMultipartUploadResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.commitMultipartUpload, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CommitMultipartUploadResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) commitMultipartUpload(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/u/{objectName}")
	if err != nil {
		return nil, err
	}

	var response CommitMultipartUploadResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateBucket Creates a bucket in the given namespace with a bucket name and optional user-defined metadata.
// To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClient) CreateBucket(ctx context.Context, request CreateBucketRequest) (response CreateBucketResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createBucket, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateBucketResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) createBucket(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/n/{namespaceName}/b/")
	if err != nil {
		return nil, err
	}

	var response CreateBucketResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateMultipartUpload Starts a new multipart upload to a specific object in the given bucket in the given namespace.
func (client ObjectStorageClient) CreateMultipartUpload(ctx context.Context, request CreateMultipartUploadRequest) (response CreateMultipartUploadResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createMultipartUpload, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateMultipartUploadResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) createMultipartUpload(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/u")
	if err != nil {
		return nil, err
	}

	var response CreateMultipartUploadResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreatePreauthenticatedRequest Create a pre-authenticated request specific to the bucket
func (client ObjectStorageClient) CreatePreauthenticatedRequest(ctx context.Context, request CreatePreauthenticatedRequestRequest) (response CreatePreauthenticatedRequestResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createPreauthenticatedRequest, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreatePreauthenticatedRequestResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) createPreauthenticatedRequest(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/p/")
	if err != nil {
		return nil, err
	}

	var response CreatePreauthenticatedRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteBucket Deletes a bucket if it is already empty. If the bucket is not empty, use DeleteObject first.
func (client ObjectStorageClient) DeleteBucket(ctx context.Context, request DeleteBucketRequest) (response DeleteBucketResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteBucket, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteBucketResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) deleteBucket(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/")
	if err != nil {
		return nil, err
	}

	var response DeleteBucketResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteObject Deletes an object.
func (client ObjectStorageClient) DeleteObject(ctx context.Context, request DeleteObjectRequest) (response DeleteObjectResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteObject, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteObjectResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) deleteObject(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/o/{objectName}")
	if err != nil {
		return nil, err
	}

	var response DeleteObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeletePreauthenticatedRequest Deletes the bucket level pre-authenticateted request
func (client ObjectStorageClient) DeletePreauthenticatedRequest(ctx context.Context, request DeletePreauthenticatedRequestRequest) (response DeletePreauthenticatedRequestResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deletePreauthenticatedRequest, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeletePreauthenticatedRequestResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) deletePreauthenticatedRequest(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/n/{namespaceName}/b/{bucketName}/p/{parId}")
	if err != nil {
		return nil, err
	}

	var response DeletePreauthenticatedRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetBucket Gets the current representation of the given bucket in the given namespace.
func (client ObjectStorageClient) GetBucket(ctx context.Context, request GetBucketRequest) (response GetBucketResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getBucket, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetBucketResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) getBucket(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/")
	if err != nil {
		return nil, err
	}

	var response GetBucketResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetNamespace Gets the name of the namespace for the user making the request. An account name must be unique, must start with a
// letter, and can have up to 15 lowercase letters and numbers. You cannot use spaces or special characters.
func (client ObjectStorageClient) GetNamespace(ctx context.Context, request GetNamespaceRequest) (response GetNamespaceResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getNamespace, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetNamespaceResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) getNamespace(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/n/")
	if err != nil {
		return nil, err
	}

	var response GetNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetObject Gets the metadata and body of an object.
func (client ObjectStorageClient) GetObject(ctx context.Context, request GetObjectRequest) (response GetObjectResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getObject, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetObjectResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) getObject(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/o/{objectName}")
	if err != nil {
		return nil, err
	}

	var response GetObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetPreauthenticatedRequest Get the bucket level pre-authenticateted request
func (client ObjectStorageClient) GetPreauthenticatedRequest(ctx context.Context, request GetPreauthenticatedRequestRequest) (response GetPreauthenticatedRequestResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getPreauthenticatedRequest, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetPreauthenticatedRequestResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) getPreauthenticatedRequest(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/p/{parId}")
	if err != nil {
		return nil, err
	}

	var response GetPreauthenticatedRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// HeadBucket Efficiently checks if a bucket exists and gets the current ETag for the bucket.
func (client ObjectStorageClient) HeadBucket(ctx context.Context, request HeadBucketRequest) (response HeadBucketResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.headBucket, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(HeadBucketResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) headBucket(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodHead, "/n/{namespaceName}/b/{bucketName}/")
	if err != nil {
		return nil, err
	}

	var response HeadBucketResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// HeadObject Gets the user-defined metadata and entity tag for an object.
func (client ObjectStorageClient) HeadObject(ctx context.Context, request HeadObjectRequest) (response HeadObjectResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.headObject, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(HeadObjectResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) headObject(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodHead, "/n/{namespaceName}/b/{bucketName}/o/{objectName}")
	if err != nil {
		return nil, err
	}

	var response HeadObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListBuckets Gets a list of all `BucketSummary`s in a compartment. A `BucketSummary` contains only summary fields for the bucket
// and does not contain fields like the user-defined metadata.
// To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClient) ListBuckets(ctx context.Context, request ListBucketsRequest) (response ListBucketsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listBuckets, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListBucketsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) listBuckets(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/n/{namespaceName}/b/")
	if err != nil {
		return nil, err
	}

	var response ListBucketsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListMultipartUploadParts Lists the parts of an in-progress multipart upload.
func (client ObjectStorageClient) ListMultipartUploadParts(ctx context.Context, request ListMultipartUploadPartsRequest) (response ListMultipartUploadPartsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listMultipartUploadParts, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListMultipartUploadPartsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) listMultipartUploadParts(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/u/{objectName}")
	if err != nil {
		return nil, err
	}

	var response ListMultipartUploadPartsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListMultipartUploads Lists all in-progress multipart uploads for the given bucket in the given namespace.
func (client ObjectStorageClient) ListMultipartUploads(ctx context.Context, request ListMultipartUploadsRequest) (response ListMultipartUploadsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listMultipartUploads, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListMultipartUploadsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) listMultipartUploads(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/u")
	if err != nil {
		return nil, err
	}

	var response ListMultipartUploadsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListObjects Lists the objects in a bucket.
// To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClient) ListObjects(ctx context.Context, request ListObjectsRequest) (response ListObjectsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listObjects, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListObjectsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) listObjects(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/o")
	if err != nil {
		return nil, err
	}

	var response ListObjectsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListPreauthenticatedRequests List pre-authenticated requests for the bucket
func (client ObjectStorageClient) ListPreauthenticatedRequests(ctx context.Context, request ListPreauthenticatedRequestsRequest) (response ListPreauthenticatedRequestsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listPreauthenticatedRequests, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListPreauthenticatedRequestsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) listPreauthenticatedRequests(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/n/{namespaceName}/b/{bucketName}/p/")
	if err != nil {
		return nil, err
	}

	var response ListPreauthenticatedRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// PutObject Creates a new object or overwrites an existing one.
// To use this and other API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
func (client ObjectStorageClient) PutObject(ctx context.Context, request PutObjectRequest) (response PutObjectResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.putObject, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(PutObjectResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) putObject(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/n/{namespaceName}/b/{bucketName}/o/{objectName}")
	if err != nil {
		return nil, err
	}

	var response PutObjectResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateBucket Performs a partial or full update of a bucket's user-defined metadata.
func (client ObjectStorageClient) UpdateBucket(ctx context.Context, request UpdateBucketRequest) (response UpdateBucketResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateBucket, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateBucketResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) updateBucket(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/n/{namespaceName}/b/{bucketName}/")
	if err != nil {
		return nil, err
	}

	var response UpdateBucketResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UploadPart Uploads a single part of a multipart upload.
func (client ObjectStorageClient) UploadPart(ctx context.Context, request UploadPartRequest) (response UploadPartResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.uploadPart, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UploadPartResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client ObjectStorageClient) uploadPart(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/n/{namespaceName}/b/{bucketName}/u/{objectName}")
	if err != nil {
		return nil, err
	}

	var response UploadPartResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}
