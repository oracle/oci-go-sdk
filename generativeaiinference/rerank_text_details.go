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

// RerankTextDetails Details required for a rerank request.
type RerankTextDetails struct {

	// Input query for search in the documents.
	Input *string `mandatory:"true" json:"input"`

	// The OCID of the compartment to call into the Generative AI service LLMs.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ServingMode ServingMode `mandatory:"true" json:"servingMode"`

	// A list of document strings to rerank based on the query asked.
	Documents []string `mandatory:"true" json:"documents"`

	// The number of most relevant documents or indices to return. Defaults to the length of the documents.
	TopN *int `mandatory:"false" json:"topN"`

	// Whether or not to return the documents in the response.
	IsEcho *bool `mandatory:"false" json:"isEcho"`

	// The maximum number of chunks to produce internally from a document.
	MaxChunksPerDocument *int `mandatory:"false" json:"maxChunksPerDocument"`

	// Used to truncate the long documents with the specified no of tokens.
	MaxTokensPerDocument *int `mandatory:"false" json:"maxTokensPerDocument"`
}

func (m RerankTextDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RerankTextDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *RerankTextDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TopN                 *int        `json:"topN"`
		IsEcho               *bool       `json:"isEcho"`
		MaxChunksPerDocument *int        `json:"maxChunksPerDocument"`
		MaxTokensPerDocument *int        `json:"maxTokensPerDocument"`
		Input                *string     `json:"input"`
		CompartmentId        *string     `json:"compartmentId"`
		ServingMode          servingmode `json:"servingMode"`
		Documents            []string    `json:"documents"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TopN = model.TopN

	m.IsEcho = model.IsEcho

	m.MaxChunksPerDocument = model.MaxChunksPerDocument

	m.MaxTokensPerDocument = model.MaxTokensPerDocument

	m.Input = model.Input

	m.CompartmentId = model.CompartmentId

	nn, e = model.ServingMode.UnmarshalPolymorphicJSON(model.ServingMode.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ServingMode = nn.(ServingMode)
	} else {
		m.ServingMode = nil
	}

	m.Documents = make([]string, len(model.Documents))
	copy(m.Documents, model.Documents)
	return
}
