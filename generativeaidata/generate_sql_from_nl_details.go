// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service NL2SQL API
//
// A description of the ReferenceService API. in progress
//

package generativeaidata

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GenerateSqlFromNlDetails The data to create a GenerateSqlFromNlJob.
type GenerateSqlFromNlDetails struct {

	// A user-provided query or instruction written in plain, conversational language.
	// This input is intended to capture the user's intent, question, or command without requiring technical syntax or structured formatting.
	// The query should clearly express what the user wants to know or accomplish, allowing the system to interpret and respond appropriately.
	InputNaturalLanguageQuery *string `mandatory:"true" json:"inputNaturalLanguageQuery"`

	// An optional description of the GenerateSqlFromNlJob.
	Description *string `mandatory:"false" json:"description"`

	// A user-friendly display name. It does not have to be unique and can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The generative AI modelId to use for Generate SQL. You can use the ListModels API to list the available models. https://docs.oracle.com/en-us/iaas/api/#/en/generative-ai/20231130/ModelCollection/ListModels
	ModelId *string `mandatory:"false" json:"modelId"`

	// Controls whether GenerateSqlFromNl should be accepted as a background job or wait for completion.
	// BACKGROUND_JOB accepts the request for background processing and returns a pollable job.
	// WAIT_FOR_COMPLETION waits for completion within the service-defined timeout.
	CompletionMode CompletionModeEnum `mandatory:"false" json:"completionMode,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m GenerateSqlFromNlDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenerateSqlFromNlDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCompletionModeEnum(string(m.CompletionMode)); !ok && m.CompletionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CompletionMode: %s. Supported values are: %s.", m.CompletionMode, strings.Join(GetCompletionModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
