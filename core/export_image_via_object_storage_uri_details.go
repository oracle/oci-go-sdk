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

type ExportImageViaObjectStorageUriDetails struct {

	// The destination type. Use `objectStorageTuple` when specifying the namespace, bucket name, and object name.
	// Use `objectStorageUri` when specifying the Object Storage URL.
	DestinationType *string `mandatory:"true" json:"destinationType,omitempty"`

	// The Object Storage URL to export the image to. See [Object Storage URLs]({{DOC_SERVER_URL}}/Content/Compute/Tasks/imageimportexport.htm#URLs)
	// and [pre-authenticated requests]({{DOC_SERVER_URL}}/Content/Object/Tasks/managingaccess.htm#pre-auth) for constructing URLs for image import/export.
	DestinationUri *string `mandatory:"true" json:"destinationUri,omitempty"`
}

func (model ExportImageViaObjectStorageUriDetails) String() string {
	return common.PointerString(model)
}
