// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Inference API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service inference API to access your custom model endpoints, or to try the out-of-the-box models to Chat, GenerateText, SummarizeText, and EmbedText.
// To use a Generative AI custom model for inference, you must first create an endpoint for that model. Use the Generative AI service management API (https://docs.cloud.oracle.com/#/en/generative-ai/latest/) to Model by fine-tuning an out-of-the-box model, or a previous version of a custom model, using your own data. Fine-tune the custom model on a  DedicatedAiCluster. Then, create a DedicatedAiCluster with an Endpoint to host your custom model. For resource management in the Generative AI service, use the Generative AI service management API (https://docs.cloud.oracle.com/#/en/generative-ai/latest/).
// To learn more about the service, see the Generative AI documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeaiinference

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ChatChoice Represents a single instance of the chat response.
type ChatChoice struct {

	// The index of the chat.
	Index *int `mandatory:"true" json:"index"`

	Message Message `mandatory:"true" json:"message"`

	// The reason why the model stopped generating tokens.
	// Stops if the model hits a natural stop point or a provided stop sequence. Returns the length if the tokens reach the specified maximum number of tokens.
	FinishReason *string `mandatory:"true" json:"finishReason"`

	Logprobs *Logprobs `mandatory:"false" json:"logprobs"`
}

func (m ChatChoice) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChatChoice) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ChatChoice) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Logprobs     *Logprobs `json:"logprobs"`
		Index        *int      `json:"index"`
		Message      message   `json:"message"`
		FinishReason *string   `json:"finishReason"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Logprobs = model.Logprobs

	m.Index = model.Index

	nn, e = model.Message.UnmarshalPolymorphicJSON(model.Message.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Message = nn.(Message)
	} else {
		m.Message = nil
	}

	m.FinishReason = model.FinishReason

	return
}
