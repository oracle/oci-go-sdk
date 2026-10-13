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

// CcdOffer The created CCD offer details.
type CcdOffer struct {

	// The CCD offer OCID.
	Id *string `mandatory:"true" json:"id"`

	// The CCD offer name.
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

	// The CCD offer type.
	TypeOfOffer CcdOfferTypeOfOfferEnum `mandatory:"false" json:"typeOfOffer,omitempty"`

	// The commercial workflow that created the CCD offer/commitment. This value is immutable.
	SourceType CcdOfferSourceTypeEnum `mandatory:"false" json:"sourceType,omitempty"`

	// The reclaim term for an offer, measured in `reclaimTermUnit`.
	ReclaimTerm *int64 `mandatory:"false" json:"reclaimTerm"`

	// The unit used to measure `reclaimTerm`.
	ReclaimTermUnit CcdOfferReclaimTermUnitEnum `mandatory:"false" json:"reclaimTermUnit,omitempty"`

	// Entity/User who lastly updated the ccdOffer.
	LastUpdatedBy *string `mandatory:"false" json:"lastUpdatedBy"`

	// Entity/User who created the ccdOffer.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// Sales order document number.
	OrderDocumentNumber *string `mandatory:"false" json:"orderDocumentNumber"`

	// The targeted customer segment.
	CustomerSegment CcdOfferCustomerSegmentEnum `mandatory:"false" json:"customerSegment,omitempty"`

	// Enable or disable filler usage if the customer usage quantity is below the minimum quantity, and you want metering to fill
	// in usages to meet the minimum quantity.
	IsFillerUsageEnable *bool `mandatory:"false" json:"isFillerUsageEnable"`

	// Defines where the filler usage is reported
	ReportingRegion *string `mandatory:"false" json:"reportingRegion"`

	// The maximum quantity of the primary SKU committed to the CCD offer.
	// For FLEX offers if provided, it must equal minQuantity.
	MaxQuantity *int64 `mandatory:"false" json:"maxQuantity"`

	// The term units.
	Unit CcdOfferUnitEnum `mandatory:"false" json:"unit,omitempty"`

	// The SKU list price in all currency codes.
	ExternalPriceOfSku []SkuExternalPriceModel `mandatory:"false" json:"externalPriceOfSku"`

	OfferRestriction *OfferRestriction `mandatory:"false" json:"offerRestriction"`

	// The polymorphic bundled SKU definitions. Every bundled SKU in an offer must use the same `quantityRule`; don't mix
	// `FIXED_RATIO` and `CUSTOMER_SELECTED` bundled SKUs in one offer. An offer can contain at most five SKUs in total,
	// including the primary SKU.
	BundleSkus []BundleSku `mandatory:"false" json:"bundleSkus"`

	// Always defaults to `Active`. For the accurate resource state,
	// refer to the `lifecycleDetails` property value.
	LifecycleState CcdOfferLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The CCD offer's detailed current state. The state value is determined by `timeCcdOfferStarts`,
	// `timeCcdOfferEnds`, and `timeTerminated` with respect to the current time.
	LifecycleDetails CcdOfferLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// Date and time the CCD was created. The format is defined by RFC 3339.
	// Example: `2022-01-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Date and time the CCD was terminated. The format is defined by RFC 3339.
	// Example: `2022-01-25T21:10:29.600Z`
	TimeTerminated *common.SDKTime `mandatory:"false" json:"timeTerminated"`

	// Date and time the CCD was updated lastly. The format is defined by RFC 3339.
	// Example: `2025-01-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

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

	// The inclusive time when the CCD offer is no longer in effect, and must be after `timeCcdOfferStarts`.
	// The format is defined by RFC 3339.
	TimeCcdOfferEnds *common.SDKTime `mandatory:"false" json:"timeCcdOfferEnds"`

	// Currency code of listPrice associated with primary product/sku. This will be always null.
	// Please refer `externalPriceOfSku` to know product/sku's associated list price in its currency code.
	CurrencyCode *string `mandatory:"false" json:"currencyCode"`

	// ListPrice of a sku.
	// Please refer `externalPriceOfSku` to know product/sku's associated list price in its currency code.
	ListPrice *string `mandatory:"false" json:"listPrice"`

	// The subscription OCID if a this is dedicated CCD offer for a customer.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// The root compartment OCID of the customer's tenancy if offer is dedicated to a particular customer only.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m CcdOffer) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CcdOffer) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCcdOfferTypeOfOfferEnum(string(m.TypeOfOffer)); !ok && m.TypeOfOffer != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TypeOfOffer: %s. Supported values are: %s.", m.TypeOfOffer, strings.Join(GetCcdOfferTypeOfOfferEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdOfferSourceTypeEnum(string(m.SourceType)); !ok && m.SourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceType: %s. Supported values are: %s.", m.SourceType, strings.Join(GetCcdOfferSourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdOfferReclaimTermUnitEnum(string(m.ReclaimTermUnit)); !ok && m.ReclaimTermUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReclaimTermUnit: %s. Supported values are: %s.", m.ReclaimTermUnit, strings.Join(GetCcdOfferReclaimTermUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdOfferCustomerSegmentEnum(string(m.CustomerSegment)); !ok && m.CustomerSegment != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CustomerSegment: %s. Supported values are: %s.", m.CustomerSegment, strings.Join(GetCcdOfferCustomerSegmentEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCcdOfferUnitEnum(string(m.Unit)); !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetCcdOfferUnitEnumStringValues(), ",")))
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
func (m *CcdOffer) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TypeOfOffer         CcdOfferTypeOfOfferEnum           `json:"typeOfOffer"`
		SourceType          CcdOfferSourceTypeEnum            `json:"sourceType"`
		ReclaimTerm         *int64                            `json:"reclaimTerm"`
		ReclaimTermUnit     CcdOfferReclaimTermUnitEnum       `json:"reclaimTermUnit"`
		LastUpdatedBy       *string                           `json:"lastUpdatedBy"`
		CreatedBy           *string                           `json:"createdBy"`
		OrderDocumentNumber *string                           `json:"orderDocumentNumber"`
		CustomerSegment     CcdOfferCustomerSegmentEnum       `json:"customerSegment"`
		IsFillerUsageEnable *bool                             `json:"isFillerUsageEnable"`
		ReportingRegion     *string                           `json:"reportingRegion"`
		MaxQuantity         *int64                            `json:"maxQuantity"`
		Unit                CcdOfferUnitEnum                  `json:"unit"`
		ExternalPriceOfSku  []SkuExternalPriceModel           `json:"externalPriceOfSku"`
		OfferRestriction    *OfferRestriction                 `json:"offerRestriction"`
		BundleSkus          []bundlesku                       `json:"bundleSkus"`
		LifecycleState      CcdOfferLifecycleStateEnum        `json:"lifecycleState"`
		LifecycleDetails    CcdOfferLifecycleDetailsEnum      `json:"lifecycleDetails"`
		TimeCreated         *common.SDKTime                   `json:"timeCreated"`
		TimeTerminated      *common.SDKTime                   `json:"timeTerminated"`
		TimeUpdated         *common.SDKTime                   `json:"timeUpdated"`
		FreeformTags        map[string]string                 `json:"freeformTags"`
		DefinedTags         map[string]map[string]interface{} `json:"definedTags"`
		SystemTags          map[string]map[string]interface{} `json:"systemTags"`
		TimeCcdOfferStarts  *common.SDKTime                   `json:"timeCcdOfferStarts"`
		TimeCcdOfferEnds    *common.SDKTime                   `json:"timeCcdOfferEnds"`
		CurrencyCode        *string                           `json:"currencyCode"`
		ListPrice           *string                           `json:"listPrice"`
		SubscriptionId      *string                           `json:"subscriptionId"`
		CompartmentId       *string                           `json:"compartmentId"`
		Id                  *string                           `json:"id"`
		DisplayName         *string                           `json:"displayName"`
		Sku                 *string                           `json:"sku"`
		SkuDisplayName      *string                           `json:"skuDisplayName"`
		SkuMetric           *string                           `json:"skuMetric"`
		MinQuantity         *int64                            `json:"minQuantity"`
		Term                *int                              `json:"term"`
		DiscountPct         *string                           `json:"discountPct"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TypeOfOffer = model.TypeOfOffer

	m.SourceType = model.SourceType

	m.ReclaimTerm = model.ReclaimTerm

	m.ReclaimTermUnit = model.ReclaimTermUnit

	m.LastUpdatedBy = model.LastUpdatedBy

	m.CreatedBy = model.CreatedBy

	m.OrderDocumentNumber = model.OrderDocumentNumber

	m.CustomerSegment = model.CustomerSegment

	m.IsFillerUsageEnable = model.IsFillerUsageEnable

	m.ReportingRegion = model.ReportingRegion

	m.MaxQuantity = model.MaxQuantity

	m.Unit = model.Unit

	m.ExternalPriceOfSku = make([]SkuExternalPriceModel, len(model.ExternalPriceOfSku))
	copy(m.ExternalPriceOfSku, model.ExternalPriceOfSku)
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
	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.TimeCreated = model.TimeCreated

	m.TimeTerminated = model.TimeTerminated

	m.TimeUpdated = model.TimeUpdated

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.TimeCcdOfferStarts = model.TimeCcdOfferStarts

	m.TimeCcdOfferEnds = model.TimeCcdOfferEnds

	m.CurrencyCode = model.CurrencyCode

	m.ListPrice = model.ListPrice

	m.SubscriptionId = model.SubscriptionId

	m.CompartmentId = model.CompartmentId

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

// CcdOfferTypeOfOfferEnum Enum with underlying type: string
type CcdOfferTypeOfOfferEnum string

// Set of constants representing the allowable values for CcdOfferTypeOfOfferEnum
const (
	CcdOfferTypeOfOfferStandard    CcdOfferTypeOfOfferEnum = "STANDARD"
	CcdOfferTypeOfOfferCustom      CcdOfferTypeOfOfferEnum = "CUSTOM"
	CcdOfferTypeOfOfferPrivate     CcdOfferTypeOfOfferEnum = "PRIVATE"
	CcdOfferTypeOfOfferContractual CcdOfferTypeOfOfferEnum = "CONTRACTUAL"
	CcdOfferTypeOfOfferFlex        CcdOfferTypeOfOfferEnum = "FLEX"
)

var mappingCcdOfferTypeOfOfferEnum = map[string]CcdOfferTypeOfOfferEnum{
	"STANDARD":    CcdOfferTypeOfOfferStandard,
	"CUSTOM":      CcdOfferTypeOfOfferCustom,
	"PRIVATE":     CcdOfferTypeOfOfferPrivate,
	"CONTRACTUAL": CcdOfferTypeOfOfferContractual,
	"FLEX":        CcdOfferTypeOfOfferFlex,
}

var mappingCcdOfferTypeOfOfferEnumLowerCase = map[string]CcdOfferTypeOfOfferEnum{
	"standard":    CcdOfferTypeOfOfferStandard,
	"custom":      CcdOfferTypeOfOfferCustom,
	"private":     CcdOfferTypeOfOfferPrivate,
	"contractual": CcdOfferTypeOfOfferContractual,
	"flex":        CcdOfferTypeOfOfferFlex,
}

// GetCcdOfferTypeOfOfferEnumValues Enumerates the set of values for CcdOfferTypeOfOfferEnum
func GetCcdOfferTypeOfOfferEnumValues() []CcdOfferTypeOfOfferEnum {
	values := make([]CcdOfferTypeOfOfferEnum, 0)
	for _, v := range mappingCcdOfferTypeOfOfferEnum {
		values = append(values, v)
	}
	return values
}

// GetCcdOfferTypeOfOfferEnumStringValues Enumerates the set of values in String for CcdOfferTypeOfOfferEnum
func GetCcdOfferTypeOfOfferEnumStringValues() []string {
	return []string{
		"STANDARD",
		"CUSTOM",
		"PRIVATE",
		"CONTRACTUAL",
		"FLEX",
	}
}

// GetMappingCcdOfferTypeOfOfferEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCcdOfferTypeOfOfferEnum(val string) (CcdOfferTypeOfOfferEnum, bool) {
	enum, ok := mappingCcdOfferTypeOfOfferEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CcdOfferSourceTypeEnum Enum with underlying type: string
type CcdOfferSourceTypeEnum string

// Set of constants representing the allowable values for CcdOfferSourceTypeEnum
const (
	CcdOfferSourceTypeContractCommitment CcdOfferSourceTypeEnum = "CONTRACT_COMMITMENT"
	CcdOfferSourceTypeContractOffer      CcdOfferSourceTypeEnum = "CONTRACT_OFFER"
	CcdOfferSourceTypeFlex               CcdOfferSourceTypeEnum = "FLEX"
	CcdOfferSourceTypeOciGenerated       CcdOfferSourceTypeEnum = "OCI_GENERATED"
)

var mappingCcdOfferSourceTypeEnum = map[string]CcdOfferSourceTypeEnum{
	"CONTRACT_COMMITMENT": CcdOfferSourceTypeContractCommitment,
	"CONTRACT_OFFER":      CcdOfferSourceTypeContractOffer,
	"FLEX":                CcdOfferSourceTypeFlex,
	"OCI_GENERATED":       CcdOfferSourceTypeOciGenerated,
}

var mappingCcdOfferSourceTypeEnumLowerCase = map[string]CcdOfferSourceTypeEnum{
	"contract_commitment": CcdOfferSourceTypeContractCommitment,
	"contract_offer":      CcdOfferSourceTypeContractOffer,
	"flex":                CcdOfferSourceTypeFlex,
	"oci_generated":       CcdOfferSourceTypeOciGenerated,
}

// GetCcdOfferSourceTypeEnumValues Enumerates the set of values for CcdOfferSourceTypeEnum
func GetCcdOfferSourceTypeEnumValues() []CcdOfferSourceTypeEnum {
	values := make([]CcdOfferSourceTypeEnum, 0)
	for _, v := range mappingCcdOfferSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCcdOfferSourceTypeEnumStringValues Enumerates the set of values in String for CcdOfferSourceTypeEnum
func GetCcdOfferSourceTypeEnumStringValues() []string {
	return []string{
		"CONTRACT_COMMITMENT",
		"CONTRACT_OFFER",
		"FLEX",
		"OCI_GENERATED",
	}
}

// GetMappingCcdOfferSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCcdOfferSourceTypeEnum(val string) (CcdOfferSourceTypeEnum, bool) {
	enum, ok := mappingCcdOfferSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CcdOfferReclaimTermUnitEnum Enum with underlying type: string
type CcdOfferReclaimTermUnitEnum string

// Set of constants representing the allowable values for CcdOfferReclaimTermUnitEnum
const (
	CcdOfferReclaimTermUnitDays CcdOfferReclaimTermUnitEnum = "DAYS"
)

var mappingCcdOfferReclaimTermUnitEnum = map[string]CcdOfferReclaimTermUnitEnum{
	"DAYS": CcdOfferReclaimTermUnitDays,
}

var mappingCcdOfferReclaimTermUnitEnumLowerCase = map[string]CcdOfferReclaimTermUnitEnum{
	"days": CcdOfferReclaimTermUnitDays,
}

// GetCcdOfferReclaimTermUnitEnumValues Enumerates the set of values for CcdOfferReclaimTermUnitEnum
func GetCcdOfferReclaimTermUnitEnumValues() []CcdOfferReclaimTermUnitEnum {
	values := make([]CcdOfferReclaimTermUnitEnum, 0)
	for _, v := range mappingCcdOfferReclaimTermUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetCcdOfferReclaimTermUnitEnumStringValues Enumerates the set of values in String for CcdOfferReclaimTermUnitEnum
func GetCcdOfferReclaimTermUnitEnumStringValues() []string {
	return []string{
		"DAYS",
	}
}

// GetMappingCcdOfferReclaimTermUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCcdOfferReclaimTermUnitEnum(val string) (CcdOfferReclaimTermUnitEnum, bool) {
	enum, ok := mappingCcdOfferReclaimTermUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CcdOfferCustomerSegmentEnum Enum with underlying type: string
type CcdOfferCustomerSegmentEnum string

// Set of constants representing the allowable values for CcdOfferCustomerSegmentEnum
const (
	CcdOfferCustomerSegmentInternal CcdOfferCustomerSegmentEnum = "INTERNAL"
	CcdOfferCustomerSegmentExternal CcdOfferCustomerSegmentEnum = "EXTERNAL"
)

var mappingCcdOfferCustomerSegmentEnum = map[string]CcdOfferCustomerSegmentEnum{
	"INTERNAL": CcdOfferCustomerSegmentInternal,
	"EXTERNAL": CcdOfferCustomerSegmentExternal,
}

var mappingCcdOfferCustomerSegmentEnumLowerCase = map[string]CcdOfferCustomerSegmentEnum{
	"internal": CcdOfferCustomerSegmentInternal,
	"external": CcdOfferCustomerSegmentExternal,
}

// GetCcdOfferCustomerSegmentEnumValues Enumerates the set of values for CcdOfferCustomerSegmentEnum
func GetCcdOfferCustomerSegmentEnumValues() []CcdOfferCustomerSegmentEnum {
	values := make([]CcdOfferCustomerSegmentEnum, 0)
	for _, v := range mappingCcdOfferCustomerSegmentEnum {
		values = append(values, v)
	}
	return values
}

// GetCcdOfferCustomerSegmentEnumStringValues Enumerates the set of values in String for CcdOfferCustomerSegmentEnum
func GetCcdOfferCustomerSegmentEnumStringValues() []string {
	return []string{
		"INTERNAL",
		"EXTERNAL",
	}
}

// GetMappingCcdOfferCustomerSegmentEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCcdOfferCustomerSegmentEnum(val string) (CcdOfferCustomerSegmentEnum, bool) {
	enum, ok := mappingCcdOfferCustomerSegmentEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CcdOfferUnitEnum Enum with underlying type: string
type CcdOfferUnitEnum string

// Set of constants representing the allowable values for CcdOfferUnitEnum
const (
	CcdOfferUnitMonths CcdOfferUnitEnum = "MONTHS"
	CcdOfferUnitYears  CcdOfferUnitEnum = "YEARS"
	CcdOfferUnitDays   CcdOfferUnitEnum = "DAYS"
	CcdOfferUnitHours  CcdOfferUnitEnum = "HOURS"
)

var mappingCcdOfferUnitEnum = map[string]CcdOfferUnitEnum{
	"MONTHS": CcdOfferUnitMonths,
	"YEARS":  CcdOfferUnitYears,
	"DAYS":   CcdOfferUnitDays,
	"HOURS":  CcdOfferUnitHours,
}

var mappingCcdOfferUnitEnumLowerCase = map[string]CcdOfferUnitEnum{
	"months": CcdOfferUnitMonths,
	"years":  CcdOfferUnitYears,
	"days":   CcdOfferUnitDays,
	"hours":  CcdOfferUnitHours,
}

// GetCcdOfferUnitEnumValues Enumerates the set of values for CcdOfferUnitEnum
func GetCcdOfferUnitEnumValues() []CcdOfferUnitEnum {
	values := make([]CcdOfferUnitEnum, 0)
	for _, v := range mappingCcdOfferUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetCcdOfferUnitEnumStringValues Enumerates the set of values in String for CcdOfferUnitEnum
func GetCcdOfferUnitEnumStringValues() []string {
	return []string{
		"MONTHS",
		"YEARS",
		"DAYS",
		"HOURS",
	}
}

// GetMappingCcdOfferUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCcdOfferUnitEnum(val string) (CcdOfferUnitEnum, bool) {
	enum, ok := mappingCcdOfferUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CcdOfferLifecycleStateEnum Enum with underlying type: string
type CcdOfferLifecycleStateEnum string

// Set of constants representing the allowable values for CcdOfferLifecycleStateEnum
const (
	CcdOfferLifecycleStateActive CcdOfferLifecycleStateEnum = "ACTIVE"
)

var mappingCcdOfferLifecycleStateEnum = map[string]CcdOfferLifecycleStateEnum{
	"ACTIVE": CcdOfferLifecycleStateActive,
}

var mappingCcdOfferLifecycleStateEnumLowerCase = map[string]CcdOfferLifecycleStateEnum{
	"active": CcdOfferLifecycleStateActive,
}

// GetCcdOfferLifecycleStateEnumValues Enumerates the set of values for CcdOfferLifecycleStateEnum
func GetCcdOfferLifecycleStateEnumValues() []CcdOfferLifecycleStateEnum {
	values := make([]CcdOfferLifecycleStateEnum, 0)
	for _, v := range mappingCcdOfferLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCcdOfferLifecycleStateEnumStringValues Enumerates the set of values in String for CcdOfferLifecycleStateEnum
func GetCcdOfferLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
	}
}

// GetMappingCcdOfferLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCcdOfferLifecycleStateEnum(val string) (CcdOfferLifecycleStateEnum, bool) {
	enum, ok := mappingCcdOfferLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CcdOfferLifecycleDetailsEnum Enum with underlying type: string
type CcdOfferLifecycleDetailsEnum string

// Set of constants representing the allowable values for CcdOfferLifecycleDetailsEnum
const (
	CcdOfferLifecycleDetailsScheduled CcdOfferLifecycleDetailsEnum = "SCHEDULED"
	CcdOfferLifecycleDetailsCurrent   CcdOfferLifecycleDetailsEnum = "CURRENT"
	CcdOfferLifecycleDetailsCanceled  CcdOfferLifecycleDetailsEnum = "CANCELED"
	CcdOfferLifecycleDetailsEnded     CcdOfferLifecycleDetailsEnum = "ENDED"
)

var mappingCcdOfferLifecycleDetailsEnum = map[string]CcdOfferLifecycleDetailsEnum{
	"SCHEDULED": CcdOfferLifecycleDetailsScheduled,
	"CURRENT":   CcdOfferLifecycleDetailsCurrent,
	"CANCELED":  CcdOfferLifecycleDetailsCanceled,
	"ENDED":     CcdOfferLifecycleDetailsEnded,
}

var mappingCcdOfferLifecycleDetailsEnumLowerCase = map[string]CcdOfferLifecycleDetailsEnum{
	"scheduled": CcdOfferLifecycleDetailsScheduled,
	"current":   CcdOfferLifecycleDetailsCurrent,
	"canceled":  CcdOfferLifecycleDetailsCanceled,
	"ended":     CcdOfferLifecycleDetailsEnded,
}

// GetCcdOfferLifecycleDetailsEnumValues Enumerates the set of values for CcdOfferLifecycleDetailsEnum
func GetCcdOfferLifecycleDetailsEnumValues() []CcdOfferLifecycleDetailsEnum {
	values := make([]CcdOfferLifecycleDetailsEnum, 0)
	for _, v := range mappingCcdOfferLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetCcdOfferLifecycleDetailsEnumStringValues Enumerates the set of values in String for CcdOfferLifecycleDetailsEnum
func GetCcdOfferLifecycleDetailsEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"CURRENT",
		"CANCELED",
		"ENDED",
	}
}

// GetMappingCcdOfferLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCcdOfferLifecycleDetailsEnum(val string) (CcdOfferLifecycleDetailsEnum, bool) {
	enum, ok := mappingCcdOfferLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
