// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// FastConnectProviderService. A service offering from a supported provider. For more information,
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
}

func (model FastConnectProviderService) String() string {
	return common.PointerString(model)
}

type FastConnectProviderServicePrivatePeeringBgpManagementEnum string
type FastConnectProviderServicePrivatePeeringBgpManagement struct{}

const (
	FAST_CONNECT_PROVIDER_SERVICE_PRIVATE_PEERING_BGP_MANAGEMENT_CUSTOMER_MANAGED FastConnectProviderServicePrivatePeeringBgpManagementEnum = "CUSTOMER_MANAGED"
	FAST_CONNECT_PROVIDER_SERVICE_PRIVATE_PEERING_BGP_MANAGEMENT_PROVIDER_MANAGED FastConnectProviderServicePrivatePeeringBgpManagementEnum = "PROVIDER_MANAGED"
	FAST_CONNECT_PROVIDER_SERVICE_PRIVATE_PEERING_BGP_MANAGEMENT_ORACLE_MANAGED   FastConnectProviderServicePrivatePeeringBgpManagementEnum = "ORACLE_MANAGED"
	FAST_CONNECT_PROVIDER_SERVICE_PRIVATE_PEERING_BGP_MANAGEMENT_UNKNOWN          FastConnectProviderServicePrivatePeeringBgpManagementEnum = "UNKNOWN"
)

var mapping_fastconnectproviderservice_privatePeeringBgpManagement = map[string]FastConnectProviderServicePrivatePeeringBgpManagementEnum{
	"CUSTOMER_MANAGED": FAST_CONNECT_PROVIDER_SERVICE_PRIVATE_PEERING_BGP_MANAGEMENT_CUSTOMER_MANAGED,
	"PROVIDER_MANAGED": FAST_CONNECT_PROVIDER_SERVICE_PRIVATE_PEERING_BGP_MANAGEMENT_PROVIDER_MANAGED,
	"ORACLE_MANAGED":   FAST_CONNECT_PROVIDER_SERVICE_PRIVATE_PEERING_BGP_MANAGEMENT_ORACLE_MANAGED,
	"UNKNOWN":          FAST_CONNECT_PROVIDER_SERVICE_PRIVATE_PEERING_BGP_MANAGEMENT_UNKNOWN,
}

func (receiver FastConnectProviderServicePrivatePeeringBgpManagement) Values() []FastConnectProviderServicePrivatePeeringBgpManagementEnum {
	values := make([]FastConnectProviderServicePrivatePeeringBgpManagementEnum, 0)
	for _, v := range mapping_fastconnectproviderservice_privatePeeringBgpManagement {
		if v != FAST_CONNECT_PROVIDER_SERVICE_PRIVATE_PEERING_BGP_MANAGEMENT_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

func (receiver FastConnectProviderServicePrivatePeeringBgpManagement) IsValid(toBeChecked string) bool {
	for _, v := range receiver.Values() {
		if FastConnectProviderServicePrivatePeeringBgpManagementEnum(toBeChecked) == v {
			return true
		}
	}
	return false
}

func (receiver FastConnectProviderServicePrivatePeeringBgpManagement) From(toBeConverted string) FastConnectProviderServicePrivatePeeringBgpManagementEnum {
	if val, ok := mapping_fastconnectproviderservice_privatePeeringBgpManagement[toBeConverted]; ok {
		return val
	}
	return FAST_CONNECT_PROVIDER_SERVICE_PRIVATE_PEERING_BGP_MANAGEMENT_UNKNOWN
}

type FastConnectProviderServicePublicPeeringBgpManagementEnum string
type FastConnectProviderServicePublicPeeringBgpManagement struct{}

const (
	FAST_CONNECT_PROVIDER_SERVICE_PUBLIC_PEERING_BGP_MANAGEMENT_CUSTOMER_MANAGED FastConnectProviderServicePublicPeeringBgpManagementEnum = "CUSTOMER_MANAGED"
	FAST_CONNECT_PROVIDER_SERVICE_PUBLIC_PEERING_BGP_MANAGEMENT_PROVIDER_MANAGED FastConnectProviderServicePublicPeeringBgpManagementEnum = "PROVIDER_MANAGED"
	FAST_CONNECT_PROVIDER_SERVICE_PUBLIC_PEERING_BGP_MANAGEMENT_ORACLE_MANAGED   FastConnectProviderServicePublicPeeringBgpManagementEnum = "ORACLE_MANAGED"
	FAST_CONNECT_PROVIDER_SERVICE_PUBLIC_PEERING_BGP_MANAGEMENT_UNKNOWN          FastConnectProviderServicePublicPeeringBgpManagementEnum = "UNKNOWN"
)

var mapping_fastconnectproviderservice_publicPeeringBgpManagement = map[string]FastConnectProviderServicePublicPeeringBgpManagementEnum{
	"CUSTOMER_MANAGED": FAST_CONNECT_PROVIDER_SERVICE_PUBLIC_PEERING_BGP_MANAGEMENT_CUSTOMER_MANAGED,
	"PROVIDER_MANAGED": FAST_CONNECT_PROVIDER_SERVICE_PUBLIC_PEERING_BGP_MANAGEMENT_PROVIDER_MANAGED,
	"ORACLE_MANAGED":   FAST_CONNECT_PROVIDER_SERVICE_PUBLIC_PEERING_BGP_MANAGEMENT_ORACLE_MANAGED,
	"UNKNOWN":          FAST_CONNECT_PROVIDER_SERVICE_PUBLIC_PEERING_BGP_MANAGEMENT_UNKNOWN,
}

func (receiver FastConnectProviderServicePublicPeeringBgpManagement) Values() []FastConnectProviderServicePublicPeeringBgpManagementEnum {
	values := make([]FastConnectProviderServicePublicPeeringBgpManagementEnum, 0)
	for _, v := range mapping_fastconnectproviderservice_publicPeeringBgpManagement {
		if v != FAST_CONNECT_PROVIDER_SERVICE_PUBLIC_PEERING_BGP_MANAGEMENT_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

func (receiver FastConnectProviderServicePublicPeeringBgpManagement) IsValid(toBeChecked string) bool {
	for _, v := range receiver.Values() {
		if FastConnectProviderServicePublicPeeringBgpManagementEnum(toBeChecked) == v {
			return true
		}
	}
	return false
}

func (receiver FastConnectProviderServicePublicPeeringBgpManagement) From(toBeConverted string) FastConnectProviderServicePublicPeeringBgpManagementEnum {
	if val, ok := mapping_fastconnectproviderservice_publicPeeringBgpManagement[toBeConverted]; ok {
		return val
	}
	return FAST_CONNECT_PROVIDER_SERVICE_PUBLIC_PEERING_BGP_MANAGEMENT_UNKNOWN
}

type FastConnectProviderServiceType_Enum string
type FastConnectProviderServiceType_ struct{}

const (
	FAST_CONNECT_PROVIDER_SERVICE_TYPE__LAYER2  FastConnectProviderServiceType_Enum = "LAYER2"
	FAST_CONNECT_PROVIDER_SERVICE_TYPE__LAYER3  FastConnectProviderServiceType_Enum = "LAYER3"
	FAST_CONNECT_PROVIDER_SERVICE_TYPE__UNKNOWN FastConnectProviderServiceType_Enum = "UNKNOWN"
)

var mapping_fastconnectproviderservice_type = map[string]FastConnectProviderServiceType_Enum{
	"LAYER2":  FAST_CONNECT_PROVIDER_SERVICE_TYPE__LAYER2,
	"LAYER3":  FAST_CONNECT_PROVIDER_SERVICE_TYPE__LAYER3,
	"UNKNOWN": FAST_CONNECT_PROVIDER_SERVICE_TYPE__UNKNOWN,
}

func (receiver FastConnectProviderServiceType_) Values() []FastConnectProviderServiceType_Enum {
	values := make([]FastConnectProviderServiceType_Enum, 0)
	for _, v := range mapping_fastconnectproviderservice_type {
		if v != FAST_CONNECT_PROVIDER_SERVICE_TYPE__UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

func (receiver FastConnectProviderServiceType_) IsValid(toBeChecked string) bool {
	for _, v := range receiver.Values() {
		if FastConnectProviderServiceType_Enum(toBeChecked) == v {
			return true
		}
	}
	return false
}

func (receiver FastConnectProviderServiceType_) From(toBeConverted string) FastConnectProviderServiceType_Enum {
	if val, ok := mapping_fastconnectproviderservice_type[toBeConverted]; ok {
		return val
	}
	return FAST_CONNECT_PROVIDER_SERVICE_TYPE__UNKNOWN
}
