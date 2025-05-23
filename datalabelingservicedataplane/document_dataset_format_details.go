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

// DocumentDatasetFormatDetails It indicates the dataset is comprised of document files.  It is open for further configurability.
type DocumentDatasetFormatDetails struct {
}

func (m DocumentDatasetFormatDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DocumentDatasetFormatDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DocumentDatasetFormatDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDocumentDatasetFormatDetails DocumentDatasetFormatDetails
	s := struct {
		DiscriminatorParam string `json:"formatType"`
		MarshalTypeDocumentDatasetFormatDetails
	}{
		"DOCUMENT",
		(MarshalTypeDocumentDatasetFormatDetails)(m),
	}

	return json.Marshal(&s)
}
