// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DTS API
//
// A description of the DTS API
//

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateTransferApplianceEntitlementDetails The representation of CreateTransferApplianceEntitlementDetails
type CreateTransferApplianceEntitlementDetails struct {
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	RequestorName *string `mandatory:"false" json:"requestorName"`

	RequestorEmail *string `mandatory:"false" json:"requestorEmail"`
}

func (m CreateTransferApplianceEntitlementDetails) String() string {
	return common.PointerString(m)
}
