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

// Publisher The model for a publisher details.
type Publisher struct {

	// Unique OCID identifier for the publisher.
	Id *string `mandatory:"true" json:"id"`

	// The root compartment of the Publisher.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The namespace for the publisher registry to persist artifacts.
	RegistryNamespace *string `mandatory:"true" json:"registryNamespace"`

	// The name of the publisher.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The public email address of the publisher for customers.
	ContactEmail *string `mandatory:"true" json:"contactEmail"`

	// The phone number of the publisher in E.164 format.
	ContactPhone *string `mandatory:"true" json:"contactPhone"`

	// publisher type.
	PublisherType PublisherPublisherTypeEnum `mandatory:"true" json:"publisherType"`

	// The time the publisher was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the publisher was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// publisher status.
	PublisherStatus PublisherPublisherStatusEnum `mandatory:"true" json:"publisherStatus"`

	// Unique legacy service identifier for the publisher.
	LegacyId *string `mandatory:"false" json:"legacyId"`

	// A description of the publisher.
	Description *string `mandatory:"false" json:"description"`

	// The year the publisher's company or organization was founded.
	YearFounded *int64 `mandatory:"false" json:"yearFounded"`

	// The publisher's website.
	WebsiteUrl *string `mandatory:"false" json:"websiteUrl"`

	// The address of the publisher's headquarters.
	HqAddress *string `mandatory:"false" json:"hqAddress"`

	Logo *UploadData `mandatory:"false" json:"logo"`

	// Publisher's Facebook URL
	FacebookUrl *string `mandatory:"false" json:"facebookUrl"`

	// Publisher's Twitter URL
	TwitterUrl *string `mandatory:"false" json:"twitterUrl"`

	// Publisher's LinkedIn URL
	LinkedinUrl *string `mandatory:"false" json:"linkedinUrl"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The private email address of the publisher product team.
	NotificationEmail *string `mandatory:"false" json:"notificationEmail"`

	OpnMembership *OpnMembership `mandatory:"false" json:"opnMembership"`

	PrivateOfferAccountDetails *PrivateOfferAccountDetails `mandatory:"false" json:"privateOfferAccountDetails"`
}

func (m Publisher) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Publisher) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPublisherPublisherTypeEnum(string(m.PublisherType)); !ok && m.PublisherType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PublisherType: %s. Supported values are: %s.", m.PublisherType, strings.Join(GetPublisherPublisherTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPublisherPublisherStatusEnum(string(m.PublisherStatus)); !ok && m.PublisherStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PublisherStatus: %s. Supported values are: %s.", m.PublisherStatus, strings.Join(GetPublisherPublisherStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PublisherPublisherTypeEnum Enum with underlying type: string
type PublisherPublisherTypeEnum string

// Set of constants representing the allowable values for PublisherPublisherTypeEnum
const (
	PublisherPublisherTypeInternal PublisherPublisherTypeEnum = "INTERNAL"
	PublisherPublisherTypeExternal PublisherPublisherTypeEnum = "EXTERNAL"
)

var mappingPublisherPublisherTypeEnum = map[string]PublisherPublisherTypeEnum{
	"INTERNAL": PublisherPublisherTypeInternal,
	"EXTERNAL": PublisherPublisherTypeExternal,
}

var mappingPublisherPublisherTypeEnumLowerCase = map[string]PublisherPublisherTypeEnum{
	"internal": PublisherPublisherTypeInternal,
	"external": PublisherPublisherTypeExternal,
}

// GetPublisherPublisherTypeEnumValues Enumerates the set of values for PublisherPublisherTypeEnum
func GetPublisherPublisherTypeEnumValues() []PublisherPublisherTypeEnum {
	values := make([]PublisherPublisherTypeEnum, 0)
	for _, v := range mappingPublisherPublisherTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPublisherPublisherTypeEnumStringValues Enumerates the set of values in String for PublisherPublisherTypeEnum
func GetPublisherPublisherTypeEnumStringValues() []string {
	return []string{
		"INTERNAL",
		"EXTERNAL",
	}
}

// GetMappingPublisherPublisherTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPublisherPublisherTypeEnum(val string) (PublisherPublisherTypeEnum, bool) {
	enum, ok := mappingPublisherPublisherTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PublisherPublisherStatusEnum Enum with underlying type: string
type PublisherPublisherStatusEnum string

// Set of constants representing the allowable values for PublisherPublisherStatusEnum
const (
	PublisherPublisherStatusNew       PublisherPublisherStatusEnum = "NEW"
	PublisherPublisherStatusApproved  PublisherPublisherStatusEnum = "APPROVED"
	PublisherPublisherStatusSuspended PublisherPublisherStatusEnum = "SUSPENDED"
	PublisherPublisherStatusRemoved   PublisherPublisherStatusEnum = "REMOVED"
	PublisherPublisherStatusRejected  PublisherPublisherStatusEnum = "REJECTED"
	PublisherPublisherStatusContacted PublisherPublisherStatusEnum = "CONTACTED"
)

var mappingPublisherPublisherStatusEnum = map[string]PublisherPublisherStatusEnum{
	"NEW":       PublisherPublisherStatusNew,
	"APPROVED":  PublisherPublisherStatusApproved,
	"SUSPENDED": PublisherPublisherStatusSuspended,
	"REMOVED":   PublisherPublisherStatusRemoved,
	"REJECTED":  PublisherPublisherStatusRejected,
	"CONTACTED": PublisherPublisherStatusContacted,
}

var mappingPublisherPublisherStatusEnumLowerCase = map[string]PublisherPublisherStatusEnum{
	"new":       PublisherPublisherStatusNew,
	"approved":  PublisherPublisherStatusApproved,
	"suspended": PublisherPublisherStatusSuspended,
	"removed":   PublisherPublisherStatusRemoved,
	"rejected":  PublisherPublisherStatusRejected,
	"contacted": PublisherPublisherStatusContacted,
}

// GetPublisherPublisherStatusEnumValues Enumerates the set of values for PublisherPublisherStatusEnum
func GetPublisherPublisherStatusEnumValues() []PublisherPublisherStatusEnum {
	values := make([]PublisherPublisherStatusEnum, 0)
	for _, v := range mappingPublisherPublisherStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPublisherPublisherStatusEnumStringValues Enumerates the set of values in String for PublisherPublisherStatusEnum
func GetPublisherPublisherStatusEnumStringValues() []string {
	return []string{
		"NEW",
		"APPROVED",
		"SUSPENDED",
		"REMOVED",
		"REJECTED",
		"CONTACTED",
	}
}

// GetMappingPublisherPublisherStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPublisherPublisherStatusEnum(val string) (PublisherPublisherStatusEnum, bool) {
	enum, ok := mappingPublisherPublisherStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
