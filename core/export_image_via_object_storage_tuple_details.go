// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"encoding/json"
)

// ExportImageViaObjectStorageTupleDetails.
type ExportImageViaObjectStorageTupleDetails struct {

	// The Object Storage bucket to export the image to.
	BucketName *string `mandatory:"false" json:"bucketName,omitempty"`

	// The Object Storage namespace to export the image to.
	NamespaceName *string `mandatory:"false" json:"namespaceName,omitempty"`

	// The Object Storage object name for the exported image.
	ObjectName *string `mandatory:"false" json:"objectName,omitempty"`
}

func (m ExportImageViaObjectStorageTupleDetails) String() string {
	return common.PointerString(m)
}

func (m ExportImageViaObjectStorageTupleDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExportImageViaObjectStorageTupleDetails ExportImageViaObjectStorageTupleDetails
	s := struct {
		DiscriminatorParam string `json:"destinationType"`
		MarshalTypeExportImageViaObjectStorageTupleDetails
	}{
		"objectStorageTuple",
		(MarshalTypeExportImageViaObjectStorageTupleDetails)(m),
	}

	return json.Marshal(&s)
}
