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

// GeoLocation Support details based on geographic location
type GeoLocation struct {

	// Region for geographic location
	GeographicRegion GeoLocationGeographicRegionEnum `mandatory:"true" json:"geographicRegion"`

	// Support details for specific region
	Details *string `mandatory:"true" json:"details"`
}

func (m GeoLocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GeoLocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGeoLocationGeographicRegionEnum(string(m.GeographicRegion)); !ok && m.GeographicRegion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GeographicRegion: %s. Supported values are: %s.", m.GeographicRegion, strings.Join(GetGeoLocationGeographicRegionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GeoLocationGeographicRegionEnum Enum with underlying type: string
type GeoLocationGeographicRegionEnum string

// Set of constants representing the allowable values for GeoLocationGeographicRegionEnum
const (
	GeoLocationGeographicRegionAfrica       GeoLocationGeographicRegionEnum = "AFRICA"
	GeoLocationGeographicRegionAsiaPacific  GeoLocationGeographicRegionEnum = "ASIA_PACIFIC"
	GeoLocationGeographicRegionEurope       GeoLocationGeographicRegionEnum = "EUROPE"
	GeoLocationGeographicRegionJapan        GeoLocationGeographicRegionEnum = "JAPAN"
	GeoLocationGeographicRegionLatinAmerica GeoLocationGeographicRegionEnum = "LATIN_AMERICA"
	GeoLocationGeographicRegionMiddleEast   GeoLocationGeographicRegionEnum = "MIDDLE_EAST"
	GeoLocationGeographicRegionNorthAmerica GeoLocationGeographicRegionEnum = "NORTH_AMERICA"
)

var mappingGeoLocationGeographicRegionEnum = map[string]GeoLocationGeographicRegionEnum{
	"AFRICA":        GeoLocationGeographicRegionAfrica,
	"ASIA_PACIFIC":  GeoLocationGeographicRegionAsiaPacific,
	"EUROPE":        GeoLocationGeographicRegionEurope,
	"JAPAN":         GeoLocationGeographicRegionJapan,
	"LATIN_AMERICA": GeoLocationGeographicRegionLatinAmerica,
	"MIDDLE_EAST":   GeoLocationGeographicRegionMiddleEast,
	"NORTH_AMERICA": GeoLocationGeographicRegionNorthAmerica,
}

var mappingGeoLocationGeographicRegionEnumLowerCase = map[string]GeoLocationGeographicRegionEnum{
	"africa":        GeoLocationGeographicRegionAfrica,
	"asia_pacific":  GeoLocationGeographicRegionAsiaPacific,
	"europe":        GeoLocationGeographicRegionEurope,
	"japan":         GeoLocationGeographicRegionJapan,
	"latin_america": GeoLocationGeographicRegionLatinAmerica,
	"middle_east":   GeoLocationGeographicRegionMiddleEast,
	"north_america": GeoLocationGeographicRegionNorthAmerica,
}

// GetGeoLocationGeographicRegionEnumValues Enumerates the set of values for GeoLocationGeographicRegionEnum
func GetGeoLocationGeographicRegionEnumValues() []GeoLocationGeographicRegionEnum {
	values := make([]GeoLocationGeographicRegionEnum, 0)
	for _, v := range mappingGeoLocationGeographicRegionEnum {
		values = append(values, v)
	}
	return values
}

// GetGeoLocationGeographicRegionEnumStringValues Enumerates the set of values in String for GeoLocationGeographicRegionEnum
func GetGeoLocationGeographicRegionEnumStringValues() []string {
	return []string{
		"AFRICA",
		"ASIA_PACIFIC",
		"EUROPE",
		"JAPAN",
		"LATIN_AMERICA",
		"MIDDLE_EAST",
		"NORTH_AMERICA",
	}
}

// GetMappingGeoLocationGeographicRegionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGeoLocationGeographicRegionEnum(val string) (GeoLocationGeographicRegionEnum, bool) {
	enum, ok := mappingGeoLocationGeographicRegionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
