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

// VirtualCircuitPublicPrefix. A public IP prefix and its details. With a public virtual circuit, the customer
// specifies the customer-owned public IP prefixes to advertise across the connection.
// For more information, see [FastConnect Overview]({{DOC_SERVER_URL}}/Content/Network/Concepts/fastconnect.htm).
type VirtualCircuitPublicPrefix struct {

	// Publix IP prefix (CIDR) that the customer specified.
	CidrBlock *string `mandatory:"true" json:"cidrBlock,omitempty"`

	// Oracle must verify that the customer owns the public IP prefix before traffic
	// for that prefix can flow across the virtual circuit. Verification can take a
	// few business days. `IN_PROGRESS` means Oracle is verifying the prefix. `COMPLETED`
	// means verification succeeded. `FAILED` means verification failed and traffic for
	// this prefix will not flow across the connection.
	VerificationState VirtualCircuitPublicPrefixVerificationStateEnum `mandatory:"true" json:"verificationState,omitempty"`
}

func (model VirtualCircuitPublicPrefix) String() string {
	return common.PointerString(model)
}

type VirtualCircuitPublicPrefixVerificationStateEnum string

const (
	VIRTUAL_CIRCUIT_PUBLIC_PREFIX_VERIFICATION_STATE_IN_PROGRESS VirtualCircuitPublicPrefixVerificationStateEnum = "IN_PROGRESS"
	VIRTUAL_CIRCUIT_PUBLIC_PREFIX_VERIFICATION_STATE_COMPLETED   VirtualCircuitPublicPrefixVerificationStateEnum = "COMPLETED"
	VIRTUAL_CIRCUIT_PUBLIC_PREFIX_VERIFICATION_STATE_FAILED      VirtualCircuitPublicPrefixVerificationStateEnum = "FAILED"
	VIRTUAL_CIRCUIT_PUBLIC_PREFIX_VERIFICATION_STATE_UNKNOWN     VirtualCircuitPublicPrefixVerificationStateEnum = "UNKNOWN"
)

var mapping_virtualcircuitpublicprefix_verificationState = map[string]VirtualCircuitPublicPrefixVerificationStateEnum{
	"IN_PROGRESS": VIRTUAL_CIRCUIT_PUBLIC_PREFIX_VERIFICATION_STATE_IN_PROGRESS,
	"COMPLETED":   VIRTUAL_CIRCUIT_PUBLIC_PREFIX_VERIFICATION_STATE_COMPLETED,
	"FAILED":      VIRTUAL_CIRCUIT_PUBLIC_PREFIX_VERIFICATION_STATE_FAILED,
	"UNKNOWN":     VIRTUAL_CIRCUIT_PUBLIC_PREFIX_VERIFICATION_STATE_UNKNOWN,
}

func GetVirtualCircuitPublicPrefixVerificationStateEnumValues() []VirtualCircuitPublicPrefixVerificationStateEnum {
	values := make([]VirtualCircuitPublicPrefixVerificationStateEnum, 0)
	for _, v := range mapping_virtualcircuitpublicprefix_verificationState {
		if v != VIRTUAL_CIRCUIT_PUBLIC_PREFIX_VERIFICATION_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
