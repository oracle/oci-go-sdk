// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Speech API
//
// The OCI Speech Service harnesses the power of spoken language by allowing developers to easily convert file-based data containing human speech into highly accurate text transcriptions.
//

package aispeech

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EntityListDataset Entity List Dataset
type EntityListDataset struct {

	// List of sentences referencing 1 or more entityType matching those defined in the linked entityLists, used to improve accuracy by providing model training context of where/how an entity may appear in a sentence.
	// EntityTypes referenced in sentences should be written in all caps surrounded by angled braces (i.e "<PATIENT>" if entityType=patient)
	ReferenceExamples []string `mandatory:"false" json:"referenceExamples"`

	// Array of entityLists
	EntityList []EntityList `mandatory:"false" json:"entityList"`
}

func (m EntityListDataset) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EntityListDataset) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m EntityListDataset) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEntityListDataset EntityListDataset
	s := struct {
		DiscriminatorParam string `json:"datasetType"`
		MarshalTypeEntityListDataset
	}{
		"ENTITY_LIST",
		(MarshalTypeEntityListDataset)(m),
	}

	return json.Marshal(&s)
}
