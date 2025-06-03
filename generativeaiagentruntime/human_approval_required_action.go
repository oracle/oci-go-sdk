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

// HumanApprovalRequiredAction An object describing human confirmation of tool execution that is required from the user.
type HumanApprovalRequiredAction struct {

	// The unique identifier for the action to be performed.
	ActionId *string `mandatory:"true" json:"actionId"`

	// Message accompanying the human input request asking for approval or denial of a tool execution.
	Message *string `mandatory:"true" json:"message"`

	// The options presented to the user approving and denying execution of the tool.
	Options []HumanApprovalRequiredActionOptionsEnum `mandatory:"true" json:"options"`
}

// GetActionId returns ActionId
func (m HumanApprovalRequiredAction) GetActionId() *string {
	return m.ActionId
}

func (m HumanApprovalRequiredAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HumanApprovalRequiredAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.Options {
		if _, ok := GetMappingHumanApprovalRequiredActionOptionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Options: %s. Supported values are: %s.", val, strings.Join(GetHumanApprovalRequiredActionOptionsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HumanApprovalRequiredAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHumanApprovalRequiredAction HumanApprovalRequiredAction
	s := struct {
		DiscriminatorParam string `json:"requiredActionType"`
		MarshalTypeHumanApprovalRequiredAction
	}{
		"HUMAN_APPROVAL_REQUIRED_ACTION",
		(MarshalTypeHumanApprovalRequiredAction)(m),
	}

	return json.Marshal(&s)
}

// HumanApprovalRequiredActionOptionsEnum Enum with underlying type: string
type HumanApprovalRequiredActionOptionsEnum string

// Set of constants representing the allowable values for HumanApprovalRequiredActionOptionsEnum
const (
	HumanApprovalRequiredActionOptionsApprove HumanApprovalRequiredActionOptionsEnum = "APPROVE"
	HumanApprovalRequiredActionOptionsDeny    HumanApprovalRequiredActionOptionsEnum = "DENY"
)

var mappingHumanApprovalRequiredActionOptionsEnum = map[string]HumanApprovalRequiredActionOptionsEnum{
	"APPROVE": HumanApprovalRequiredActionOptionsApprove,
	"DENY":    HumanApprovalRequiredActionOptionsDeny,
}

var mappingHumanApprovalRequiredActionOptionsEnumLowerCase = map[string]HumanApprovalRequiredActionOptionsEnum{
	"approve": HumanApprovalRequiredActionOptionsApprove,
	"deny":    HumanApprovalRequiredActionOptionsDeny,
}

// GetHumanApprovalRequiredActionOptionsEnumValues Enumerates the set of values for HumanApprovalRequiredActionOptionsEnum
func GetHumanApprovalRequiredActionOptionsEnumValues() []HumanApprovalRequiredActionOptionsEnum {
	values := make([]HumanApprovalRequiredActionOptionsEnum, 0)
	for _, v := range mappingHumanApprovalRequiredActionOptionsEnum {
		values = append(values, v)
	}
	return values
}

// GetHumanApprovalRequiredActionOptionsEnumStringValues Enumerates the set of values in String for HumanApprovalRequiredActionOptionsEnum
func GetHumanApprovalRequiredActionOptionsEnumStringValues() []string {
	return []string{
		"APPROVE",
		"DENY",
	}
}

// GetMappingHumanApprovalRequiredActionOptionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHumanApprovalRequiredActionOptionsEnum(val string) (HumanApprovalRequiredActionOptionsEnum, bool) {
	enum, ok := mappingHumanApprovalRequiredActionOptionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
