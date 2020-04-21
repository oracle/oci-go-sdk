// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ConnectionConfiguration Configuration details for the connection between the client and backend servers.
type ConnectionConfiguration struct {

	// The maximum idle time, in seconds, allowed between two successive receive or two successive send operations
	// between the client and backend servers. A send operation does not reset the timer for receive operations. A
	// receive operation does not reset the timer for send operations.
	// For more information, see Connection Configuration (https://docs.cloud.oracle.com/Content/Balance/Reference/connectionreuse.htm#ConnectionConfiguration).
	// Example: `1200`
	IdleTimeout *int64 `mandatory:"true" json:"idleTimeout"`

	// The backend TCP Proxy Protocol version.
	// Example: `1`
	BackendTcpProxyProtocolVersion *int `mandatory:"false" json:"backendTcpProxyProtocolVersion"`
}

func (m ConnectionConfiguration) String() string {
	return common.PointerString(m)
}
