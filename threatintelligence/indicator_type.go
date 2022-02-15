// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Threat Intelligence API
//
// Use the Threat Intelligence API to view indicators of compromise and related items. For more information, see Overview of Threat Intelligence (https://docs.cloud.oracle.com/Content/ThreatIntelligence/Concepts/threatintelligenceoverview.htm).
//

package threatintelligence

import (
	"strings"
)

// IndicatorTypeEnum Enum with underlying type: string
type IndicatorTypeEnum string

// Set of constants representing the allowable values for IndicatorTypeEnum
const (
	IndicatorTypeDomainName IndicatorTypeEnum = "DOMAIN_NAME"
	IndicatorTypeFileName   IndicatorTypeEnum = "FILE_NAME"
	IndicatorTypeMd5Hash    IndicatorTypeEnum = "MD5_HASH"
	IndicatorTypeSha1Hash   IndicatorTypeEnum = "SHA1_HASH"
	IndicatorTypeSha256Hash IndicatorTypeEnum = "SHA256_HASH"
	IndicatorTypeIpAddress  IndicatorTypeEnum = "IP_ADDRESS"
	IndicatorTypeUrl        IndicatorTypeEnum = "URL"
)

var mappingIndicatorTypeEnum = map[string]IndicatorTypeEnum{
	"DOMAIN_NAME": IndicatorTypeDomainName,
	"FILE_NAME":   IndicatorTypeFileName,
	"MD5_HASH":    IndicatorTypeMd5Hash,
	"SHA1_HASH":   IndicatorTypeSha1Hash,
	"SHA256_HASH": IndicatorTypeSha256Hash,
	"IP_ADDRESS":  IndicatorTypeIpAddress,
	"URL":         IndicatorTypeUrl,
}

// GetIndicatorTypeEnumValues Enumerates the set of values for IndicatorTypeEnum
func GetIndicatorTypeEnumValues() []IndicatorTypeEnum {
	values := make([]IndicatorTypeEnum, 0)
	for _, v := range mappingIndicatorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetIndicatorTypeEnumStringValues Enumerates the set of values in String for IndicatorTypeEnum
func GetIndicatorTypeEnumStringValues() []string {
	return []string{
		"DOMAIN_NAME",
		"FILE_NAME",
		"MD5_HASH",
		"SHA1_HASH",
		"SHA256_HASH",
		"IP_ADDRESS",
		"URL",
	}
}

// GetMappingIndicatorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIndicatorTypeEnum(val string) (IndicatorTypeEnum, bool) {
	mappingIndicatorTypeEnumIgnoreCase := make(map[string]IndicatorTypeEnum)
	for k, v := range mappingIndicatorTypeEnum {
		mappingIndicatorTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingIndicatorTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
