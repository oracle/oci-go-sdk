// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Limits Increase API
//
// Use the Limits Increase API to work with limit increase requests.
// For more information, see
// Working with Limit Increase Requests (https://docs.oracle.com/iaas/Content/General/service-limits/requests.htm).
//

package limitsincrease

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LimitsIncreaseQuestionSummary A summary of properties for a question to be asked for a certain limit.
// Example questions include database version for some limits in Autonomous AI Database or email domain for some limits in Email Delivery.
// For more information, see
// Creating a Limit Increase Request (https://docs.oracle.com/iaas/Content/General/service-limits/create-request.htm).
type LimitsIncreaseQuestionSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the question.
	Id *string `mandatory:"true" json:"id"`

	// The name of the service for the question.
	Service *string `mandatory:"true" json:"service"`

	// The text for the question.
	QuestionText *string `mandatory:"true" json:"questionText"`

	// The type of question.
	QuestionType LimitsIncreaseQuestionSummaryQuestionTypeEnum `mandatory:"true" json:"questionType"`

	// When `true`, requires an answer to the question.
	IsRequired *bool `mandatory:"true" json:"isRequired"`

	// The name of the limit for the question (empty if the question is for the service).
	LimitName *string `mandatory:"false" json:"limitName"`

	// Set options for the question. Applies to questions that aren't free text.
	Options map[string]string `mandatory:"false" json:"options"`
}

func (m LimitsIncreaseQuestionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LimitsIncreaseQuestionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLimitsIncreaseQuestionSummaryQuestionTypeEnum(string(m.QuestionType)); !ok && m.QuestionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for QuestionType: %s. Supported values are: %s.", m.QuestionType, strings.Join(GetLimitsIncreaseQuestionSummaryQuestionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LimitsIncreaseQuestionSummaryQuestionTypeEnum Enum with underlying type: string
type LimitsIncreaseQuestionSummaryQuestionTypeEnum string

// Set of constants representing the allowable values for LimitsIncreaseQuestionSummaryQuestionTypeEnum
const (
	LimitsIncreaseQuestionSummaryQuestionTypeText     LimitsIncreaseQuestionSummaryQuestionTypeEnum = "TEXT"
	LimitsIncreaseQuestionSummaryQuestionTypeNumber   LimitsIncreaseQuestionSummaryQuestionTypeEnum = "NUMBER"
	LimitsIncreaseQuestionSummaryQuestionTypeRadio    LimitsIncreaseQuestionSummaryQuestionTypeEnum = "RADIO"
	LimitsIncreaseQuestionSummaryQuestionTypeCheckbox LimitsIncreaseQuestionSummaryQuestionTypeEnum = "CHECKBOX"
)

var mappingLimitsIncreaseQuestionSummaryQuestionTypeEnum = map[string]LimitsIncreaseQuestionSummaryQuestionTypeEnum{
	"TEXT":     LimitsIncreaseQuestionSummaryQuestionTypeText,
	"NUMBER":   LimitsIncreaseQuestionSummaryQuestionTypeNumber,
	"RADIO":    LimitsIncreaseQuestionSummaryQuestionTypeRadio,
	"CHECKBOX": LimitsIncreaseQuestionSummaryQuestionTypeCheckbox,
}

var mappingLimitsIncreaseQuestionSummaryQuestionTypeEnumLowerCase = map[string]LimitsIncreaseQuestionSummaryQuestionTypeEnum{
	"text":     LimitsIncreaseQuestionSummaryQuestionTypeText,
	"number":   LimitsIncreaseQuestionSummaryQuestionTypeNumber,
	"radio":    LimitsIncreaseQuestionSummaryQuestionTypeRadio,
	"checkbox": LimitsIncreaseQuestionSummaryQuestionTypeCheckbox,
}

// GetLimitsIncreaseQuestionSummaryQuestionTypeEnumValues Enumerates the set of values for LimitsIncreaseQuestionSummaryQuestionTypeEnum
func GetLimitsIncreaseQuestionSummaryQuestionTypeEnumValues() []LimitsIncreaseQuestionSummaryQuestionTypeEnum {
	values := make([]LimitsIncreaseQuestionSummaryQuestionTypeEnum, 0)
	for _, v := range mappingLimitsIncreaseQuestionSummaryQuestionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLimitsIncreaseQuestionSummaryQuestionTypeEnumStringValues Enumerates the set of values in String for LimitsIncreaseQuestionSummaryQuestionTypeEnum
func GetLimitsIncreaseQuestionSummaryQuestionTypeEnumStringValues() []string {
	return []string{
		"TEXT",
		"NUMBER",
		"RADIO",
		"CHECKBOX",
	}
}

// GetMappingLimitsIncreaseQuestionSummaryQuestionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLimitsIncreaseQuestionSummaryQuestionTypeEnum(val string) (LimitsIncreaseQuestionSummaryQuestionTypeEnum, bool) {
	enum, ok := mappingLimitsIncreaseQuestionSummaryQuestionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
