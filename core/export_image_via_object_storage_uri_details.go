// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"encoding/json"
)

// ExportImageViaObjectStorageUriDetails.
type ExportImageViaObjectStorageUriDetails struct {

	// The Object Storage URL to export the image to. See [Object Storage URLs]({{DOC_SERVER_URL}}/Content/Compute/Tasks/imageimportexport.htm#URLs)
	// and [pre-authenticated requests]({{DOC_SERVER_URL}}/Content/Object/Tasks/managingaccess.htm#pre-auth) for constructing URLs for image import/export.
	DestinationUri *string `mandatory:"true" json:"destinationUri,omitempty"`
}

func (m ExportImageViaObjectStorageUriDetails) String() string {
	return common.PointerString(m)
}

func (m ExportImageViaObjectStorageUriDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExportImageViaObjectStorageUriDetails ExportImageViaObjectStorageUriDetails
	s := struct {
		DiscriminatorParam string `json:"destinationType"`
		MarshalTypeExportImageViaObjectStorageUriDetails
	}{
		"objectStorageUri",
		(MarshalTypeExportImageViaObjectStorageUriDetails)(m),
	}

	return json.Marshal(&s)
}
