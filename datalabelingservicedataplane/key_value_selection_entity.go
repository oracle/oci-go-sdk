// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling API
//
// Use Data Labeling API to create Annotations on Images, Texts & Documents, and generate snapshots.
//

package datalabelingservicedataplane

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KeyValueSelectionEntity This allows the labeler to apply label the highlighted text from OCR, this includes labelled and unlabelled data.
type KeyValueSelectionEntity struct {

	// Entity Name.
	Text *string `mandatory:"true" json:"text"`

	BoundingPolygon *BoundingPolygon `mandatory:"true" json:"boundingPolygon"`

	// float value, score from OCR.
	Confidence *float32 `mandatory:"true" json:"confidence"`

	// A collection of label entities.
	Labels []Label `mandatory:"false" json:"labels"`

	// Integer value.
	Rotation *float32 `mandatory:"false" json:"rotation"`

	// Integer value.
	PageNumber *float32 `mandatory:"false" json:"pageNumber"`
}

func (m KeyValueSelectionEntity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KeyValueSelectionEntity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m KeyValueSelectionEntity) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeKeyValueSelectionEntity KeyValueSelectionEntity
	s := struct {
		DiscriminatorParam string `json:"entityType"`
		MarshalTypeKeyValueSelectionEntity
	}{
		"KEYVALUESELECTION",
		(MarshalTypeKeyValueSelectionEntity)(m),
	}

	return json.Marshal(&s)
}
