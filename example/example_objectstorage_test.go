// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for Object Storage Service API

package example

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
	"github.com/oracle/oci-go-sdk/v65/objectstorage"
	"github.com/oracle/oci-go-sdk/v65/objectstorage/transfer"
)

// Example_objectStorage_UploadFile shows how to create a bucket and upload a file
func Example_objectStorage_UploadFile() {
	c, clerr := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)

	ctx := context.Background()
	bname := helpers.GetRandomString(8)
	namespace := getNamespace(ctx, c)

	createBucket(ctx, c, namespace, bname)
	defer deleteBucket(ctx, c, namespace, bname)

	contentlen := 1024 * 1000
	fpath, filesize := helpers.WriteTempFileOfSize(int64(contentlen))
	filename := filepath.Base(fpath)
	defer func() {
		os.Remove(filename)
	}()

	file, e := os.Open(fpath)

	if e != nil {
		file.Close()
		helpers.FatalIfError(e)
	} else {
		defer file.Close()
	}

	e = putObject(ctx, c, namespace, bname, filename, filesize, file, nil)
	helpers.FatalIfError(e)
	defer deleteObject(ctx, c, namespace, bname, filename)

	// Output:
	// get namespace
	// create bucket
	// put object
	// delete object
	// delete bucket
}

func Example_objectStorage_UploadManager_UploadFile() {
	c, clerr := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)
	// Disable timeout to support big file upload(Once need to specify the os client for Upload Manager)
	c.HTTPClient = &http.Client{}

	ctx := context.Background()
	bname := "bname"
	namespace := getNamespace(ctx, c)

	createBucket(ctx, c, namespace, bname)
	defer deleteBucket(ctx, c, namespace, bname)

	contentlen := 1024 * 1024 * 300 // 300MB
	fpath, _ := helpers.WriteTempFileOfSize(int64(contentlen))
	filename := filepath.Base(fpath)
	defer os.Remove(filename)

	uploadManager := transfer.NewUploadManager()
	objectName := "sampleFileUploadObj"

	req := transfer.UploadFileRequest{
		UploadRequest: transfer.UploadRequest{
			NamespaceName:                       common.String(namespace),
			BucketName:                          common.String(bname),
			ObjectName:                          common.String(objectName),
			PartSize:                            common.Int64(128 * 1024 * 1024),
			CallBack:                            callBack,
			ObjectStorageClient:                 &c,
			EnableMultipartChecksumVerification: common.Bool(true),
		},
		FilePath: fpath,
	}

	// if you want to overwrite default value, you can do it
	// as: transfer.UploadRequest.AllowMultipartUploads = common.Bool(false) // default is true
	// or: transfer.UploadRequest.AllowParrallelUploads = common.Bool(false) // default is true
	resp, err := uploadManager.UploadFile(ctx, req)

	if err != nil && resp.IsResumable() {
		resp, err = uploadManager.ResumeUploadFile(ctx, *resp.MultipartUploadResponse.UploadID)
		if err != nil {
			fmt.Println(resp)
		}
	}

	defer deleteObject(ctx, c, namespace, bname, objectName)
	fmt.Println("file uploaded")

	// Output:
	// get namespace
	// create bucket
	// One example of progress bar could be the above comment content.
	// One example of progress bar could be the above comment content.
	// One example of progress bar could be the above comment content.
	// file uploaded
	// delete object
	// delete bucket
}

func callBack(multiPartUploadPart transfer.MultiPartUploadPart) {
	if nil == multiPartUploadPart.Err {
		// Please refer this as the progress bar print content.
		// fmt.Printf("Part: %d / %d is uploaded.\n", multiPartUploadPart.PartNum, multiPartUploadPart.TotalParts)
		fmt.Printf("One example of progress bar could be the above comment content.\n")
		// Please refer following fmt to get each part opc-md5 res.
		// fmt.Printf("and this part opcMD5(64BasedEncoding) is: %s.\n", *multiPartUploadPart.OpcMD5 )
	}
}

func Example_objectStorage_UploadManager_Stream() {
	c, clerr := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)

	ctx := context.Background()
	bname := "bname"
	namespace := getNamespace(ctx, c)

	createBucket(ctx, c, namespace, bname)
	defer deleteBucket(ctx, c, namespace, bname)

	contentlen := 1024 * 1000 * 130 // 130MB
	fpath, _ := helpers.WriteTempFileOfSize(int64(contentlen))
	filename := filepath.Base(fpath)
	defer func() {
		os.Remove(filename)
	}()

	uploadManager := transfer.NewUploadManager()
	objectName := "sampleStreamUploadObj"

	file, _ := os.Open(fpath)
	defer file.Close()

	req := transfer.UploadStreamRequest{
		UploadRequest: transfer.UploadRequest{
			NamespaceName:                       common.String(namespace),
			BucketName:                          common.String(bname),
			ObjectName:                          common.String(objectName),
			EnableMultipartChecksumVerification: common.Bool(true),
		},
		StreamReader: file, // any struct implements the io.Reader interface
	}

	// if you want to overwrite default value, you can do it
	// as: transfer.UploadRequest.AllowMultipartUploads = common.Bool(false) // default is true
	// or: transfer.UploadRequest.AllowParallelUploads = common.Bool(false) // default is true
	_, err := uploadManager.UploadStream(context.Background(), req)

	if err != nil {
		fmt.Println(err)
	}

	defer deleteObject(ctx, c, namespace, bname, objectName)
	fmt.Println("stream uploaded")

	// Output:
	// get namespace
	// create bucket
	// stream uploaded
	// delete object
	// delete bucket
}

// Example for getting Object Storage namespace of a tenancy that is not their own. This
// is useful in cross-tenant Object Storage operations. Object Storage namespace can be retrieved using the
// compartment id of the target tenancy if the user has necessary permissions to access that tenancy.
//
// For example if Tenant A wants to access Tenant B's object storage namespace then Tenant A has to define
// a policy similar to following:
//
// DEFINE TENANCY TenantB AS <TenantB OCID>
// ENDORSE GROUP <TenantA user group name> TO {OBJECTSTORAGE_NAMESPACE_READ} IN TENANCY TenantB
//
// and Tenant B should add a policy similar to following:
//
// DEFINE TENANCY TenantA AS <TenantA OCID>
// DEFINE GROUP TenantAGroup AS <TenantA user group OCID>
// ADMIT GROUP TenantAGroup OF TENANCY TenantA TO {OBJECTSTORAGE_NAMESPACE_READ} IN TENANCY
//
// This example covers only GetNamespace operation across tenants. Additional permissions
// will be required to perform more Object Storage operations.
//
// Example_objectStorage_GetNamespace shows how to get namespace providing compartmentId.
func Example_objectStorage_GetNamespace() {
	c, clerr := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)

	ctx := context.Background()

	request := objectstorage.GetNamespaceRequest{}
	request.CompartmentId = helpers.CompartmentID()

	r, err := c.GetNamespace(ctx, request)
	helpers.FatalIfError(err)

	log.Printf("Namespace for compartment %s is: %s", *request.CompartmentId, *r.Value)

	fmt.Println("Namespace retrieved")

	// Output:
	// Namespace retrieved
}

func Example_objectStorage_GetObjectUsingRealmSpecificEndpoint() {
	// This example shows how to use realm specific endpoint to get object.
	// You can select either this environment variable or the customClientConfiguration to enable realm specific endpoint.
	os.Setenv("OCI_REALM_SPECIFIC_SERVICE_ENDPOINT_TEMPLATE_ENABLED", "true")

	c, clerr := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)

	// An alternative way to enable realm specific endpoint is to use the following code.
	c.SetCustomClientConfiguration(common.CustomClientConfiguration{
		RealmSpecificServiceEndpointTemplateEnabled: common.Bool(true),
	})
	ctx := context.Background()
	bname := helpers.GetRandomString(8)
	namespace := getNamespace(ctx, c)
	getRequest := objectstorage.GetObjectRequest{
		NamespaceName: common.String(namespace),
		BucketName:    common.String(bname),
		ObjectName:    common.String("Example_objectStorage_GetObjectUsingRealmSpecificEndpoint"),
	}

	response, err := c.GetObject(context.Background(), getRequest)
	if err != nil {
		fmt.Println("404")
		return
	}
	fmt.Println(response)
	// Output:
	// 404
}

func Example_objectStorage_EnableDualStackEndpoints() {

	c, clerr := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)

	// Dual stack endpoints can be enabled at the client level as shown, or through the
	// environment variable OCI_DUAL_STACK_ENDPOINT_ENABLED
	c.EnableDualStackEndpoints(true)

	ctx := context.Background()
	bname := helpers.GetRandomString(8)
	namespace := getNamespace(ctx, c)

	createBucket(ctx, c, namespace, bname)

	defer deleteBucket(ctx, c, namespace, bname)

	// Output:
	// create bucket
	// delete bucket
}

func getNamespace(ctx context.Context, c objectstorage.ObjectStorageClient) string {
	request := objectstorage.GetNamespaceRequest{}
	r, err := c.GetNamespace(ctx, request)
	helpers.FatalIfError(err)
	fmt.Println("get namespace")
	return *r.Value
}

func putObject(ctx context.Context, c objectstorage.ObjectStorageClient, namespace, bucketname, objectname string, contentLen int64, content io.ReadCloser, metadata map[string]string) error {
	request := objectstorage.PutObjectRequest{
		NamespaceName: &namespace,
		BucketName:    &bucketname,
		ObjectName:    &objectname,
		ContentLength: &contentLen,
		PutObjectBody: content,
		OpcMeta:       metadata,
	}
	_, err := c.PutObject(ctx, request)
	fmt.Println("put object")
	return err
}

func deleteObject(ctx context.Context, c objectstorage.ObjectStorageClient, namespace, bucketname, objectname string) (err error) {
	request := objectstorage.DeleteObjectRequest{
		NamespaceName: &namespace,
		BucketName:    &bucketname,
		ObjectName:    &objectname,
	}
	_, err = c.DeleteObject(ctx, request)
	helpers.FatalIfError(err)
	fmt.Println("delete object")
	return
}

func createBucket(ctx context.Context, c objectstorage.ObjectStorageClient, namespace, name string) {
	request := objectstorage.CreateBucketRequest{
		NamespaceName: &namespace,
	}
	request.CompartmentId = helpers.CompartmentID()
	request.Name = &name
	request.Metadata = make(map[string]string)
	request.PublicAccessType = objectstorage.CreateBucketDetailsPublicAccessTypeNopublicaccess
	_, err := c.CreateBucket(ctx, request)
	helpers.FatalIfError(err)

	fmt.Println("create bucket")
}

func deleteBucket(ctx context.Context, c objectstorage.ObjectStorageClient, namespace, name string) (err error) {
	request := objectstorage.DeleteBucketRequest{
		NamespaceName: &namespace,
		BucketName:    &name,
	}
	_, err = c.DeleteBucket(ctx, request)
	helpers.FatalIfError(err)

	fmt.Println("delete bucket")
	return
}
