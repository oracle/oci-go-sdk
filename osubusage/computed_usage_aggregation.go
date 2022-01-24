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

// ComputedUsageAggregation Computed Usage Aggregation object
type ComputedUsageAggregation struct {

	// Total Quantity that was used for computation
	Quantity *string `mandatory:"false" json:"quantity"`

	Product *Product `mandatory:"false" json:"product"`

	// Data Center Attribute as sent by MQS to SPM.
	DataCenter *string `mandatory:"false" json:"dataCenter"`

	// Metered Service date , expressed in RFC 3339 timestamp format.
	TimeMeteredOn *common.SDKTime `mandatory:"false" json:"timeMeteredOn"`

	// Net Unit Price for the product in consideration.
	NetUnitPrice *string `mandatory:"false" json:"netUnitPrice"`

	// Sum of Computed Line Amount unrounded
	CostUnrounded *string `mandatory:"false" json:"costUnrounded"`

	// Sum of Computed Line Amount rounded
	Cost *string `mandatory:"false" json:"cost"`

	// Usage compute type in SPM.
	Type ComputedUsageAggregationTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m ComputedUsageAggregation) String() string {
	return common.PointerString(m)
}

// ComputedUsageAggregationTypeEnum Enum with underlying type: string
type ComputedUsageAggregationTypeEnum string

// Set of constants representing the allowable values for ComputedUsageAggregationTypeEnum
const (
	ComputedUsageAggregationTypePromotion                     ComputedUsageAggregationTypeEnum = "PROMOTION"
	ComputedUsageAggregationTypeDoNotBill                     ComputedUsageAggregationTypeEnum = "DO_NOT_BILL"
	ComputedUsageAggregationTypeUsage                         ComputedUsageAggregationTypeEnum = "USAGE"
	ComputedUsageAggregationTypeCommit                        ComputedUsageAggregationTypeEnum = "COMMIT"
	ComputedUsageAggregationTypeOverage                       ComputedUsageAggregationTypeEnum = "OVERAGE"
	ComputedUsageAggregationTypePayAsYouGo                    ComputedUsageAggregationTypeEnum = "PAY_AS_YOU_GO"
	ComputedUsageAggregationTypeMonthlyMinimum                ComputedUsageAggregationTypeEnum = "MONTHLY_MINIMUM"
	ComputedUsageAggregationTypeDelayedUsageInvoiceTiming     ComputedUsageAggregationTypeEnum = "DELAYED_USAGE_INVOICE_TIMING"
	ComputedUsageAggregationTypeDelayedUsageCommitmentExp     ComputedUsageAggregationTypeEnum = "DELAYED_USAGE_COMMITMENT_EXP"
	ComputedUsageAggregationTypeOnAccountCredit               ComputedUsageAggregationTypeEnum = "ON_ACCOUNT_CREDIT"
	ComputedUsageAggregationTypeServiceCredit                 ComputedUsageAggregationTypeEnum = "SERVICE_CREDIT"
	ComputedUsageAggregationTypeCommitmentExpiration          ComputedUsageAggregationTypeEnum = "COMMITMENT_EXPIRATION"
	ComputedUsageAggregationTypeFundedAllocation              ComputedUsageAggregationTypeEnum = "FUNDED_ALLOCATION"
	ComputedUsageAggregationTypeDonotBillUsagePostTermination ComputedUsageAggregationTypeEnum = "DONOT_BILL_USAGE_POST_TERMINATION"
	ComputedUsageAggregationTypeDelayedUsagePostTermination   ComputedUsageAggregationTypeEnum = "DELAYED_USAGE_POST_TERMINATION"
)

var mappingComputedUsageAggregationType = map[string]ComputedUsageAggregationTypeEnum{
	"PROMOTION":                         ComputedUsageAggregationTypePromotion,
	"DO_NOT_BILL":                       ComputedUsageAggregationTypeDoNotBill,
	"USAGE":                             ComputedUsageAggregationTypeUsage,
	"COMMIT":                            ComputedUsageAggregationTypeCommit,
	"OVERAGE":                           ComputedUsageAggregationTypeOverage,
	"PAY_AS_YOU_GO":                     ComputedUsageAggregationTypePayAsYouGo,
	"MONTHLY_MINIMUM":                   ComputedUsageAggregationTypeMonthlyMinimum,
	"DELAYED_USAGE_INVOICE_TIMING":      ComputedUsageAggregationTypeDelayedUsageInvoiceTiming,
	"DELAYED_USAGE_COMMITMENT_EXP":      ComputedUsageAggregationTypeDelayedUsageCommitmentExp,
	"ON_ACCOUNT_CREDIT":                 ComputedUsageAggregationTypeOnAccountCredit,
	"SERVICE_CREDIT":                    ComputedUsageAggregationTypeServiceCredit,
	"COMMITMENT_EXPIRATION":             ComputedUsageAggregationTypeCommitmentExpiration,
	"FUNDED_ALLOCATION":                 ComputedUsageAggregationTypeFundedAllocation,
	"DONOT_BILL_USAGE_POST_TERMINATION": ComputedUsageAggregationTypeDonotBillUsagePostTermination,
	"DELAYED_USAGE_POST_TERMINATION":    ComputedUsageAggregationTypeDelayedUsagePostTermination,
}

// GetComputedUsageAggregationTypeEnumValues Enumerates the set of values for ComputedUsageAggregationTypeEnum
func GetComputedUsageAggregationTypeEnumValues() []ComputedUsageAggregationTypeEnum {
	values := make([]ComputedUsageAggregationTypeEnum, 0)
	for _, v := range mappingComputedUsageAggregationType {
		values = append(values, v)
	}
	return values
}
