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

// Listener. The listener's configuration.
// For more information on backend set configuration, see
// [Managing Load Balancer Listeners]({{DOC_SERVER_URL}}/Content/Balance/tasks/managinglisteners.htm).
type Listener struct {

	// The name of the associated backend set.
	DefaultBackendSetName *string `mandatory:"true" json:"defaultBackendSetName,omitempty"`

	// A friendly name for the listener. It must be unique and it cannot be changed.
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

	SslConfiguration *SslConfiguration `mandatory:"false" json:"sslConfiguration,omitempty"`
}

func (model Listener) String() string {
	return common.PointerString(model)
}
