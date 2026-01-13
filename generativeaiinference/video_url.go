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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VideoUrl The base64 encoded video data or a video uri if it's supported.
type VideoUrl struct {

	// The base64 encoded video data or a video uri if it's supported.
	// Example for an mp4 video:
	//   `{
	//     "type": "VIDEO",
	//     "videoUrl": {
	//       "url": "data:video/mp4;base64,<base64 encoded video content>"
	//     }
	//   }`
	// Example with a video uri:
	//   `{
	//     "type": "VIDEO",
	//     "videoUrl": {
	//       "url": "data:video/mp4;uri,<video uri>"
	//     }
	//   }`
	Url *string `mandatory:"true" json:"url"`

	// The default value is AUTO and only AUTO is supported. This option controls how to convert the base64 encoded video to tokens.
	Detail VideoUrlDetailEnum `mandatory:"false" json:"detail,omitempty"`
}

func (m VideoUrl) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VideoUrl) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingVideoUrlDetailEnum(string(m.Detail)); !ok && m.Detail != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Detail: %s. Supported values are: %s.", m.Detail, strings.Join(GetVideoUrlDetailEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VideoUrlDetailEnum Enum with underlying type: string
type VideoUrlDetailEnum string

// Set of constants representing the allowable values for VideoUrlDetailEnum
const (
	VideoUrlDetailAuto VideoUrlDetailEnum = "AUTO"
	VideoUrlDetailHigh VideoUrlDetailEnum = "HIGH"
	VideoUrlDetailLow  VideoUrlDetailEnum = "LOW"
)

var mappingVideoUrlDetailEnum = map[string]VideoUrlDetailEnum{
	"AUTO": VideoUrlDetailAuto,
	"HIGH": VideoUrlDetailHigh,
	"LOW":  VideoUrlDetailLow,
}

var mappingVideoUrlDetailEnumLowerCase = map[string]VideoUrlDetailEnum{
	"auto": VideoUrlDetailAuto,
	"high": VideoUrlDetailHigh,
	"low":  VideoUrlDetailLow,
}

// GetVideoUrlDetailEnumValues Enumerates the set of values for VideoUrlDetailEnum
func GetVideoUrlDetailEnumValues() []VideoUrlDetailEnum {
	values := make([]VideoUrlDetailEnum, 0)
	for _, v := range mappingVideoUrlDetailEnum {
		values = append(values, v)
	}
	return values
}

// GetVideoUrlDetailEnumStringValues Enumerates the set of values in String for VideoUrlDetailEnum
func GetVideoUrlDetailEnumStringValues() []string {
	return []string{
		"AUTO",
		"HIGH",
		"LOW",
	}
}

// GetMappingVideoUrlDetailEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVideoUrlDetailEnum(val string) (VideoUrlDetailEnum, bool) {
	enum, ok := mappingVideoUrlDetailEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
