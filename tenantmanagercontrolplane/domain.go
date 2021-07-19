// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// Domain The domain model that is associated with a tenancy.
type Domain struct {

	// The OCID of the domain.
	Id *string `mandatory:"true" json:"id"`

	// The domain name.
	DomainName *string `mandatory:"true" json:"domainName"`

	// The OCID of the tenancy that has started the registration process for this domain.
	OwnerId *string `mandatory:"true" json:"ownerId"`

	// Lifecycle state of the domain.
	LifecycleState DomainLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Status of the domain.
	Status DomainStatusEnum `mandatory:"true" json:"status"`

	// The code that the owner of the domain will need to add as a TXT record to their domain.
	TxtRecord *string `mandatory:"true" json:"txtRecord"`

	// Date-time when this domain was created. An RFC3339-formatted date and time string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Date-time when this domain was last updated. An RFC3339-formatted date and time string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Domain) String() string {
	return common.PointerString(m)
}

// DomainLifecycleStateEnum Enum with underlying type: string
type DomainLifecycleStateEnum string

// Set of constants representing the allowable values for DomainLifecycleStateEnum
const (
	DomainLifecycleStateActive  DomainLifecycleStateEnum = "ACTIVE"
	DomainLifecycleStateDeleted DomainLifecycleStateEnum = "DELETED"
	DomainLifecycleStateFailed  DomainLifecycleStateEnum = "FAILED"
)

var mappingDomainLifecycleState = map[string]DomainLifecycleStateEnum{
	"ACTIVE":  DomainLifecycleStateActive,
	"DELETED": DomainLifecycleStateDeleted,
	"FAILED":  DomainLifecycleStateFailed,
}

// GetDomainLifecycleStateEnumValues Enumerates the set of values for DomainLifecycleStateEnum
func GetDomainLifecycleStateEnumValues() []DomainLifecycleStateEnum {
	values := make([]DomainLifecycleStateEnum, 0)
	for _, v := range mappingDomainLifecycleState {
		values = append(values, v)
	}
	return values
}

// DomainStatusEnum Enum with underlying type: string
type DomainStatusEnum string

// Set of constants representing the allowable values for DomainStatusEnum
const (
	DomainStatusPending   DomainStatusEnum = "PENDING"
	DomainStatusReleasing DomainStatusEnum = "RELEASING"
	DomainStatusReleased  DomainStatusEnum = "RELEASED"
	DomainStatusExpiring  DomainStatusEnum = "EXPIRING"
	DomainStatusRevoking  DomainStatusEnum = "REVOKING"
	DomainStatusRevoked   DomainStatusEnum = "REVOKED"
	DomainStatusActive    DomainStatusEnum = "ACTIVE"
	DomainStatusFailed    DomainStatusEnum = "FAILED"
)

var mappingDomainStatus = map[string]DomainStatusEnum{
	"PENDING":   DomainStatusPending,
	"RELEASING": DomainStatusReleasing,
	"RELEASED":  DomainStatusReleased,
	"EXPIRING":  DomainStatusExpiring,
	"REVOKING":  DomainStatusRevoking,
	"REVOKED":   DomainStatusRevoked,
	"ACTIVE":    DomainStatusActive,
	"FAILED":    DomainStatusFailed,
}

// GetDomainStatusEnumValues Enumerates the set of values for DomainStatusEnum
func GetDomainStatusEnumValues() []DomainStatusEnum {
	values := make([]DomainStatusEnum, 0)
	for _, v := range mappingDomainStatus {
		values = append(values, v)
	}
	return values
}
