// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// PublisherStatusEnum Enum with underlying type: string
type PublisherStatusEnum string

// Set of constants representing the allowable values for PublisherStatusEnum
const (
	PublisherStatusNew       PublisherStatusEnum = "NEW"
	PublisherStatusApproved  PublisherStatusEnum = "APPROVED"
	PublisherStatusSuspended PublisherStatusEnum = "SUSPENDED"
	PublisherStatusRemoved   PublisherStatusEnum = "REMOVED"
	PublisherStatusRejected  PublisherStatusEnum = "REJECTED"
	PublisherStatusContacted PublisherStatusEnum = "CONTACTED"
)

var mappingPublisherStatusEnum = map[string]PublisherStatusEnum{
	"NEW":       PublisherStatusNew,
	"APPROVED":  PublisherStatusApproved,
	"SUSPENDED": PublisherStatusSuspended,
	"REMOVED":   PublisherStatusRemoved,
	"REJECTED":  PublisherStatusRejected,
	"CONTACTED": PublisherStatusContacted,
}

var mappingPublisherStatusEnumLowerCase = map[string]PublisherStatusEnum{
	"new":       PublisherStatusNew,
	"approved":  PublisherStatusApproved,
	"suspended": PublisherStatusSuspended,
	"removed":   PublisherStatusRemoved,
	"rejected":  PublisherStatusRejected,
	"contacted": PublisherStatusContacted,
}

// GetPublisherStatusEnumValues Enumerates the set of values for PublisherStatusEnum
func GetPublisherStatusEnumValues() []PublisherStatusEnum {
	values := make([]PublisherStatusEnum, 0)
	for _, v := range mappingPublisherStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPublisherStatusEnumStringValues Enumerates the set of values in String for PublisherStatusEnum
func GetPublisherStatusEnumStringValues() []string {
	return []string{
		"NEW",
		"APPROVED",
		"SUSPENDED",
		"REMOVED",
		"REJECTED",
		"CONTACTED",
	}
}

// GetMappingPublisherStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPublisherStatusEnum(val string) (PublisherStatusEnum, bool) {
	enum, ok := mappingPublisherStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
