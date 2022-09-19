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

// IndicatorSummary Summary of a data signature observed on a network or host that indicates a potential security threat.
type IndicatorSummary struct {

	// The OCID of the indicator.
	Id *string `mandatory:"true" json:"id"`

	// The type of indicator.
	Type IndicatorTypeEnum `mandatory:"true" json:"type"`

	// The indicator data value.
	Value *string `mandatory:"true" json:"value"`

	// Characteristics of the threat indicator based on previous observations or behavior. May include related tactics, techniques, and procedures.
	ThreatTypes []string `mandatory:"true" json:"threatTypes"`

	// A map of attributes with additional information about the indicator.
	// Each attribute has a name (string), value (string), and attribution (supporting data).
	Attributes []IndicatorAttributeSummary `mandatory:"true" json:"attributes"`

	// The date and time that the indicator was first detected. An RFC3339 formatted string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time that this indicator was last updated by the system.  Updates can include new reports or regular updates in confidence. The value is the same as `timeCreated` for a new indicator. An RFC3339 formatted string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The date and time that this indicator was last seen. The value is the same as `timeCreated` for a new indicator. An RFC3339 formatted string.
	TimeLastSeen *common.SDKTime `mandatory:"true" json:"timeLastSeen"`

	Geodata *GeodataDetails `mandatory:"true" json:"geodata"`

	// An integer from 0 to 100 that represents how certain we are that the indicator is malicious and a potential threat if it is detected communicating with your cloud resources. This confidence value is aggregated from the confidence in the threat types, attributes, and relationships to create an overall value for the indicator.
	Confidence *int `mandatory:"false" json:"confidence"`

	// The OCID of the compartment that contains this indicator.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The state of the indicator. It will always be `ACTIVE`.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m IndicatorSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IndicatorSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingIndicatorTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetIndicatorTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
