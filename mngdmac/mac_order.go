// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Managed Services for Mac API
//
// Use the OCI Managed Services for Mac API to create and manage orders for dedicated, partially-managed Mac servers hosted in an OCI Data Center. For more information, see the user guide documentation for the [OCI Managed Services for Mac]()
//

package mngdmac

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MacOrder A description of a MacOrder resource.
type MacOrder struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the compartment to which the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	OrderDescription *string `mandatory:"true" json:"orderDescription"`

	// Number of macs requested in this MacOrder.
	OrderSize *int `mandatory:"true" json:"orderSize"`

	// Checkbox value that indicates whether the customer completed docusign process.
	IsDocusigned *bool `mandatory:"true" json:"isDocusigned"`

	// Enum indicating the requested shape for the MacDevices.
	Shape MacOrderShapeEnum `mandatory:"true" json:"shape"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Enum indicating the agreed commitment term for the requested MacDevices.
	CommitmentTerm MacOrderCommitmentTermEnum `mandatory:"true" json:"commitmentTerm"`

	// The current status of the MacOrder.
	OrderStatus MacOrderOrderStatusEnum `mandatory:"true" json:"orderStatus"`

	// The current state of the MacOrder.
	LifecycleState MacOrderLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The IP Range specified by the customer for this order.
	IpRange *string `mandatory:"false" json:"ipRange"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// An RFC3339-formatted datetime string containing the date and time that billing for this MacOrder begins.
	TimeBillingStarted *common.SDKTime `mandatory:"false" json:"timeBillingStarted"`

	// An RFC3339-formatted datetime string containing the date and time that billing for this MacOrder ends.
	TimeBillingEnded *common.SDKTime `mandatory:"false" json:"timeBillingEnded"`

	// A message that describes the current state of the MacOrder in more detail.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The reason for the MacOrder cancellation.
	CancelReason *string `mandatory:"false" json:"cancelReason"`

	// An RFC3339-formatted datetime string containing the time this MacOrder was cancelled.
	TimeCanceled *common.SDKTime `mandatory:"false" json:"timeCanceled"`
}

func (m MacOrder) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MacOrder) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMacOrderShapeEnum(string(m.Shape)); !ok && m.Shape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shape: %s. Supported values are: %s.", m.Shape, strings.Join(GetMacOrderShapeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMacOrderCommitmentTermEnum(string(m.CommitmentTerm)); !ok && m.CommitmentTerm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CommitmentTerm: %s. Supported values are: %s.", m.CommitmentTerm, strings.Join(GetMacOrderCommitmentTermEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMacOrderOrderStatusEnum(string(m.OrderStatus)); !ok && m.OrderStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OrderStatus: %s. Supported values are: %s.", m.OrderStatus, strings.Join(GetMacOrderOrderStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMacOrderLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMacOrderLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MacOrderShapeEnum Enum with underlying type: string
type MacOrderShapeEnum string

// Set of constants representing the allowable values for MacOrderShapeEnum
const (
	MacOrderShapeM2ProMacMini32gb1tb      MacOrderShapeEnum = "M2_PRO_MAC_MINI_32GB_1TB"
	MacOrderShapeM2UltraMacStudio64gb4tb  MacOrderShapeEnum = "M2_ULTRA_MAC_STUDIO_64GB_4TB"
	MacOrderShapeM2UltraMacStudio192gb4tb MacOrderShapeEnum = "M2_ULTRA_MAC_STUDIO_192GB_4TB"
	MacOrderShapeM4ProMacMini64gb2tb      MacOrderShapeEnum = "M4_PRO_MAC_MINI_64GB_2TB"
	MacOrderShapeM4ProMacMini64gb4tb      MacOrderShapeEnum = "M4_PRO_MAC_MINI_64GB_4TB"
)

var mappingMacOrderShapeEnum = map[string]MacOrderShapeEnum{
	"M2_PRO_MAC_MINI_32GB_1TB":      MacOrderShapeM2ProMacMini32gb1tb,
	"M2_ULTRA_MAC_STUDIO_64GB_4TB":  MacOrderShapeM2UltraMacStudio64gb4tb,
	"M2_ULTRA_MAC_STUDIO_192GB_4TB": MacOrderShapeM2UltraMacStudio192gb4tb,
	"M4_PRO_MAC_MINI_64GB_2TB":      MacOrderShapeM4ProMacMini64gb2tb,
	"M4_PRO_MAC_MINI_64GB_4TB":      MacOrderShapeM4ProMacMini64gb4tb,
}

var mappingMacOrderShapeEnumLowerCase = map[string]MacOrderShapeEnum{
	"m2_pro_mac_mini_32gb_1tb":      MacOrderShapeM2ProMacMini32gb1tb,
	"m2_ultra_mac_studio_64gb_4tb":  MacOrderShapeM2UltraMacStudio64gb4tb,
	"m2_ultra_mac_studio_192gb_4tb": MacOrderShapeM2UltraMacStudio192gb4tb,
	"m4_pro_mac_mini_64gb_2tb":      MacOrderShapeM4ProMacMini64gb2tb,
	"m4_pro_mac_mini_64gb_4tb":      MacOrderShapeM4ProMacMini64gb4tb,
}

// GetMacOrderShapeEnumValues Enumerates the set of values for MacOrderShapeEnum
func GetMacOrderShapeEnumValues() []MacOrderShapeEnum {
	values := make([]MacOrderShapeEnum, 0)
	for _, v := range mappingMacOrderShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetMacOrderShapeEnumStringValues Enumerates the set of values in String for MacOrderShapeEnum
func GetMacOrderShapeEnumStringValues() []string {
	return []string{
		"M2_PRO_MAC_MINI_32GB_1TB",
		"M2_ULTRA_MAC_STUDIO_64GB_4TB",
		"M2_ULTRA_MAC_STUDIO_192GB_4TB",
		"M4_PRO_MAC_MINI_64GB_2TB",
		"M4_PRO_MAC_MINI_64GB_4TB",
	}
}

// GetMappingMacOrderShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMacOrderShapeEnum(val string) (MacOrderShapeEnum, bool) {
	enum, ok := mappingMacOrderShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MacOrderCommitmentTermEnum Enum with underlying type: string
type MacOrderCommitmentTermEnum string

// Set of constants representing the allowable values for MacOrderCommitmentTermEnum
const (
	MacOrderCommitmentTermYears3 MacOrderCommitmentTermEnum = "YEARS_3"
)

var mappingMacOrderCommitmentTermEnum = map[string]MacOrderCommitmentTermEnum{
	"YEARS_3": MacOrderCommitmentTermYears3,
}

var mappingMacOrderCommitmentTermEnumLowerCase = map[string]MacOrderCommitmentTermEnum{
	"years_3": MacOrderCommitmentTermYears3,
}

// GetMacOrderCommitmentTermEnumValues Enumerates the set of values for MacOrderCommitmentTermEnum
func GetMacOrderCommitmentTermEnumValues() []MacOrderCommitmentTermEnum {
	values := make([]MacOrderCommitmentTermEnum, 0)
	for _, v := range mappingMacOrderCommitmentTermEnum {
		values = append(values, v)
	}
	return values
}

// GetMacOrderCommitmentTermEnumStringValues Enumerates the set of values in String for MacOrderCommitmentTermEnum
func GetMacOrderCommitmentTermEnumStringValues() []string {
	return []string{
		"YEARS_3",
	}
}

// GetMappingMacOrderCommitmentTermEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMacOrderCommitmentTermEnum(val string) (MacOrderCommitmentTermEnum, bool) {
	enum, ok := mappingMacOrderCommitmentTermEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MacOrderOrderStatusEnum Enum with underlying type: string
type MacOrderOrderStatusEnum string

// Set of constants representing the allowable values for MacOrderOrderStatusEnum
const (
	MacOrderOrderStatusSubmitted        MacOrderOrderStatusEnum = "SUBMITTED"
	MacOrderOrderStatusPendingDocusign  MacOrderOrderStatusEnum = "PENDING_DOCUSIGN"
	MacOrderOrderStatusOrderReview      MacOrderOrderStatusEnum = "ORDER_REVIEW"
	MacOrderOrderStatusCustomerReview   MacOrderOrderStatusEnum = "CUSTOMER_REVIEW"
	MacOrderOrderStatusCustomerApproved MacOrderOrderStatusEnum = "CUSTOMER_APPROVED"
	MacOrderOrderStatusFundingApproved  MacOrderOrderStatusEnum = "FUNDING_APPROVED"
	MacOrderOrderStatusProvisioning     MacOrderOrderStatusEnum = "PROVISIONING"
	MacOrderOrderStatusCustomerAccepted MacOrderOrderStatusEnum = "CUSTOMER_ACCEPTED"
	MacOrderOrderStatusCompleted        MacOrderOrderStatusEnum = "COMPLETED"
	MacOrderOrderStatusCanceled         MacOrderOrderStatusEnum = "CANCELED"
)

var mappingMacOrderOrderStatusEnum = map[string]MacOrderOrderStatusEnum{
	"SUBMITTED":         MacOrderOrderStatusSubmitted,
	"PENDING_DOCUSIGN":  MacOrderOrderStatusPendingDocusign,
	"ORDER_REVIEW":      MacOrderOrderStatusOrderReview,
	"CUSTOMER_REVIEW":   MacOrderOrderStatusCustomerReview,
	"CUSTOMER_APPROVED": MacOrderOrderStatusCustomerApproved,
	"FUNDING_APPROVED":  MacOrderOrderStatusFundingApproved,
	"PROVISIONING":      MacOrderOrderStatusProvisioning,
	"CUSTOMER_ACCEPTED": MacOrderOrderStatusCustomerAccepted,
	"COMPLETED":         MacOrderOrderStatusCompleted,
	"CANCELED":          MacOrderOrderStatusCanceled,
}

var mappingMacOrderOrderStatusEnumLowerCase = map[string]MacOrderOrderStatusEnum{
	"submitted":         MacOrderOrderStatusSubmitted,
	"pending_docusign":  MacOrderOrderStatusPendingDocusign,
	"order_review":      MacOrderOrderStatusOrderReview,
	"customer_review":   MacOrderOrderStatusCustomerReview,
	"customer_approved": MacOrderOrderStatusCustomerApproved,
	"funding_approved":  MacOrderOrderStatusFundingApproved,
	"provisioning":      MacOrderOrderStatusProvisioning,
	"customer_accepted": MacOrderOrderStatusCustomerAccepted,
	"completed":         MacOrderOrderStatusCompleted,
	"canceled":          MacOrderOrderStatusCanceled,
}

// GetMacOrderOrderStatusEnumValues Enumerates the set of values for MacOrderOrderStatusEnum
func GetMacOrderOrderStatusEnumValues() []MacOrderOrderStatusEnum {
	values := make([]MacOrderOrderStatusEnum, 0)
	for _, v := range mappingMacOrderOrderStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetMacOrderOrderStatusEnumStringValues Enumerates the set of values in String for MacOrderOrderStatusEnum
func GetMacOrderOrderStatusEnumStringValues() []string {
	return []string{
		"SUBMITTED",
		"PENDING_DOCUSIGN",
		"ORDER_REVIEW",
		"CUSTOMER_REVIEW",
		"CUSTOMER_APPROVED",
		"FUNDING_APPROVED",
		"PROVISIONING",
		"CUSTOMER_ACCEPTED",
		"COMPLETED",
		"CANCELED",
	}
}

// GetMappingMacOrderOrderStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMacOrderOrderStatusEnum(val string) (MacOrderOrderStatusEnum, bool) {
	enum, ok := mappingMacOrderOrderStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MacOrderLifecycleStateEnum Enum with underlying type: string
type MacOrderLifecycleStateEnum string

// Set of constants representing the allowable values for MacOrderLifecycleStateEnum
const (
	MacOrderLifecycleStateCreating       MacOrderLifecycleStateEnum = "CREATING"
	MacOrderLifecycleStateUpdating       MacOrderLifecycleStateEnum = "UPDATING"
	MacOrderLifecycleStateActive         MacOrderLifecycleStateEnum = "ACTIVE"
	MacOrderLifecycleStateNeedsAttention MacOrderLifecycleStateEnum = "NEEDS_ATTENTION"
	MacOrderLifecycleStateDeleting       MacOrderLifecycleStateEnum = "DELETING"
	MacOrderLifecycleStateDeleted        MacOrderLifecycleStateEnum = "DELETED"
	MacOrderLifecycleStateFailed         MacOrderLifecycleStateEnum = "FAILED"
)

var mappingMacOrderLifecycleStateEnum = map[string]MacOrderLifecycleStateEnum{
	"CREATING":        MacOrderLifecycleStateCreating,
	"UPDATING":        MacOrderLifecycleStateUpdating,
	"ACTIVE":          MacOrderLifecycleStateActive,
	"NEEDS_ATTENTION": MacOrderLifecycleStateNeedsAttention,
	"DELETING":        MacOrderLifecycleStateDeleting,
	"DELETED":         MacOrderLifecycleStateDeleted,
	"FAILED":          MacOrderLifecycleStateFailed,
}

var mappingMacOrderLifecycleStateEnumLowerCase = map[string]MacOrderLifecycleStateEnum{
	"creating":        MacOrderLifecycleStateCreating,
	"updating":        MacOrderLifecycleStateUpdating,
	"active":          MacOrderLifecycleStateActive,
	"needs_attention": MacOrderLifecycleStateNeedsAttention,
	"deleting":        MacOrderLifecycleStateDeleting,
	"deleted":         MacOrderLifecycleStateDeleted,
	"failed":          MacOrderLifecycleStateFailed,
}

// GetMacOrderLifecycleStateEnumValues Enumerates the set of values for MacOrderLifecycleStateEnum
func GetMacOrderLifecycleStateEnumValues() []MacOrderLifecycleStateEnum {
	values := make([]MacOrderLifecycleStateEnum, 0)
	for _, v := range mappingMacOrderLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMacOrderLifecycleStateEnumStringValues Enumerates the set of values in String for MacOrderLifecycleStateEnum
func GetMacOrderLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingMacOrderLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMacOrderLifecycleStateEnum(val string) (MacOrderLifecycleStateEnum, bool) {
	enum, ok := mappingMacOrderLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
