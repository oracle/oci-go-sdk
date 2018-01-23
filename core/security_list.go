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

// SecurityList A set of virtual firewall rules for your VCN. Security lists are configured at the subnet
// level, but the rules are applied to the ingress and egress traffic for the individual instances
// in the subnet. The rules can be stateful or stateless. For more information, see
// https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/securitylists.htm.
// **Important:** Oracle Cloud Infrastructure Compute service images automatically include firewall rules (for example,
// Linux iptables, Windows firewall). If there are issues with some type of access to an instance,
// make sure both the security lists associated with the instance's subnet and the instance's
// firewall rules are set correctly.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm.
type SecurityList struct {

	// The OCID of the compartment containing the security list.
	CompartmentId *string `mandatory:"true" json:"compartmentId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName,omitempty"`

	// Rules for allowing egress IP packets.
	EgressSecurityRules []EgressSecurityRule `mandatory:"true" json:"egressSecurityRules,omitempty"`

	// The security list's Oracle Cloud ID (OCID).
	Id *string `mandatory:"true" json:"id,omitempty"`

	// Rules for allowing ingress IP packets.
	IngressSecurityRules []IngressSecurityRule `mandatory:"true" json:"ingressSecurityRules,omitempty"`

	// The security list's current state.
	LifecycleState SecurityListLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The date and time the security list was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The OCID of the VCN the security list belongs to.
	VcnId *string `mandatory:"true" json:"vcnId,omitempty"`
}

func (m SecurityList) String() string {
	return common.PointerString(m)
}

// SecurityListLifecycleStateEnum Enum with underlying type: string
type SecurityListLifecycleStateEnum string

// Set of constants representing the allowable values for SecurityListLifecycleState
const (
	SecurityListLifecycleStateProvisioning SecurityListLifecycleStateEnum = "PROVISIONING"
	SecurityListLifecycleStateAvailable    SecurityListLifecycleStateEnum = "AVAILABLE"
	SecurityListLifecycleStateTerminating  SecurityListLifecycleStateEnum = "TERMINATING"
	SecurityListLifecycleStateTerminated   SecurityListLifecycleStateEnum = "TERMINATED"
	SecurityListLifecycleStateUnknown      SecurityListLifecycleStateEnum = "UNKNOWN"
)

var mappingSecurityListLifecycleState = map[string]SecurityListLifecycleStateEnum{
	"PROVISIONING": SecurityListLifecycleStateProvisioning,
	"AVAILABLE":    SecurityListLifecycleStateAvailable,
	"TERMINATING":  SecurityListLifecycleStateTerminating,
	"TERMINATED":   SecurityListLifecycleStateTerminated,
	"UNKNOWN":      SecurityListLifecycleStateUnknown,
}

// GetSecurityListLifecycleStateEnumValues Enumerates the set of values for SecurityListLifecycleState
func GetSecurityListLifecycleStateEnumValues() []SecurityListLifecycleStateEnum {
	values := make([]SecurityListLifecycleStateEnum, 0)
	for _, v := range mappingSecurityListLifecycleState {
		if v != SecurityListLifecycleStateUnknown {
			values = append(values, v)
		}
	}
	return values
}
