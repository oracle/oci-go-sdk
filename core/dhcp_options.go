// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// DhcpOptions. A set of DHCP options. Used by the VCN to automatically provide configuration
// information to the instances when they boot up. There are two options you can set:
// - DhcpDnsOption: Lets you specify how DNS (hostname resolution) is
// handled in the subnets in your VCN.
// - DhcpSearchDomainOption: Lets you specify
// a search domain name to use for DNS queries.
// For more information, see  [DNS in Your Virtual Cloud Network]({{DOC_SERVER_URL}}/Content/Network/Concepts/dns.htm)
// and [DHCP Options]({{DOC_SERVER_URL}}/Content/Network/Tasks/managingDHCP.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type DhcpOptions struct {

	// The OCID of the compartment containing the set of DHCP options.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// Oracle ID (OCID) for the set of DHCP options.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The current state of the set of DHCP options.
	LifecycleState DhcpOptionsLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The collection of individual DHCP options.
	Options []DhcpOption `mandatory:"true" json:"options,omitempty"`

	// Date and time the set of DHCP options was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The OCID of the VCN the set of DHCP options belongs to.
	VcnID *string `mandatory:"true" json:"vcnId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
}

func (model DhcpOptions) String() string {
	return common.PointerString(model)
}

func (model *DhcpOptions) UnmarshalJSON(data []byte) (e error) {
	m := struct {
		DisplayName    *string                       `mandatory:"true" json:"displayName,omitempty"`
		CompartmentID  *string                       `mandatory:"true" json:"compartmentId,omitempty"`
		ID             *string                       `mandatory:"true" json:"id,omitempty"`
		LifecycleState DhcpOptionsLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`
		Options        []dhcpoption                  `mandatory:"true" json:"options,omitempty"`
		TimeCreated    *common.SDKTime               `mandatory:"true" json:"timeCreated,omitempty"`
		VcnID          *string                       `mandatory:"true" json:"vcnId,omitempty"`
	}{}

	e = json.Unmarshal(data, &m)
	if e != nil {
		return
	}
	model.DisplayName = m.DisplayName
	model.CompartmentID = m.CompartmentID
	model.ID = m.ID
	model.LifecycleState = m.LifecycleState
	model.Options = make([]DhcpOption, len(m.Options))
	for i, n := range m.Options {
		nn, err := n.UnmarshalPolymorphicJSON(n.JsonData)
		if err != nil {
			return err
		}
		model.Options[i] = nn
	}
	model.TimeCreated = m.TimeCreated
	model.VcnID = m.VcnID
	return
}

type DhcpOptionsLifecycleStateEnum string

const (
	DHCP_OPTIONS_LIFECYCLE_STATE_PROVISIONING DhcpOptionsLifecycleStateEnum = "PROVISIONING"
	DHCP_OPTIONS_LIFECYCLE_STATE_AVAILABLE    DhcpOptionsLifecycleStateEnum = "AVAILABLE"
	DHCP_OPTIONS_LIFECYCLE_STATE_TERMINATING  DhcpOptionsLifecycleStateEnum = "TERMINATING"
	DHCP_OPTIONS_LIFECYCLE_STATE_TERMINATED   DhcpOptionsLifecycleStateEnum = "TERMINATED"
	DHCP_OPTIONS_LIFECYCLE_STATE_UNKNOWN      DhcpOptionsLifecycleStateEnum = "UNKNOWN"
)

var mapping_dhcpoptions_lifecycleState = map[string]DhcpOptionsLifecycleStateEnum{
	"PROVISIONING": DHCP_OPTIONS_LIFECYCLE_STATE_PROVISIONING,
	"AVAILABLE":    DHCP_OPTIONS_LIFECYCLE_STATE_AVAILABLE,
	"TERMINATING":  DHCP_OPTIONS_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":   DHCP_OPTIONS_LIFECYCLE_STATE_TERMINATED,
	"UNKNOWN":      DHCP_OPTIONS_LIFECYCLE_STATE_UNKNOWN,
}

func GetDhcpOptionsLifecycleStateEnumValues() []DhcpOptionsLifecycleStateEnum {
	values := make([]DhcpOptionsLifecycleStateEnum, 0)
	for _, v := range mapping_dhcpoptions_lifecycleState {
		if v != DHCP_OPTIONS_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
