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

type ImageSourceViaObjectStorageUriDetails struct {

	// The source type for the image. Use `objectStorageTuple` when specifying the namespace,
	// bucket name, and object name. Use `objectStorageUri` when specifying the Object Storage URL.
	SourceType *string `mandatory:"true" json:"sourceType,omitempty"`

	// The Object Storage URL for the image.
	SourceUri *string `mandatory:"true" json:"sourceUri,omitempty"`
}

func (model ImageSourceViaObjectStorageUriDetails) String() string {
	return common.PointerString(model)
}
