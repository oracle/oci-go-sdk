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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ChatDetails Details of the conversation for the model to respond.
type ChatDetails struct {

	// The OCID of compartment that the user is authorized to use to call into the Generative AI service.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ServingMode ServingMode `mandatory:"true" json:"servingMode"`

	ChatRequest BaseChatRequest `mandatory:"false" json:"chatRequest"`
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
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ChatDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ChatRequest   basechatrequest `json:"chatRequest"`
		CompartmentId *string         `json:"compartmentId"`
		ServingMode   servingmode     `json:"servingMode"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ChatRequest.UnmarshalPolymorphicJSON(model.ChatRequest.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ChatRequest = nn.(BaseChatRequest)
	} else {
		m.ChatRequest = nil
	}

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

	return
}
