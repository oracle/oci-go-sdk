// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ObjectStorageDataset Object Storage Dataset
type ObjectStorageDataset struct {
	LocationDetails LocationDetails `mandatory:"true" json:"locationDetails"`

	// Entity Type categorizing the following list of words.
	EntityType *string `mandatory:"false" json:"entityType"`
}

func (m ObjectStorageDataset) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageDataset) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ObjectStorageDataset) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageDataset ObjectStorageDataset
	s := struct {
		DiscriminatorParam string `json:"datasetType"`
		MarshalTypeObjectStorageDataset
	}{
		"OBJECT_STORAGE",
		(MarshalTypeObjectStorageDataset)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ObjectStorageDataset) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		EntityType      *string         `json:"entityType"`
		LocationDetails locationdetails `json:"locationDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.EntityType = model.EntityType

	nn, e = model.LocationDetails.UnmarshalPolymorphicJSON(model.LocationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.LocationDetails = nn.(LocationDetails)
	} else {
		m.LocationDetails = nil
	}

	return
}
