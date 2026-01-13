// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// DisbursementReportRecord The model for disbursement report item
type DisbursementReportRecord struct {

	// Enterprise Quarter
	EnterpriseQuarter *string `mandatory:"false" json:"enterpriseQuarter"`

	// Enterprise Period
	EnterprisePeriod *string `mandatory:"false" json:"enterprisePeriod"`

	// Listing Name
	ListingName *string `mandatory:"false" json:"listingName"`

	// Listing ID
	ListingId *string `mandatory:"false" json:"listingId"`

	// Private Offer ID
	PrivateOfferId *string `mandatory:"false" json:"privateOfferId"`

	// Private Offer Name
	PrivateOfferName *string `mandatory:"false" json:"privateOfferName"`

	// SKU
	Sku *string `mandatory:"false" json:"sku"`

	// Transaction Reference ID
	TransactionReferenceId *string `mandatory:"false" json:"transactionReferenceId"`

	// Customer ID
	CustomerId *string `mandatory:"false" json:"customerId"`

	// Customer Name
	CustomerName *string `mandatory:"false" json:"customerName"`

	// Customer Domain
	CustomerDomain *string `mandatory:"false" json:"customerDomain"`

	// End User Customer ID
	EndUserCustomerId *string `mandatory:"false" json:"endUserCustomerId"`

	// EEnd User Customer Name
	EndUserCustomerName *string `mandatory:"false" json:"endUserCustomerName"`

	// Country
	Country *string `mandatory:"false" json:"country"`

	// State
	State *string `mandatory:"false" json:"state"`

	// City
	City *string `mandatory:"false" json:"city"`

	// Zip
	Zip *string `mandatory:"false" json:"zip"`

	// Currency Code
	CurrencyCode *string `mandatory:"false" json:"currencyCode"`

	// Usage Quantity
	UsageQuantity *string `mandatory:"false" json:"usageQuantity"`

	// Used Amount (LC)
	UsedAmountLc *string `mandatory:"false" json:"usedAmountLc"`

	// LC Amount for Disbursement Calculations
	LcAmountForDisbursementCalculations *string `mandatory:"false" json:"lcAmountForDisbursementCalculations"`

	// LC Base Amount for Tax Calculations
	LcBaseAmountForTaxCalculations *string `mandatory:"false" json:"lcBaseAmountForTaxCalculations"`

	// Withholding Tax Rate
	WithholdingTaxRate *string `mandatory:"false" json:"withholdingTaxRate"`

	// Withholding Tax Amount
	WithholdingTaxAmount *string `mandatory:"false" json:"withholdingTaxAmount"`

	// Transfer Tax Base Amount
	TransferTaxBaseAmount *string `mandatory:"false" json:"transferTaxBaseAmount"`

	// Transfer Tax Rate
	TransferTaxRate *string `mandatory:"false" json:"transferTaxRate"`

	// Transfer Tax Amount
	TransferTaxAmount *string `mandatory:"false" json:"transferTaxAmount"`

	// LC Amount for Disbursement after Tax
	LcAmountForDisbursementAfterTax *string `mandatory:"false" json:"lcAmountForDisbursementAfterTax"`

	// LC Partner Payment
	LcPartnerPayment *string `mandatory:"false" json:"lcPartnerPayment"`

	// FX Rate (LC to USD)
	FxRateLcToUsd *string `mandatory:"false" json:"fxRateLcToUsd"`

	// Disbursement Amount (USD)
	DisbursementAmountUsd *string `mandatory:"false" json:"disbursementAmountUsd"`

	// Partner Name
	PartnerName *string `mandatory:"false" json:"partnerName"`

	// Partner OCID
	PartnerId *string `mandatory:"false" json:"partnerId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DisbursementReportRecord) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DisbursementReportRecord) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
