// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

type ExportImageViaObjectStorageTupleDetails struct {

	// The destination type. Use `objectStorageTuple` when specifying the namespace, bucket name, and object name.
	// Use `objectStorageUri` when specifying the Object Storage URL.
	DestinationType *string `mandatory:"true" json:"destinationType,omitempty"`

	// The Object Storage bucket to export the image to.
	BucketName *string `mandatory:"false" json:"bucketName,omitempty"`

	// The Object Storage namespace to export the image to.
	NamespaceName *string `mandatory:"false" json:"namespaceName,omitempty"`

	// The Object Storage object name for the exported image.
	ObjectName *string `mandatory:"false" json:"objectName,omitempty"`
}

func (model ExportImageViaObjectStorageTupleDetails) String() string {
	return common.PointerString(model)
}
