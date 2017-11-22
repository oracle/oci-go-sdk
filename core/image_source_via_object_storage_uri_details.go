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

// ImageSourceViaObjectStorageUriDetails.
type ImageSourceViaObjectStorageUriDetails struct {

	// The Object Storage URL for the image.
	SourceUri *string `mandatory:"true" json:"sourceUri,omitempty"`
}

func (m ImageSourceViaObjectStorageUriDetails) String() string {
	return common.PointerString(m)
}

func (m ImageSourceViaObjectStorageUriDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeImageSourceViaObjectStorageUriDetails ImageSourceViaObjectStorageUriDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeImageSourceViaObjectStorageUriDetails
	}{
		"objectStorageUri",
		(MarshalTypeImageSourceViaObjectStorageUriDetails)(m),
	}

	return json.Marshal(&s)
}
