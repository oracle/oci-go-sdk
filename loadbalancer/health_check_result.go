// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// HealthCheckResult. Information about a single backend server health check result reported by a load balancer.
type HealthCheckResult struct {

	// The result of the most recent health check.
	HealthCheckStatus HealthCheckResultHealthCheckStatusEnum `mandatory:"true" json:"healthCheckStatus,omitempty"`

	// The IP address of the health check status report provider. This identifier helps you differentiate same-subnet
	// (private) load balancers that report health check status.
	// Example: `10.2.0.1`
	SourceIpAddress *string `mandatory:"true" json:"sourceIpAddress,omitempty"`

	// The OCID of the subnet hosting the load balancer that reported this health check status.
	SubnetID *string `mandatory:"true" json:"subnetId,omitempty"`

	// The date and time the data was retrieved, in the format defined by RFC3339.
	// Example: `2017-06-02T18:28:11+00:00`
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp,omitempty"`
}

func (model HealthCheckResult) String() string {
	return common.PointerString(model)
}

type HealthCheckResultHealthCheckStatusEnum string

const (
	HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_OK                  HealthCheckResultHealthCheckStatusEnum = "OK"
	HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_INVALID_STATUS_CODE HealthCheckResultHealthCheckStatusEnum = "INVALID_STATUS_CODE"
	HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_TIMED_OUT           HealthCheckResultHealthCheckStatusEnum = "TIMED_OUT"
	HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_REGEX_MISMATCH      HealthCheckResultHealthCheckStatusEnum = "REGEX_MISMATCH"
	HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_CONNECT_FAILED      HealthCheckResultHealthCheckStatusEnum = "CONNECT_FAILED"
	HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_IO_ERROR            HealthCheckResultHealthCheckStatusEnum = "IO_ERROR"
	HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_OFFLINE             HealthCheckResultHealthCheckStatusEnum = "OFFLINE"
	HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_UNKNOWN             HealthCheckResultHealthCheckStatusEnum = "UNKNOWN"
	HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_UNKNOWN             HealthCheckResultHealthCheckStatusEnum = "UNKNOWN"
)

var mapping_healthcheckresult_healthCheckStatus = map[string]HealthCheckResultHealthCheckStatusEnum{
	"OK": HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_OK,
	"INVALID_STATUS_CODE": HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_INVALID_STATUS_CODE,
	"TIMED_OUT":           HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_TIMED_OUT,
	"REGEX_MISMATCH":      HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_REGEX_MISMATCH,
	"CONNECT_FAILED":      HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_CONNECT_FAILED,
	"IO_ERROR":            HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_IO_ERROR,
	"OFFLINE":             HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_OFFLINE,
	"UNKNOWN":             HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_UNKNOWN,
	"UNKNOWN":             HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_UNKNOWN,
}

func GetHealthCheckResultHealthCheckStatusEnumValues() []HealthCheckResultHealthCheckStatusEnum {
	values := make([]HealthCheckResultHealthCheckStatusEnum, 0)
	for _, v := range mapping_healthcheckresult_healthCheckStatus {
		if v != HEALTH_CHECK_RESULT_HEALTH_CHECK_STATUS_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
