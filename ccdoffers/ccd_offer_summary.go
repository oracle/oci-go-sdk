// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Continuous Consumption Discounts API
//
// Use the Continuous Consumption Discounts API to manage continuous consumption discount operations.
//

package ccdoffers

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CcdOfferSummary CCD offer details.
type CcdOfferSummary struct {

	// The CCD offer resource OCID.
	Id *string `mandatory:"true" json:"id"`

	// The CCD offer name. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The SKU part number (prefaced with "B") the CCD is associated with.
	Sku *string `mandatory:"true" json:"sku"`

	// The canonical SKU display name.
	SkuDisplayName *string `mandatory:"true" json:"skuDisplayName"`

	// The canonical billing metric.
	SkuMetric *string `mandatory:"true" json:"skuMetric"`

	// The minimum quantity of the primary SKU committed to the CCD offer.
	// For FLEX offers, minQuantity and maxQuantity must both be equal if maxQuantity provided.
	MinQuantity *int64 `mandatory:"true" json:"minQuantity"`

	// The number of terms the offer is applicable to.
	Term *int `mandatory:"true" json:"term"`

	// The discount percentage applicable to the CCD offer, represented as a decimal string
	// from 0 through 100 inclusive.
	DiscountPct *string `mandatory:"true" json:"discountPct"`

	// The subscription OCID for which the dedicated CCD offer has been created.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// The root compartment OCID of the customer's tenancy.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The created CCD offer type.
	TypeOfOffer CcdOfferSummaryTypeOfOfferEnum `mandatory:"false" json:"typeOfOffer,omitempty"`

	// The commercial workflow that created the CCD offer.
	SourceType CcdOfferSourceTypeEnum `mandatory:"false" json:"sourceType,omitempty"`

	// The reclaim term for an offer, measured in `reclaimTermUnit`.
	ReclaimTerm *int64 `mandatory:"false" json:"reclaimTerm"`

	// The unit used to measure `reclaimTerm`.
	ReclaimTermUnit CcdOfferReclaimTermUnitEnum `mandatory:"false" json:"reclaimTermUnit,omitempty"`

	// The targeted customer segment.
	CustomerSegment CcdOfferSummaryCustomerSegmentEnum `mandatory:"false" json:"customerSegment,omitempty"`

	// Enable or disable filler usage if the customer usage quantity is below the minimum quantity, and you want metering to fill
	// in usages to meet the minimum quantity.
	IsFillerUsageEnable *bool `mandatory:"false" json:"isFillerUsageEnable"`

	// Entity/User who lastly updated the ccdOffer.
	LastUpdatedBy *string `mandatory:"false" json:"lastUpdatedBy"`

	// Entity/User who created the ccdOffer.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// Sales order document number.
	OrderDocumentNumber *string `mandatory:"false" json:"orderDocumentNumber"`

	// Defines where the filler usage is reported
	ReportingRegion *string `mandatory:"false" json:"reportingRegion"`

	// The maximum quantity of the primary SKU committed to the CCD offer.
	// For FLEX offers if provided, it must equal minQuantity.
	MaxQuantity *int64 `mandatory:"false" json:"maxQuantity"`

	// The term units.
	Unit CcdOfferSummaryUnitEnum `mandatory:"false" json:"unit,omitempty"`

	OfferRestriction *OfferRestriction `mandatory:"false" json:"offerRestriction"`

	// The polymorphic bundled SKU definitions. Every bundled SKU in an offer must use the same `quantityRule`; don't mix
	// `FIXED_RATIO` and `CUSTOMER_SELECTED` bundled SKUs in one offer. An offer can contain at most five SKUs in total,
	// including the primary SKU.
	BundleSkus []BundleSku `mandatory:"false" json:"bundleSkus"`

	// The SKU list price in all currency codes.
	ExternalPriceOfSku []SkuExternalPriceModel `mandatory:"false" json:"externalPriceOfSku"`

	// The CCD offer's current state. This always defaults to `Active`. For the accurate resource state,
	// refer to the `lifecycleDetails` property value.
	LifecycleState CcdOfferLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The CCD offer's detailed current state. The state value is determined by `timeCcdOfferStarts`,
	// `timeCcdOfferEnds`, and `timeTerminated` with respect to the current time.
	LifecycleDetails CcdOfferLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

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

	// The inclusive date the CCD offer will be effective. The format is defined by RFC 3339.
	TimeCcdOfferStarts *common.SDKTime `mandatory:"false" json:"timeCcdOfferStarts"`

	// The inclusive time when the CCD offer is no longer in effect. The format is defined by RFC 3339.
	TimeCcdOfferEnds *common.SDKTime `mandatory:"false" json:"timeCcdOfferEnds"`

	// Currency code of listPrice associated with primary product/sku. This will be always null.
	// Please refer `externalPriceOfSku` to know product/sku's associated list price in its currency code.
	CurrencyCode *string `mandatory:"false" json:"currencyCode"`

	// ListPrice of a sku.
	// Please refer `externalPriceOfSku` to know product/sku's associated list price in its currency code.
	ListPrice *string `mandatory:"false" json:"listPrice"`
}

func (m CcdOfferSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CcdOfferSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCcdOfferSummaryTypeOfOfferEnum(string(m.TypeOfOffer)); !ok && m.TypeOfOffer != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TypeOfOffer: %s. Supported values are: %s.", m.TypeOfOffer, strings.Join(GetCcdOfferSummaryTypeOfOfferEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdOfferSourceTypeEnum(string(m.SourceType)); !ok && m.SourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceType: %s. Supported values are: %s.", m.SourceType, strings.Join(GetCcdOfferSourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdOfferReclaimTermUnitEnum(string(m.ReclaimTermUnit)); !ok && m.ReclaimTermUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReclaimTermUnit: %s. Supported values are: %s.", m.ReclaimTermUnit, strings.Join(GetCcdOfferReclaimTermUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdOfferSummaryCustomerSegmentEnum(string(m.CustomerSegment)); !ok && m.CustomerSegment != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CustomerSegment: %s. Supported values are: %s.", m.CustomerSegment, strings.Join(GetCcdOfferSummaryCustomerSegmentEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdOfferSummaryUnitEnum(string(m.Unit)); !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetCcdOfferSummaryUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdOfferLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCcdOfferLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdOfferLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetCcdOfferLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CcdOfferSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SubscriptionId      *string                            `json:"subscriptionId"`
		CompartmentId       *string                            `json:"compartmentId"`
		TypeOfOffer         CcdOfferSummaryTypeOfOfferEnum     `json:"typeOfOffer"`
		SourceType          CcdOfferSourceTypeEnum             `json:"sourceType"`
		ReclaimTerm         *int64                             `json:"reclaimTerm"`
		ReclaimTermUnit     CcdOfferReclaimTermUnitEnum        `json:"reclaimTermUnit"`
		CustomerSegment     CcdOfferSummaryCustomerSegmentEnum `json:"customerSegment"`
		IsFillerUsageEnable *bool                              `json:"isFillerUsageEnable"`
		LastUpdatedBy       *string                            `json:"lastUpdatedBy"`
		CreatedBy           *string                            `json:"createdBy"`
		OrderDocumentNumber *string                            `json:"orderDocumentNumber"`
		ReportingRegion     *string                            `json:"reportingRegion"`
		MaxQuantity         *int64                             `json:"maxQuantity"`
		Unit                CcdOfferSummaryUnitEnum            `json:"unit"`
		OfferRestriction    *OfferRestriction                  `json:"offerRestriction"`
		BundleSkus          []bundlesku                        `json:"bundleSkus"`
		ExternalPriceOfSku  []SkuExternalPriceModel            `json:"externalPriceOfSku"`
		LifecycleState      CcdOfferLifecycleStateEnum         `json:"lifecycleState"`
		LifecycleDetails    CcdOfferLifecycleDetailsEnum       `json:"lifecycleDetails"`
		TimeCreated         *common.SDKTime                    `json:"timeCreated"`
		TimeUpdated         *common.SDKTime                    `json:"timeUpdated"`
		TimeTerminated      *common.SDKTime                    `json:"timeTerminated"`
		FreeformTags        map[string]string                  `json:"freeformTags"`
		DefinedTags         map[string]map[string]interface{}  `json:"definedTags"`
		SystemTags          map[string]map[string]interface{}  `json:"systemTags"`
		TimeCcdOfferStarts  *common.SDKTime                    `json:"timeCcdOfferStarts"`
		TimeCcdOfferEnds    *common.SDKTime                    `json:"timeCcdOfferEnds"`
		CurrencyCode        *string                            `json:"currencyCode"`
		ListPrice           *string                            `json:"listPrice"`
		Id                  *string                            `json:"id"`
		DisplayName         *string                            `json:"displayName"`
		Sku                 *string                            `json:"sku"`
		SkuDisplayName      *string                            `json:"skuDisplayName"`
		SkuMetric           *string                            `json:"skuMetric"`
		MinQuantity         *int64                             `json:"minQuantity"`
		Term                *int                               `json:"term"`
		DiscountPct         *string                            `json:"discountPct"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.SubscriptionId = model.SubscriptionId

	m.CompartmentId = model.CompartmentId

	m.TypeOfOffer = model.TypeOfOffer

	m.SourceType = model.SourceType

	m.ReclaimTerm = model.ReclaimTerm

	m.ReclaimTermUnit = model.ReclaimTermUnit

	m.CustomerSegment = model.CustomerSegment

	m.IsFillerUsageEnable = model.IsFillerUsageEnable

	m.LastUpdatedBy = model.LastUpdatedBy

	m.CreatedBy = model.CreatedBy

	m.OrderDocumentNumber = model.OrderDocumentNumber

	m.ReportingRegion = model.ReportingRegion

	m.MaxQuantity = model.MaxQuantity

	m.Unit = model.Unit

	m.OfferRestriction = model.OfferRestriction

	m.BundleSkus = make([]BundleSku, len(model.BundleSkus))
	for i, n := range model.BundleSkus {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.BundleSkus[i] = nn.(BundleSku)
		} else {
			m.BundleSkus[i] = nil
		}
	}
	m.ExternalPriceOfSku = make([]SkuExternalPriceModel, len(model.ExternalPriceOfSku))
	copy(m.ExternalPriceOfSku, model.ExternalPriceOfSku)
	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.TimeTerminated = model.TimeTerminated

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.TimeCcdOfferStarts = model.TimeCcdOfferStarts

	m.TimeCcdOfferEnds = model.TimeCcdOfferEnds

	m.CurrencyCode = model.CurrencyCode

	m.ListPrice = model.ListPrice

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.Sku = model.Sku

	m.SkuDisplayName = model.SkuDisplayName

	m.SkuMetric = model.SkuMetric

	m.MinQuantity = model.MinQuantity

	m.Term = model.Term

	m.DiscountPct = model.DiscountPct

	return
}

// CcdOfferSummaryTypeOfOfferEnum Enum with underlying type: string
type CcdOfferSummaryTypeOfOfferEnum string

// Set of constants representing the allowable values for CcdOfferSummaryTypeOfOfferEnum
const (
	CcdOfferSummaryTypeOfOfferStandard    CcdOfferSummaryTypeOfOfferEnum = "STANDARD"
	CcdOfferSummaryTypeOfOfferCustom      CcdOfferSummaryTypeOfOfferEnum = "CUSTOM"
	CcdOfferSummaryTypeOfOfferPrivate     CcdOfferSummaryTypeOfOfferEnum = "PRIVATE"
	CcdOfferSummaryTypeOfOfferContractual CcdOfferSummaryTypeOfOfferEnum = "CONTRACTUAL"
	CcdOfferSummaryTypeOfOfferFlex        CcdOfferSummaryTypeOfOfferEnum = "FLEX"
)

var mappingCcdOfferSummaryTypeOfOfferEnum = map[string]CcdOfferSummaryTypeOfOfferEnum{
	"STANDARD":    CcdOfferSummaryTypeOfOfferStandard,
	"CUSTOM":      CcdOfferSummaryTypeOfOfferCustom,
	"PRIVATE":     CcdOfferSummaryTypeOfOfferPrivate,
	"CONTRACTUAL": CcdOfferSummaryTypeOfOfferContractual,
	"FLEX":        CcdOfferSummaryTypeOfOfferFlex,
}

var mappingCcdOfferSummaryTypeOfOfferEnumLowerCase = map[string]CcdOfferSummaryTypeOfOfferEnum{
	"standard":    CcdOfferSummaryTypeOfOfferStandard,
	"custom":      CcdOfferSummaryTypeOfOfferCustom,
	"private":     CcdOfferSummaryTypeOfOfferPrivate,
	"contractual": CcdOfferSummaryTypeOfOfferContractual,
	"flex":        CcdOfferSummaryTypeOfOfferFlex,
}

// GetCcdOfferSummaryTypeOfOfferEnumValues Enumerates the set of values for CcdOfferSummaryTypeOfOfferEnum
func GetCcdOfferSummaryTypeOfOfferEnumValues() []CcdOfferSummaryTypeOfOfferEnum {
	values := make([]CcdOfferSummaryTypeOfOfferEnum, 0)
	for _, v := range mappingCcdOfferSummaryTypeOfOfferEnum {
		values = append(values, v)
	}
	return values
}

// GetCcdOfferSummaryTypeOfOfferEnumStringValues Enumerates the set of values in String for CcdOfferSummaryTypeOfOfferEnum
func GetCcdOfferSummaryTypeOfOfferEnumStringValues() []string {
	return []string{
		"STANDARD",
		"CUSTOM",
		"PRIVATE",
		"CONTRACTUAL",
		"FLEX",
	}
}

// GetMappingCcdOfferSummaryTypeOfOfferEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCcdOfferSummaryTypeOfOfferEnum(val string) (CcdOfferSummaryTypeOfOfferEnum, bool) {
	enum, ok := mappingCcdOfferSummaryTypeOfOfferEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CcdOfferSummaryCustomerSegmentEnum Enum with underlying type: string
type CcdOfferSummaryCustomerSegmentEnum string

// Set of constants representing the allowable values for CcdOfferSummaryCustomerSegmentEnum
const (
	CcdOfferSummaryCustomerSegmentInternal CcdOfferSummaryCustomerSegmentEnum = "INTERNAL"
	CcdOfferSummaryCustomerSegmentExternal CcdOfferSummaryCustomerSegmentEnum = "EXTERNAL"
)

var mappingCcdOfferSummaryCustomerSegmentEnum = map[string]CcdOfferSummaryCustomerSegmentEnum{
	"INTERNAL": CcdOfferSummaryCustomerSegmentInternal,
	"EXTERNAL": CcdOfferSummaryCustomerSegmentExternal,
}

var mappingCcdOfferSummaryCustomerSegmentEnumLowerCase = map[string]CcdOfferSummaryCustomerSegmentEnum{
	"internal": CcdOfferSummaryCustomerSegmentInternal,
	"external": CcdOfferSummaryCustomerSegmentExternal,
}

// GetCcdOfferSummaryCustomerSegmentEnumValues Enumerates the set of values for CcdOfferSummaryCustomerSegmentEnum
func GetCcdOfferSummaryCustomerSegmentEnumValues() []CcdOfferSummaryCustomerSegmentEnum {
	values := make([]CcdOfferSummaryCustomerSegmentEnum, 0)
	for _, v := range mappingCcdOfferSummaryCustomerSegmentEnum {
		values = append(values, v)
	}
	return values
}

// GetCcdOfferSummaryCustomerSegmentEnumStringValues Enumerates the set of values in String for CcdOfferSummaryCustomerSegmentEnum
func GetCcdOfferSummaryCustomerSegmentEnumStringValues() []string {
	return []string{
		"INTERNAL",
		"EXTERNAL",
	}
}

// GetMappingCcdOfferSummaryCustomerSegmentEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCcdOfferSummaryCustomerSegmentEnum(val string) (CcdOfferSummaryCustomerSegmentEnum, bool) {
	enum, ok := mappingCcdOfferSummaryCustomerSegmentEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CcdOfferSummaryUnitEnum Enum with underlying type: string
type CcdOfferSummaryUnitEnum string

// Set of constants representing the allowable values for CcdOfferSummaryUnitEnum
const (
	CcdOfferSummaryUnitMonths CcdOfferSummaryUnitEnum = "MONTHS"
	CcdOfferSummaryUnitYears  CcdOfferSummaryUnitEnum = "YEARS"
	CcdOfferSummaryUnitDays   CcdOfferSummaryUnitEnum = "DAYS"
	CcdOfferSummaryUnitHours  CcdOfferSummaryUnitEnum = "HOURS"
)

var mappingCcdOfferSummaryUnitEnum = map[string]CcdOfferSummaryUnitEnum{
	"MONTHS": CcdOfferSummaryUnitMonths,
	"YEARS":  CcdOfferSummaryUnitYears,
	"DAYS":   CcdOfferSummaryUnitDays,
	"HOURS":  CcdOfferSummaryUnitHours,
}

var mappingCcdOfferSummaryUnitEnumLowerCase = map[string]CcdOfferSummaryUnitEnum{
	"months": CcdOfferSummaryUnitMonths,
	"years":  CcdOfferSummaryUnitYears,
	"days":   CcdOfferSummaryUnitDays,
	"hours":  CcdOfferSummaryUnitHours,
}

// GetCcdOfferSummaryUnitEnumValues Enumerates the set of values for CcdOfferSummaryUnitEnum
func GetCcdOfferSummaryUnitEnumValues() []CcdOfferSummaryUnitEnum {
	values := make([]CcdOfferSummaryUnitEnum, 0)
	for _, v := range mappingCcdOfferSummaryUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetCcdOfferSummaryUnitEnumStringValues Enumerates the set of values in String for CcdOfferSummaryUnitEnum
func GetCcdOfferSummaryUnitEnumStringValues() []string {
	return []string{
		"MONTHS",
		"YEARS",
		"DAYS",
		"HOURS",
	}
}

// GetMappingCcdOfferSummaryUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCcdOfferSummaryUnitEnum(val string) (CcdOfferSummaryUnitEnum, bool) {
	enum, ok := mappingCcdOfferSummaryUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
