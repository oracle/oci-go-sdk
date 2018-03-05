// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateListenerDetails The configuration details for updating a listener.
type UpdateListenerDetails struct {

	// The name of the associated backend set.
	DefaultBackendSetName *string `mandatory:"true" json:"defaultBackendSetName"`

	// The communication port for the listener.
	// Example: `80`
	Port *int `mandatory:"true" json:"port"`

	// The protocol on which the listener accepts connection requests.
	// To get a list of valid protocols, use the ListProtocols
	// operation.
	// Example: `HTTP`
	Protocol *string `mandatory:"true" json:"protocol"`

	// The name of the associated path-based routes applied to this listener's traffic.
	PathRouteSetName *string `mandatory:"false" json:"pathRouteSetName"`

	// A HTTP hostname for this listener. TODO: This description needs work.
	// Example: `app.example.com`
	ServerName *string `mandatory:"false" json:"serverName"`

	SslConfiguration *SslConfigurationDetails `mandatory:"false" json:"sslConfiguration"`
}

func (m UpdateListenerDetails) String() string {
	return common.PointerString(m)
}
