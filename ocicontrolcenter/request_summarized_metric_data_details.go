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

// RequestSummarizedMetricDataDetails The request details for retrieving aggregated data. Use the query and optional properties to filter the returned results.
type RequestSummarizedMetricDataDetails struct {

	// The source service or application to use when searching for metric data points to aggregate. For a list of valid namespaces, see ListNamespaces.
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	// The name of a metric for retrieving aggregated data. For a list of valid metrics for a given namespace, see ListMetricProperties.
	MetricName *string `mandatory:"true" json:"metricName"`

	// The OCID of the compartment to use for authorization to read metrics. To use the root compartment, provide the tenancyId.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Qualifiers to use when searching for metric data. For a list of valid dimensions for a given metric, see ListMetricProperties.
	Dimensions map[string]DimensionValue `mandatory:"false" json:"dimensions"`

	// The beginning of the sampled time range to use when searching for metric data points. Format is defined by <a href="https://www.rfc-editor.org/rfc/rfc3339">RFC3339</a>. The response includes metric data points for the sampled time. Example 2019-02-01T02:02:29.600Z
	StartTime *common.SDKTime `mandatory:"false" json:"startTime"`

	// The end of the sampled time range to use when searching for metric data points. Format is defined by <a href="https://www.rfc-editor.org/rfc/rfc3339">RFC3339</a>. The response excludes metric data points for sampled time. Example 2019-02-01T02:02:29.600Z
	EndTime *common.SDKTime `mandatory:"false" json:"endTime"`
}

func (m RequestSummarizedMetricDataDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RequestSummarizedMetricDataDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
