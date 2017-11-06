// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// CrossConnectGroup. For use with Oracle Cloud Infrastructure FastConnect. A cross-connect group
// is a link aggregation group (LAG), which can contain one or more
// CrossConnect. Customers who are colocated with
// Oracle in a FastConnect location create and use cross-connect groups. For more
// information, see [FastConnect Overview]({{DOC_SERVER_URL}}/Content/Network/Concepts/fastconnect.htm).
// **Note:** If you're a provider who is setting up a physical connection to Oracle so customers
// can use FastConnect over the connection, be aware that your connection is modeled the
// same way as a colocated customer's (with `CrossConnect` and `CrossConnectGroup` objects, and so on).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type CrossConnectGroup struct {

	// The OCID of the compartment containing the cross-connect group.
	CompartmentID *string `mandatory:"false" json:"compartmentId,omitempty"`

	// The display name of A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// The cross-connect group's Oracle ID (OCID).
	ID *string `mandatory:"false" json:"id,omitempty"`

	// The cross-connect group's current state.
	LifecycleState CrossConnectGroupLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the cross-connect group was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`
}

func (model CrossConnectGroup) String() string {
	return common.PointerString(model)
}

type CrossConnectGroupLifecycleStateEnum string

const (
	CROSS_CONNECT_GROUP_LIFECYCLE_STATE_PROVISIONING CrossConnectGroupLifecycleStateEnum = "PROVISIONING"
	CROSS_CONNECT_GROUP_LIFECYCLE_STATE_PROVISIONED  CrossConnectGroupLifecycleStateEnum = "PROVISIONED"
	CROSS_CONNECT_GROUP_LIFECYCLE_STATE_INACTIVE     CrossConnectGroupLifecycleStateEnum = "INACTIVE"
	CROSS_CONNECT_GROUP_LIFECYCLE_STATE_TERMINATING  CrossConnectGroupLifecycleStateEnum = "TERMINATING"
	CROSS_CONNECT_GROUP_LIFECYCLE_STATE_TERMINATED   CrossConnectGroupLifecycleStateEnum = "TERMINATED"
	CROSS_CONNECT_GROUP_LIFECYCLE_STATE_UNKNOWN      CrossConnectGroupLifecycleStateEnum = "UNKNOWN"
)

var mapping_crossconnectgroup_lifecycleState = map[string]CrossConnectGroupLifecycleStateEnum{
	"PROVISIONING": CROSS_CONNECT_GROUP_LIFECYCLE_STATE_PROVISIONING,
	"PROVISIONED":  CROSS_CONNECT_GROUP_LIFECYCLE_STATE_PROVISIONED,
	"INACTIVE":     CROSS_CONNECT_GROUP_LIFECYCLE_STATE_INACTIVE,
	"TERMINATING":  CROSS_CONNECT_GROUP_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":   CROSS_CONNECT_GROUP_LIFECYCLE_STATE_TERMINATED,
	"UNKNOWN":      CROSS_CONNECT_GROUP_LIFECYCLE_STATE_UNKNOWN,
}

func GetCrossConnectGroupLifecycleStateEnumValues() []CrossConnectGroupLifecycleStateEnum {
	values := make([]CrossConnectGroupLifecycleStateEnum, 0)
	for _, v := range mapping_crossconnectgroup_lifecycleState {
		if v != CROSS_CONNECT_GROUP_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
