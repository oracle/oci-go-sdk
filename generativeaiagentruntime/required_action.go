// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Agents Client API
//
// OCI Generative AI Agents is a fully managed service that combines the power of large language models (LLMs) with an intelligent retrieval system to create contextually relevant answers by searching your knowledge base, making your AI applications smart and efficient.
// OCI Generative AI Agents supports several ways to onboard your data and then allows you and your customers to interact with your data using a chat interface or API.
// Use the Generative AI Agents Client API to create and manage client chat sessions. A session represents an interactive conversation initiated by a user through an API to engage with an agent. It involves a series of exchanges where the user sends queries or prompts, and the agent responds with relevant information, actions, or assistance based on the user's input. The session persists for the duration of the interaction, maintaining context and continuity to provide coherent and meaningful responses throughout the conversation.
// For creating and managing agents, knowledge bases, data sources, endpoints, and data ingestion jobs see the /EN/generative-ai-agents/latest/.
// To learn more about the service, see the Generative AI Agents documentation (https://docs.oracle.com/iaas/Content/generative-ai-agents/home.htm).
//

package generativeaiagentruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RequiredAction Represents an action that needs to be performed by the user or client.
type RequiredAction interface {

	// The unique identifier for the action to be performed.
	GetActionId() *string
}

type requiredaction struct {
	JsonData           []byte
	ActionId           *string `mandatory:"true" json:"actionId"`
	RequiredActionType string  `json:"requiredActionType"`
}

// UnmarshalJSON unmarshals json
func (m *requiredaction) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerrequiredaction requiredaction
	s := struct {
		Model Unmarshalerrequiredaction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ActionId = s.Model.ActionId
	m.RequiredActionType = s.Model.RequiredActionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *requiredaction) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RequiredActionType {
	case "HUMAN_APPROVAL_REQUIRED_ACTION":
		mm := HumanApprovalRequiredAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FUNCTION_CALLING_REQUIRED_ACTION":
		mm := FunctionCallingRequiredAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for RequiredAction: %s.", m.RequiredActionType)
		return *m, nil
	}
}

// GetActionId returns ActionId
func (m requiredaction) GetActionId() *string {
	return m.ActionId
}

func (m requiredaction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m requiredaction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RequiredActionRequiredActionTypeEnum Enum with underlying type: string
type RequiredActionRequiredActionTypeEnum string

// Set of constants representing the allowable values for RequiredActionRequiredActionTypeEnum
const (
	RequiredActionRequiredActionTypeHumanApprovalRequiredAction   RequiredActionRequiredActionTypeEnum = "HUMAN_APPROVAL_REQUIRED_ACTION"
	RequiredActionRequiredActionTypeFunctionCallingRequiredAction RequiredActionRequiredActionTypeEnum = "FUNCTION_CALLING_REQUIRED_ACTION"
)

var mappingRequiredActionRequiredActionTypeEnum = map[string]RequiredActionRequiredActionTypeEnum{
	"HUMAN_APPROVAL_REQUIRED_ACTION":   RequiredActionRequiredActionTypeHumanApprovalRequiredAction,
	"FUNCTION_CALLING_REQUIRED_ACTION": RequiredActionRequiredActionTypeFunctionCallingRequiredAction,
}

var mappingRequiredActionRequiredActionTypeEnumLowerCase = map[string]RequiredActionRequiredActionTypeEnum{
	"human_approval_required_action":   RequiredActionRequiredActionTypeHumanApprovalRequiredAction,
	"function_calling_required_action": RequiredActionRequiredActionTypeFunctionCallingRequiredAction,
}

// GetRequiredActionRequiredActionTypeEnumValues Enumerates the set of values for RequiredActionRequiredActionTypeEnum
func GetRequiredActionRequiredActionTypeEnumValues() []RequiredActionRequiredActionTypeEnum {
	values := make([]RequiredActionRequiredActionTypeEnum, 0)
	for _, v := range mappingRequiredActionRequiredActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRequiredActionRequiredActionTypeEnumStringValues Enumerates the set of values in String for RequiredActionRequiredActionTypeEnum
func GetRequiredActionRequiredActionTypeEnumStringValues() []string {
	return []string{
		"HUMAN_APPROVAL_REQUIRED_ACTION",
		"FUNCTION_CALLING_REQUIRED_ACTION",
	}
}

// GetMappingRequiredActionRequiredActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequiredActionRequiredActionTypeEnum(val string) (RequiredActionRequiredActionTypeEnum, bool) {
	enum, ok := mappingRequiredActionRequiredActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
