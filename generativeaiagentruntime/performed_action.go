// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// PerformedAction The base structure for an action that has already been performed.
type PerformedAction interface {

	// The unique identifier for the action that has been performed.
	GetActionId() *string
}

type performedaction struct {
	JsonData            []byte
	ActionId            *string `mandatory:"true" json:"actionId"`
	PerformedActionType string  `json:"performedActionType"`
}

// UnmarshalJSON unmarshals json
func (m *performedaction) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerperformedaction performedaction
	s := struct {
		Model Unmarshalerperformedaction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ActionId = s.Model.ActionId
	m.PerformedActionType = s.Model.PerformedActionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *performedaction) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PerformedActionType {
	case "FUNCTION_CALLING_PERFORMED_ACTION":
		mm := FunctionCallingPerformedAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HUMAN_APPROVAL_PERFORMED_ACTION":
		mm := HumanApprovalPerformedAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for PerformedAction: %s.", m.PerformedActionType)
		return *m, nil
	}
}

// GetActionId returns ActionId
func (m performedaction) GetActionId() *string {
	return m.ActionId
}

func (m performedaction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m performedaction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PerformedActionPerformedActionTypeEnum Enum with underlying type: string
type PerformedActionPerformedActionTypeEnum string

// Set of constants representing the allowable values for PerformedActionPerformedActionTypeEnum
const (
	PerformedActionPerformedActionTypeHumanApprovalPerformedAction   PerformedActionPerformedActionTypeEnum = "HUMAN_APPROVAL_PERFORMED_ACTION"
	PerformedActionPerformedActionTypeFunctionCallingPerformedAction PerformedActionPerformedActionTypeEnum = "FUNCTION_CALLING_PERFORMED_ACTION"
)

var mappingPerformedActionPerformedActionTypeEnum = map[string]PerformedActionPerformedActionTypeEnum{
	"HUMAN_APPROVAL_PERFORMED_ACTION":   PerformedActionPerformedActionTypeHumanApprovalPerformedAction,
	"FUNCTION_CALLING_PERFORMED_ACTION": PerformedActionPerformedActionTypeFunctionCallingPerformedAction,
}

var mappingPerformedActionPerformedActionTypeEnumLowerCase = map[string]PerformedActionPerformedActionTypeEnum{
	"human_approval_performed_action":   PerformedActionPerformedActionTypeHumanApprovalPerformedAction,
	"function_calling_performed_action": PerformedActionPerformedActionTypeFunctionCallingPerformedAction,
}

// GetPerformedActionPerformedActionTypeEnumValues Enumerates the set of values for PerformedActionPerformedActionTypeEnum
func GetPerformedActionPerformedActionTypeEnumValues() []PerformedActionPerformedActionTypeEnum {
	values := make([]PerformedActionPerformedActionTypeEnum, 0)
	for _, v := range mappingPerformedActionPerformedActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPerformedActionPerformedActionTypeEnumStringValues Enumerates the set of values in String for PerformedActionPerformedActionTypeEnum
func GetPerformedActionPerformedActionTypeEnumStringValues() []string {
	return []string{
		"HUMAN_APPROVAL_PERFORMED_ACTION",
		"FUNCTION_CALLING_PERFORMED_ACTION",
	}
}

// GetMappingPerformedActionPerformedActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPerformedActionPerformedActionTypeEnum(val string) (PerformedActionPerformedActionTypeEnum, bool) {
	enum, ok := mappingPerformedActionPerformedActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
