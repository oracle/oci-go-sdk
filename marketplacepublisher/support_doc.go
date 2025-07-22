// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MarketplacePublisherService API
//
// Use the Marketplace Publisher API to manage the publishing of applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplacepublisher

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SupportDoc The model for the supportDoc details.
type SupportDoc struct {

	// The name for the supportDoc.
	Name *string `mandatory:"true" json:"name"`

	// The code for the supportDoc.
	Code *string `mandatory:"true" json:"code"`

	// The supportDoc group for the supportDoc.
	SupportDocGroup *string `mandatory:"true" json:"supportDocGroup"`

	// The current state for the supportDoc.
	LifecycleState SupportDocLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the supportDoc was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the supportDoc was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The description of the supportDoc.
	Description *string `mandatory:"false" json:"description"`

	// The URL of the specified attachment.
	ContentUrl *string `mandatory:"false" json:"contentUrl"`

	// The MIME type of the attachment.
	MimeType *string `mandatory:"false" json:"mimeType"`

	// The file name of the attachment.
	FileName *string `mandatory:"false" json:"fileName"`
}

func (m SupportDoc) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SupportDoc) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSupportDocLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSupportDocLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SupportDocLifecycleStateEnum Enum with underlying type: string
type SupportDocLifecycleStateEnum string

// Set of constants representing the allowable values for SupportDocLifecycleStateEnum
const (
	SupportDocLifecycleStateActive   SupportDocLifecycleStateEnum = "ACTIVE"
	SupportDocLifecycleStateInactive SupportDocLifecycleStateEnum = "INACTIVE"
)

var mappingSupportDocLifecycleStateEnum = map[string]SupportDocLifecycleStateEnum{
	"ACTIVE":   SupportDocLifecycleStateActive,
	"INACTIVE": SupportDocLifecycleStateInactive,
}

var mappingSupportDocLifecycleStateEnumLowerCase = map[string]SupportDocLifecycleStateEnum{
	"active":   SupportDocLifecycleStateActive,
	"inactive": SupportDocLifecycleStateInactive,
}

// GetSupportDocLifecycleStateEnumValues Enumerates the set of values for SupportDocLifecycleStateEnum
func GetSupportDocLifecycleStateEnumValues() []SupportDocLifecycleStateEnum {
	values := make([]SupportDocLifecycleStateEnum, 0)
	for _, v := range mappingSupportDocLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSupportDocLifecycleStateEnumStringValues Enumerates the set of values in String for SupportDocLifecycleStateEnum
func GetSupportDocLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingSupportDocLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSupportDocLifecycleStateEnum(val string) (SupportDocLifecycleStateEnum, bool) {
	enum, ok := mappingSupportDocLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
