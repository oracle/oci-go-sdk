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

// ExportImageDetails. The destination details for the image export.
// Set `destinationType` to `objectStorageTuple`
// and use ExportImageViaObjectStorageTupleDetails
// when specifying the namespace, bucket name, and object name.
// Set `destinationType` to `objectStorageUri` and
// use ExportImageViaObjectStorageUriDetails
// when specifying the Object Storage URL.
type ExportImageDetails struct {

	// The destination type. Use `objectStorageTuple` when specifying the namespace, bucket name, and object name.
	// Use `objectStorageUri` when specifying the Object Storage URL.
	DestinationType *string `mandatory:"true" json:"destinationType,omitempty"`
}

func (model ExportImageDetails) String() string {
	return common.PointerString(model)
}
