// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Threat Intelligence API
//
// Use the Threat Intelligence API to search for information about known threat indicators, including suspicious IP addresses, domain names, and other digital fingerprints. Threat Intelligence is a managed database of curated threat intelligence that comes from first party Oracle security insights, open source feeds, and vendor-procured data. For more information, see the Threat Intelligence documentation (https://docs.oracle.com/iaas/Content/threat-intel/home.htm).
//

package threatintelligence

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IndicatorCountDimensions The indicator dimension that was counted, such as the indicator type.
type IndicatorCountDimensions struct {

	// The compartment OCID that contains the indicator type.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The indicator type that was counted.
	Type IndicatorTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m IndicatorCountDimensions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IndicatorCountDimensions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingIndicatorTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetIndicatorTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
