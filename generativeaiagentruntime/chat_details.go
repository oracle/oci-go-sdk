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

// ChatDetails Chat details for managing user interactions and tool executions.
type ChatDetails struct {

	// The input user message content for the chat.
	UserMessage *string `mandatory:"false" json:"userMessage"`

	// Whether to stream the response.
	ShouldStream *bool `mandatory:"false" json:"shouldStream"`

	// Optional sessionId. If not provided, will chat without any prior context.
	SessionId *string `mandatory:"false" json:"sessionId"`

	// A map where each key is a toolId and the value contains tool type and additional dynamic parameters. This field is deprecated and will be removed after July 02 2026.
	ToolParameters map[string]string `mandatory:"false" json:"toolParameters"`

	// Array of tool input objects, each specifying a tool's ID, type, and corresponding input parameters required for execution.
	ToolInputs []ToolInput `mandatory:"false" json:"toolInputs"`

	// A list of actions that have been performed based on prior required actions.
	PerformedActions []PerformedAction `mandatory:"false" json:"performedActions"`
}

func (m ChatDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChatDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ChatDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		UserMessage      *string           `json:"userMessage"`
		ShouldStream     *bool             `json:"shouldStream"`
		SessionId        *string           `json:"sessionId"`
		ToolParameters   map[string]string `json:"toolParameters"`
		ToolInputs       []toolinput       `json:"toolInputs"`
		PerformedActions []performedaction `json:"performedActions"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.UserMessage = model.UserMessage

	m.ShouldStream = model.ShouldStream

	m.SessionId = model.SessionId

	m.ToolParameters = model.ToolParameters

	m.ToolInputs = make([]ToolInput, len(model.ToolInputs))
	for i, n := range model.ToolInputs {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.ToolInputs[i] = nn.(ToolInput)
		} else {
			m.ToolInputs[i] = nil
		}
	}
	m.PerformedActions = make([]PerformedAction, len(model.PerformedActions))
	for i, n := range model.PerformedActions {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.PerformedActions[i] = nn.(PerformedAction)
		} else {
			m.PerformedActions[i] = nil
		}
	}
	return
}
