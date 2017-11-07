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

// UpdateListenerDetails. The configuration details for updating a listener.
type UpdateListenerDetails struct {

	// The name of the associated backend set.
	DefaultBackendSetName *string `mandatory:"true" json:"defaultBackendSetName,omitempty"`

	// The communication port for the listener.
	// Example: `80`
	Port *int `mandatory:"true" json:"port,omitempty"`

	// The protocol on which the listener accepts connection requests.
	// To get a list of valid protocols, use the ListProtocols
	// operation.
	// Example: `HTTP`
	Protocol *string `mandatory:"true" json:"protocol,omitempty"`

	SslConfiguration *SslConfigurationDetails `mandatory:"false" json:"sslConfiguration,omitempty"`
}

func (model UpdateListenerDetails) String() string {
	return common.PointerString(model)
}
