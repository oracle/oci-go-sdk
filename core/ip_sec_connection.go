// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// IpSecConnection A connection between a DRG and CPE. This connection consists of multiple IPSec
// tunnels. Creating this connection is one of the steps required when setting up
// an IPSec VPN. For more information, see
// https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/overview.htm.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm.
type IpSecConnection struct {

	// The OCID of the compartment containing the IPSec connection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the CPE.
	CpeId *string `mandatory:"true" json:"cpeId"`

	// The OCID of the DRG.
	DrgId *string `mandatory:"true" json:"drgId"`

	// The IPSec connection's Oracle ID (OCID).
	Id *string `mandatory:"true" json:"id"`

	// The IPSec connection's current state.
	LifecycleState IpSecConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Static routes to the CPE. At least one route must be included. The CIDR must not be a
	// multicast address or class E address.
	// Example: `10.0.1.0/24`
	StaticRoutes []string `mandatory:"true" json:"staticRoutes"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The date and time the IPSec connection was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m IpSecConnection) String() string {
	return common.PointerString(m)
}

// IpSecConnectionLifecycleStateEnum Enum with underlying type: string
type IpSecConnectionLifecycleStateEnum string

// Set of constants representing the allowable values for IpSecConnectionLifecycleState
const (
	IpSecConnectionLifecycleStateProvisioning IpSecConnectionLifecycleStateEnum = "PROVISIONING"
	IpSecConnectionLifecycleStateAvailable    IpSecConnectionLifecycleStateEnum = "AVAILABLE"
	IpSecConnectionLifecycleStateTerminating  IpSecConnectionLifecycleStateEnum = "TERMINATING"
	IpSecConnectionLifecycleStateTerminated   IpSecConnectionLifecycleStateEnum = "TERMINATED"
	IpSecConnectionLifecycleStateUnknown      IpSecConnectionLifecycleStateEnum = "UNKNOWN"
)

var mappingIpSecConnectionLifecycleState = map[string]IpSecConnectionLifecycleStateEnum{
	"PROVISIONING": IpSecConnectionLifecycleStateProvisioning,
	"AVAILABLE":    IpSecConnectionLifecycleStateAvailable,
	"TERMINATING":  IpSecConnectionLifecycleStateTerminating,
	"TERMINATED":   IpSecConnectionLifecycleStateTerminated,
	"UNKNOWN":      IpSecConnectionLifecycleStateUnknown,
}

// GetIpSecConnectionLifecycleStateEnumValues Enumerates the set of values for IpSecConnectionLifecycleState
func GetIpSecConnectionLifecycleStateEnumValues() []IpSecConnectionLifecycleStateEnum {
	values := make([]IpSecConnectionLifecycleStateEnum, 0)
	for _, v := range mappingIpSecConnectionLifecycleState {
		if v != IpSecConnectionLifecycleStateUnknown {
			values = append(values, v)
		}
	}
	return values
}
