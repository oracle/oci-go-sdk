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

// FastConnectProviderService A service offering from a supported provider. For more information,
// see [FastConnect Overview]({{DOC_SERVER_URL}}/Content/Network/Concepts/fastconnect.htm).
type FastConnectProviderService struct {

	// The OCID of the service offered by the provider.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// Private peering BGP management.
	PrivatePeeringBgpManagement FastConnectProviderServicePrivatePeeringBgpManagementEnum `mandatory:"true" json:"privatePeeringBgpManagement,omitempty"`

	// The name of the provider.
	ProviderName *string `mandatory:"true" json:"providerName,omitempty"`

	// The name of the service offered by the provider.
	ProviderServiceName *string `mandatory:"true" json:"providerServiceName,omitempty"`

	// Public peering BGP management.
	PublicPeeringBgpManagement FastConnectProviderServicePublicPeeringBgpManagementEnum `mandatory:"true" json:"publicPeeringBgpManagement,omitempty"`

	// Provider service type.
	Type_ FastConnectProviderServiceType_Enum `mandatory:"true" json:"type,omitempty"`

	// A description of the service offered by the provider.
	Description *string `mandatory:"false" json:"description,omitempty"`

	// An array of virtual circuit types supported by this service.
	SupportedVirtualCircuitTypes []FastConnectProviderServiceSupportedVirtualCircuitTypesEnum `mandatory:"false" json:"supportedVirtualCircuitTypes,omitempty"`
}

func (model FastConnectProviderService) String() string {
	return common.PointerString(model)
}

// FastConnectProviderServicePrivatePeeringBgpManagementEnum Enum with underlying type: string
type FastConnectProviderServicePrivatePeeringBgpManagementEnum string

// Set of constants representing the allowable values for FastConnectProviderServicePrivatePeeringBgpManagement
const (
	FastConnectProviderServicePrivatePeeringBgpManagementCustomerManaged FastConnectProviderServicePrivatePeeringBgpManagementEnum = "CUSTOMER_MANAGED"
	FastConnectProviderServicePrivatePeeringBgpManagementProviderManaged FastConnectProviderServicePrivatePeeringBgpManagementEnum = "PROVIDER_MANAGED"
	FastConnectProviderServicePrivatePeeringBgpManagementOracleManaged   FastConnectProviderServicePrivatePeeringBgpManagementEnum = "ORACLE_MANAGED"
	FastConnectProviderServicePrivatePeeringBgpManagementUnknown         FastConnectProviderServicePrivatePeeringBgpManagementEnum = "UNKNOWN"
)

var mappingFastConnectProviderServicePrivatePeeringBgpManagement = map[string]FastConnectProviderServicePrivatePeeringBgpManagementEnum{
	"CUSTOMER_MANAGED": FastConnectProviderServicePrivatePeeringBgpManagementCustomerManaged,
	"PROVIDER_MANAGED": FastConnectProviderServicePrivatePeeringBgpManagementProviderManaged,
	"ORACLE_MANAGED":   FastConnectProviderServicePrivatePeeringBgpManagementOracleManaged,
	"UNKNOWN":          FastConnectProviderServicePrivatePeeringBgpManagementUnknown,
}

// GetFastConnectProviderServicePrivatePeeringBgpManagementEnumValues Enumerates the set of values for FastConnectProviderServicePrivatePeeringBgpManagement
func GetFastConnectProviderServicePrivatePeeringBgpManagementEnumValues() []FastConnectProviderServicePrivatePeeringBgpManagementEnum {
	values := make([]FastConnectProviderServicePrivatePeeringBgpManagementEnum, 0)
	for _, v := range mappingFastConnectProviderServicePrivatePeeringBgpManagement {
		if v != FastConnectProviderServicePrivatePeeringBgpManagementUnknown {
			values = append(values, v)
		}
	}
	return values
}

// FastConnectProviderServicePublicPeeringBgpManagementEnum Enum with underlying type: string
type FastConnectProviderServicePublicPeeringBgpManagementEnum string

// Set of constants representing the allowable values for FastConnectProviderServicePublicPeeringBgpManagement
const (
	FastConnectProviderServicePublicPeeringBgpManagementCustomerManaged FastConnectProviderServicePublicPeeringBgpManagementEnum = "CUSTOMER_MANAGED"
	FastConnectProviderServicePublicPeeringBgpManagementProviderManaged FastConnectProviderServicePublicPeeringBgpManagementEnum = "PROVIDER_MANAGED"
	FastConnectProviderServicePublicPeeringBgpManagementOracleManaged   FastConnectProviderServicePublicPeeringBgpManagementEnum = "ORACLE_MANAGED"
	FastConnectProviderServicePublicPeeringBgpManagementUnknown         FastConnectProviderServicePublicPeeringBgpManagementEnum = "UNKNOWN"
)

var mappingFastConnectProviderServicePublicPeeringBgpManagement = map[string]FastConnectProviderServicePublicPeeringBgpManagementEnum{
	"CUSTOMER_MANAGED": FastConnectProviderServicePublicPeeringBgpManagementCustomerManaged,
	"PROVIDER_MANAGED": FastConnectProviderServicePublicPeeringBgpManagementProviderManaged,
	"ORACLE_MANAGED":   FastConnectProviderServicePublicPeeringBgpManagementOracleManaged,
	"UNKNOWN":          FastConnectProviderServicePublicPeeringBgpManagementUnknown,
}

// GetFastConnectProviderServicePublicPeeringBgpManagementEnumValues Enumerates the set of values for FastConnectProviderServicePublicPeeringBgpManagement
func GetFastConnectProviderServicePublicPeeringBgpManagementEnumValues() []FastConnectProviderServicePublicPeeringBgpManagementEnum {
	values := make([]FastConnectProviderServicePublicPeeringBgpManagementEnum, 0)
	for _, v := range mappingFastConnectProviderServicePublicPeeringBgpManagement {
		if v != FastConnectProviderServicePublicPeeringBgpManagementUnknown {
			values = append(values, v)
		}
	}
	return values
}

// FastConnectProviderServiceSupportedVirtualCircuitTypesEnum Enum with underlying type: string
type FastConnectProviderServiceSupportedVirtualCircuitTypesEnum string

// Set of constants representing the allowable values for FastConnectProviderServiceSupportedVirtualCircuitTypes
const (
	FastConnectProviderServiceSupportedVirtualCircuitTypesPublic  FastConnectProviderServiceSupportedVirtualCircuitTypesEnum = "PUBLIC"
	FastConnectProviderServiceSupportedVirtualCircuitTypesPrivate FastConnectProviderServiceSupportedVirtualCircuitTypesEnum = "PRIVATE"
	FastConnectProviderServiceSupportedVirtualCircuitTypesUnknown FastConnectProviderServiceSupportedVirtualCircuitTypesEnum = "UNKNOWN"
)

var mappingFastConnectProviderServiceSupportedVirtualCircuitTypes = map[string]FastConnectProviderServiceSupportedVirtualCircuitTypesEnum{
	"PUBLIC":  FastConnectProviderServiceSupportedVirtualCircuitTypesPublic,
	"PRIVATE": FastConnectProviderServiceSupportedVirtualCircuitTypesPrivate,
	"UNKNOWN": FastConnectProviderServiceSupportedVirtualCircuitTypesUnknown,
}

// GetFastConnectProviderServiceSupportedVirtualCircuitTypesEnumValues Enumerates the set of values for FastConnectProviderServiceSupportedVirtualCircuitTypes
func GetFastConnectProviderServiceSupportedVirtualCircuitTypesEnumValues() []FastConnectProviderServiceSupportedVirtualCircuitTypesEnum {
	values := make([]FastConnectProviderServiceSupportedVirtualCircuitTypesEnum, 0)
	for _, v := range mappingFastConnectProviderServiceSupportedVirtualCircuitTypes {
		if v != FastConnectProviderServiceSupportedVirtualCircuitTypesUnknown {
			values = append(values, v)
		}
	}
	return values
}

// FastConnectProviderServiceType_Enum Enum with underlying type: string
type FastConnectProviderServiceType_Enum string

// Set of constants representing the allowable values for FastConnectProviderServiceType_
const (
	FastConnectProviderServiceType_Layer2  FastConnectProviderServiceType_Enum = "LAYER2"
	FastConnectProviderServiceType_Layer3  FastConnectProviderServiceType_Enum = "LAYER3"
	FastConnectProviderServiceType_Unknown FastConnectProviderServiceType_Enum = "UNKNOWN"
)

var mappingFastConnectProviderServiceType_ = map[string]FastConnectProviderServiceType_Enum{
	"LAYER2":  FastConnectProviderServiceType_Layer2,
	"LAYER3":  FastConnectProviderServiceType_Layer3,
	"UNKNOWN": FastConnectProviderServiceType_Unknown,
}

// GetFastConnectProviderServiceType_EnumValues Enumerates the set of values for FastConnectProviderServiceType_
func GetFastConnectProviderServiceType_EnumValues() []FastConnectProviderServiceType_Enum {
	values := make([]FastConnectProviderServiceType_Enum, 0)
	for _, v := range mappingFastConnectProviderServiceType_ {
		if v != FastConnectProviderServiceType_Unknown {
			values = append(values, v)
		}
	}
	return values
}
