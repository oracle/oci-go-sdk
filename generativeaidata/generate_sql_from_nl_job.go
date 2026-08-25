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

// GenerateSqlFromNlJob Object representing to GenerateSqlFromNlJob.
// ocidEntityType: generativeaisemanticstorejob
// adLocality: regional
type GenerateSqlFromNlJob struct {

	// The OCID of the Semantic Store job.
	Id *string `mandatory:"true" json:"id"`

	// Owning SemanticStore OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for a GenerateSqlFromNlJob.
	SemanticStoreId *string `mandatory:"true" json:"semanticStoreId"`

	// The date and time that the GenerateSqlFromNlJob was accepted in the format of an RFC3339 datetime string.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The current state of GenerateSqlFromNlJob.
	// - ACCEPTED: Job has been created but not yet started.
	// - IN_PROGRESS: Job is currently running.
	// - SUCCEEDED: Job completed successfully. The result is available in jobOutput.
	// - FAILED: Job failed. See lifecycleDetails for error information.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A message describing the current state in more detail that can provide actionable information.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The generative AI modelId used for Generate SQL. You can use the ListModels API to list the available models. https://docs.oracle.com/en-us/iaas/api/#/en/generative-ai/20231130/ModelCollection/ListModels
	ModelId *string `mandatory:"true" json:"modelId"`

	// An optional description of the GenerateSqlFromNlJob.
	Description *string `mandatory:"false" json:"description"`

	// A user-friendly display name. It does not have to be unique and can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The date and time that the GenerateSqlFromNlJob was started in the format of an RFC3339 datetime string.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time that the GenerateSqlFromNlJob was finished in the format of an RFC3339 datetime string.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// A user-provided query or instruction written in plain, conversational language.
	// This input is intended to capture the user's intent, question, or command without requiring technical syntax or structured formatting.
	// The query should clearly express what the user wants to know or accomplish, allowing the system to interpret and respond appropriately.
	InputNaturalLanguageQuery *string `mandatory:"false" json:"inputNaturalLanguageQuery"`

	JobOutput JobOutput `mandatory:"false" json:"jobOutput"`

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

func (m GenerateSqlFromNlJob) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateSqlFromNlJob) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *GenerateSqlFromNlJob) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description               *string                           `json:"description"`
		DisplayName               *string                           `json:"displayName"`
		TimeStarted               *common.SDKTime                   `json:"timeStarted"`
		TimeFinished              *common.SDKTime                   `json:"timeFinished"`
		InputNaturalLanguageQuery *string                           `json:"inputNaturalLanguageQuery"`
		JobOutput                 joboutput                         `json:"jobOutput"`
		FreeformTags              map[string]string                 `json:"freeformTags"`
		DefinedTags               map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                map[string]map[string]interface{} `json:"systemTags"`
		Id                        *string                           `json:"id"`
		SemanticStoreId           *string                           `json:"semanticStoreId"`
		TimeAccepted              *common.SDKTime                   `json:"timeAccepted"`
		LifecycleState            LifecycleStateEnum                `json:"lifecycleState"`
		LifecycleDetails          *string                           `json:"lifecycleDetails"`
		ModelId                   *string                           `json:"modelId"`
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

	m.InputNaturalLanguageQuery = model.InputNaturalLanguageQuery

	nn, e = model.JobOutput.UnmarshalPolymorphicJSON(model.JobOutput.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.JobOutput = nn.(JobOutput)
	} else {
		m.JobOutput = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.SemanticStoreId = model.SemanticStoreId

	m.TimeAccepted = model.TimeAccepted

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.ModelId = model.ModelId

	return
}
