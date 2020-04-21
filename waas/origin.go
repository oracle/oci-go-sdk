// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Origin A detailed description of your web application's origin host server. An origin must be defined to set up WAF rules.
type Origin struct {

	// The URI of the origin. Does not support paths. Port numbers should be specified in the `httpPort` and `httpsPort` fields.
	Uri *string `mandatory:"true" json:"uri"`

	// The HTTP port on the origin that the web application listens on. If unspecified, defaults to `80`.
	HttpPort *int `mandatory:"false" json:"httpPort"`

	// The HTTPS port on the origin that the web application listens on. If unspecified, defaults to `443`.
	HttpsPort *int `mandatory:"false" json:"httpsPort"`

	// A list of HTTP headers to forward to your origin.
	CustomHeaders []Header `mandatory:"false" json:"customHeaders"`
}

func (m Origin) String() string {
	return common.PointerString(m)
}
