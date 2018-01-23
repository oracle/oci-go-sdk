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

// AttachVnicDetails The representation of AttachVnicDetails
type AttachVnicDetails struct {

	// Details for creating a new VNIC.
	CreateVnicDetails *CreateVnicDetails `mandatory:"true" json:"createVnicDetails,omitempty"`

	// The OCID of the instance.
	InstanceId *string `mandatory:"true" json:"instanceId,omitempty"`

	// A user-friendly name for the attachment. Does not have to be unique, and it cannot be changed.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// Which physical network interface card (NIC) the VNIC will use. Defaults to 0.
	// Certain bare metal instance shapes have two active physical NICs (0 and 1). If
	// you add a secondary VNIC to one of these instances, you can specify which NIC
	// the VNIC will use. For more information, see
	// https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingVNICs.htm.
	NicIndex *int `mandatory:"false" json:"nicIndex,omitempty"`
}

func (m AttachVnicDetails) String() string {
	return common.PointerString(m)
}
