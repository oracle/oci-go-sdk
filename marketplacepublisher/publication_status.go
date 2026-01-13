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

// PublicationStatusEnum Enum with underlying type: string
type PublicationStatusEnum string

// Set of constants representing the allowable values for PublicationStatusEnum
const (
	PublicationStatusPublicationInProgress PublicationStatusEnum = "PUBLICATION_IN_PROGRESS"
	PublicationStatusPublicationCompleted  PublicationStatusEnum = "PUBLICATION_COMPLETED"
	PublicationStatusPublicationFailed     PublicationStatusEnum = "PUBLICATION_FAILED"
)

var mappingPublicationStatusEnum = map[string]PublicationStatusEnum{
	"PUBLICATION_IN_PROGRESS": PublicationStatusPublicationInProgress,
	"PUBLICATION_COMPLETED":   PublicationStatusPublicationCompleted,
	"PUBLICATION_FAILED":      PublicationStatusPublicationFailed,
}

var mappingPublicationStatusEnumLowerCase = map[string]PublicationStatusEnum{
	"publication_in_progress": PublicationStatusPublicationInProgress,
	"publication_completed":   PublicationStatusPublicationCompleted,
	"publication_failed":      PublicationStatusPublicationFailed,
}

// GetPublicationStatusEnumValues Enumerates the set of values for PublicationStatusEnum
func GetPublicationStatusEnumValues() []PublicationStatusEnum {
	values := make([]PublicationStatusEnum, 0)
	for _, v := range mappingPublicationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPublicationStatusEnumStringValues Enumerates the set of values in String for PublicationStatusEnum
func GetPublicationStatusEnumStringValues() []string {
	return []string{
		"PUBLICATION_IN_PROGRESS",
		"PUBLICATION_COMPLETED",
		"PUBLICATION_FAILED",
	}
}

// GetMappingPublicationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPublicationStatusEnum(val string) (PublicationStatusEnum, bool) {
	enum, ok := mappingPublicationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
