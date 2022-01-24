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

// ComputedUsageSummary Computed Usage Summary object
type ComputedUsageSummary struct {

	// SPM Internal computed usage Id , 32 character string
	ComputedUsageId *string `mandatory:"true" json:"computedUsageId"`

	// Computed Usage created time, expressed in RFC 3339 timestamp format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Computed Usage updated time, expressed in RFC 3339 timestamp format.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Subscribed service line parent id
	ParentSubscribedServiceId *string `mandatory:"false" json:"parentSubscribedServiceId"`

	ParentProduct *Product `mandatory:"false" json:"parentProduct"`

	// Subscription plan number
	PlanNumber *string `mandatory:"false" json:"planNumber"`

	// Currency code
	CurrencyCode *string `mandatory:"false" json:"currencyCode"`

	// References the tier in the ratecard for that usage (OCI will be using the same reference to cross-reference for correctness on the usage csv report), comes from Entity OBSCNTR_IPT_PRODUCTTIER.
	RateCardTierdId *string `mandatory:"false" json:"rateCardTierdId"`

	// Ratecard Id at subscribed service level
	RateCardId *string `mandatory:"false" json:"rateCardId"`

	// SPM Internal compute records source .
	ComputeSource *string `mandatory:"false" json:"computeSource"`

	// Data Center Attribute as sent by MQS to SPM.
	DataCenter *string `mandatory:"false" json:"dataCenter"`

	// MQS Identfier send to SPM , SPM does not transform this attribute and is received as is.
	MqsMessageId *string `mandatory:"false" json:"mqsMessageId"`

	// Total Quantity that was used for computation
	Quantity *string `mandatory:"false" json:"quantity"`

	// SPM Internal usage Line number identifier in SPM coming from Metered Services entity.
	UsageNumber *string `mandatory:"false" json:"usageNumber"`

	// SPM Internal Original usage Line number identifier in SPM coming from Metered Services entity.
	OriginalUsageNumber *string `mandatory:"false" json:"originalUsageNumber"`

	// Subscribed service commitmentId.
	CommitmentServiceId *string `mandatory:"false" json:"commitmentServiceId"`

	// Invoicing status for the aggregated compute usage
	IsInvoiced *bool `mandatory:"false" json:"isInvoiced"`

	// Usage compute type in SPM.
	Type ComputedUsageSummaryTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Usae computation date, expressed in RFC 3339 timestamp format.
	TimeOfArrival *common.SDKTime `mandatory:"false" json:"timeOfArrival"`

	// Metered Service date, expressed in RFC 3339 timestamp format.
	TimeMeteredOn *common.SDKTime `mandatory:"false" json:"timeMeteredOn"`

	// Net Unit Price for the product in consideration, price actual.
	NetUnitPrice *string `mandatory:"false" json:"netUnitPrice"`

	// Computed Line Amount rounded.
	CostRounded *string `mandatory:"false" json:"costRounded"`

	// Computed Line Amount not rounded
	Cost *string `mandatory:"false" json:"cost"`

	Product *Product `mandatory:"false" json:"product"`

	// Unit of Messure
	UnitOfMeasure *string `mandatory:"false" json:"unitOfMeasure"`
}

func (m ComputedUsageSummary) String() string {
	return common.PointerString(m)
}

// ComputedUsageSummaryTypeEnum Enum with underlying type: string
type ComputedUsageSummaryTypeEnum string

// Set of constants representing the allowable values for ComputedUsageSummaryTypeEnum
const (
	ComputedUsageSummaryTypePromotion                     ComputedUsageSummaryTypeEnum = "PROMOTION"
	ComputedUsageSummaryTypeDoNotBill                     ComputedUsageSummaryTypeEnum = "DO_NOT_BILL"
	ComputedUsageSummaryTypeUsage                         ComputedUsageSummaryTypeEnum = "USAGE"
	ComputedUsageSummaryTypeCommit                        ComputedUsageSummaryTypeEnum = "COMMIT"
	ComputedUsageSummaryTypeOverage                       ComputedUsageSummaryTypeEnum = "OVERAGE"
	ComputedUsageSummaryTypePayAsYouGo                    ComputedUsageSummaryTypeEnum = "PAY_AS_YOU_GO"
	ComputedUsageSummaryTypeMonthlyMinimum                ComputedUsageSummaryTypeEnum = "MONTHLY_MINIMUM"
	ComputedUsageSummaryTypeDelayedUsageInvoiceTiming     ComputedUsageSummaryTypeEnum = "DELAYED_USAGE_INVOICE_TIMING"
	ComputedUsageSummaryTypeDelayedUsageCommitmentExp     ComputedUsageSummaryTypeEnum = "DELAYED_USAGE_COMMITMENT_EXP"
	ComputedUsageSummaryTypeOnAccountCredit               ComputedUsageSummaryTypeEnum = "ON_ACCOUNT_CREDIT"
	ComputedUsageSummaryTypeServiceCredit                 ComputedUsageSummaryTypeEnum = "SERVICE_CREDIT"
	ComputedUsageSummaryTypeCommitmentExpiration          ComputedUsageSummaryTypeEnum = "COMMITMENT_EXPIRATION"
	ComputedUsageSummaryTypeFundedAllocation              ComputedUsageSummaryTypeEnum = "FUNDED_ALLOCATION"
	ComputedUsageSummaryTypeDonotBillUsagePostTermination ComputedUsageSummaryTypeEnum = "DONOT_BILL_USAGE_POST_TERMINATION"
	ComputedUsageSummaryTypeDelayedUsagePostTermination   ComputedUsageSummaryTypeEnum = "DELAYED_USAGE_POST_TERMINATION"
)

var mappingComputedUsageSummaryType = map[string]ComputedUsageSummaryTypeEnum{
	"PROMOTION":                         ComputedUsageSummaryTypePromotion,
	"DO_NOT_BILL":                       ComputedUsageSummaryTypeDoNotBill,
	"USAGE":                             ComputedUsageSummaryTypeUsage,
	"COMMIT":                            ComputedUsageSummaryTypeCommit,
	"OVERAGE":                           ComputedUsageSummaryTypeOverage,
	"PAY_AS_YOU_GO":                     ComputedUsageSummaryTypePayAsYouGo,
	"MONTHLY_MINIMUM":                   ComputedUsageSummaryTypeMonthlyMinimum,
	"DELAYED_USAGE_INVOICE_TIMING":      ComputedUsageSummaryTypeDelayedUsageInvoiceTiming,
	"DELAYED_USAGE_COMMITMENT_EXP":      ComputedUsageSummaryTypeDelayedUsageCommitmentExp,
	"ON_ACCOUNT_CREDIT":                 ComputedUsageSummaryTypeOnAccountCredit,
	"SERVICE_CREDIT":                    ComputedUsageSummaryTypeServiceCredit,
	"COMMITMENT_EXPIRATION":             ComputedUsageSummaryTypeCommitmentExpiration,
	"FUNDED_ALLOCATION":                 ComputedUsageSummaryTypeFundedAllocation,
	"DONOT_BILL_USAGE_POST_TERMINATION": ComputedUsageSummaryTypeDonotBillUsagePostTermination,
	"DELAYED_USAGE_POST_TERMINATION":    ComputedUsageSummaryTypeDelayedUsagePostTermination,
}

// GetComputedUsageSummaryTypeEnumValues Enumerates the set of values for ComputedUsageSummaryTypeEnum
func GetComputedUsageSummaryTypeEnumValues() []ComputedUsageSummaryTypeEnum {
	values := make([]ComputedUsageSummaryTypeEnum, 0)
	for _, v := range mappingComputedUsageSummaryType {
		values = append(values, v)
	}
	return values
}
