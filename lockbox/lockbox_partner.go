// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Managed Access API
//
// Use the Managed Access API to approve access requests, create and manage templates, and manage resource approval settings. For more information, see Managed Access Overview (https://docs.oracle.com/iaas/Content/managed-access/home.htm).
// Use the table of contents and search tool to explore the Managed Access API.
//

package lockbox

import (
	"strings"
)

// LockboxPartnerEnum Enum with underlying type: string
type LockboxPartnerEnum string

// Set of constants representing the allowable values for LockboxPartnerEnum
const (
	LockboxPartnerFaaas  LockboxPartnerEnum = "FAAAS"
	LockboxPartnerCanary LockboxPartnerEnum = "CANARY"
)

var mappingLockboxPartnerEnum = map[string]LockboxPartnerEnum{
	"FAAAS":  LockboxPartnerFaaas,
	"CANARY": LockboxPartnerCanary,
}

var mappingLockboxPartnerEnumLowerCase = map[string]LockboxPartnerEnum{
	"faaas":  LockboxPartnerFaaas,
	"canary": LockboxPartnerCanary,
}

// GetLockboxPartnerEnumValues Enumerates the set of values for LockboxPartnerEnum
func GetLockboxPartnerEnumValues() []LockboxPartnerEnum {
	values := make([]LockboxPartnerEnum, 0)
	for _, v := range mappingLockboxPartnerEnum {
		values = append(values, v)
	}
	return values
}

// GetLockboxPartnerEnumStringValues Enumerates the set of values in String for LockboxPartnerEnum
func GetLockboxPartnerEnumStringValues() []string {
	return []string{
		"FAAAS",
		"CANARY",
	}
}

// GetMappingLockboxPartnerEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLockboxPartnerEnum(val string) (LockboxPartnerEnum, bool) {
	enum, ok := mappingLockboxPartnerEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
