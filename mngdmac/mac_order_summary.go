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

// MacOrderSummary Summary information about a MacOrder.
type MacOrderSummary struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the compartment to which the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	OrderDescription *string `mandatory:"true" json:"orderDescription"`

	// Number of macs requested in this MacOrder.
	OrderSize *int `mandatory:"true" json:"orderSize"`

	// Checkbox value that indicates whether the customer completed docusign process.
	IsDocusigned *bool `mandatory:"true" json:"isDocusigned"`

	// The requested shape for Macs in this MacOrder.
	Shape MacOrderShapeEnum `mandatory:"true" json:"shape"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Enum indicating the agreed commitment term on the MacOrder.
	CommitmentTerm MacOrderCommitmentTermEnum `mandatory:"true" json:"commitmentTerm"`

	// The current status of the MacOrder.
	OrderStatus MacOrderOrderStatusEnum `mandatory:"true" json:"orderStatus"`

	// The current state of the MacOrder.
	LifecycleState MacOrderLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The IP Range specified by the customer for this order.
	IpRange *string `mandatory:"false" json:"ipRange"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// An RFC3339-formatted datetime string containing the date and time this MacOrder begins.
	TimeBillingStarted *common.SDKTime `mandatory:"false" json:"timeBillingStarted"`

	// An RFC3339-formatted datetime string containing the date and time this MacOrder begins.
	TimeBillingEnded *common.SDKTime `mandatory:"false" json:"timeBillingEnded"`

	// A message that describes the current state of the MacOrder in more detail.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m MacOrderSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MacOrderSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMacOrderShapeEnum(string(m.Shape)); !ok && m.Shape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shape: %s. Supported values are: %s.", m.Shape, strings.Join(GetMacOrderShapeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMacOrderCommitmentTermEnum(string(m.CommitmentTerm)); !ok && m.CommitmentTerm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CommitmentTerm: %s. Supported values are: %s.", m.CommitmentTerm, strings.Join(GetMacOrderCommitmentTermEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMacOrderOrderStatusEnum(string(m.OrderStatus)); !ok && m.OrderStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OrderStatus: %s. Supported values are: %s.", m.OrderStatus, strings.Join(GetMacOrderOrderStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMacOrderLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMacOrderLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
