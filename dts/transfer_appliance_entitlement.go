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

// TransferApplianceEntitlement The representation of TransferApplianceEntitlement
type TransferApplianceEntitlement struct {
	TenantId *string `mandatory:"true" json:"tenantId"`

	Status TransferApplianceEntitlementStatusEnum `mandatory:"true" json:"status"`

	RequestorName *string `mandatory:"false" json:"requestorName"`

	RequestorEmail *string `mandatory:"false" json:"requestorEmail"`

	CreationTime *common.SDKTime `mandatory:"false" json:"creationTime"`

	UpdateTime *common.SDKTime `mandatory:"false" json:"updateTime"`
}

func (m TransferApplianceEntitlement) String() string {
	return common.PointerString(m)
}

// TransferApplianceEntitlementStatusEnum Enum with underlying type: string
type TransferApplianceEntitlementStatusEnum string

// Set of constants representing the allowable values for TransferApplianceEntitlementStatusEnum
const (
	TransferApplianceEntitlementStatusRequested       TransferApplianceEntitlementStatusEnum = "REQUESTED"
	TransferApplianceEntitlementStatusPendingSigning  TransferApplianceEntitlementStatusEnum = "PENDING_SIGNING"
	TransferApplianceEntitlementStatusPendingApproval TransferApplianceEntitlementStatusEnum = "PENDING_APPROVAL"
	TransferApplianceEntitlementStatusTermsExpired    TransferApplianceEntitlementStatusEnum = "TERMS_EXPIRED"
	TransferApplianceEntitlementStatusApproved        TransferApplianceEntitlementStatusEnum = "APPROVED"
	TransferApplianceEntitlementStatusRejected        TransferApplianceEntitlementStatusEnum = "REJECTED"
	TransferApplianceEntitlementStatusCancelled       TransferApplianceEntitlementStatusEnum = "CANCELLED"
)

var mappingTransferApplianceEntitlementStatus = map[string]TransferApplianceEntitlementStatusEnum{
	"REQUESTED":        TransferApplianceEntitlementStatusRequested,
	"PENDING_SIGNING":  TransferApplianceEntitlementStatusPendingSigning,
	"PENDING_APPROVAL": TransferApplianceEntitlementStatusPendingApproval,
	"TERMS_EXPIRED":    TransferApplianceEntitlementStatusTermsExpired,
	"APPROVED":         TransferApplianceEntitlementStatusApproved,
	"REJECTED":         TransferApplianceEntitlementStatusRejected,
	"CANCELLED":        TransferApplianceEntitlementStatusCancelled,
}

// GetTransferApplianceEntitlementStatusEnumValues Enumerates the set of values for TransferApplianceEntitlementStatusEnum
func GetTransferApplianceEntitlementStatusEnumValues() []TransferApplianceEntitlementStatusEnum {
	values := make([]TransferApplianceEntitlementStatusEnum, 0)
	for _, v := range mappingTransferApplianceEntitlementStatus {
		values = append(values, v)
	}
	return values
}
