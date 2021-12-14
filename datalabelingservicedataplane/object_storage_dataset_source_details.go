// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DlsDataPlane API
//
// A description of the DlsDataPlane API.
//

package datalabelingservicedataplane

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v54/common"
)

// ObjectStorageDatasetSourceDetails Specifies the dataset location in object storage. This requires that all records are in this bucket, and under this prefix. We do not support a dataset with objects in arbitrary locations across buckets or prefixes.
type ObjectStorageDatasetSourceDetails struct {

	// Namespace of the bucket that contains the dataset data source
	Namespace *string `mandatory:"true" json:"namespace"`

	// The object storage bucket that contains the dataset data source
	Bucket *string `mandatory:"true" json:"bucket"`

	// A common path prefix shared by the objects that make up the dataset.
	Prefix *string `mandatory:"false" json:"prefix"`
}

func (m ObjectStorageDatasetSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ObjectStorageDatasetSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageDatasetSourceDetails ObjectStorageDatasetSourceDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeObjectStorageDatasetSourceDetails
	}{
		"OBJECT_STORAGE",
		(MarshalTypeObjectStorageDatasetSourceDetails)(m),
	}

	return json.Marshal(&s)
}
