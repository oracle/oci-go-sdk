// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Threat Intelligence API
//
// Use the Threat Intelligence API to view indicators of compromise and related items. For more information, see Overview of Threat Intelligence (https://docs.cloud.oracle.com/Content/ThreatIntelligence/Concepts/threatintelligenceoverview.htm).
//

package threatintelligence

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v62/common"
	"strings"
)

// DataVisibility The visibility level associated with data and an associated TLP (https://www.cisa.gov/tlp) level.
type DataVisibility struct {

	// The name of the visibility level.
	Name *string `mandatory:"true" json:"name"`

	// The Traffic Light Protocol (TLP) name of the visibility level.
	TlpName DataVisibilityTlpNameEnum `mandatory:"true" json:"tlpName"`
}

func (m DataVisibility) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataVisibility) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataVisibilityTlpNameEnum(string(m.TlpName)); !ok && m.TlpName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TlpName: %s. Supported values are: %s.", m.TlpName, strings.Join(GetDataVisibilityTlpNameEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataVisibilityTlpNameEnum Enum with underlying type: string
type DataVisibilityTlpNameEnum string

// Set of constants representing the allowable values for DataVisibilityTlpNameEnum
const (
	DataVisibilityTlpNameInternalAudit DataVisibilityTlpNameEnum = "TLP_INTERNAL_AUDIT"
	DataVisibilityTlpNameWhite         DataVisibilityTlpNameEnum = "TLP_WHITE"
	DataVisibilityTlpNameGreen         DataVisibilityTlpNameEnum = "TLP_GREEN"
	DataVisibilityTlpNameAmber         DataVisibilityTlpNameEnum = "TLP_AMBER"
	DataVisibilityTlpNameRed           DataVisibilityTlpNameEnum = "TLP_RED"
)

var mappingDataVisibilityTlpNameEnum = map[string]DataVisibilityTlpNameEnum{
	"TLP_INTERNAL_AUDIT": DataVisibilityTlpNameInternalAudit,
	"TLP_WHITE":          DataVisibilityTlpNameWhite,
	"TLP_GREEN":          DataVisibilityTlpNameGreen,
	"TLP_AMBER":          DataVisibilityTlpNameAmber,
	"TLP_RED":            DataVisibilityTlpNameRed,
}

var mappingDataVisibilityTlpNameEnumLowerCase = map[string]DataVisibilityTlpNameEnum{
	"tlp_internal_audit": DataVisibilityTlpNameInternalAudit,
	"tlp_white":          DataVisibilityTlpNameWhite,
	"tlp_green":          DataVisibilityTlpNameGreen,
	"tlp_amber":          DataVisibilityTlpNameAmber,
	"tlp_red":            DataVisibilityTlpNameRed,
}

// GetDataVisibilityTlpNameEnumValues Enumerates the set of values for DataVisibilityTlpNameEnum
func GetDataVisibilityTlpNameEnumValues() []DataVisibilityTlpNameEnum {
	values := make([]DataVisibilityTlpNameEnum, 0)
	for _, v := range mappingDataVisibilityTlpNameEnum {
		values = append(values, v)
	}
	return values
}

// GetDataVisibilityTlpNameEnumStringValues Enumerates the set of values in String for DataVisibilityTlpNameEnum
func GetDataVisibilityTlpNameEnumStringValues() []string {
	return []string{
		"TLP_INTERNAL_AUDIT",
		"TLP_WHITE",
		"TLP_GREEN",
		"TLP_AMBER",
		"TLP_RED",
	}
}

// GetMappingDataVisibilityTlpNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataVisibilityTlpNameEnum(val string) (DataVisibilityTlpNameEnum, bool) {
	enum, ok := mappingDataVisibilityTlpNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
