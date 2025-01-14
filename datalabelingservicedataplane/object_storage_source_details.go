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

// ObjectStorageSourceDetails Object Storage Source Details.
type ObjectStorageSourceDetails struct {

	// The path relative to the prefix specified in the dataset source details (file name).
	RelativePath *string `mandatory:"true" json:"relativePath"`

	// The full path of the file this record belongs to.
	Path *string `mandatory:"true" json:"path"`

	// The offset into the file containing the content.
	Offset *float32 `mandatory:"false" json:"offset"`

	// The length from the offset into the file containing the content.
	Length *float32 `mandatory:"false" json:"length"`
}

func (m ObjectStorageSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ObjectStorageSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageSourceDetails ObjectStorageSourceDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeObjectStorageSourceDetails
	}{
		"OBJECT_STORAGE",
		(MarshalTypeObjectStorageSourceDetails)(m),
	}

	return json.Marshal(&s)
}
