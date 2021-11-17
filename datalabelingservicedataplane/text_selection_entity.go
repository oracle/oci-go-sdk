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
	"github.com/oracle/oci-go-sdk/v52/common"
)

// TextSelectionEntity This allows the labeler to highlight text by specifying an offset and a length and apply labels to it.
type TextSelectionEntity struct {

	// Collection of Label entities
	Labels []Label `mandatory:"true" json:"labels"`

	TextSpan *TextSpan `mandatory:"true" json:"textSpan"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	ExtendedMetadata map[string]string `mandatory:"false" json:"extendedMetadata"`
}

func (m TextSelectionEntity) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m TextSelectionEntity) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTextSelectionEntity TextSelectionEntity
	s := struct {
		DiscriminatorParam string `json:"entityType"`
		MarshalTypeTextSelectionEntity
	}{
		"TEXTSELECTION",
		(MarshalTypeTextSelectionEntity)(m),
	}

	return json.Marshal(&s)
}
