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

type ImageSourceDetails struct {

	// The source type for the image. Use `objectStorageTuple` when specifying the namespace,
	// bucket name, and object name. Use `objectStorageUri` when specifying the Object Storage URL.
	SourceType *string `mandatory:"true" json:"sourceType,omitempty"`
}

func (model ImageSourceDetails) String() string {
	return common.PointerString(model)
}
