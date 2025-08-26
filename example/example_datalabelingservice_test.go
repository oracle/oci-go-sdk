// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package example

import (
	"context"
	"fmt"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/datalabelingservice"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
)

// Example_createDataset to create a dataset in a given compartment Id.
func Example_createDataset() {
	compartment := "REPLACE_WITH_COMPARTMENT_OCID"
	annotationFormat := "REPLACE_WITH_ANNOTATION_FORMAT"
	datasetFormat := "REPLACE_WITH_DATASET_FORMAT"
	namespace := "REPLACE_WITH_OBJECT_STORAGE_NAMESPACE"
	bucketName := "REPLACE_WITH_OBJECT_STORAGE_BUCKET_NAME"
	displayName := "REPLACE_WITH_DATASET_DISPLAY_NAME"
	labelString := "REPLACE_WITH_LABEL_NAME"
	description := "REPLACE_WITH_DATASET_DESCRIPTION"

	client, err := datalabelingservice.NewDataLabelingManagementClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		panic(err)
	}

	fmt.Println("Creating dataset")

	var addLabels []datalabelingservice.Label

	addLabels = append(addLabels, datalabelingservice.Label{Name: common.String(labelString)})

	// createDatasetRequest for the CreateDataset operation
	createDatasetRequest := datalabelingservice.CreateDatasetRequest{
		CreateDatasetDetails: datalabelingservice.CreateDatasetDetails{
			CompartmentId:                        common.String(compartment),
			AnnotationFormat:                     common.String(annotationFormat),
			DatasetSourceDetails:                 map[string]interface{}{"sourceType": "OBJECT_STORAGE", "namespace": namespace, "bucket": bucketName},
			DatasetFormatDetails:                 map[string]interface{}{"formatType": datasetFormat},
			LabelSet:                             &datalabelingservice.LabelSet{Items: addLabels},
			DisplayName:                          common.String(displayName),
			Description:                          common.String(description),
			InitialRecordGenerationConfiguration: nil,
			FreeformTags:                         nil,
		},
	}

	// Send the request using the service client
	_, datasetErr := client.CreateDataset(context.Background(), createDatasetRequest)
	helpers.FatalIfError(datasetErr)

	fmt.Println("Dataset creation completed")
	fmt.Println("Done")
}

// Example_getDataset to get dataset details from a given dataset Id.
func Example_getDataset() {
	datasetId := "REPLACE_WITH_DATASET_OCID"

	client, err := datalabelingservice.NewDataLabelingManagementClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		panic(err)
	}

	fmt.Println("Getting dataset")
	getDatasetRequest := datalabelingservice.GetDatasetRequest{
		DatasetId: common.String(datasetId),
	}

	// Send the request using the service client
	_, datasetErr := client.GetDataset(context.Background(), getDatasetRequest)
	helpers.FatalIfError(datasetErr)

	fmt.Println("Done.")
}

// Example_listDataset is to list all dataset in a given compartment Id.
func Example_listDataset() {
	compartment := "REPLACE_WITH_COMPARTMENT_OCID"

	client, err := datalabelingservice.NewDataLabelingManagementClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		panic(err)
	}

	fmt.Println("Listing all datasets")

	req := datalabelingservice.ListDatasetsRequest{
		CompartmentId:  common.String(compartment),
		Limit:          common.Int(500),
		LifecycleState: datalabelingservice.DatasetLifecycleStateActive,
	}

	// Send the request using the service client
	_, datasetErr := client.ListDatasets(context.Background(), req)
	helpers.FatalIfError(datasetErr)
	fmt.Println("Listing datasets completed")
}

// Example_deleteDataset to  delete a existing dataset.
func Example_deleteDataset() {
	datasetId := "REPLACE_WITH_DATASET_OCID"

	client, err := datalabelingservice.NewDataLabelingManagementClientWithConfigurationProvider(common.DefaultConfigProvider())
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleting dataset")

	// deleteDatasetRequest create request for deletion of dataset
	deleteDatasetRequest := datalabelingservice.DeleteDatasetRequest{
		DatasetId: common.String(datasetId),
	}

	// Send the request using the service client
	_, datasetErr := client.DeleteDataset(context.Background(), deleteDatasetRequest)
	helpers.FatalIfError(datasetErr)
	fmt.Println("Dataset deleted")

}
