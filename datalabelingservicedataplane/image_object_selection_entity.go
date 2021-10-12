// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DlsDataPlane API
//
// A description of the DlsDataPlane API.
//

package datalabelingservicedataplane

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v49/common"
)

// ImageObjectSelectionEntity This allows the labeler to use specify a series of coordinates in the image to represent an object and apply labels to it.  The coordinates will be connected in the order that they are provided and the last coordinate in the array will be connected to the first.
type ImageObjectSelectionEntity struct {

	// Collection of Label entities
	Labels []Label `mandatory:"true" json:"labels"`

	BoundingPolygon *BoundingPolygon `mandatory:"true" json:"boundingPolygon"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	ExtendedMetadata map[string]string `mandatory:"false" json:"extendedMetadata"`
}

func (m ImageObjectSelectionEntity) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ImageObjectSelectionEntity) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeImageObjectSelectionEntity ImageObjectSelectionEntity
	s := struct {
		DiscriminatorParam string `json:"entityType"`
		MarshalTypeImageObjectSelectionEntity
	}{
		"IMAGEOBJECTSELECTION",
		(MarshalTypeImageObjectSelectionEntity)(m),
	}

	return json.Marshal(&s)
}
