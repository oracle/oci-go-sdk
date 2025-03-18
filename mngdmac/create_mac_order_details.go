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

// CreateMacOrderDetails The data to create a new MacOrder.
type CreateMacOrderDetails struct {

	// OCID of the compartment to which the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	OrderDescription *string `mandatory:"true" json:"orderDescription"`

	// Number of macs requested in this MacOrder.
	OrderSize *int `mandatory:"true" json:"orderSize"`

	// The requested shape of the MacDevices in the MacOrder.
	Shape MacOrderShapeEnum `mandatory:"true" json:"shape"`

	// Enum that indicates the agreed commitment term for the MacDevices.
	CommitmentTerm MacOrderCommitmentTermEnum `mandatory:"true" json:"commitmentTerm"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The IP Range specified by the customer for this order.
	IpRange *string `mandatory:"false" json:"ipRange"`
}

func (m CreateMacOrderDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMacOrderDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMacOrderShapeEnum(string(m.Shape)); !ok && m.Shape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shape: %s. Supported values are: %s.", m.Shape, strings.Join(GetMacOrderShapeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMacOrderCommitmentTermEnum(string(m.CommitmentTerm)); !ok && m.CommitmentTerm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CommitmentTerm: %s. Supported values are: %s.", m.CommitmentTerm, strings.Join(GetMacOrderCommitmentTermEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
