// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Agents Client API
//
// OCI Generative AI Agents is a fully managed service that combines the power of large language models (LLMs) with an intelligent retrieval system to create contextually relevant answers by searching your knowledge base, making your AI applications smart and efficient.
// OCI Generative AI Agents supports several ways to onboard your data and then allows you and your customers to interact with your data using a chat interface or API.
// Use the Generative AI Agents Client API to create and manage client chat sessions. A session represents an interactive conversation initiated by a user through an API to engage with an agent. It involves a series of exchanges where the user sends queries or prompts, and the agent responds with relevant information, actions, or assistance based on the user's input. The session persists for the duration of the interaction, maintaining context and continuity to provide coherent and meaningful responses throughout the conversation.
// For creating and managing agents, knowledge bases, data sources, endpoints, and data ingestion jobs see the /EN/generative-ai-agents/latest/.
// To learn more about the service, see the Generative AI Agents documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai-agents/home.htm).
//

package generativeaiagentruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OciDatabaseSourceLocation The location of the OCI Database that the agent will use.
type OciDatabaseSourceLocation struct {

	// The OCID of the OCI Database Tools Connection (https://docs.cloud.oracle.com/iaas/database-tools/home.htm).
	Id *string `mandatory:"false" json:"id"`

	// The name of the function that returns source text.
	FunctionName *string `mandatory:"false" json:"functionName"`

	// The URL of the retrieved document, if available.
	Url *string `mandatory:"false" json:"url"`
}

func (m OciDatabaseSourceLocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciDatabaseSourceLocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OciDatabaseSourceLocation) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOciDatabaseSourceLocation OciDatabaseSourceLocation
	s := struct {
		DiscriminatorParam string `json:"sourceLocationType"`
		MarshalTypeOciDatabaseSourceLocation
	}{
		"OCI_DATABASE",
		(MarshalTypeOciDatabaseSourceLocation)(m),
	}

	return json.Marshal(&s)
}