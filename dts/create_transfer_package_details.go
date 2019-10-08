// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateTransferPackageDetails The representation of CreateTransferPackageDetails
type CreateTransferPackageDetails struct {
	OriginalPackageDeliveryTrackingNumber *string `mandatory:"false" json:"originalPackageDeliveryTrackingNumber"`

	ReturnPackageDeliveryTrackingNumber *string `mandatory:"false" json:"returnPackageDeliveryTrackingNumber"`

	PackageDeliveryVendor *string `mandatory:"false" json:"packageDeliveryVendor"`
}

func (m CreateTransferPackageDetails) String() string {
	return common.PointerString(m)
}
