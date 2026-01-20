// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Inference API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service inference API to access your custom model endpoints, or to try the out-of-the-box models to /EN/generative-ai-inference/latest/ChatResult/Chat, /EN/generative-ai-inference/latest/GenerateTextResult/GenerateText, /EN/generative-ai-inference/latest/SummarizeTextResult/SummarizeText, and /EN/generative-ai-inference/latest/EmbedTextResult/EmbedText.
// To use a Generative AI custom model for inference, you must first create an endpoint for that model. Use the /EN/generative-ai/latest/ to /EN/generative-ai/latest/Model/ by fine-tuning an out-of-the-box model, or a previous version of a custom model, using your own data. Fine-tune the custom model on a /EN/generative-ai/latest/DedicatedAiCluster/. Then, create a /EN/generative-ai/latest/DedicatedAiCluster/ with an Endpoint to host your custom model. For resource management in the Generative AI service, use the /EN/generative-ai/latest/.
// To learn more about the service, see the Generative AI documentation (https://docs.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeaiinference

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CohereUserMessageV2 A message that represents a single chat dialog as USER role.
type CohereUserMessageV2 struct {

	// Contents of the chat message.
	Content []CohereContentV2 `mandatory:"true" json:"content"`
}

// GetContent returns Content
func (m CohereUserMessageV2) GetContent() []CohereContentV2 {
	return m.Content
}

func (m CohereUserMessageV2) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CohereUserMessageV2) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CohereUserMessageV2) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCohereUserMessageV2 CohereUserMessageV2
	s := struct {
		DiscriminatorParam string `json:"role"`
		MarshalTypeCohereUserMessageV2
	}{
		"USER",
		(MarshalTypeCohereUserMessageV2)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CohereUserMessageV2) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Content []coherecontentv2 `json:"content"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Content = make([]CohereContentV2, len(model.Content))
	for i, n := range model.Content {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Content[i] = nn.(CohereContentV2)
		} else {
			m.Content[i] = nil
		}
	}
	return
}
