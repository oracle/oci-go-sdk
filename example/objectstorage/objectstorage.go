package objectstorage

import (
	"context"
	"fmt"
	"os"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/objectstorage"
)

// CreateBucket does create a new buket
func CreateBucket(namespaceName, bucketName, compartmentId string) (err error) {
	// create object storage client from default config provider
	client, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())

	if err != nil {
		fmt.Printf("unable to initialize object storage client with error: %s \n", err)
		return
	}

	createBucketRequest := objectstorage.CreateBucketRequest{
		NamespaceName: &namespaceName,
	}
	createBucketRequest.CompartmentId = &compartmentId
	createBucketRequest.Name = &bucketName

	// call the CreateBucket API
	response, err := client.CreateBucket(context.Background(), createBucketRequest)

	if response.RawResponse.StatusCode != 200 {
		fmt.Printf("CreateBucket failed with error: %s \n", err)
		if response.RawResponse.StatusCode == 409 {
			return nil // bucket alrady been created, ignore the error here
		}
	}

	fmt.Println("CreateBucket succeed")
	return
}

// UploadFileToBucket does upload an file to bucket
func UploadFileToBucket(namespaceName, bucketName, objectName, filePath string) {
	// create object storage client from default config provider
	client, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())

	if err != nil {
		fmt.Printf("unable to initialize object storage client with error: %s \n", err)
		return
	}

	// Path of a a file to be uploaded
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		fmt.Printf("unable to open the file with error: %s \n", err)
		return
	}

	fileInfo, _ := file.Stat()

	// The content length of the body.
	contentLength := int(fileInfo.Size())

	putObjectRequest := objectstorage.PutObjectRequest{
		NamespaceName: &namespaceName,
		BucketName:    &bucketName,
		ObjectName:    &objectName, // The name of the object. Avoid entering confidential information.
		ContentLength: &contentLength,
		PutObjectBody: file,
	}

	putObjectError := client.PutObject(context.Background(), putObjectRequest)

	if putObjectError != nil {
		fmt.Printf("PutObject failed with error: %s \n", putObjectError)
		return
	}

	fmt.Println("Upload file succeed.")
}
