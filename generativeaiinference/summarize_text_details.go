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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SummarizeTextDetails Details for the request to summarize text.
type SummarizeTextDetails struct {

	// The input string to be summarized.
	Input *string `mandatory:"true" json:"input"`

	ServingMode ServingMode `mandatory:"true" json:"servingMode"`

	// The OCID of compartment in which to call the Generative AI service to summarize text.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Whether or not to include the original inputs in the response.
	IsEcho *bool `mandatory:"false" json:"isEcho"`

	// A number that sets the randomness of the generated output. Lower temperatures mean less random generations.
	// Use lower numbers for tasks with a correct answer such as question answering or summarizing. High temperatures can generate hallucinations or factually incorrect information. Start with temperatures lower than 1.0, and increase the temperature for more creative outputs, as you regenerate the prompts to refine the outputs.
	Temperature *float64 `mandatory:"false" json:"temperature"`

	// A free-form instruction for modifying how the summaries get generated. Should complete the sentence "Generate a summary _". For example, "focusing on the next steps" or "written by Yoda".
	AdditionalCommand *string `mandatory:"false" json:"additionalCommand"`

	// Indicates the approximate length of the summary. If "AUTO" is selected, the best option will be picked based on the input text.
	Length SummarizeTextDetailsLengthEnum `mandatory:"false" json:"length,omitempty"`

	// Indicates the style in which the summary will be delivered - in a free form paragraph or in bullet points. If "AUTO" is selected, the best option will be picked based on the input text.
	Format SummarizeTextDetailsFormatEnum `mandatory:"false" json:"format,omitempty"`

	// Controls how close to the original text the summary is. High extractiveness summaries will lean towards reusing sentences verbatim, while low extractiveness summaries will tend to paraphrase more.
	Extractiveness SummarizeTextDetailsExtractivenessEnum `mandatory:"false" json:"extractiveness,omitempty"`
}

func (m SummarizeTextDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeTextDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSummarizeTextDetailsLengthEnum(string(m.Length)); !ok && m.Length != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Length: %s. Supported values are: %s.", m.Length, strings.Join(GetSummarizeTextDetailsLengthEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeTextDetailsFormatEnum(string(m.Format)); !ok && m.Format != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Format: %s. Supported values are: %s.", m.Format, strings.Join(GetSummarizeTextDetailsFormatEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeTextDetailsExtractivenessEnum(string(m.Extractiveness)); !ok && m.Extractiveness != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Extractiveness: %s. Supported values are: %s.", m.Extractiveness, strings.Join(GetSummarizeTextDetailsExtractivenessEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SummarizeTextDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		IsEcho            *bool                                  `json:"isEcho"`
		Temperature       *float64                               `json:"temperature"`
		AdditionalCommand *string                                `json:"additionalCommand"`
		Length            SummarizeTextDetailsLengthEnum         `json:"length"`
		Format            SummarizeTextDetailsFormatEnum         `json:"format"`
		Extractiveness    SummarizeTextDetailsExtractivenessEnum `json:"extractiveness"`
		Input             *string                                `json:"input"`
		ServingMode       servingmode                            `json:"servingMode"`
		CompartmentId     *string                                `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.IsEcho = model.IsEcho

	m.Temperature = model.Temperature

	m.AdditionalCommand = model.AdditionalCommand

	m.Length = model.Length

	m.Format = model.Format

	m.Extractiveness = model.Extractiveness

	m.Input = model.Input

	nn, e = model.ServingMode.UnmarshalPolymorphicJSON(model.ServingMode.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ServingMode = nn.(ServingMode)
	} else {
		m.ServingMode = nil
	}

	m.CompartmentId = model.CompartmentId

	return
}

// SummarizeTextDetailsLengthEnum Enum with underlying type: string
type SummarizeTextDetailsLengthEnum string

// Set of constants representing the allowable values for SummarizeTextDetailsLengthEnum
const (
	SummarizeTextDetailsLengthShort  SummarizeTextDetailsLengthEnum = "SHORT"
	SummarizeTextDetailsLengthMedium SummarizeTextDetailsLengthEnum = "MEDIUM"
	SummarizeTextDetailsLengthLong   SummarizeTextDetailsLengthEnum = "LONG"
	SummarizeTextDetailsLengthAuto   SummarizeTextDetailsLengthEnum = "AUTO"
)

var mappingSummarizeTextDetailsLengthEnum = map[string]SummarizeTextDetailsLengthEnum{
	"SHORT":  SummarizeTextDetailsLengthShort,
	"MEDIUM": SummarizeTextDetailsLengthMedium,
	"LONG":   SummarizeTextDetailsLengthLong,
	"AUTO":   SummarizeTextDetailsLengthAuto,
}

var mappingSummarizeTextDetailsLengthEnumLowerCase = map[string]SummarizeTextDetailsLengthEnum{
	"short":  SummarizeTextDetailsLengthShort,
	"medium": SummarizeTextDetailsLengthMedium,
	"long":   SummarizeTextDetailsLengthLong,
	"auto":   SummarizeTextDetailsLengthAuto,
}

// GetSummarizeTextDetailsLengthEnumValues Enumerates the set of values for SummarizeTextDetailsLengthEnum
func GetSummarizeTextDetailsLengthEnumValues() []SummarizeTextDetailsLengthEnum {
	values := make([]SummarizeTextDetailsLengthEnum, 0)
	for _, v := range mappingSummarizeTextDetailsLengthEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeTextDetailsLengthEnumStringValues Enumerates the set of values in String for SummarizeTextDetailsLengthEnum
func GetSummarizeTextDetailsLengthEnumStringValues() []string {
	return []string{
		"SHORT",
		"MEDIUM",
		"LONG",
		"AUTO",
	}
}

// GetMappingSummarizeTextDetailsLengthEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeTextDetailsLengthEnum(val string) (SummarizeTextDetailsLengthEnum, bool) {
	enum, ok := mappingSummarizeTextDetailsLengthEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeTextDetailsFormatEnum Enum with underlying type: string
type SummarizeTextDetailsFormatEnum string

// Set of constants representing the allowable values for SummarizeTextDetailsFormatEnum
const (
	SummarizeTextDetailsFormatParagraph SummarizeTextDetailsFormatEnum = "PARAGRAPH"
	SummarizeTextDetailsFormatBullets   SummarizeTextDetailsFormatEnum = "BULLETS"
	SummarizeTextDetailsFormatAuto      SummarizeTextDetailsFormatEnum = "AUTO"
)

var mappingSummarizeTextDetailsFormatEnum = map[string]SummarizeTextDetailsFormatEnum{
	"PARAGRAPH": SummarizeTextDetailsFormatParagraph,
	"BULLETS":   SummarizeTextDetailsFormatBullets,
	"AUTO":      SummarizeTextDetailsFormatAuto,
}

var mappingSummarizeTextDetailsFormatEnumLowerCase = map[string]SummarizeTextDetailsFormatEnum{
	"paragraph": SummarizeTextDetailsFormatParagraph,
	"bullets":   SummarizeTextDetailsFormatBullets,
	"auto":      SummarizeTextDetailsFormatAuto,
}

// GetSummarizeTextDetailsFormatEnumValues Enumerates the set of values for SummarizeTextDetailsFormatEnum
func GetSummarizeTextDetailsFormatEnumValues() []SummarizeTextDetailsFormatEnum {
	values := make([]SummarizeTextDetailsFormatEnum, 0)
	for _, v := range mappingSummarizeTextDetailsFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeTextDetailsFormatEnumStringValues Enumerates the set of values in String for SummarizeTextDetailsFormatEnum
func GetSummarizeTextDetailsFormatEnumStringValues() []string {
	return []string{
		"PARAGRAPH",
		"BULLETS",
		"AUTO",
	}
}

// GetMappingSummarizeTextDetailsFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeTextDetailsFormatEnum(val string) (SummarizeTextDetailsFormatEnum, bool) {
	enum, ok := mappingSummarizeTextDetailsFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeTextDetailsExtractivenessEnum Enum with underlying type: string
type SummarizeTextDetailsExtractivenessEnum string

// Set of constants representing the allowable values for SummarizeTextDetailsExtractivenessEnum
const (
	SummarizeTextDetailsExtractivenessLow    SummarizeTextDetailsExtractivenessEnum = "LOW"
	SummarizeTextDetailsExtractivenessMedium SummarizeTextDetailsExtractivenessEnum = "MEDIUM"
	SummarizeTextDetailsExtractivenessHigh   SummarizeTextDetailsExtractivenessEnum = "HIGH"
	SummarizeTextDetailsExtractivenessAuto   SummarizeTextDetailsExtractivenessEnum = "AUTO"
)

var mappingSummarizeTextDetailsExtractivenessEnum = map[string]SummarizeTextDetailsExtractivenessEnum{
	"LOW":    SummarizeTextDetailsExtractivenessLow,
	"MEDIUM": SummarizeTextDetailsExtractivenessMedium,
	"HIGH":   SummarizeTextDetailsExtractivenessHigh,
	"AUTO":   SummarizeTextDetailsExtractivenessAuto,
}

var mappingSummarizeTextDetailsExtractivenessEnumLowerCase = map[string]SummarizeTextDetailsExtractivenessEnum{
	"low":    SummarizeTextDetailsExtractivenessLow,
	"medium": SummarizeTextDetailsExtractivenessMedium,
	"high":   SummarizeTextDetailsExtractivenessHigh,
	"auto":   SummarizeTextDetailsExtractivenessAuto,
}

// GetSummarizeTextDetailsExtractivenessEnumValues Enumerates the set of values for SummarizeTextDetailsExtractivenessEnum
func GetSummarizeTextDetailsExtractivenessEnumValues() []SummarizeTextDetailsExtractivenessEnum {
	values := make([]SummarizeTextDetailsExtractivenessEnum, 0)
	for _, v := range mappingSummarizeTextDetailsExtractivenessEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeTextDetailsExtractivenessEnumStringValues Enumerates the set of values in String for SummarizeTextDetailsExtractivenessEnum
func GetSummarizeTextDetailsExtractivenessEnumStringValues() []string {
	return []string{
		"LOW",
		"MEDIUM",
		"HIGH",
		"AUTO",
	}
}

// GetMappingSummarizeTextDetailsExtractivenessEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeTextDetailsExtractivenessEnum(val string) (SummarizeTextDetailsExtractivenessEnum, bool) {
	enum, ok := mappingSummarizeTextDetailsExtractivenessEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
