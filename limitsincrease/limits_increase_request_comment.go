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

// LimitsIncreaseRequestComment The properties that define a comment in a limit increase request.
type LimitsIncreaseRequestComment struct {

	// The comment message.
	Message *string `mandatory:"true" json:"message"`

	// The time that the comment was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The comment sender type.
	SenderType LimitsIncreaseRequestCommentSenderTypeEnum `mandatory:"true" json:"senderType"`
}

func (m LimitsIncreaseRequestComment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LimitsIncreaseRequestComment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLimitsIncreaseRequestCommentSenderTypeEnum(string(m.SenderType)); !ok && m.SenderType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SenderType: %s. Supported values are: %s.", m.SenderType, strings.Join(GetLimitsIncreaseRequestCommentSenderTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LimitsIncreaseRequestCommentSenderTypeEnum Enum with underlying type: string
type LimitsIncreaseRequestCommentSenderTypeEnum string

// Set of constants representing the allowable values for LimitsIncreaseRequestCommentSenderTypeEnum
const (
	LimitsIncreaseRequestCommentSenderTypeCustomer LimitsIncreaseRequestCommentSenderTypeEnum = "CUSTOMER"
	LimitsIncreaseRequestCommentSenderTypeOracle   LimitsIncreaseRequestCommentSenderTypeEnum = "ORACLE"
)

var mappingLimitsIncreaseRequestCommentSenderTypeEnum = map[string]LimitsIncreaseRequestCommentSenderTypeEnum{
	"CUSTOMER": LimitsIncreaseRequestCommentSenderTypeCustomer,
	"ORACLE":   LimitsIncreaseRequestCommentSenderTypeOracle,
}

var mappingLimitsIncreaseRequestCommentSenderTypeEnumLowerCase = map[string]LimitsIncreaseRequestCommentSenderTypeEnum{
	"customer": LimitsIncreaseRequestCommentSenderTypeCustomer,
	"oracle":   LimitsIncreaseRequestCommentSenderTypeOracle,
}

// GetLimitsIncreaseRequestCommentSenderTypeEnumValues Enumerates the set of values for LimitsIncreaseRequestCommentSenderTypeEnum
func GetLimitsIncreaseRequestCommentSenderTypeEnumValues() []LimitsIncreaseRequestCommentSenderTypeEnum {
	values := make([]LimitsIncreaseRequestCommentSenderTypeEnum, 0)
	for _, v := range mappingLimitsIncreaseRequestCommentSenderTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLimitsIncreaseRequestCommentSenderTypeEnumStringValues Enumerates the set of values in String for LimitsIncreaseRequestCommentSenderTypeEnum
func GetLimitsIncreaseRequestCommentSenderTypeEnumStringValues() []string {
	return []string{
		"CUSTOMER",
		"ORACLE",
	}
}

// GetMappingLimitsIncreaseRequestCommentSenderTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLimitsIncreaseRequestCommentSenderTypeEnum(val string) (LimitsIncreaseRequestCommentSenderTypeEnum, bool) {
	enum, ok := mappingLimitsIncreaseRequestCommentSenderTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
