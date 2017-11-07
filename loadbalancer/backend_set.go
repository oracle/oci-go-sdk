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

// BackendSet. The configuration of a load balancer backend set.
// For more information on backend set configuration, see
// [Managing Backend Sets]({{DOC_SERVER_URL}}/Content/Balance/tasks/managingbackendsets.htm).
type BackendSet struct {
	Backends *[]Backend `mandatory:"true" json:"backends,omitempty"`

	HealthChecker *HealthChecker `mandatory:"true" json:"healthChecker,omitempty"`

	// A friendly name for the backend set. It must be unique and it cannot be changed.
	// Valid backend set names include only alphanumeric characters, dashes, and underscores. Backend set names cannot
	// contain spaces. Avoid entering confidential information.
	// Example: `My_backend_set`
	Name *string `mandatory:"true" json:"name,omitempty"`

	// The load balancer policy for the backend set. To get a list of available policies, use the
	// ListPolicies operation.
	// Example: `LEAST_CONNECTIONS`
	Policy *string `mandatory:"true" json:"policy,omitempty"`

	SessionPersistenceConfiguration *SessionPersistenceConfigurationDetails `mandatory:"false" json:"sessionPersistenceConfiguration,omitempty"`

	SslConfiguration *SslConfiguration `mandatory:"false" json:"sslConfiguration,omitempty"`
}

func (model BackendSet) String() string {
	return common.PointerString(model)
}
