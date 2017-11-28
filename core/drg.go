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

// Drg. A Dynamic Routing Gateway (DRG), which is a virtual router that provides a path for private
// network traffic between your VCN and your existing network. You use it with other Networking
// Service components to create an IPSec VPN or a connection that uses
// Oracle Cloud Infrastructure FastConnect. For more information, see
// [Overview of the Networking Service]({{DOC_SERVER_URL}}/Content/Network/Concepts/overview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type Drg struct {

	// The OCID of the compartment containing the DRG.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The DRG's Oracle ID (OCID).
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The DRG's current state.
	LifecycleState DrgLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// The date and time the DRG was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`
}

func (model Drg) String() string {
	return common.PointerString(model)
}

type DrgLifecycleStateEnum string

const (
	DRG_LIFECYCLE_STATE_PROVISIONING DrgLifecycleStateEnum = "PROVISIONING"
	DRG_LIFECYCLE_STATE_AVAILABLE    DrgLifecycleStateEnum = "AVAILABLE"
	DRG_LIFECYCLE_STATE_TERMINATING  DrgLifecycleStateEnum = "TERMINATING"
	DRG_LIFECYCLE_STATE_TERMINATED   DrgLifecycleStateEnum = "TERMINATED"
	DRG_LIFECYCLE_STATE_UNKNOWN      DrgLifecycleStateEnum = "UNKNOWN"
)

var mapping_drg_lifecycleState = map[string]DrgLifecycleStateEnum{
	"PROVISIONING": DRG_LIFECYCLE_STATE_PROVISIONING,
	"AVAILABLE":    DRG_LIFECYCLE_STATE_AVAILABLE,
	"TERMINATING":  DRG_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":   DRG_LIFECYCLE_STATE_TERMINATED,
	"UNKNOWN":      DRG_LIFECYCLE_STATE_UNKNOWN,
}

func GetDrgLifecycleStateEnumValues() []DrgLifecycleStateEnum {
	values := make([]DrgLifecycleStateEnum, 0)
	for _, v := range mapping_drg_lifecycleState {
		if v != DRG_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
