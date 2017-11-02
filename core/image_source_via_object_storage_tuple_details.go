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

type ImageSourceViaObjectStorageTupleDetails struct {

	// The source type for the image. Use `objectStorageTuple` when specifying the namespace,
	// bucket name, and object name. Use `objectStorageUri` when specifying the Object Storage URL.
	SourceType *string `mandatory:"true" json:"sourceType,omitempty"`

	// The Object Storage bucket for the image.
	BucketName *string `mandatory:"true" json:"bucketName,omitempty"`

	// The Object Storage namespace for the image.
	NamespaceName *string `mandatory:"true" json:"namespaceName,omitempty"`

	// The Object Storage name for the image.
	ObjectName *string `mandatory:"true" json:"objectName,omitempty"`
}

func (model ImageSourceViaObjectStorageTupleDetails) String() string {
	return common.PointerString(model)
}
