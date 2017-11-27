// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// IpSecConnection. A connection between a DRG and CPE. This connection consists of multiple IPSec
// tunnels. Creating this connection is one of the steps required when setting up
// an IPSec VPN. For more information, see
// [Overview of the Networking Service]({{DOC_SERVER_URL}}/Content/Network/Concepts/overview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type IpSecConnection struct {

	// The OCID of the compartment containing the IPSec connection.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The OCID of the CPE.
	CpeID *string `mandatory:"true" json:"cpeId,omitempty"`

	// The OCID of the DRG.
	DrgID *string `mandatory:"true" json:"drgId,omitempty"`

	// The IPSec connection's Oracle ID (OCID).
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The IPSec connection's current state.
	LifecycleState IpSecConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// Static routes to the CPE. At least one route must be included. The CIDR must not be a
	// multicast address or class E address.
	// Example: `10.0.1.0/24`
	StaticRoutes *[]string `mandatory:"true" json:"staticRoutes,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// The date and time the IPSec connection was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`
}

func (model IpSecConnection) String() string {
	return common.PointerString(model)
}

type IpSecConnectionLifecycleStateEnum string

const (
	IP_SEC_CONNECTION_LIFECYCLE_STATE_PROVISIONING IpSecConnectionLifecycleStateEnum = "PROVISIONING"
	IP_SEC_CONNECTION_LIFECYCLE_STATE_AVAILABLE    IpSecConnectionLifecycleStateEnum = "AVAILABLE"
	IP_SEC_CONNECTION_LIFECYCLE_STATE_TERMINATING  IpSecConnectionLifecycleStateEnum = "TERMINATING"
	IP_SEC_CONNECTION_LIFECYCLE_STATE_TERMINATED   IpSecConnectionLifecycleStateEnum = "TERMINATED"
	IP_SEC_CONNECTION_LIFECYCLE_STATE_UNKNOWN      IpSecConnectionLifecycleStateEnum = "UNKNOWN"
)

var mapping_ipsecconnection_lifecycleState = map[string]IpSecConnectionLifecycleStateEnum{
	"PROVISIONING": IP_SEC_CONNECTION_LIFECYCLE_STATE_PROVISIONING,
	"AVAILABLE":    IP_SEC_CONNECTION_LIFECYCLE_STATE_AVAILABLE,
	"TERMINATING":  IP_SEC_CONNECTION_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":   IP_SEC_CONNECTION_LIFECYCLE_STATE_TERMINATED,
	"UNKNOWN":      IP_SEC_CONNECTION_LIFECYCLE_STATE_UNKNOWN,
}

func GetIpSecConnectionLifecycleStateEnumValues() []IpSecConnectionLifecycleStateEnum {
	values := make([]IpSecConnectionLifecycleStateEnum, 0)
	for _, v := range mapping_ipsecconnection_lifecycleState {
		if v != IP_SEC_CONNECTION_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
