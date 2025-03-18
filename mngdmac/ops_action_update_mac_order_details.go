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

// OpsActionUpdateMacOrderDetails The data to update an order in MAC_ORDER bucket.
type OpsActionUpdateMacOrderDetails struct {

	// The new status of the MacOrder.
	OrderStatus MacOrderOrderStatusEnum `mandatory:"false" json:"orderStatus,omitempty"`

	// The IP Range specified by the customer for this order.
	IpRange *string `mandatory:"false" json:"ipRange"`

	// The date and time this mac order is Active from. An RFC3339 formatted datetime string.
	TimeBillingStarted *common.SDKTime `mandatory:"false" json:"timeBillingStarted"`

	// The date and time this mac order until which this mac is Active. An RFC3339 formatted datetime string.
	TimeBillingEnded *common.SDKTime `mandatory:"false" json:"timeBillingEnded"`
}

func (m OpsActionUpdateMacOrderDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OpsActionUpdateMacOrderDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMacOrderOrderStatusEnum(string(m.OrderStatus)); !ok && m.OrderStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OrderStatus: %s. Supported values are: %s.", m.OrderStatus, strings.Join(GetMacOrderOrderStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
