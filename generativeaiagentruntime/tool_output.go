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

// ToolOutput Base schema for tool outputs. Identified by `toolOutputType`, which determines the format of the output content.
type ToolOutput interface {

	// Specifies the unique OCID of the tool.
	GetToolId() *string

	// Specifies the display name of the tool.
	GetToolName() *string
}

type tooloutput struct {
	JsonData       []byte
	ToolName       *string `mandatory:"false" json:"toolName"`
	ToolId         *string `mandatory:"true" json:"toolId"`
	ToolOutputType string  `json:"toolOutputType"`
}

// UnmarshalJSON unmarshals json
func (m *tooloutput) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertooloutput tooloutput
	s := struct {
		Model Unmarshalertooloutput
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ToolId = s.Model.ToolId
	m.ToolName = s.Model.ToolName
	m.ToolOutputType = s.Model.ToolOutputType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *tooloutput) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ToolOutputType {
	case "GENERIC_TOOL_OUTPUT":
		mm := GenericToolOutput{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SQL_TOOL_OUTPUT":
		mm := SqlToolOutput{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RAG_TOOL_OUTPUT":
		mm := RagToolOutput{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ToolOutput: %s.", m.ToolOutputType)
		return *m, nil
	}
}

// GetToolName returns ToolName
func (m tooloutput) GetToolName() *string {
	return m.ToolName
}

// GetToolId returns ToolId
func (m tooloutput) GetToolId() *string {
	return m.ToolId
}

func (m tooloutput) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m tooloutput) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ToolOutputToolOutputTypeEnum Enum with underlying type: string
type ToolOutputToolOutputTypeEnum string

// Set of constants representing the allowable values for ToolOutputToolOutputTypeEnum
const (
	ToolOutputToolOutputTypeGenericToolOutput ToolOutputToolOutputTypeEnum = "GENERIC_TOOL_OUTPUT"
	ToolOutputToolOutputTypeSqlToolOutput     ToolOutputToolOutputTypeEnum = "SQL_TOOL_OUTPUT"
	ToolOutputToolOutputTypeRagToolOutput     ToolOutputToolOutputTypeEnum = "RAG_TOOL_OUTPUT"
)

var mappingToolOutputToolOutputTypeEnum = map[string]ToolOutputToolOutputTypeEnum{
	"GENERIC_TOOL_OUTPUT": ToolOutputToolOutputTypeGenericToolOutput,
	"SQL_TOOL_OUTPUT":     ToolOutputToolOutputTypeSqlToolOutput,
	"RAG_TOOL_OUTPUT":     ToolOutputToolOutputTypeRagToolOutput,
}

var mappingToolOutputToolOutputTypeEnumLowerCase = map[string]ToolOutputToolOutputTypeEnum{
	"generic_tool_output": ToolOutputToolOutputTypeGenericToolOutput,
	"sql_tool_output":     ToolOutputToolOutputTypeSqlToolOutput,
	"rag_tool_output":     ToolOutputToolOutputTypeRagToolOutput,
}

// GetToolOutputToolOutputTypeEnumValues Enumerates the set of values for ToolOutputToolOutputTypeEnum
func GetToolOutputToolOutputTypeEnumValues() []ToolOutputToolOutputTypeEnum {
	values := make([]ToolOutputToolOutputTypeEnum, 0)
	for _, v := range mappingToolOutputToolOutputTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetToolOutputToolOutputTypeEnumStringValues Enumerates the set of values in String for ToolOutputToolOutputTypeEnum
func GetToolOutputToolOutputTypeEnumStringValues() []string {
	return []string{
		"GENERIC_TOOL_OUTPUT",
		"SQL_TOOL_OUTPUT",
		"RAG_TOOL_OUTPUT",
	}
}

// GetMappingToolOutputToolOutputTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingToolOutputToolOutputTypeEnum(val string) (ToolOutputToolOutputTypeEnum, bool) {
	enum, ok := mappingToolOutputToolOutputTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
