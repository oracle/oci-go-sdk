// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// DataAttribution The confidence, source information, and visibility for a particular sighting or observation of some data associated with a threat indicator. This associated data can be the indicator's threat type, attribute, or relationship.
type DataAttribution struct {

	// An integer from 0 to 100 that provides a measure of our certainty in the maliciousness of data attributed to an indicator. For example, if the source of the data being attributed is the Tor Project, our confidence that the associated indicator is a tor exit node would be 100.
	Confidence *int `mandatory:"true" json:"confidence"`

	Source *IndicatorSourceSummary `mandatory:"true" json:"source"`

	Visibility *DataVisibility `mandatory:"true" json:"visibility"`

	// The last date and time the attribution data was seen for this entity. An RFC3339 formatted string.
	TimeLastSeen *common.SDKTime `mandatory:"true" json:"timeLastSeen"`

	// The date and time the attribution data was first seen for this entity. If the data source does not provide this information, it is set to the last time it was seen. An RFC3339 formatted string.
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
