// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MarketplacePublisherService API
//
// Use the Marketplace Publisher API to manage the publishing of applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplacepublisher

import (
	"strings"
)

// ValidationStatusEnum Enum with underlying type: string
type ValidationStatusEnum string

// Set of constants representing the allowable values for ValidationStatusEnum
const (
	ValidationStatusValidationInProgress ValidationStatusEnum = "VALIDATION_IN_PROGRESS"
	ValidationStatusValidationFailed     ValidationStatusEnum = "VALIDATION_FAILED"
	ValidationStatusValidationCompleted  ValidationStatusEnum = "VALIDATION_COMPLETED"
)

var mappingValidationStatusEnum = map[string]ValidationStatusEnum{
	"VALIDATION_IN_PROGRESS": ValidationStatusValidationInProgress,
	"VALIDATION_FAILED":      ValidationStatusValidationFailed,
	"VALIDATION_COMPLETED":   ValidationStatusValidationCompleted,
}

var mappingValidationStatusEnumLowerCase = map[string]ValidationStatusEnum{
	"validation_in_progress": ValidationStatusValidationInProgress,
	"validation_failed":      ValidationStatusValidationFailed,
	"validation_completed":   ValidationStatusValidationCompleted,
}

// GetValidationStatusEnumValues Enumerates the set of values for ValidationStatusEnum
func GetValidationStatusEnumValues() []ValidationStatusEnum {
	values := make([]ValidationStatusEnum, 0)
	for _, v := range mappingValidationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetValidationStatusEnumStringValues Enumerates the set of values in String for ValidationStatusEnum
func GetValidationStatusEnumStringValues() []string {
	return []string{
		"VALIDATION_IN_PROGRESS",
		"VALIDATION_FAILED",
		"VALIDATION_COMPLETED",
	}
}

// GetMappingValidationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingValidationStatusEnum(val string) (ValidationStatusEnum, bool) {
	enum, ok := mappingValidationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
