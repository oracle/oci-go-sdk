// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Threat Intelligence API
//
// Use the Threat Intelligence API to view indicators of compromise and related items. For more information, see Overview of Threat Intelligence (https://docs.cloud.oracle.com/Content/ThreatIntelligence/Concepts/threatintelligenceoverview.htm).
//

package threatintelligence

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v63/common"
	"strings"
)

// IndicatorReference a reference to an Indicator resource
type IndicatorReference struct {

	// the OCID of the referenced Indicator
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
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
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
