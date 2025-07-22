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

// ToolInput Base object for tool input parameters. The 'toolInputType' discriminator determines the specific input structure to be used.
type ToolInput interface {

	// Unique OCID of the tool.
	GetToolId() *string
}

type toolinput struct {
	JsonData      []byte
	ToolId        *string `mandatory:"true" json:"toolId"`
	ToolInputType string  `json:"toolInputType"`
}

// UnmarshalJSON unmarshals json
func (m *toolinput) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertoolinput toolinput
	s := struct {
		Model Unmarshalertoolinput
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ToolId = s.Model.ToolId
	m.ToolInputType = s.Model.ToolInputType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *toolinput) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ToolInputType {
	case "GENERIC_TOOL_INPUT":
		mm := GenericToolInput{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ToolInput: %s.", m.ToolInputType)
		return *m, nil
	}
}

// GetToolId returns ToolId
func (m toolinput) GetToolId() *string {
	return m.ToolId
}

func (m toolinput) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m toolinput) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ToolInputToolInputTypeEnum Enum with underlying type: string
type ToolInputToolInputTypeEnum string

// Set of constants representing the allowable values for ToolInputToolInputTypeEnum
const (
	ToolInputToolInputTypeGenericToolInput ToolInputToolInputTypeEnum = "GENERIC_TOOL_INPUT"
)

var mappingToolInputToolInputTypeEnum = map[string]ToolInputToolInputTypeEnum{
	"GENERIC_TOOL_INPUT": ToolInputToolInputTypeGenericToolInput,
}

var mappingToolInputToolInputTypeEnumLowerCase = map[string]ToolInputToolInputTypeEnum{
	"generic_tool_input": ToolInputToolInputTypeGenericToolInput,
}

// GetToolInputToolInputTypeEnumValues Enumerates the set of values for ToolInputToolInputTypeEnum
func GetToolInputToolInputTypeEnumValues() []ToolInputToolInputTypeEnum {
	values := make([]ToolInputToolInputTypeEnum, 0)
	for _, v := range mappingToolInputToolInputTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetToolInputToolInputTypeEnumStringValues Enumerates the set of values in String for ToolInputToolInputTypeEnum
func GetToolInputToolInputTypeEnumStringValues() []string {
	return []string{
		"GENERIC_TOOL_INPUT",
	}
}

// GetMappingToolInputToolInputTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingToolInputToolInputTypeEnum(val string) (ToolInputToolInputTypeEnum, bool) {
	enum, ok := mappingToolInputToolInputTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
