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

// SecurityList. A set of virtual firewall rules for your VCN. Security lists are configured at the subnet
// level, but the rules are applied to the ingress and egress traffic for the individual instances
// in the subnet. The rules can be stateful or stateless. For more information, see
// [Security Lists]({{DOC_SERVER_URL}}/Content/Network/Concepts/securitylists.htm).
// **Important:** Oracle Cloud Infrastructure Compute service images automatically include firewall rules (for example,
// Linux iptables, Windows firewall). If there are issues with some type of access to an instance,
// make sure both the security lists associated with the instance's subnet and the instance's
// firewall rules are set correctly.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type SecurityList struct {

	// The OCID of the compartment containing the security list.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName,omitempty"`

	// Rules for allowing egress IP packets.
	EgressSecurityRules *[]EgressSecurityRule `mandatory:"true" json:"egressSecurityRules,omitempty"`

	// The security list's Oracle Cloud ID (OCID).
	ID *string `mandatory:"true" json:"id,omitempty"`

	// Rules for allowing ingress IP packets.
	IngressSecurityRules *[]IngressSecurityRule `mandatory:"true" json:"ingressSecurityRules,omitempty"`

	// The security list's current state.
	LifecycleState SecurityListLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The date and time the security list was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The OCID of the VCN the security list belongs to.
	VcnID *string `mandatory:"true" json:"vcnId,omitempty"`
}

func (model SecurityList) String() string {
	return common.PointerString(model)
}

type SecurityListLifecycleStateEnum string

const (
	SECURITY_LIST_LIFECYCLE_STATE_PROVISIONING SecurityListLifecycleStateEnum = "PROVISIONING"
	SECURITY_LIST_LIFECYCLE_STATE_AVAILABLE    SecurityListLifecycleStateEnum = "AVAILABLE"
	SECURITY_LIST_LIFECYCLE_STATE_TERMINATING  SecurityListLifecycleStateEnum = "TERMINATING"
	SECURITY_LIST_LIFECYCLE_STATE_TERMINATED   SecurityListLifecycleStateEnum = "TERMINATED"
	SECURITY_LIST_LIFECYCLE_STATE_UNKNOWN      SecurityListLifecycleStateEnum = "UNKNOWN"
)

var mapping_securitylist_lifecycleState = map[string]SecurityListLifecycleStateEnum{
	"PROVISIONING": SECURITY_LIST_LIFECYCLE_STATE_PROVISIONING,
	"AVAILABLE":    SECURITY_LIST_LIFECYCLE_STATE_AVAILABLE,
	"TERMINATING":  SECURITY_LIST_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":   SECURITY_LIST_LIFECYCLE_STATE_TERMINATED,
	"UNKNOWN":      SECURITY_LIST_LIFECYCLE_STATE_UNKNOWN,
}

func GetSecurityListLifecycleStateEnumValues() []SecurityListLifecycleStateEnum {
	values := make([]SecurityListLifecycleStateEnum, 0)
	for _, v := range mapping_securitylist_lifecycleState {
		if v != SECURITY_LIST_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
