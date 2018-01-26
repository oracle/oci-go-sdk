// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// APIs for managing buckets and objects.
//

package integtest

import (
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/objectstorage"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"io/ioutil"
	"io"
	"os"
	"path"
	"crypto/sha256"
	"encoding/hex"
	"compress/gzip"
	"bytes"
	"time"
)


func getNamespace(t *testing.T) string {
	c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	request := objectstorage.GetNamespaceRequest{}
	r, err := c.GetNamespace(context.Background(), request)
	failIfError(t, err)
	return *r.Value
}

func getObject(t *testing.T, namespace, bucketname, objectname string) (objectstorage.GetObjectResponse, error){
	c, _ := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	request := objectstorage.GetObjectRequest{
		NamespaceName:      &namespace,
		BucketName:         &bucketname,
		ObjectName:         &objectname,
	}

	return c.GetObject(context.Background(), request)
}

func putObject(t *testing.T, namespace, bucketname, objectname string, contentLen int, content io.ReadCloser) error {
	c, _ := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	request := objectstorage.PutObjectRequest{
		NamespaceName: &namespace,
		BucketName: &bucketname,
		ObjectName: &objectname,
		ContentLength:&contentLen,
		PutObjectBody:content,
	}
	return c.PutObject(context.Background(), request)
}

func createBucket(t *testing.T, namespace, compartment, name string){
	c, _ := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	request := objectstorage.CreateBucketRequest{
		NamespaceName:&namespace,

	}
	request.CompartmentId = &compartment
	request.Name = &name
	request.Metadata = make(map[string]string)
	request.PublicAccessType = objectstorage.CreateBucketDetailsPublicAccessTypeNopublicaccess
	_, err := c.CreateBucket(context.Background(), request)
	failIfError(t, err)
	return
}

func deleteBucket(t *testing.T, namespace, name string)(err error){
	c, _ := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	request := objectstorage.DeleteBucketRequest{
		NamespaceName:&namespace,
		BucketName:&name,
	}
	err = c.DeleteBucket(context.Background(), request)
	failIfError(t, err)
	return
}

func deleteObject(t *testing.T, namespace, bucketname, objectname string)(err error){
	c,_ := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	request := objectstorage.DeleteObjectRequest{
		NamespaceName:&namespace,
		BucketName: &bucketname,
		ObjectName:&objectname,
	}
	err = c.DeleteObject(context.Background(), request)
	failIfError(t, err)
	return
}

func TestObjectStorageClient_GetNamespace(t *testing.T) {
	namespace := getNamespace(t)
	assert.NotEmpty(t, namespace)
	return
}


func TestObjectStorageClient_BigFile(t *testing.T) {
	bname := getUniqueName("largeBucket")
	namespace := getNamespace(t)

	createBucket(t, getNamespace(t), getTenancyID(), bname)
	defer deleteBucket(t, namespace, bname)

	contentlen := 1024 * 10000
	filepath, filesize, expectedHash := writeTempFileOfSize(int64(contentlen))
	filename := path.Base(filepath)
	fmt.Println("uploading ", filepath)
	defer removeFileFn(filepath)
	file, e := os.Open(filepath)
	defer file.Close()
	failIfError(t, e)

	e = putObject(t, namespace , bname, filename, int(filesize), file)
	failIfError(t, e)
	fmt.Println(expectedHash)
	rGet, e := getObject(t, namespace, bname, filename)
	failIfError(t, e)
	defer deleteObject(t, namespace, bname, filename)

	h := sha256.New()
	_, e = io.Copy(h, rGet.Content)
	rGet.Content.Close()
	actualHash := hex.EncodeToString(h.Sum(nil))
	assert.NoError(t, e)
	assert.Equal(t, filesize, int64(*rGet.ContentLength))
	assert.Equal(t, "application/octet-stream", *rGet.ContentType)
	assert.Equal(t, expectedHash, actualHash)
}

func TestObjectStorage_GzipFileEncoding(t *testing.T) {
	bname := getUniqueName("BucketWithZip")
	objname := getUniqueName("gzippedcontent")
	namespace := getNamespace(t)

	createBucket(t, getNamespace(t), getTenancyID(), bname)
	defer deleteBucket(t, namespace, bname)


	message := " some random content that will get gzipped"
	zBytes := bytes.Buffer{}
	gz := gzip.NewWriter(&zBytes)
	gz.Write([]byte(message))
	gz.Close()
	c, _ := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	request := objectstorage.PutObjectRequest{
		NamespaceName: &namespace,
		BucketName: &bname,
		ObjectName: &objname,
		ContentLength:common.Int(zBytes.Len()),
		PutObjectBody:ioutil.NopCloser(&zBytes),
		ContentType:common.String("text/plain"),
		ContentEncoding:common.String("gzip"),
	}
	e := c.PutObject(context.Background(), request)
	defer deleteObject(t, namespace, bname, objname)
	failIfError(t, e)

	response, e := getObject(t, namespace, bname, objname)
	failIfError(t, e)

	//The file is served back by its Content-Type, thus it should not be gzipped anymore
	content, e := ioutil.ReadAll(response.Content)
	failIfError(t, e)
	assert.Equal(t, message, string(content))
}

func TestObjectStorageClient_Object(t *testing.T) {
	bname := getUniqueName("bucket")
	data := "some temp data"
	namespace := getNamespace(t)

	createBucket(t, getNamespace(t), getTenancyID(), bname)
	defer deleteBucket(t, namespace, bname)

	contentlen := len([]byte(data))
	filepath := writeTempFile(data)
	filename := path.Base(filepath)
	defer removeFileFn(filepath)
	file, e := os.Open(filepath)
	defer file.Close()
	failIfError(t, e)
	e = putObject(t, namespace , bname, filename, contentlen, file)
	failIfError(t, e)

	r, e := getObject(t, namespace, bname, filename)
	failIfError(t, e)
	defer deleteObject(t, namespace, bname, filename)
	defer r.Content.Close()
	bytes, e := ioutil.ReadAll(r.Content)
	failIfError(t, e)
	assert.Equal(t, contentlen, *r.ContentLength)
	assert.Equal(t, data, string(bytes))
	return
}


func TestObjectStorageClient_AbortUpload(t *testing.T) {
	bname := getUniqueName("abortUpload")
	namespace := getNamespace(t)

	createBucket(t, getNamespace(t), getTenancyID(), bname)
	defer deleteBucket(t, namespace, bname)

	contentlen := 1024 * 100
	filepath, filesize, _ := writeTempFileOfSize(int64(contentlen))
	filename := path.Base(filepath)
	defer removeFileFn(filepath)
	file, e := os.Open(filepath)
	defer file.Close()
	failIfError(t, e)

	c, _ := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	request := objectstorage.PutObjectRequest{
		NamespaceName: &namespace,
		BucketName: &bname,
		ObjectName: &filename,
		ContentLength:common.Int(int(filesize)),
		PutObjectBody:file,
	}
	ctx, cancelFn := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelFn()
	e = c.PutObject(ctx, request)
	assert.Error(t, e)
}

func TestObjectStorageClient_AbortMultipartUpload(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, e)
	request := objectstorage.AbortMultipartUploadRequest{}
	err := c.AbortMultipartUpload(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_CommitMultipartUpload(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, e)
	request := objectstorage.CommitMultipartUploadRequest{}
	err := c.CommitMultipartUpload(context.Background(), request)
	assert.NoError(t, err)
	return
}


func TestObjectStorageClient_CreateMultipartUpload(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, e)
	request := objectstorage.CreateMultipartUploadRequest{}
	r, err := c.CreateMultipartUpload(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_CreatePreauthenticatedRequest(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, e)
	request := objectstorage.CreatePreauthenticatedRequestRequest{}
	r, err := c.CreatePreauthenticatedRequest(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_DeletePreauthenticatedRequest(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, e)
	request := objectstorage.DeletePreauthenticatedRequestRequest{}
	err := c.DeletePreauthenticatedRequest(context.Background(), request)
	assert.NoError(t, err)
	return
}


func TestObjectStorageClient_GetPreauthenticatedRequest(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, e)
	request := objectstorage.GetPreauthenticatedRequestRequest{}
	r, err := c.GetPreauthenticatedRequest(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_HeadBucket(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, e)
	request := objectstorage.HeadBucketRequest{}
	err := c.HeadBucket(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_HeadObject(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, e)
	request := objectstorage.HeadObjectRequest{}
	err := c.HeadObject(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_ListBuckets(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, e)
	request := objectstorage.ListBucketsRequest{}
	r, err := c.ListBuckets(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_ListMultipartUploadParts(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, e)
	request := objectstorage.ListMultipartUploadPartsRequest{}
	r, err := c.ListMultipartUploadParts(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_ListMultipartUploads(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, e)
	request := objectstorage.ListMultipartUploadsRequest{}
	r, err := c.ListMultipartUploads(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_ListObjects(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, e)
	request := objectstorage.ListObjectsRequest{}
	r, err := c.ListObjects(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_ListPreauthenticatedRequests(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, e)
	request := objectstorage.ListPreauthenticatedRequestsRequest{}
	r, err := c.ListPreauthenticatedRequests(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}


func TestObjectStorageClient_UpdateBucket(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, e)
	request := objectstorage.UpdateBucketRequest{}
	r, err := c.UpdateBucket(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestObjectStorageClient_UploadPart(t *testing.T) {
	t.Skip("Not implemented")
	c, e := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, e)
	request := objectstorage.UploadPartRequest{}
	err := c.UploadPart(context.Background(), request)
	assert.NoError(t, err)
	return
}
