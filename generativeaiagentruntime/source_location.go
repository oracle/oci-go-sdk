// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// SourceLocation The location of the data files that the agent will use.
type SourceLocation interface {
}

type sourcelocation struct {
	JsonData           []byte
	SourceLocationType string `json:"sourceLocationType"`
}

// UnmarshalJSON unmarshals json
func (m *sourcelocation) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersourcelocation sourcelocation
	s := struct {
		Model Unmarshalersourcelocation
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceLocationType = s.Model.SourceLocationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sourcelocation) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceLocationType {
	case "OCI_OBJECT_STORAGE":
		mm := OciObjectStorageSourceLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_OPEN_SEARCH":
		mm := OciOpenSearchSourceLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATABASE":
		mm := OciDatabaseSourceLocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SourceLocation: %s.", m.SourceLocationType)
		return *m, nil
	}
}

func (m sourcelocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sourcelocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SourceLocationSourceLocationTypeEnum Enum with underlying type: string
type SourceLocationSourceLocationTypeEnum string

// Set of constants representing the allowable values for SourceLocationSourceLocationTypeEnum
const (
	SourceLocationSourceLocationTypeObjectStorage SourceLocationSourceLocationTypeEnum = "OCI_OBJECT_STORAGE"
	SourceLocationSourceLocationTypeOpenSearch    SourceLocationSourceLocationTypeEnum = "OCI_OPEN_SEARCH"
	SourceLocationSourceLocationTypeDatabase      SourceLocationSourceLocationTypeEnum = "OCI_DATABASE"
)

var mappingSourceLocationSourceLocationTypeEnum = map[string]SourceLocationSourceLocationTypeEnum{
	"OCI_OBJECT_STORAGE": SourceLocationSourceLocationTypeObjectStorage,
	"OCI_OPEN_SEARCH":    SourceLocationSourceLocationTypeOpenSearch,
	"OCI_DATABASE":       SourceLocationSourceLocationTypeDatabase,
}

var mappingSourceLocationSourceLocationTypeEnumLowerCase = map[string]SourceLocationSourceLocationTypeEnum{
	"oci_object_storage": SourceLocationSourceLocationTypeObjectStorage,
	"oci_open_search":    SourceLocationSourceLocationTypeOpenSearch,
	"oci_database":       SourceLocationSourceLocationTypeDatabase,
}

// GetSourceLocationSourceLocationTypeEnumValues Enumerates the set of values for SourceLocationSourceLocationTypeEnum
func GetSourceLocationSourceLocationTypeEnumValues() []SourceLocationSourceLocationTypeEnum {
	values := make([]SourceLocationSourceLocationTypeEnum, 0)
	for _, v := range mappingSourceLocationSourceLocationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSourceLocationSourceLocationTypeEnumStringValues Enumerates the set of values in String for SourceLocationSourceLocationTypeEnum
func GetSourceLocationSourceLocationTypeEnumStringValues() []string {
	return []string{
		"OCI_OBJECT_STORAGE",
		"OCI_OPEN_SEARCH",
		"OCI_DATABASE",
	}
}

// GetMappingSourceLocationSourceLocationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSourceLocationSourceLocationTypeEnum(val string) (SourceLocationSourceLocationTypeEnum, bool) {
	enum, ok := mappingSourceLocationSourceLocationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
