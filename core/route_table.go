// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oci-go-sdk/common"
)

// RouteTable. A collection of `RouteRule` objects, which are used to route packets
// based on destination IP to a particular network entity. For more information, see
// [Overview of the Networking Service]({{DOC_SERVER_URL}}/Content/Network/Concepts/overview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type RouteTable struct {

	// The OCID of the compartment containing the route table.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The route table's Oracle ID (OCID).
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The route table's current state.
	LifecycleState RouteTableLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The collection of rules for routing destination IPs to network devices.
	RouteRules *[]RouteRule `mandatory:"true" json:"routeRules,omitempty"`

	// The OCID of the VCN the route table list belongs to.
	VcnID *string `mandatory:"true" json:"vcnId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// The date and time the route table was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`
}

func (model RouteTable) String() string {
	return common.PointerString(model)
}

type RouteTableLifecycleStateEnum string

const (
	ROUTE_TABLE_LIFECYCLE_STATE_PROVISIONING RouteTableLifecycleStateEnum = "PROVISIONING"
	ROUTE_TABLE_LIFECYCLE_STATE_AVAILABLE    RouteTableLifecycleStateEnum = "AVAILABLE"
	ROUTE_TABLE_LIFECYCLE_STATE_TERMINATING  RouteTableLifecycleStateEnum = "TERMINATING"
	ROUTE_TABLE_LIFECYCLE_STATE_TERMINATED   RouteTableLifecycleStateEnum = "TERMINATED"
	ROUTE_TABLE_LIFECYCLE_STATE_UNKNOWN      RouteTableLifecycleStateEnum = "UNKNOWN"
)

var mapping_routetable_lifecycleState = map[string]RouteTableLifecycleStateEnum{
	"PROVISIONING": ROUTE_TABLE_LIFECYCLE_STATE_PROVISIONING,
	"AVAILABLE":    ROUTE_TABLE_LIFECYCLE_STATE_AVAILABLE,
	"TERMINATING":  ROUTE_TABLE_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":   ROUTE_TABLE_LIFECYCLE_STATE_TERMINATED,
	"UNKNOWN":      ROUTE_TABLE_LIFECYCLE_STATE_UNKNOWN,
}

func GetRouteTableLifecycleStateEnumValues() []RouteTableLifecycleStateEnum {
	values := make([]RouteTableLifecycleStateEnum, 0)
	for _, v := range mapping_routetable_lifecycleState {
		if v != ROUTE_TABLE_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
