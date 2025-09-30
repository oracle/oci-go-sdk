// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AudioUrl Provide a base64 encoded audio or an audio uri if it's supported.
type AudioUrl struct {

	// The base64 encoded audio data or an audio uri if it's supported.
	// Example for an mp3 audio:
	//   `{
	//     "type": "AUDIO",
	//     "audioUrl": {
	//       "url": "data:audio/mp3;base64,<base64 encoded audio content>"
	//     }
	//   }`
	// Example with an audio uri:
	//   `{
	//     "type": "AUDIO",
	//     "audioUrl": {
	//       "url": "data:audio/mp3;uri,<audio uri>"
	//     }
	//   }`
	Url *string `mandatory:"true" json:"url"`

	// The default value is AUTO and only AUTO is supported. This option controls how to convert the base64 encoded audio to tokens.
	Detail AudioUrlDetailEnum `mandatory:"false" json:"detail,omitempty"`
}

func (m AudioUrl) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AudioUrl) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAudioUrlDetailEnum(string(m.Detail)); !ok && m.Detail != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Detail: %s. Supported values are: %s.", m.Detail, strings.Join(GetAudioUrlDetailEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AudioUrlDetailEnum Enum with underlying type: string
type AudioUrlDetailEnum string

// Set of constants representing the allowable values for AudioUrlDetailEnum
const (
	AudioUrlDetailAuto AudioUrlDetailEnum = "AUTO"
	AudioUrlDetailHigh AudioUrlDetailEnum = "HIGH"
	AudioUrlDetailLow  AudioUrlDetailEnum = "LOW"
)

var mappingAudioUrlDetailEnum = map[string]AudioUrlDetailEnum{
	"AUTO": AudioUrlDetailAuto,
	"HIGH": AudioUrlDetailHigh,
	"LOW":  AudioUrlDetailLow,
}

var mappingAudioUrlDetailEnumLowerCase = map[string]AudioUrlDetailEnum{
	"auto": AudioUrlDetailAuto,
	"high": AudioUrlDetailHigh,
	"low":  AudioUrlDetailLow,
}

// GetAudioUrlDetailEnumValues Enumerates the set of values for AudioUrlDetailEnum
func GetAudioUrlDetailEnumValues() []AudioUrlDetailEnum {
	values := make([]AudioUrlDetailEnum, 0)
	for _, v := range mappingAudioUrlDetailEnum {
		values = append(values, v)
	}
	return values
}

// GetAudioUrlDetailEnumStringValues Enumerates the set of values in String for AudioUrlDetailEnum
func GetAudioUrlDetailEnumStringValues() []string {
	return []string{
		"AUTO",
		"HIGH",
		"LOW",
	}
}

// GetMappingAudioUrlDetailEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAudioUrlDetailEnum(val string) (AudioUrlDetailEnum, bool) {
	enum, ok := mappingAudioUrlDetailEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
