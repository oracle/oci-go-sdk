// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling API
//
// Use Data Labeling API to create Annotations on Images, Texts & Documents, and generate snapshots.
//

package datalabelingservicedataplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v62/common"
	"strings"
)

// TextSpan A wrapper class for offset and length, which together, represent a span of text in a text document.
type TextSpan struct {

	// The offset of the selected text within the entire text.
	Offset *float32 `mandatory:"false" json:"offset"`

	// The length of the selected text.
	Length *float32 `mandatory:"false" json:"length"`
}

func (m TextSpan) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TextSpan) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
