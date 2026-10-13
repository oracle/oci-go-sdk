// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Continuous Consumption Discounts API
//
// Use the Continuous Consumption Discounts API to manage continuous consumption discount operations.
//

package ccdoffers

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CcdCommitment The created CCD commitment details.
type CcdCommitment struct {

	// The CCD commitment resource OCID.
	Id *string `mandatory:"true" json:"id"`

	// The hourly quantity that the customer is committing to a CCD offer's main SKU.
	// For example, if the quantity is 1000, then you are committing to a 1000 usage quantity value per hour.
	Quantity *int64 `mandatory:"true" json:"quantity"`

	// The CCD commitment name. Ensure the name is unique for a given subscription ID. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The CCD offer OCID.
	OfferId *string `mandatory:"true" json:"offerId"`

	// The commercial workflow of the associated CCD offer.
	SourceType CcdOfferSourceTypeEnum `mandatory:"true" json:"sourceType"`

	// The date and time the CCD commitment will be effective.The format is defined by RFC 3339.
	TimeCcdCommitmentStarts *common.SDKTime `mandatory:"true" json:"timeCcdCommitmentStarts"`

	// The date and time the CCD commitment will be deactivated. The date and time is calculated based on `timeCcdCommitmentStarts`, the term, and units
	// of the CCD offer. For example: `2025-06-01T03:00:00`. The format is defined by RFC 3339.
	TimeCcdCommitmentEnds *common.SDKTime `mandatory:"true" json:"timeCcdCommitmentEnds"`

	// The subscription ID the CCD offer is applicable to.
	SubscriptionId *string `mandatory:"true" json:"subscriptionId"`

	// The root compartment ID. The CCD Commitment is associated with the tenancy root compartment ID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The reclaim term copied from the offer when the commitment is created. The value is measured in
	// `reclaimTermUnit`.
	ReclaimTerm *int64 `mandatory:"false" json:"reclaimTerm"`

	// The unit used to measure `reclaimTerm`, copied from the offer when the commitment is created.
	ReclaimTermUnit CcdOfferReclaimTermUnitEnum `mandatory:"false" json:"reclaimTermUnit,omitempty"`

	// The tenancy OCID that you want to apply filler usage to.
	FillerTenancyId *string `mandatory:"false" json:"fillerTenancyId"`

	// The filler tenancy home region.
	FillerTenancyHomeRegion *string `mandatory:"false" json:"fillerTenancyHomeRegion"`

	UsageEligibility *UsageEligibility `mandatory:"false" json:"usageEligibility"`

	// Secondary SKUs associated with the main SKU.
	SecondarySkus []SecondarySkus `mandatory:"false" json:"secondarySkus"`

	// This state always defaults to `Active`. For a detailed and exact state,
	// check the `lifecycleDetails` state.
	LifecycleState CcdCommitmentLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The CCD commitment's current state. A commitment that's expected to get a dedicated capacity will be in the
	// Awaiting Provisioning state unless the commitment is ended or canceled. Otherwise, the value is determined by
	// `timeCcdCommitmentStarts`, `timeCcdCommitmentEnds`, and `timeTerminated`, with respect to the current time.
	LifecycleDetails CcdCommitmentLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// The list of SKUs associated with the commitment and their prices in all currency codes.
	SkuQuantityPriceModel []CommittedSkuQuantityPriceModel `mandatory:"false" json:"skuQuantityPriceModel"`

	// Date and time the CCD was created. The format is defined by RFC 3339.
	// Example: `2022-01-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Date and time the CCD was updated lastly. The format is defined by RFC 3339.
	// Example: `2025-01-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Date and time the CCD was terminated. The format is defined by RFC 3339.
	// Example: `2022-01-25T21:10:29.600Z`
	TimeTerminated *common.SDKTime `mandatory:"false" json:"timeTerminated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Entity/User who lastly updated the ccdCommitment.
	LastUpdatedBy *string `mandatory:"false" json:"lastUpdatedBy"`

	// The entity that created the CCD commitment.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// Entity/User who approved the ccdCommitment.
	ApprovedBy *string `mandatory:"false" json:"approvedBy"`
}

func (m CcdCommitment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CcdCommitment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCcdOfferSourceTypeEnum(string(m.SourceType)); !ok && m.SourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceType: %s. Supported values are: %s.", m.SourceType, strings.Join(GetCcdOfferSourceTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCcdOfferReclaimTermUnitEnum(string(m.ReclaimTermUnit)); !ok && m.ReclaimTermUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReclaimTermUnit: %s. Supported values are: %s.", m.ReclaimTermUnit, strings.Join(GetCcdOfferReclaimTermUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdCommitmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCcdCommitmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdCommitmentLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetCcdCommitmentLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CcdCommitmentLifecycleStateEnum Enum with underlying type: string
type CcdCommitmentLifecycleStateEnum string

// Set of constants representing the allowable values for CcdCommitmentLifecycleStateEnum
const (
	CcdCommitmentLifecycleStateActive CcdCommitmentLifecycleStateEnum = "ACTIVE"
)

var mappingCcdCommitmentLifecycleStateEnum = map[string]CcdCommitmentLifecycleStateEnum{
	"ACTIVE": CcdCommitmentLifecycleStateActive,
}

var mappingCcdCommitmentLifecycleStateEnumLowerCase = map[string]CcdCommitmentLifecycleStateEnum{
	"active": CcdCommitmentLifecycleStateActive,
}

// GetCcdCommitmentLifecycleStateEnumValues Enumerates the set of values for CcdCommitmentLifecycleStateEnum
func GetCcdCommitmentLifecycleStateEnumValues() []CcdCommitmentLifecycleStateEnum {
	values := make([]CcdCommitmentLifecycleStateEnum, 0)
	for _, v := range mappingCcdCommitmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCcdCommitmentLifecycleStateEnumStringValues Enumerates the set of values in String for CcdCommitmentLifecycleStateEnum
func GetCcdCommitmentLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
	}
}

// GetMappingCcdCommitmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCcdCommitmentLifecycleStateEnum(val string) (CcdCommitmentLifecycleStateEnum, bool) {
	enum, ok := mappingCcdCommitmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CcdCommitmentLifecycleDetailsEnum Enum with underlying type: string
type CcdCommitmentLifecycleDetailsEnum string

// Set of constants representing the allowable values for CcdCommitmentLifecycleDetailsEnum
const (
	CcdCommitmentLifecycleDetailsCurrent              CcdCommitmentLifecycleDetailsEnum = "CURRENT"
	CcdCommitmentLifecycleDetailsEnded                CcdCommitmentLifecycleDetailsEnum = "ENDED"
	CcdCommitmentLifecycleDetailsCanceled             CcdCommitmentLifecycleDetailsEnum = "CANCELED"
	CcdCommitmentLifecycleDetailsScheduled            CcdCommitmentLifecycleDetailsEnum = "SCHEDULED"
	CcdCommitmentLifecycleDetailsAwaitingProvisioning CcdCommitmentLifecycleDetailsEnum = "AWAITING_PROVISIONING"
)

var mappingCcdCommitmentLifecycleDetailsEnum = map[string]CcdCommitmentLifecycleDetailsEnum{
	"CURRENT":               CcdCommitmentLifecycleDetailsCurrent,
	"ENDED":                 CcdCommitmentLifecycleDetailsEnded,
	"CANCELED":              CcdCommitmentLifecycleDetailsCanceled,
	"SCHEDULED":             CcdCommitmentLifecycleDetailsScheduled,
	"AWAITING_PROVISIONING": CcdCommitmentLifecycleDetailsAwaitingProvisioning,
}

var mappingCcdCommitmentLifecycleDetailsEnumLowerCase = map[string]CcdCommitmentLifecycleDetailsEnum{
	"current":               CcdCommitmentLifecycleDetailsCurrent,
	"ended":                 CcdCommitmentLifecycleDetailsEnded,
	"canceled":              CcdCommitmentLifecycleDetailsCanceled,
	"scheduled":             CcdCommitmentLifecycleDetailsScheduled,
	"awaiting_provisioning": CcdCommitmentLifecycleDetailsAwaitingProvisioning,
}

// GetCcdCommitmentLifecycleDetailsEnumValues Enumerates the set of values for CcdCommitmentLifecycleDetailsEnum
func GetCcdCommitmentLifecycleDetailsEnumValues() []CcdCommitmentLifecycleDetailsEnum {
	values := make([]CcdCommitmentLifecycleDetailsEnum, 0)
	for _, v := range mappingCcdCommitmentLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetCcdCommitmentLifecycleDetailsEnumStringValues Enumerates the set of values in String for CcdCommitmentLifecycleDetailsEnum
func GetCcdCommitmentLifecycleDetailsEnumStringValues() []string {
	return []string{
		"CURRENT",
		"ENDED",
		"CANCELED",
		"SCHEDULED",
		"AWAITING_PROVISIONING",
	}
}

// GetMappingCcdCommitmentLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCcdCommitmentLifecycleDetailsEnum(val string) (CcdCommitmentLifecycleDetailsEnum, bool) {
	enum, ok := mappingCcdCommitmentLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
