// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests.
// For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm).
// **Note**: Before you can create service requests with this API,
// you need to have an Oracle Single Sign On (SSO) account,
// and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Contact Contact details for the customer. Avoid entering confidential information.
type Contact struct {

	// The name of the contact person.
	ContactName *string `mandatory:"false" json:"contactName"`

	// The email of the contact person.
	ContactEmail *string `mandatory:"false" json:"contactEmail"`

	// The email of the contact person.
	Email *string `mandatory:"false" json:"email"`

	// The phone number of the contact person.
	ContactPhone *string `mandatory:"false" json:"contactPhone"`

	// The type of contact, such as primary or alternate.
	ContactType ContactContactTypeEnum `mandatory:"false" json:"contactType,omitempty"`
}

func (m Contact) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Contact) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingContactContactTypeEnum(string(m.ContactType)); !ok && m.ContactType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ContactType: %s. Supported values are: %s.", m.ContactType, strings.Join(GetContactContactTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ContactContactTypeEnum Enum with underlying type: string
type ContactContactTypeEnum string

// Set of constants representing the allowable values for ContactContactTypeEnum
const (
	ContactContactTypePrimary   ContactContactTypeEnum = "PRIMARY"
	ContactContactTypeAlternate ContactContactTypeEnum = "ALTERNATE"
	ContactContactTypeSecondary ContactContactTypeEnum = "SECONDARY"
	ContactContactTypeAdmin     ContactContactTypeEnum = "ADMIN"
	ContactContactTypeManager   ContactContactTypeEnum = "MANAGER"
)

var mappingContactContactTypeEnum = map[string]ContactContactTypeEnum{
	"PRIMARY":   ContactContactTypePrimary,
	"ALTERNATE": ContactContactTypeAlternate,
	"SECONDARY": ContactContactTypeSecondary,
	"ADMIN":     ContactContactTypeAdmin,
	"MANAGER":   ContactContactTypeManager,
}

var mappingContactContactTypeEnumLowerCase = map[string]ContactContactTypeEnum{
	"primary":   ContactContactTypePrimary,
	"alternate": ContactContactTypeAlternate,
	"secondary": ContactContactTypeSecondary,
	"admin":     ContactContactTypeAdmin,
	"manager":   ContactContactTypeManager,
}

// GetContactContactTypeEnumValues Enumerates the set of values for ContactContactTypeEnum
func GetContactContactTypeEnumValues() []ContactContactTypeEnum {
	values := make([]ContactContactTypeEnum, 0)
	for _, v := range mappingContactContactTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetContactContactTypeEnumStringValues Enumerates the set of values in String for ContactContactTypeEnum
func GetContactContactTypeEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"ALTERNATE",
		"SECONDARY",
		"ADMIN",
		"MANAGER",
	}
}

// GetMappingContactContactTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContactContactTypeEnum(val string) (ContactContactTypeEnum, bool) {
	enum, ok := mappingContactContactTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
