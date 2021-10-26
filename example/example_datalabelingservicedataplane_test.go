// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for Database API

package example

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v50/common"
	"github.com/oracle/oci-go-sdk/v50/datalabelingservicedataplane"
	"github.com/oracle/oci-go-sdk/v50/example/helpers"
)

func ExampleCreateRecord() {
	datasetId := "REPLACE_WITH_DATASET_OCID"
	compartment := "REPLACE_WITH_COMPARTMENT_OCID"
	namespace := "REPLACE_WITH_OBJECT_STORAGE_NAMESPACE"
	bucketName := "REPLACE_WITH_OBJECT_STORAGE_BUCKET_NAME"
	objectName := "REPLACE_WITH_OBJECT_NAME"

	client, err := datalabelingservicedataplane.NewDataLabelingClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		panic(err)
	}

	fmt.Println("Record creation")

	objectPath := fmt.Sprintf("/n/%s/b/%s/o/%s", namespace, bucketName, objectName)

	sourceDetails := datalabelingservicedataplane.ObjectStorageSourceDetails{
		RelativePath: common.String(objectName),
		Path:         common.String(objectPath),
	}
	name := common.String(objectName)

	createRecordRequest := datalabelingservicedataplane.CreateRecordRequest{
		CreateRecordDetails: datalabelingservicedataplane.CreateRecordDetails{
			Name:          name,
			DatasetId:     common.String(datasetId),
			CompartmentId: common.String(compartment),
			SourceDetails: sourceDetails,
			FreeformTags:  nil,
			DefinedTags:   nil,
		},
		OpcRetryToken:   nil,
		OpcRequestId:    nil,
		RequestMetadata: common.RequestMetadata{},
	}

	_, recordErr := client.CreateRecord(context.Background(), createRecordRequest)
	helpers.FatalIfError(recordErr)

	fmt.Println("Record creation succeeded")

	// Output:
	// Record creation.
	// Record creation succeeded.
}

func ExampleGetRecords() {
	recordId := "REPLACE_WITH_RECORD_OCID"

	client, err := datalabelingservicedataplane.NewDataLabelingClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		panic(err)
	}
	fmt.Println("Getting record.")

	// getRecordRequest creating a request body
	getRecordRequest := datalabelingservicedataplane.GetRecordRequest{
		RecordId:        common.String(recordId),
		OpcRequestId:    nil,
		RequestMetadata: common.RequestMetadata{},
	}

	// Send the request using the service client
	_, recordErr := client.GetRecord(context.Background(), getRecordRequest)
	helpers.FatalIfError(recordErr)

	fmt.Println("Get record succeeded.")
	fmt.Println("Done")

	// Output:
	// Getting record.
	// Get record succeeded.
	// Done
}

func ExampleListRecords() {
	datasetId := "REPLACE_WITH_DATASET_OCID"
	compartment := "REPLACE_WITH_COMPARTMENT_OCID"

	client, err := datalabelingservicedataplane.NewDataLabelingClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		panic(err)
	}
	fmt.Println("Listing records.")

	listRecordsRequest := datalabelingservicedataplane.ListRecordsRequest{
		CompartmentId: common.String(compartment),
		DatasetId:     common.String(datasetId),
		Limit:         common.Int(500),
		SortBy:        datalabelingservicedataplane.ListRecordsSortByTimecreated,
		SortOrder:     datalabelingservicedataplane.ListRecordsSortOrderDesc,
	}

	// Send the request using the service client
	_, recordErr := client.ListRecords(context.Background(), listRecordsRequest)
	helpers.FatalIfError(recordErr)

	fmt.Println("Record listing succeeded.")
	fmt.Println("Done")

	// Output:
	// Listing records.
	// Record listing succeeded.
	// Done
}

func ExampleDeleteRecord() {
	recordId := "REPLACE_WITH_RECORD_OCID"

	client, err := datalabelingservicedataplane.NewDataLabelingClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleting record.")

	deleteRecordRequest := datalabelingservicedataplane.DeleteRecordRequest{
		RecordId: common.String(recordId),
	}
	_, recordErr := client.DeleteRecord(context.Background(), deleteRecordRequest)
	helpers.FatalIfError(recordErr)

	fmt.Println("Record deletion succeeded.")
	fmt.Println("Done")

	// Output:
	// Deleting record.
	// Record deletion succeeded.
	// Done
}
