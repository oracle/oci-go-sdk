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

// CcdCommitmentSummary The CCD commitment summary.
type CcdCommitmentSummary struct {

	// The CCD commitment resource OCID.
	Id *string `mandatory:"true" json:"id"`

	// The hourly quantity that the customer is committing to a CCD offer's main SKU.
	// For example, if the quantity is 1000, then you are committing to a 1000 usage quantity value per hour.
	Quantity *int64 `mandatory:"true" json:"quantity"`

	// The CCD offer OCID.
	OfferId *string `mandatory:"true" json:"offerId"`

	// The date and time, rounded down to an hour, that the CCD commitment will be effective.
	// The format conforms to RFC 3339.
	TimeCcdCommitmentStarts *common.SDKTime `mandatory:"true" json:"timeCcdCommitmentStarts"`

	// The date and time the CCD commitment ends, which is calculated from the `timeCcdCommitmentStarts` of the `ccdCommitment`,
	// the term, and `ccdOffer` units. The format must conform to RFC 3339.
	TimeCcdCommitmentEnds *common.SDKTime `mandatory:"true" json:"timeCcdCommitmentEnds"`

	// The subscription OCID the CCD commitment is applicable to.
	SubscriptionId *string `mandatory:"true" json:"subscriptionId"`

	// The root compartment ID associated with the CCD offer the customer has signed up for.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The CCD commitment's current state. A commitment that's expected to get a dedicated capacity will be in the
	// Awaiting Provisioning state unless the commitment is ended or canceled. Otherwise, the value is determined by
	// `timeCcdCommitmentStarts`, `timeCcdCommitmentEnds`, and `timeTerminated`, with respect to the current time.
	LifecycleDetails CcdCommitmentLifecycleDetailsEnum `mandatory:"true" json:"lifecycleDetails"`

	// The commercial workflow of the associated CCD offer.
	SourceType CcdOfferSourceTypeEnum `mandatory:"false" json:"sourceType,omitempty"`

	// The reclaim term copied from the offer when the commitment is created. The value is measured in `reclaimTermUnit`.
	ReclaimTerm *int64 `mandatory:"false" json:"reclaimTerm"`

	// The unit used to measure `reclaimTerm`, copied from the offer when the commitment is created.
	ReclaimTermUnit CcdOfferReclaimTermUnitEnum `mandatory:"false" json:"reclaimTermUnit,omitempty"`

	// The CCD commitment name. The name will be unique across the subscription for which the CCD commitment has been created. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The tenancy OCID that you want to apply filler usage to.
	FillerTenancyId *string `mandatory:"false" json:"fillerTenancyId"`

	// The filler tenancy home region.
	FillerTenancyHomeRegion *string `mandatory:"false" json:"fillerTenancyHomeRegion"`

	UsageEligibility *UsageEligibility `mandatory:"false" json:"usageEligibility"`

	// The CCD commitment's current state, which always defaults to `Active`. For a more accurate detailed state,
	// refer to `lifecycleDetails`.
	LifecycleState CcdCommitmentLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The list of SKUs associated with the commitment and their prices in all currency codes.
	SkuQuantityPriceModel []CommittedSkuQuantityPriceModel `mandatory:"false" json:"skuQuantityPriceModel"`

	// The secondary SKUs associated with the main SKU.
	SecondarySkus []SecondarySkus `mandatory:"false" json:"secondarySkus"`

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

func (m CcdCommitmentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CcdCommitmentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCcdCommitmentLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetCcdCommitmentLifecycleDetailsEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCcdOfferSourceTypeEnum(string(m.SourceType)); !ok && m.SourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceType: %s. Supported values are: %s.", m.SourceType, strings.Join(GetCcdOfferSourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdOfferReclaimTermUnitEnum(string(m.ReclaimTermUnit)); !ok && m.ReclaimTermUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReclaimTermUnit: %s. Supported values are: %s.", m.ReclaimTermUnit, strings.Join(GetCcdOfferReclaimTermUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdCommitmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCcdCommitmentLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
