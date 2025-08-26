// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Threat Intelligence API
//
// Use the Threat Intelligence API to search for information about known threat indicators, including suspicious IP addresses, domain names, and other digital fingerprints. Threat Intelligence is a managed database of curated threat intelligence that comes from first party Oracle security insights, open source feeds, and vendor-procured data. For more information, see the Threat Intelligence documentation (https://docs.oracle.com/iaas/Content/threat-intel/home.htm).
//

package threatintelligence

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IndicatorReference A reference to a threat indicator resource.
type IndicatorReference struct {

	// The unique OCID of the referenced threat indicator.
	IndicatorId *string `mandatory:"true" json:"indicatorId"`
}

func (m IndicatorReference) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IndicatorReference) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m IndicatorReference) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeIndicatorReference IndicatorReference
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeIndicatorReference
	}{
		"INDICATOR",
		(MarshalTypeIndicatorReference)(m),
	}

	return json.Marshal(&s)
}
