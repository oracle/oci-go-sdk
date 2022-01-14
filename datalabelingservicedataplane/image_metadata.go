// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling API
//
// Use Data Labeling API to create Annotations on Images, Texts & Documents, and generate snapshots.
//

package datalabelingservicedataplane

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// ImageMetadata Collection of metadata related to image record.
type ImageMetadata struct {

	// Height of the image record.
	Height *int `mandatory:"false" json:"height"`

	// Width of the image record.
	Width *int `mandatory:"false" json:"width"`

	// Depth of the image record.
	Depth *int `mandatory:"false" json:"depth"`
}

func (m ImageMetadata) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ImageMetadata) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeImageMetadata ImageMetadata
	s := struct {
		DiscriminatorParam string `json:"recordType"`
		MarshalTypeImageMetadata
	}{
		"IMAGE_METADATA",
		(MarshalTypeImageMetadata)(m),
	}

	return json.Marshal(&s)
}
