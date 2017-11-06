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

// InternetGateway. Represents a router that connects the edge of a VCN with the Internet. For an example scenario
// that uses an Internet Gateway, see
// [Typical Networking Service Scenarios]({{DOC_SERVER_URL}}/Content/Network/Concepts/overview.htm#scenarios).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type InternetGateway struct {

	// The OCID of the compartment containing the Internet Gateway.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The Internet Gateway's Oracle ID (OCID).
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The Internet Gateway's current state.
	LifecycleState InternetGatewayLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The OCID of the VCN the Internet Gateway belongs to.
	VcnID *string `mandatory:"true" json:"vcnId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// Whether the gateway is enabled. When the gateway is disabled, traffic is not
	// routed to/from the Internet, regardless of route rules.
	IsEnabled *bool `mandatory:"false" json:"isEnabled,omitempty"`

	// The date and time the Internet Gateway was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`
}

func (model InternetGateway) String() string {
	return common.PointerString(model)
}

type InternetGatewayLifecycleStateEnum string

const (
	INTERNET_GATEWAY_LIFECYCLE_STATE_PROVISIONING InternetGatewayLifecycleStateEnum = "PROVISIONING"
	INTERNET_GATEWAY_LIFECYCLE_STATE_AVAILABLE    InternetGatewayLifecycleStateEnum = "AVAILABLE"
	INTERNET_GATEWAY_LIFECYCLE_STATE_TERMINATING  InternetGatewayLifecycleStateEnum = "TERMINATING"
	INTERNET_GATEWAY_LIFECYCLE_STATE_TERMINATED   InternetGatewayLifecycleStateEnum = "TERMINATED"
	INTERNET_GATEWAY_LIFECYCLE_STATE_UNKNOWN      InternetGatewayLifecycleStateEnum = "UNKNOWN"
)

var mapping_internetgateway_lifecycleState = map[string]InternetGatewayLifecycleStateEnum{
	"PROVISIONING": INTERNET_GATEWAY_LIFECYCLE_STATE_PROVISIONING,
	"AVAILABLE":    INTERNET_GATEWAY_LIFECYCLE_STATE_AVAILABLE,
	"TERMINATING":  INTERNET_GATEWAY_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":   INTERNET_GATEWAY_LIFECYCLE_STATE_TERMINATED,
	"UNKNOWN":      INTERNET_GATEWAY_LIFECYCLE_STATE_UNKNOWN,
}

func GetInternetGatewayLifecycleStateEnumValues() []InternetGatewayLifecycleStateEnum {
	values := make([]InternetGatewayLifecycleStateEnum, 0)
	for _, v := range mapping_internetgateway_lifecycleState {
		if v != INTERNET_GATEWAY_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
