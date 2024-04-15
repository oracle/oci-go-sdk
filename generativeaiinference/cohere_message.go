// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Inference API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service inference API to access your custom model endpoints, or to try the out-of-the-box models to GenerateText, SummarizeText, and EmbedText.
// To use a Generative AI custom model for inference, you must first create an endpoint for that model. Use the Generative AI service management API (https://docs.cloud.oracle.com/#/en/generative-ai/latest/) to Model by fine-tuning an out-of-the-box model, or a previous version of a custom model, using your own data. Fine-tune the custom model on a  DedicatedAiCluster. Then, create a DedicatedAiCluster with an Endpoint to host your custom model. For resource management in the Generative AI service, use the Generative AI service management API (https://docs.cloud.oracle.com/#/en/generative-ai/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeaiinference

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CohereMessage An message that represents a single dialogue of chat
type CohereMessage struct {

	// One of CHATBOT|USER to identify who the message is coming from.
	Role CohereMessageRoleEnum `mandatory:"true" json:"role"`

	// Contents of the chat message.
	Message *string `mandatory:"true" json:"message"`
}

func (m CohereMessage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CohereMessage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCohereMessageRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetCohereMessageRoleEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CohereMessageRoleEnum Enum with underlying type: string
type CohereMessageRoleEnum string

// Set of constants representing the allowable values for CohereMessageRoleEnum
const (
	CohereMessageRoleChatbot CohereMessageRoleEnum = "CHATBOT"
	CohereMessageRoleUser    CohereMessageRoleEnum = "USER"
)

var mappingCohereMessageRoleEnum = map[string]CohereMessageRoleEnum{
	"CHATBOT": CohereMessageRoleChatbot,
	"USER":    CohereMessageRoleUser,
}

var mappingCohereMessageRoleEnumLowerCase = map[string]CohereMessageRoleEnum{
	"chatbot": CohereMessageRoleChatbot,
	"user":    CohereMessageRoleUser,
}

// GetCohereMessageRoleEnumValues Enumerates the set of values for CohereMessageRoleEnum
func GetCohereMessageRoleEnumValues() []CohereMessageRoleEnum {
	values := make([]CohereMessageRoleEnum, 0)
	for _, v := range mappingCohereMessageRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetCohereMessageRoleEnumStringValues Enumerates the set of values in String for CohereMessageRoleEnum
func GetCohereMessageRoleEnumStringValues() []string {
	return []string{
		"CHATBOT",
		"USER",
	}
}

// GetMappingCohereMessageRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCohereMessageRoleEnum(val string) (CohereMessageRoleEnum, bool) {
	enum, ok := mappingCohereMessageRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
