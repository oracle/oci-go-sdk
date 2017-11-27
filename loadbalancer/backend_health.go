// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"github.com/oci-go-sdk/common"
)

// BackendHealth. The health status of the specified backend server as reported by the primary and standby load balancers.
type BackendHealth struct {

	// A list of the most recent health check results returned for the specified backend server.
	HealthCheckResults *[]HealthCheckResult `mandatory:"true" json:"healthCheckResults,omitempty"`

	// The general health status of the specified backend server as reported by the primary and standby load balancers.
	// *   **OK:** Both health checks returned `OK`.
	// *   **WARNING:** One health check returned `OK` and one did not.
	// *   **CRITICAL:** Neither health check returned `OK`.
	// *   **UNKNOWN:** One or both health checks returned `UNKNOWN`, or the system was unable to retrieve metrics at this time.
	Status BackendHealthStatusEnum `mandatory:"true" json:"status,omitempty"`
}

func (model BackendHealth) String() string {
	return common.PointerString(model)
}

type BackendHealthStatusEnum string

const (
	BACKEND_HEALTH_STATUS_OK       BackendHealthStatusEnum = "OK"
	BACKEND_HEALTH_STATUS_WARNING  BackendHealthStatusEnum = "WARNING"
	BACKEND_HEALTH_STATUS_CRITICAL BackendHealthStatusEnum = "CRITICAL"
	BACKEND_HEALTH_STATUS_UNKNOWN  BackendHealthStatusEnum = "UNKNOWN"
)

var mapping_backendhealth_status = map[string]BackendHealthStatusEnum{
	"OK":       BACKEND_HEALTH_STATUS_OK,
	"WARNING":  BACKEND_HEALTH_STATUS_WARNING,
	"CRITICAL": BACKEND_HEALTH_STATUS_CRITICAL,
	"UNKNOWN":  BACKEND_HEALTH_STATUS_UNKNOWN,
}

func GetBackendHealthStatusEnumValues() []BackendHealthStatusEnum {
	values := make([]BackendHealthStatusEnum, 0)
	for _, v := range mapping_backendhealth_status {
		values = append(values, v)
	}
	return values
}
