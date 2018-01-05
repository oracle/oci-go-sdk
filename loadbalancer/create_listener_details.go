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

// CreateListenerDetails. The configuration details for adding a listener to a backend set.
// For more information on listener configuration, see
// [Managing Load Balancer Listeners]({{DOC_SERVER_URL}}/Content/Balance/tasks/managinglisteners.htm).
type CreateListenerDetails struct {

	// The name of the associated backend set.
	DefaultBackendSetName *string `mandatory:"true" json:"defaultBackendSetName,omitempty"`

	// A friendly name for the listener. It must be unique and it cannot be changed.
	// Avoid entering confidential information.
	// Example: `My listener`
	Name *string `mandatory:"true" json:"name,omitempty"`

	// The communication port for the listener.
	// Example: `80`
	Port *int `mandatory:"true" json:"port,omitempty"`

	// The protocol on which the listener accepts connection requests.
	// To get a list of valid protocols, use the ListProtocols
	// operation.
	// Example: `HTTP`
	Protocol *string `mandatory:"true" json:"protocol,omitempty"`

	// The name of the associated path-based routes applied to this listener's traffic.
	PathRouteSetName *string `mandatory:"false" json:"pathRouteSetName,omitempty"`

	// A HTTP hostname for this listener. TODO: This description needs work.
	// Example: `app.example.com`
	ServerName *string `mandatory:"false" json:"serverName,omitempty"`

	SslConfiguration *SslConfigurationDetails `mandatory:"false" json:"sslConfiguration,omitempty"`
}

func (model CreateListenerDetails) String() string {
	return common.PointerString(model)
}
