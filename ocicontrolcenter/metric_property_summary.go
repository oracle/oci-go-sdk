// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center API
//
// OCI Control Center (OCC) service enables you to monitor the region-level cloud consumption and provides the region-level capacity data, in realms where OCC is available. Use the OCI Control Center (OCC) API to explore region-level capacity and utilization information about core services. For more information, see OCI Control Center (https://docs.oracle.com/iaas/Content/control-center/home.htm).
//

package ocicontrolcenter

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MetricPropertySummary A summary of the properties that define a metric.
type MetricPropertySummary struct {

	// The name of the metric.
	MetricName *string `mandatory:"true" json:"metricName"`

	// Qualifiers provided in a metric definition. Available dimensions vary by metric namespace.
	Dimensions map[string]DimensionValue `mandatory:"false" json:"dimensions"`
}

func (m MetricPropertySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MetricPropertySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
