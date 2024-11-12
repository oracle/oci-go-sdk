// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Inference API
//
// OCI Generative AI is a fully managed service that provides a set of state-of-the-art, customizable large language models (LLMs) that cover a wide range of use cases for text generation, summarization, and text embeddings.
// Use the Generative AI service inference API to access your custom model endpoints, or to try the out-of-the-box models to /EN/generative-ai-inference/latest/ChatResult/Chat, /EN/generative-ai-inference/latest/GenerateTextResult/GenerateText, /EN/generative-ai-inference/latest/SummarizeTextResult/SummarizeText, and /EN/generative-ai-inference/latest/EmbedTextResult/EmbedText.
// To use a Generative AI custom model for inference, you must first create an endpoint for that model. Use the /EN/generative-ai/latest/ to /EN/generative-ai/latest/Model/ by fine-tuning an out-of-the-box model, or a previous version of a custom model, using your own data. Fine-tune the custom model on a /EN/generative-ai/latest/DedicatedAiCluster/. Then, create a /EN/generative-ai/latest/DedicatedAiCluster/ with an Endpoint to host your custom model. For resource management in the Generative AI service, use the /EN/generative-ai/latest/.
// To learn more about the service, see the Generative AI documentation (https://docs.cloud.oracle.com/iaas/Content/generative-ai/home.htm).
//

package generativeaiinference

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ImageUrl Represents a single instance of chat image url.
type ImageUrl struct {

	// The URL of the image.
	Url *string `mandatory:"true" json:"url"`

	// The level of the detail.
	Detail ImageUrlDetailEnum `mandatory:"false" json:"detail,omitempty"`
}

func (m ImageUrl) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImageUrl) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingImageUrlDetailEnum(string(m.Detail)); !ok && m.Detail != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Detail: %s. Supported values are: %s.", m.Detail, strings.Join(GetImageUrlDetailEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ImageUrlDetailEnum Enum with underlying type: string
type ImageUrlDetailEnum string

// Set of constants representing the allowable values for ImageUrlDetailEnum
const (
	ImageUrlDetailAuto ImageUrlDetailEnum = "AUTO"
	ImageUrlDetailHigh ImageUrlDetailEnum = "HIGH"
	ImageUrlDetailLow  ImageUrlDetailEnum = "LOW"
)

var mappingImageUrlDetailEnum = map[string]ImageUrlDetailEnum{
	"AUTO": ImageUrlDetailAuto,
	"HIGH": ImageUrlDetailHigh,
	"LOW":  ImageUrlDetailLow,
}

var mappingImageUrlDetailEnumLowerCase = map[string]ImageUrlDetailEnum{
	"auto": ImageUrlDetailAuto,
	"high": ImageUrlDetailHigh,
	"low":  ImageUrlDetailLow,
}

// GetImageUrlDetailEnumValues Enumerates the set of values for ImageUrlDetailEnum
func GetImageUrlDetailEnumValues() []ImageUrlDetailEnum {
	values := make([]ImageUrlDetailEnum, 0)
	for _, v := range mappingImageUrlDetailEnum {
		values = append(values, v)
	}
	return values
}

// GetImageUrlDetailEnumStringValues Enumerates the set of values in String for ImageUrlDetailEnum
func GetImageUrlDetailEnumStringValues() []string {
	return []string{
		"AUTO",
		"HIGH",
		"LOW",
	}
}

// GetMappingImageUrlDetailEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingImageUrlDetailEnum(val string) (ImageUrlDetailEnum, bool) {
	enum, ok := mappingImageUrlDetailEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
