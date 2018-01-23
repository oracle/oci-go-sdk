// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateBackendSetDetails The configuration details for updating a load balancer backend set.
// For more information on backend set configuration, see
// https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/tasks/managingbackendsets.htm.
type UpdateBackendSetDetails struct {
	Backends []BackendDetails `mandatory:"true" json:"backends,omitempty"`

	HealthChecker *HealthCheckerDetails `mandatory:"true" json:"healthChecker,omitempty"`

	// The load balancer policy for the backend set. To get a list of available policies, use the
	// ListPolicies operation.
	// Example: `LEAST_CONNECTIONS`
	Policy *string `mandatory:"true" json:"policy,omitempty"`

	SessionPersistenceConfiguration *SessionPersistenceConfigurationDetails `mandatory:"false" json:"sessionPersistenceConfiguration,omitempty"`

	SslConfiguration *SslConfigurationDetails `mandatory:"false" json:"sslConfiguration,omitempty"`
}

func (m UpdateBackendSetDetails) String() string {
	return common.PointerString(m)
}
