// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Threat Intelligence API
//
// Use the Threat Intelligence API to search for information about known threat indicators, including suspicious IP addresses, domain names, and other digital fingerprints. Threat Intelligence is a managed database of curated threat intelligence that comes from first party Oracle security insights, open source feeds, and vendor-procured data. For more information, see the Threat Intelligence documentation (https://docs.cloud.oracle.com/iaas/Content/threat-intel/home.htm).
//

package threatintelligence

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SummarizeIndicatorsDetails Query parameters to filter indicators
type SummarizeIndicatorsDetails struct {

	// The type of indicator this is
	IndicatorType IndicatorTypeEnum `mandatory:"false" json:"indicatorType,omitempty"`

	// The value for the type of indicator this is
	IndicatorValue *string `mandatory:"false" json:"indicatorValue"`

	// The threat type of entites to be returned.
	ThreatTypes []string `mandatory:"false" json:"threatTypes"`

	// The minimum level of confidence to return
	ConfidenceGreaterThanOrEqualTo *int `mandatory:"false" json:"confidenceGreaterThanOrEqualTo"`

	// The oldest update time of entities to be returned.
	TimeUpdatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" json:"timeUpdatedGreaterThanOrEqualTo"`

	// The newest update time of entities to be returned.
	TimeUpdatedLessThan *common.SDKTime `mandatory:"false" json:"timeUpdatedLessThan"`

	// The oldest last seen time of entities to be returned.
	TimeLastSeenGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" json:"timeLastSeenGreaterThanOrEqualTo"`

	// The newest last seen time of entities to be returned.
	TimeLastSeenLessThan *common.SDKTime `mandatory:"false" json:"timeLastSeenLessThan"`

	// The oldest creation time of entities to be returned.
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" json:"timeCreatedGreaterThanOrEqualTo"`

	// The newest creation time of entities to be returned.
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" json:"timeCreatedLessThan"`

	// Filter to include indicators that have been seen by the provided source.
	IndicatorSeenBy *string `mandatory:"false" json:"indicatorSeenBy"`

	// Filter to include indicators associated with the provided malware.
	Malware *string `mandatory:"false" json:"malware"`

	// Filter to included indicators associated with the provided threat actor.
	ThreatActor *string `mandatory:"false" json:"threatActor"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder SortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`

	// The field to sort by. Only one field to sort by may be provided
	SortBy SummarizeIndicatorsDetailsSortByEnum `mandatory:"false" json:"sortBy,omitempty"`
}

func (m SummarizeIndicatorsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeIndicatorsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingIndicatorTypeEnum(string(m.IndicatorType)); !ok && m.IndicatorType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IndicatorType: %s. Supported values are: %s.", m.IndicatorType, strings.Join(GetIndicatorTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeIndicatorsDetailsSortByEnum(string(m.SortBy)); !ok && m.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", m.SortBy, strings.Join(GetSummarizeIndicatorsDetailsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeIndicatorsDetailsSortByEnum Enum with underlying type: string
type SummarizeIndicatorsDetailsSortByEnum string

// Set of constants representing the allowable values for SummarizeIndicatorsDetailsSortByEnum
const (
	SummarizeIndicatorsDetailsSortByConfidence   SummarizeIndicatorsDetailsSortByEnum = "CONFIDENCE"
	SummarizeIndicatorsDetailsSortByTimecreated  SummarizeIndicatorsDetailsSortByEnum = "TIMECREATED"
	SummarizeIndicatorsDetailsSortByTimeupdated  SummarizeIndicatorsDetailsSortByEnum = "TIMEUPDATED"
	SummarizeIndicatorsDetailsSortByTimelastseen SummarizeIndicatorsDetailsSortByEnum = "TIMELASTSEEN"
)

var mappingSummarizeIndicatorsDetailsSortByEnum = map[string]SummarizeIndicatorsDetailsSortByEnum{
	"CONFIDENCE":   SummarizeIndicatorsDetailsSortByConfidence,
	"TIMECREATED":  SummarizeIndicatorsDetailsSortByTimecreated,
	"TIMEUPDATED":  SummarizeIndicatorsDetailsSortByTimeupdated,
	"TIMELASTSEEN": SummarizeIndicatorsDetailsSortByTimelastseen,
}

var mappingSummarizeIndicatorsDetailsSortByEnumLowerCase = map[string]SummarizeIndicatorsDetailsSortByEnum{
	"confidence":   SummarizeIndicatorsDetailsSortByConfidence,
	"timecreated":  SummarizeIndicatorsDetailsSortByTimecreated,
	"timeupdated":  SummarizeIndicatorsDetailsSortByTimeupdated,
	"timelastseen": SummarizeIndicatorsDetailsSortByTimelastseen,
}

// GetSummarizeIndicatorsDetailsSortByEnumValues Enumerates the set of values for SummarizeIndicatorsDetailsSortByEnum
func GetSummarizeIndicatorsDetailsSortByEnumValues() []SummarizeIndicatorsDetailsSortByEnum {
	values := make([]SummarizeIndicatorsDetailsSortByEnum, 0)
	for _, v := range mappingSummarizeIndicatorsDetailsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeIndicatorsDetailsSortByEnumStringValues Enumerates the set of values in String for SummarizeIndicatorsDetailsSortByEnum
func GetSummarizeIndicatorsDetailsSortByEnumStringValues() []string {
	return []string{
		"CONFIDENCE",
		"TIMECREATED",
		"TIMEUPDATED",
		"TIMELASTSEEN",
	}
}

// GetMappingSummarizeIndicatorsDetailsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeIndicatorsDetailsSortByEnum(val string) (SummarizeIndicatorsDetailsSortByEnum, bool) {
	enum, ok := mappingSummarizeIndicatorsDetailsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
