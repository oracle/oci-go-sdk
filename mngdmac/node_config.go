// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Managed Services for Mac API
//
// Use the OCI Managed Services for Mac API to create and manage orders for dedicated, partially-managed Mac servers hosted in an OCI Data Center. For more information, see the user guide documentation for the [OCI Managed Services for Mac]()
//

package mngdmac

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NodeConfig Represents a MacDevice resource.
type NodeConfig struct {

	// The serial number of the MacDevice.
	SerialNumber *string `mandatory:"true" json:"serialNumber"`

	// The macAddress.
	MacAddress *string `mandatory:"true" json:"macAddress"`

	// The switchHostname.
	SwitchHostname *string `mandatory:"true" json:"switchHostname"`

	// The switchEthPort.
	SwitchEthPort *string `mandatory:"true" json:"switchEthPort"`

	// The ipKvmHostname.
	IpKvmHostname *string `mandatory:"true" json:"ipKvmHostname"`

	// The ipKvmPortNumber.
	IpKvmPortNumber *int `mandatory:"true" json:"ipKvmPortNumber"`

	// The pduHostname.
	PduHostname *string `mandatory:"true" json:"pduHostname"`

	// The pduPort.
	PduPort *int `mandatory:"true" json:"pduPort"`

	// The buildVlanId.
	BuildVlanId *int `mandatory:"true" json:"buildVlanId"`

	// The buildIpAddress.
	BuildIpAddress *string `mandatory:"true" json:"buildIpAddress"`

	// The prodVlanId.
	ProdVlanId *int `mandatory:"true" json:"prodVlanId"`

	// The prodIpAddress.
	ProdIpAddress *string `mandatory:"true" json:"prodIpAddress"`

	// The rackLocation.
	RackLocation *string `mandatory:"true" json:"rackLocation"`

	// The chipSetn.
	ChipSet *string `mandatory:"true" json:"chipSet"`

	// The osVersion.
	OsVersion *string `mandatory:"true" json:"osVersion"`

	// The macOrderId.
	MacOrderId *string `mandatory:"true" json:"macOrderId"`

	// The tenancyId.
	TenancyId *string `mandatory:"true" json:"tenancyId"`
}

func (m NodeConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NodeConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
