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
	"github.com/oracle/oci-go-sdk/v60/common"
	"strings"
)

// DataAttribution The confidence, source information, and visibility for a particular sighting or observation of some data associated with an indicator such as threat type, attribute or relationship.
type DataAttribution struct {

	// Confidence is an integer from 0 to 100 that provides a measure of our certainty in the maliciousness of data attributed to an indicator.  For example, if the source of the data being attributed is the Tor Project, our confidence that the associated indicator is a tor exit node would be 100.
	Confidence *int `mandatory:"true" json:"confidence"`

	Source *IndicatorSourceSummary `mandatory:"true" json:"source"`

	Visibility *DataVisibility `mandatory:"true" json:"visibility"`

	// The last time this data was seen for this entity. An RFC3339 formatted datetime string
	TimeLastSeen *common.SDKTime `mandatory:"true" json:"timeLastSeen"`

	// The time the data was first seen for this entity. Defaults to time last seen if no time first seen is provided from the data source. An RFC3339 formatted datetime string
	TimeFirstSeen *common.SDKTime `mandatory:"false" json:"timeFirstSeen"`
}

func (m DataAttribution) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataAttribution) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
