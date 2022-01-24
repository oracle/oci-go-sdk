// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription API Usage Computation
//
// OneSubscription API Common set of Subscription Plan Management (SPM) Usage Computation resources
//

package osubusage

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ComputedUsageAggregatedSummary Subscribed Service Contract details
type ComputedUsageAggregatedSummary struct {

	// Subscription Id is an identifier associated to the service used for filter the Computed Usage in SPM
	SubscriptionId *string `mandatory:"true" json:"subscriptionId"`

	// Subscribed service line parent id
	ParentSubscribedServiceId *string `mandatory:"false" json:"parentSubscribedServiceId"`

	ParentProduct *Product `mandatory:"false" json:"parentProduct"`

	// Subscribed services contract line start date, expressed in RFC 3339 timestamp format.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Subscribed services contract line end date, expressed in RFC 3339 timestamp format.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// Subscribed service asociated subscription plan number.
	PlanNumber *string `mandatory:"false" json:"planNumber"`

	// Currency code
	CurrencyCode *string `mandatory:"false" json:"currencyCode"`

	// Inernal SPM Ratecard Id at line level
	RateCardId *string `mandatory:"false" json:"rateCardId"`

	// Subscribed services pricing model
	PricingModel ComputedUsageAggregatedSummaryPricingModelEnum `mandatory:"false" json:"pricingModel,omitempty"`

	// Aggregation of computed usages for the subscribed service.
	AggregatedComputedUsages []ComputedUsageAggregation `mandatory:"false" json:"aggregatedComputedUsages"`
}

func (m ComputedUsageAggregatedSummary) String() string {
	return common.PointerString(m)
}

// ComputedUsageAggregatedSummaryPricingModelEnum Enum with underlying type: string
type ComputedUsageAggregatedSummaryPricingModelEnum string

// Set of constants representing the allowable values for ComputedUsageAggregatedSummaryPricingModelEnum
const (
	ComputedUsageAggregatedSummaryPricingModelPayAsYouGo       ComputedUsageAggregatedSummaryPricingModelEnum = "PAY_AS_YOU_GO"
	ComputedUsageAggregatedSummaryPricingModelMonthly          ComputedUsageAggregatedSummaryPricingModelEnum = "MONTHLY"
	ComputedUsageAggregatedSummaryPricingModelAnnual           ComputedUsageAggregatedSummaryPricingModelEnum = "ANNUAL"
	ComputedUsageAggregatedSummaryPricingModelPrepaid          ComputedUsageAggregatedSummaryPricingModelEnum = "PREPAID"
	ComputedUsageAggregatedSummaryPricingModelFundedAllocation ComputedUsageAggregatedSummaryPricingModelEnum = "FUNDED_ALLOCATION"
)

var mappingComputedUsageAggregatedSummaryPricingModel = map[string]ComputedUsageAggregatedSummaryPricingModelEnum{
	"PAY_AS_YOU_GO":     ComputedUsageAggregatedSummaryPricingModelPayAsYouGo,
	"MONTHLY":           ComputedUsageAggregatedSummaryPricingModelMonthly,
	"ANNUAL":            ComputedUsageAggregatedSummaryPricingModelAnnual,
	"PREPAID":           ComputedUsageAggregatedSummaryPricingModelPrepaid,
	"FUNDED_ALLOCATION": ComputedUsageAggregatedSummaryPricingModelFundedAllocation,
}

// GetComputedUsageAggregatedSummaryPricingModelEnumValues Enumerates the set of values for ComputedUsageAggregatedSummaryPricingModelEnum
func GetComputedUsageAggregatedSummaryPricingModelEnumValues() []ComputedUsageAggregatedSummaryPricingModelEnum {
	values := make([]ComputedUsageAggregatedSummaryPricingModelEnum, 0)
	for _, v := range mappingComputedUsageAggregatedSummaryPricingModel {
		values = append(values, v)
	}
	return values
}
