// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Inference API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service inference API to access your custom model endpoints, or to try the out-of-the-box models to /EN/generative-ai-inference/latest/ChatResult/Chat, /EN/generative-ai-inference/latest/GenerateTextResult/GenerateText, /EN/generative-ai-inference/latest/SummarizeTextResult/SummarizeText, and /EN/generative-ai-inference/latest/EmbedTextResult/EmbedText.
// To use a Generative AI custom model for inference, you must first create an endpoint for that model. Use the /EN/generative-ai/latest/ to /EN/generative-ai/latest/Model/ by fine-tuning an out-of-the-box model, or a previous version of a custom model, using your own data. Fine-tune the custom model on a /EN/generative-ai/latest/DedicatedAiCluster/. Then, create a /EN/generative-ai/latest/DedicatedAiCluster/ with an Endpoint to host your custom model. For resource management in the Generative AI service, use the /EN/generative-ai/latest/.
// To learn more about the service, see the Generative AI documentation (https://docs.oracle.com/iaas/Content/generative-ai/home.htm).
// **Important:** The IP addresses behind each DNS endpoint might change over time. Always use the DNS hostname listed under the following **API Endpoints** section and avoid using hard-coded fixed IP addresses.
//

package generativeaiinference

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GuardrailVersionSummary Details of a guardrail version.
type GuardrailVersionSummary struct {

	// The guardrail system version string, e.g., "1.0.0"
	GuardrailVersion *string `mandatory:"true" json:"guardrailVersion"`

	// The state of the guardrail version.
	State GuardrailVersionStateEnum `mandatory:"true" json:"state"`

	// The description of the guardrail version.
	Description *string `mandatory:"true" json:"description"`

	// The preview date of the guardrail version.
	TimePreviewed *common.SDKTime `mandatory:"false" json:"timePreviewed"`

	// The activation date of the guardrail version.
	TimeActivated *common.SDKTime `mandatory:"false" json:"timeActivated"`

	// The deprecated date of the guardrail version.
	TimeDeprecated *common.SDKTime `mandatory:"false" json:"timeDeprecated"`

	// The retired date of the guardrail version.
	TimeRetired *common.SDKTime `mandatory:"false" json:"timeRetired"`
}

func (m GuardrailVersionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GuardrailVersionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGuardrailVersionStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetGuardrailVersionStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
