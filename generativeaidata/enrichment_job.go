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

// EnrichmentJob Object representing to EnrichmentJob.
// ocidEntityType: generativeaiEnrichmentJob
// adLocality: regional
type EnrichmentJob struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Owning SemanticStore OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for an EnrichmentJob.
	SemanticStoreId *string `mandatory:"true" json:"semanticStoreId"`

	// Enrichment job type. Currently supported Full Build (All supported objects in a given schema) and Partial Build (Selected tables and/or supported objects in a given schema).
	EnrichmentJobType EnrichmentJobTypeEnum `mandatory:"true" json:"enrichmentJobType"`

	EnrichmentJobConfiguration EnrichmentJobConfiguration `mandatory:"true" json:"enrichmentJobConfiguration"`

	// The date and time that the enrichment job was accepted in the format of an RFC3339 datetime string.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// A message describing the current state in more detail that can provide actionable information.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The lifecycleState of GenerateSqlJob.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// An optional description of the EnrichmentJob.
	Description *string `mandatory:"false" json:"description"`

	// A user-friendly display name. It does not have to be unique and can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The date and time that the enrichment job was started in the format of an RFC3339 datetime string.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time that the enrichment job was finished in the format of an RFC3339 datetime string.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// The percentage of the enrichment job that has been completed.
	PercentComplete *float32 `mandatory:"false" json:"percentComplete"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m EnrichmentJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnrichmentJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEnrichmentJobTypeEnum(string(m.EnrichmentJobType)); !ok && m.EnrichmentJobType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnrichmentJobType: %s. Supported values are: %s.", m.EnrichmentJobType, strings.Join(GetEnrichmentJobTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *EnrichmentJob) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                *string                           `json:"description"`
		DisplayName                *string                           `json:"displayName"`
		TimeStarted                *common.SDKTime                   `json:"timeStarted"`
		TimeFinished               *common.SDKTime                   `json:"timeFinished"`
		PercentComplete            *float32                          `json:"percentComplete"`
		FreeformTags               map[string]string                 `json:"freeformTags"`
		DefinedTags                map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                 map[string]map[string]interface{} `json:"systemTags"`
		Id                         *string                           `json:"id"`
		SemanticStoreId            *string                           `json:"semanticStoreId"`
		EnrichmentJobType          EnrichmentJobTypeEnum             `json:"enrichmentJobType"`
		EnrichmentJobConfiguration enrichmentjobconfiguration        `json:"enrichmentJobConfiguration"`
		TimeAccepted               *common.SDKTime                   `json:"timeAccepted"`
		LifecycleDetails           *string                           `json:"lifecycleDetails"`
		LifecycleState             LifecycleStateEnum                `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.TimeStarted = model.TimeStarted

	m.TimeFinished = model.TimeFinished

	m.PercentComplete = model.PercentComplete

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.SemanticStoreId = model.SemanticStoreId

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

	m.TimeAccepted = model.TimeAccepted

	m.LifecycleDetails = model.LifecycleDetails

	m.LifecycleState = model.LifecycleState

	return
}
