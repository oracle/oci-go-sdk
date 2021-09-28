// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for Object Storage multipart download

package example

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"sync"

	"github.com/oracle/oci-go-sdk/v48/common"
	"github.com/oracle/oci-go-sdk/v48/example/helpers"
	"github.com/oracle/oci-go-sdk/v48/objectstorage"
)

// ExampleMultipartDownload shows how to use get object API to perform multi-part download operation
func ExampleMultipartDownload() {
	c, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	ctx := context.Background()
	// Change the bname and objectName to the file name and bucket you want you download from
	bname := "test_bucket_name"
	objectName := "test_download_file"
	// This value controls the size per part
	partSize := 500
	namespace := getNamespace(ctx, c)
	downloadThread := 5

	// Get the object size info from object storage
	listResponse, err := c.ListObjects(ctx, objectstorage.ListObjectsRequest{
		NamespaceName: common.String(namespace),
		BucketName:    common.String(bname),
		Prefix:        common.String(objectName),
		Fields:        common.String("name,size"),
	})
	helpers.FatalIfError(err)
	// The result will return a list of objects with the required name, we just select the first in this example
	size := int(*listResponse.Objects[0].Size)
	totalParts := size / partSize
	if size%partSize != 0 {
		totalParts++
	}

	done := make(chan struct{})
	prepareDownloadParts := splitToParts(done, totalParts, partSize, size, namespace, bname, objectName)

	downloadedParts := multipartDownload(ctx, c, downloadThread, done, prepareDownloadParts)

	// In this example, we're storing the download content in memory, please be aware of any issue with oom
	result := make([]byte, size)

	for part := range downloadedParts {
		if part.err != nil {
			// User should properly handle failure here, can be either raise an fatal error or retry to download the error part
			// For this example, we simply ignore the error handling here
			continue
		}
		for i := int64(0); i < part.size; i++ {
			result[i+part.offset] = part.partBody[i]
		}
	}
	fmt.Println(result)
}

// downloadPart contains the data downloaded from object storage and the body part info
type downloadPart struct {
	size     int64
	partBody []byte
	offset   int64
	partNum  int
	err      error
}

// prepareDownloadPart wraps an GetObjectRequest with splitted part related info
type prepareDownloadPart struct {
	request objectstorage.GetObjectRequest
	offset  int64
	partNum int
	size    int64
}

// splitToParts splits the file to the partSize and build a new struct to prepare for multipart download, this function will return a prepareDownloadPart channel
func splitToParts(done <-chan struct{}, totalParts int, partSize int, fileSize int, namespace string, bname string, objectName string) chan prepareDownloadPart {
	prepareDownloadParts := make(chan prepareDownloadPart)
	go func() {
		defer func() {
			fmt.Println("Split to parts completed, closing channel")
			close(prepareDownloadParts)
		}()

		for part := 0; part < totalParts; part++ {
			start := int64(part * partSize)
			end := int64(math.Min(float64((part+1)*partSize), float64(fileSize)) - 1)
			bytesRange := strconv.FormatInt(start, 10) + "-" + strconv.FormatInt(end, 10)
			part := prepareDownloadPart{
				request: objectstorage.GetObjectRequest{
					NamespaceName: common.String(namespace),
					BucketName:    common.String(bname),
					ObjectName:    common.String(objectName),
					// This is the parameter where you control the download size/request
					Range: common.String("bytes=" + bytesRange),
				},
				offset:  start,
				partNum: part,
				size:    end - start,
			}

			select {
			case prepareDownloadParts <- part:
			case <-done:
				return
			}
		}
	}()
	return prepareDownloadParts
}

// multipartDownload will consume prepareDownloadPart from channel and from different gorountine, it will perform multipart download and save the download result to another channel
func multipartDownload(ctx context.Context, c objectstorage.ObjectStorageClient, downloadThreads int, done <-chan struct{}, prepareDownloadParts chan prepareDownloadPart) chan downloadPart {
	result := make(chan downloadPart)
	var wg sync.WaitGroup
	wg.Add(downloadThreads)

	for i := 0; i < downloadThreads; i++ {
		go func() {
			downloadFilePart(ctx, c, done, prepareDownloadParts, result)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

// downloadFilePart wraps objectStorage GetObject API call
func downloadFilePart(ctx context.Context, c objectstorage.ObjectStorageClient, done <-chan struct{}, prepareDownloadParts chan prepareDownloadPart, result chan downloadPart) {
	for part := range prepareDownloadParts {
		resp, err := c.GetObject(ctx, part.request)
		downloadedPart := downloadPart{}
		if err != nil {
			fmt.Println("Error in downloading: ", err)
			downloadedPart.err = err
		} else {
			content, _ := ioutil.ReadAll(resp.Content)
			downloadedPart = downloadPart{
				size:     int64(len(content)),
				partBody: content,
				offset:   part.offset,
				partNum:  part.partNum,
			}
		}
		select {
		case result <- downloadedPart:
		case <-done:
			fmt.Println("downloadParts received Done")
			return
		}
	}
}
