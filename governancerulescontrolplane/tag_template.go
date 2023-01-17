// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GovernanceRulesControlPlane API
//
// A description of the GovernanceRulesControlPlane API
//

package governancerulescontrolplane

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TagTemplate Template for governance rules of type tag.
type TagTemplate struct {

	// The name of the tag namespace. It must be unique across all tag namespaces in the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// Description of the tag namespace.
	Description *string `mandatory:"false" json:"description"`

	// Represents an array of tags for tag namespace.
	Tags []Tag `mandatory:"false" json:"tags"`

	TagDefaults []TagDefault `mandatory:"false" json:"tagDefaults"`
}

func (m TagTemplate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TagTemplate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TagTemplate) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTagTemplate TagTemplate
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTagTemplate
	}{
		"TAG",
		(MarshalTypeTagTemplate)(m),
	}

	return json.Marshal(&s)
}
