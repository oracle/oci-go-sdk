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

// Trace The trace that displays the internal progression, such as reasoning and actions during an execution.
type Trace interface {

	// Unique identifier for the event (UUID).
	GetKey() *string

	// Identifier of the parent event, if applicable (UUID).
	GetParentKey() *string

	GetSource() *SourceDetails

	// The date and time that the trace was created in the format of an RFC3339 datetime string.
	GetTimeCreated() *common.SDKTime

	// Timestamp for when the event ended (In RFC 3339).
	GetTimeFinished() *common.SDKTime
}

type trace struct {
	JsonData     []byte
	Key          *string         `mandatory:"false" json:"key"`
	ParentKey    *string         `mandatory:"false" json:"parentKey"`
	Source       *SourceDetails  `mandatory:"false" json:"source"`
	TimeCreated  *common.SDKTime `mandatory:"false" json:"timeCreated"`
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
	TraceType    string          `json:"traceType"`
}

// UnmarshalJSON unmarshals json
func (m *trace) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertrace trace
	s := struct {
		Model Unmarshalertrace
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ParentKey = s.Model.ParentKey
	m.Source = s.Model.Source
	m.TimeCreated = s.Model.TimeCreated
	m.TimeFinished = s.Model.TimeFinished
	m.TraceType = s.Model.TraceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *trace) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TraceType {
	case "ERROR_TRACE":
		mm := ErrorTrace{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RETRIEVAL_TRACE":
		mm := RetrievalTrace{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXECUTION_TRACE":
		mm := ExecutionTrace{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GENERATION_TRACE":
		mm := GenerationTrace{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TOOL_INVOCATION_TRACE":
		mm := ToolInvocationTrace{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PLANNING_TRACE":
		mm := PlanningTrace{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for Trace: %s.", m.TraceType)
		return *m, nil
	}
}

// GetKey returns Key
func (m trace) GetKey() *string {
	return m.Key
}

// GetParentKey returns ParentKey
func (m trace) GetParentKey() *string {
	return m.ParentKey
}

// GetSource returns Source
func (m trace) GetSource() *SourceDetails {
	return m.Source
}

// GetTimeCreated returns TimeCreated
func (m trace) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeFinished returns TimeFinished
func (m trace) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

func (m trace) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m trace) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TraceTraceTypeEnum Enum with underlying type: string
type TraceTraceTypeEnum string

// Set of constants representing the allowable values for TraceTraceTypeEnum
const (
	TraceTraceTypeErrorTrace          TraceTraceTypeEnum = "ERROR_TRACE"
	TraceTraceTypeRetrievalTrace      TraceTraceTypeEnum = "RETRIEVAL_TRACE"
	TraceTraceTypeGenerationTrace     TraceTraceTypeEnum = "GENERATION_TRACE"
	TraceTraceTypeToolInvocationTrace TraceTraceTypeEnum = "TOOL_INVOCATION_TRACE"
	TraceTraceTypePlanningTrace       TraceTraceTypeEnum = "PLANNING_TRACE"
	TraceTraceTypeExecutionTrace      TraceTraceTypeEnum = "EXECUTION_TRACE"
)

var mappingTraceTraceTypeEnum = map[string]TraceTraceTypeEnum{
	"ERROR_TRACE":           TraceTraceTypeErrorTrace,
	"RETRIEVAL_TRACE":       TraceTraceTypeRetrievalTrace,
	"GENERATION_TRACE":      TraceTraceTypeGenerationTrace,
	"TOOL_INVOCATION_TRACE": TraceTraceTypeToolInvocationTrace,
	"PLANNING_TRACE":        TraceTraceTypePlanningTrace,
	"EXECUTION_TRACE":       TraceTraceTypeExecutionTrace,
}

var mappingTraceTraceTypeEnumLowerCase = map[string]TraceTraceTypeEnum{
	"error_trace":           TraceTraceTypeErrorTrace,
	"retrieval_trace":       TraceTraceTypeRetrievalTrace,
	"generation_trace":      TraceTraceTypeGenerationTrace,
	"tool_invocation_trace": TraceTraceTypeToolInvocationTrace,
	"planning_trace":        TraceTraceTypePlanningTrace,
	"execution_trace":       TraceTraceTypeExecutionTrace,
}

// GetTraceTraceTypeEnumValues Enumerates the set of values for TraceTraceTypeEnum
func GetTraceTraceTypeEnumValues() []TraceTraceTypeEnum {
	values := make([]TraceTraceTypeEnum, 0)
	for _, v := range mappingTraceTraceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTraceTraceTypeEnumStringValues Enumerates the set of values in String for TraceTraceTypeEnum
func GetTraceTraceTypeEnumStringValues() []string {
	return []string{
		"ERROR_TRACE",
		"RETRIEVAL_TRACE",
		"GENERATION_TRACE",
		"TOOL_INVOCATION_TRACE",
		"PLANNING_TRACE",
		"EXECUTION_TRACE",
	}
}

// GetMappingTraceTraceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTraceTraceTypeEnum(val string) (TraceTraceTypeEnum, bool) {
	enum, ok := mappingTraceTraceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
