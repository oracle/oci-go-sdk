// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service NL2SQL API
//
// A description of the ReferenceService API. in progress
//

package generativeaidata

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GenerateEnrichmentJobDetails The details required to create an EnrichmentJob.
type GenerateEnrichmentJobDetails struct {

	// Enrichment job type. Currently supported Full Build (All supported objects in a given schema) and Partial Build (Selected tables and/or supported objects in a given schema).
	EnrichmentJobType EnrichmentJobTypeEnum `mandatory:"true" json:"enrichmentJobType"`

	EnrichmentJobConfiguration EnrichmentJobConfiguration `mandatory:"true" json:"enrichmentJobConfiguration"`

	// An optional description of the EnrichmentJob.
	Description *string `mandatory:"false" json:"description"`

	// A user-friendly display name. It does not have to be unique and can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m GenerateEnrichmentJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateEnrichmentJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEnrichmentJobTypeEnum(string(m.EnrichmentJobType)); !ok && m.EnrichmentJobType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnrichmentJobType: %s. Supported values are: %s.", m.EnrichmentJobType, strings.Join(GetEnrichmentJobTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *GenerateEnrichmentJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                *string                           `json:"description"`
		DisplayName                *string                           `json:"displayName"`
		FreeformTags               map[string]string                 `json:"freeformTags"`
		DefinedTags                map[string]map[string]interface{} `json:"definedTags"`
		EnrichmentJobType          EnrichmentJobTypeEnum             `json:"enrichmentJobType"`
		EnrichmentJobConfiguration enrichmentjobconfiguration        `json:"enrichmentJobConfiguration"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.EnrichmentJobType = model.EnrichmentJobType

	nn, e = model.EnrichmentJobConfiguration.UnmarshalPolymorphicJSON(model.EnrichmentJobConfiguration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.EnrichmentJobConfiguration = nn.(EnrichmentJobConfiguration)
	} else {
		m.EnrichmentJobConfiguration = nil
	}

	return
}
