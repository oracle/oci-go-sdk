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

// EmbedTextResult The generated embedded result to return.
type EmbedTextResult struct {

	// A unique identifier for the generated result.
	Id *string `mandatory:"true" json:"id"`

	// The embeddings corresponding to float.
	Embeddings [][]float32 `mandatory:"true" json:"embeddings"`

	// The original inputs. Only present if "isEcho" is set to true.
	Inputs []string `mandatory:"false" json:"inputs"`

	// The original inputs. Only present if "isEcho" is set to true.
	EmbedContents []EmbedContent `mandatory:"false" json:"embedContents"`

	// The embeddings corresponding to embedding types input.
	EmbeddingsByType *interface{} `mandatory:"false" json:"embeddingsByType"`

	// The OCID of the model used in this inference request.
	ModelId *string `mandatory:"false" json:"modelId"`

	// The version of the model.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	Usage *Usage `mandatory:"false" json:"usage"`
}

func (m EmbedTextResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmbedTextResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *EmbedTextResult) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Inputs           []string       `json:"inputs"`
		EmbedContents    []embedcontent `json:"embedContents"`
		EmbeddingsByType *interface{}   `json:"embeddingsByType"`
		ModelId          *string        `json:"modelId"`
		ModelVersion     *string        `json:"modelVersion"`
		Usage            *Usage         `json:"usage"`
		Id               *string        `json:"id"`
		Embeddings       [][]float32    `json:"embeddings"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Inputs = make([]string, len(model.Inputs))
	copy(m.Inputs, model.Inputs)
	m.EmbedContents = make([]EmbedContent, len(model.EmbedContents))
	for i, n := range model.EmbedContents {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.EmbedContents[i] = nn.(EmbedContent)
		} else {
			m.EmbedContents[i] = nil
		}
	}
	m.EmbeddingsByType = model.EmbeddingsByType

	m.ModelId = model.ModelId

	m.ModelVersion = model.ModelVersion

	m.Usage = model.Usage

	m.Id = model.Id

	m.Embeddings = make([][]float32, len(model.Embeddings))
	for i := range model.Embeddings {
		if model.Embeddings[i] != nil {
			m.Embeddings[i] = make([]float32, len(model.Embeddings[i]))
			copy(m.Embeddings[i], model.Embeddings[i])
		}
	}
	return
}
