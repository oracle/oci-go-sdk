// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Health Checks API
//
// API for the Health Checks service. Use this API to manage endpoint probes and monitors.
// For more information, see
// Overview of the Health Checks Service (https://docs.cloud.oracle.com/iaas/Content/HealthChecks/Concepts/healthchecks.htm).
//

package healthchecks

import (
	"github.com/oracle/oci-go-sdk/common"
)

// PingProbe This model contains all of the mutable and immutable properties for a ping probe.
type PingProbe struct {

	// The OCID of the resource.
	Id *string `mandatory:"false" json:"id"`

	// A URL for fetching the probe results.
	ResultsUrl *string `mandatory:"false" json:"resultsUrl"`

	// The region where updates must be made and where results must be fetched from.
	HomeRegion *string `mandatory:"false" json:"homeRegion"`

	// The RFC 3339-formatted creation date and time of the probe.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// A list of targets (hostnames or IP addresses) of the probe.
	Targets []string `mandatory:"false" json:"targets"`

	// A list of names of vantage points from which to execute the probe.
	VantagePointNames []string `mandatory:"false" json:"vantagePointNames"`

	// The port on which to probe endpoints. If unspecified, probes will use the
	// default port of their protocol.
	Port *int `mandatory:"false" json:"port"`

	// The probe timeout in seconds. Valid values: 10, 20, 30, and 60.
	// The probe timeout must be less than or equal to `intervalInSeconds` for monitors.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	Protocol PingProbeProtocolEnum `mandatory:"false" json:"protocol,omitempty"`
}

func (m PingProbe) String() string {
	return common.PointerString(m)
}
