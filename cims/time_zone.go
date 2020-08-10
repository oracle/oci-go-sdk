// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

// TimeZoneEnum Enum with underlying type: string
type TimeZoneEnum string

// Set of constants representing the allowable values for TimeZoneEnum
const (
	TimeZoneGreenwichMeanTime                TimeZoneEnum = "GREENWICH_MEAN_TIME"
	TimeZoneEuropeanCentralTime              TimeZoneEnum = "EUROPEAN_CENTRAL_TIME"
	TimeZoneEasternEuropeanTime              TimeZoneEnum = "EASTERN_EUROPEAN_TIME"
	TimeZoneEasternAfricanTime               TimeZoneEnum = "EASTERN_AFRICAN_TIME"
	TimeZoneMiddleEastTime                   TimeZoneEnum = "MIDDLE_EAST_TIME"
	TimeZoneNearEastTime                     TimeZoneEnum = "NEAR_EAST_TIME"
	TimeZonePakistanLahoreTime               TimeZoneEnum = "PAKISTAN_LAHORE_TIME"
	TimeZoneIndiaStandardTime                TimeZoneEnum = "INDIA_STANDARD_TIME"
	TimeZoneBangladeshStandardTime           TimeZoneEnum = "BANGLADESH_STANDARD_TIME"
	TimeZoneVietnamStandardTime              TimeZoneEnum = "VIETNAM_STANDARD_TIME"
	TimeZoneChinaTaiwanTime                  TimeZoneEnum = "CHINA_TAIWAN_TIME"
	TimeZoneJapanStandardTime                TimeZoneEnum = "JAPAN_STANDARD_TIME"
	TimeZoneAustraliaCentralTime             TimeZoneEnum = "AUSTRALIA_CENTRAL_TIME"
	TimeZoneAustraliaEasternTime             TimeZoneEnum = "AUSTRALIA_EASTERN_TIME"
	TimeZoneSolomonStandardTime              TimeZoneEnum = "SOLOMON_STANDARD_TIME"
	TimeZoneNewZealandStandardTime           TimeZoneEnum = "NEW_ZEALAND_STANDARD_TIME"
	TimeZoneMidwayIslandsTime                TimeZoneEnum = "MIDWAY_ISLANDS_TIME"
	TimeZoneHawaiiStandardTime               TimeZoneEnum = "HAWAII_STANDARD_TIME"
	TimeZoneAlaskaStandardTime               TimeZoneEnum = "ALASKA_STANDARD_TIME"
	TimeZonePacificStandardTime              TimeZoneEnum = "PACIFIC_STANDARD_TIME"
	TimeZoneMountainStandardTime             TimeZoneEnum = "MOUNTAIN_STANDARD_TIME"
	TimeZoneCentralStandardTime              TimeZoneEnum = "CENTRAL_STANDARD_TIME"
	TimeZoneEasternStandardTime              TimeZoneEnum = "EASTERN_STANDARD_TIME"
	TimeZonePuertoRicoAndUsVirginIslandsTime TimeZoneEnum = "PUERTO_RICO_AND_US_VIRGIN_ISLANDS_TIME"
	TimeZoneCanadaNewfoundlandTime           TimeZoneEnum = "CANADA_NEWFOUNDLAND_TIME"
	TimeZoneArgentinaStandardTime            TimeZoneEnum = "ARGENTINA_STANDARD_TIME"
	TimeZoneMidAtlanticTime                  TimeZoneEnum = "MID_ATLANTIC_TIME"
	TimeZoneCentralAfricanTime               TimeZoneEnum = "CENTRAL_AFRICAN_TIME"
)

var mappingTimeZone = map[string]TimeZoneEnum{
	"GREENWICH_MEAN_TIME":                    TimeZoneGreenwichMeanTime,
	"EUROPEAN_CENTRAL_TIME":                  TimeZoneEuropeanCentralTime,
	"EASTERN_EUROPEAN_TIME":                  TimeZoneEasternEuropeanTime,
	"EASTERN_AFRICAN_TIME":                   TimeZoneEasternAfricanTime,
	"MIDDLE_EAST_TIME":                       TimeZoneMiddleEastTime,
	"NEAR_EAST_TIME":                         TimeZoneNearEastTime,
	"PAKISTAN_LAHORE_TIME":                   TimeZonePakistanLahoreTime,
	"INDIA_STANDARD_TIME":                    TimeZoneIndiaStandardTime,
	"BANGLADESH_STANDARD_TIME":               TimeZoneBangladeshStandardTime,
	"VIETNAM_STANDARD_TIME":                  TimeZoneVietnamStandardTime,
	"CHINA_TAIWAN_TIME":                      TimeZoneChinaTaiwanTime,
	"JAPAN_STANDARD_TIME":                    TimeZoneJapanStandardTime,
	"AUSTRALIA_CENTRAL_TIME":                 TimeZoneAustraliaCentralTime,
	"AUSTRALIA_EASTERN_TIME":                 TimeZoneAustraliaEasternTime,
	"SOLOMON_STANDARD_TIME":                  TimeZoneSolomonStandardTime,
	"NEW_ZEALAND_STANDARD_TIME":              TimeZoneNewZealandStandardTime,
	"MIDWAY_ISLANDS_TIME":                    TimeZoneMidwayIslandsTime,
	"HAWAII_STANDARD_TIME":                   TimeZoneHawaiiStandardTime,
	"ALASKA_STANDARD_TIME":                   TimeZoneAlaskaStandardTime,
	"PACIFIC_STANDARD_TIME":                  TimeZonePacificStandardTime,
	"MOUNTAIN_STANDARD_TIME":                 TimeZoneMountainStandardTime,
	"CENTRAL_STANDARD_TIME":                  TimeZoneCentralStandardTime,
	"EASTERN_STANDARD_TIME":                  TimeZoneEasternStandardTime,
	"PUERTO_RICO_AND_US_VIRGIN_ISLANDS_TIME": TimeZonePuertoRicoAndUsVirginIslandsTime,
	"CANADA_NEWFOUNDLAND_TIME":               TimeZoneCanadaNewfoundlandTime,
	"ARGENTINA_STANDARD_TIME":                TimeZoneArgentinaStandardTime,
	"MID_ATLANTIC_TIME":                      TimeZoneMidAtlanticTime,
	"CENTRAL_AFRICAN_TIME":                   TimeZoneCentralAfricanTime,
}

// GetTimeZoneEnumValues Enumerates the set of values for TimeZoneEnum
func GetTimeZoneEnumValues() []TimeZoneEnum {
	values := make([]TimeZoneEnum, 0)
	for _, v := range mappingTimeZone {
		values = append(values, v)
	}
	return values
}
