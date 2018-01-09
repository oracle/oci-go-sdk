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

// VirtualCircuit. For use with Oracle Cloud Infrastructure FastConnect.
// A virtual circuit is an isolated network path that runs over one or more physical
// network connections to provide a single, logical connection between the edge router
// on the customer's existing network and Oracle Cloud Infrastructure. *Private*
// virtual circuits support private peering, and *public* virtual circuits support
// public peering. For more information, see [FastConnect Overview]({{DOC_SERVER_URL}}/Content/Network/Concepts/fastconnect.htm).
// Each virtual circuit is made up of information shared between a customer, Oracle,
// and a provider (if the customer is using FastConnect via a provider). Who fills in
// a given property of a virtual circuit depends on whether the BGP session related to
// that virtual circuit goes from the customer's edge router to Oracle, or from the provider's
// edge router to Oracle. Also, in the case where the customer is using a provider, values
// for some of the properties may not be present immediately, but may get filled in as the
// provider and Oracle each do their part to provision the virtual circuit.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type VirtualCircuit struct {

	// The provisioned data rate of the connection.
	BandwidthShapeName *string `mandatory:"false" json:"bandwidthShapeName,omitempty"`

	// BGP management option.
	BgpManagement VirtualCircuitBgpManagementEnum `mandatory:"false" json:"bgpManagement,omitempty"`

	// The state of the BGP session associated with the virtual circuit.
	BgpSessionState VirtualCircuitBgpSessionStateEnum `mandatory:"false" json:"bgpSessionState,omitempty"`

	// The OCID of the compartment containing the virtual circuit.
	CompartmentID *string `mandatory:"false" json:"compartmentId,omitempty"`

	// An array of mappings, each containing properties for a
	// cross-connect or cross-connect group that is associated with this
	// virtual circuit.
	CrossConnectMappings []CrossConnectMapping `mandatory:"false" json:"crossConnectMappings,omitempty"`

	// The BGP ASN of the network at the other end of the BGP
	// session from Oracle. If the session is between the customer's
	// edge router and Oracle, the value is the customer's ASN. If the BGP
	// session is between the provider's edge router and Oracle, the value
	// is the provider's ASN.
	CustomerBgpAsn *int `mandatory:"false" json:"customerBgpAsn,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// The OCID of the customer's Drg
	// that this virtual circuit uses. Applicable only to private virtual circuits.
	GatewayID *string `mandatory:"false" json:"gatewayId,omitempty"`

	// The virtual circuit's Oracle ID (OCID).
	ID *string `mandatory:"false" json:"id,omitempty"`

	// The virtual circuit's current state. For information about
	// the different states, see
	// [FastConnect Overview]({{DOC_SERVER_URL}}/Content/Network/Concepts/fastconnect.htm).
	LifecycleState VirtualCircuitLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The Oracle BGP ASN.
	OracleBgpAsn *int `mandatory:"false" json:"oracleBgpAsn,omitempty"`

	// Deprecated. Instead use `providerServiceId`.
	ProviderName *string `mandatory:"false" json:"providerName,omitempty"`

	// The OCID of the service offered by the provider (if the customer is connecting via a provider).
	ProviderServiceID *string `mandatory:"false" json:"providerServiceId,omitempty"`

	// Deprecated. Instead use `providerServiceId`.
	ProviderServiceName *string `mandatory:"false" json:"providerServiceName,omitempty"`

	// The provider's state in relation to this virtual circuit (if the
	// customer is connecting via a provider). ACTIVE means
	// the provider has provisioned the virtual circuit from their end.
	// INACTIVE means the provider has not yet provisioned the virtual
	// circuit, or has de-provisioned it.
	ProviderState VirtualCircuitProviderStateEnum `mandatory:"false" json:"providerState,omitempty"`

	// For a public virtual circuit. The public IP prefixes (CIDRs) the customer wants to
	// advertise across the connection. Each prefix must be /24 or less specific.
	PublicPrefixes []string `mandatory:"false" json:"publicPrefixes,omitempty"`

	// Provider-supplied reference information about this virtual circuit
	// (if the customer is connecting via a provider).
	ReferenceComment *string `mandatory:"false" json:"referenceComment,omitempty"`

	// The Oracle Cloud Infrastructure region where this virtual
	// circuit is located.
	Region *string `mandatory:"false" json:"region,omitempty"`

	// Provider service type.
	ServiceType VirtualCircuitServiceTypeEnum `mandatory:"false" json:"serviceType,omitempty"`

	// The date and time the virtual circuit was created,
	// in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`

	// Whether the virtual circuit supports private or public peering. For more information,
	// see [FastConnect Overview]({{DOC_SERVER_URL}}/Content/Network/Concepts/fastconnect.htm).
	Type_ VirtualCircuitType_Enum `mandatory:"false" json:"type,omitempty"`
}

func (model VirtualCircuit) String() string {
	return common.PointerString(model)
}

type VirtualCircuitBgpManagementEnum string

const (
	VIRTUAL_CIRCUIT_BGP_MANAGEMENT_CUSTOMER_MANAGED VirtualCircuitBgpManagementEnum = "CUSTOMER_MANAGED"
	VIRTUAL_CIRCUIT_BGP_MANAGEMENT_PROVIDER_MANAGED VirtualCircuitBgpManagementEnum = "PROVIDER_MANAGED"
	VIRTUAL_CIRCUIT_BGP_MANAGEMENT_ORACLE_MANAGED   VirtualCircuitBgpManagementEnum = "ORACLE_MANAGED"
	VIRTUAL_CIRCUIT_BGP_MANAGEMENT_UNKNOWN          VirtualCircuitBgpManagementEnum = "UNKNOWN"
)

var mapping_virtualcircuit_bgpManagement = map[string]VirtualCircuitBgpManagementEnum{
	"CUSTOMER_MANAGED": VIRTUAL_CIRCUIT_BGP_MANAGEMENT_CUSTOMER_MANAGED,
	"PROVIDER_MANAGED": VIRTUAL_CIRCUIT_BGP_MANAGEMENT_PROVIDER_MANAGED,
	"ORACLE_MANAGED":   VIRTUAL_CIRCUIT_BGP_MANAGEMENT_ORACLE_MANAGED,
	"UNKNOWN":          VIRTUAL_CIRCUIT_BGP_MANAGEMENT_UNKNOWN,
}

func GetVirtualCircuitBgpManagementEnumValues() []VirtualCircuitBgpManagementEnum {
	values := make([]VirtualCircuitBgpManagementEnum, 0)
	for _, v := range mapping_virtualcircuit_bgpManagement {
		if v != VIRTUAL_CIRCUIT_BGP_MANAGEMENT_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type VirtualCircuitBgpSessionStateEnum string

const (
	VIRTUAL_CIRCUIT_BGP_SESSION_STATE_UP      VirtualCircuitBgpSessionStateEnum = "UP"
	VIRTUAL_CIRCUIT_BGP_SESSION_STATE_DOWN    VirtualCircuitBgpSessionStateEnum = "DOWN"
	VIRTUAL_CIRCUIT_BGP_SESSION_STATE_UNKNOWN VirtualCircuitBgpSessionStateEnum = "UNKNOWN"
)

var mapping_virtualcircuit_bgpSessionState = map[string]VirtualCircuitBgpSessionStateEnum{
	"UP":      VIRTUAL_CIRCUIT_BGP_SESSION_STATE_UP,
	"DOWN":    VIRTUAL_CIRCUIT_BGP_SESSION_STATE_DOWN,
	"UNKNOWN": VIRTUAL_CIRCUIT_BGP_SESSION_STATE_UNKNOWN,
}

func GetVirtualCircuitBgpSessionStateEnumValues() []VirtualCircuitBgpSessionStateEnum {
	values := make([]VirtualCircuitBgpSessionStateEnum, 0)
	for _, v := range mapping_virtualcircuit_bgpSessionState {
		if v != VIRTUAL_CIRCUIT_BGP_SESSION_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type VirtualCircuitLifecycleStateEnum string

const (
	VIRTUAL_CIRCUIT_LIFECYCLE_STATE_PENDING_PROVIDER VirtualCircuitLifecycleStateEnum = "PENDING_PROVIDER"
	VIRTUAL_CIRCUIT_LIFECYCLE_STATE_VERIFYING        VirtualCircuitLifecycleStateEnum = "VERIFYING"
	VIRTUAL_CIRCUIT_LIFECYCLE_STATE_PROVISIONING     VirtualCircuitLifecycleStateEnum = "PROVISIONING"
	VIRTUAL_CIRCUIT_LIFECYCLE_STATE_PROVISIONED      VirtualCircuitLifecycleStateEnum = "PROVISIONED"
	VIRTUAL_CIRCUIT_LIFECYCLE_STATE_FAILED           VirtualCircuitLifecycleStateEnum = "FAILED"
	VIRTUAL_CIRCUIT_LIFECYCLE_STATE_INACTIVE         VirtualCircuitLifecycleStateEnum = "INACTIVE"
	VIRTUAL_CIRCUIT_LIFECYCLE_STATE_TERMINATING      VirtualCircuitLifecycleStateEnum = "TERMINATING"
	VIRTUAL_CIRCUIT_LIFECYCLE_STATE_TERMINATED       VirtualCircuitLifecycleStateEnum = "TERMINATED"
	VIRTUAL_CIRCUIT_LIFECYCLE_STATE_UNKNOWN          VirtualCircuitLifecycleStateEnum = "UNKNOWN"
)

var mapping_virtualcircuit_lifecycleState = map[string]VirtualCircuitLifecycleStateEnum{
	"PENDING_PROVIDER": VIRTUAL_CIRCUIT_LIFECYCLE_STATE_PENDING_PROVIDER,
	"VERIFYING":        VIRTUAL_CIRCUIT_LIFECYCLE_STATE_VERIFYING,
	"PROVISIONING":     VIRTUAL_CIRCUIT_LIFECYCLE_STATE_PROVISIONING,
	"PROVISIONED":      VIRTUAL_CIRCUIT_LIFECYCLE_STATE_PROVISIONED,
	"FAILED":           VIRTUAL_CIRCUIT_LIFECYCLE_STATE_FAILED,
	"INACTIVE":         VIRTUAL_CIRCUIT_LIFECYCLE_STATE_INACTIVE,
	"TERMINATING":      VIRTUAL_CIRCUIT_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":       VIRTUAL_CIRCUIT_LIFECYCLE_STATE_TERMINATED,
	"UNKNOWN":          VIRTUAL_CIRCUIT_LIFECYCLE_STATE_UNKNOWN,
}

func GetVirtualCircuitLifecycleStateEnumValues() []VirtualCircuitLifecycleStateEnum {
	values := make([]VirtualCircuitLifecycleStateEnum, 0)
	for _, v := range mapping_virtualcircuit_lifecycleState {
		if v != VIRTUAL_CIRCUIT_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type VirtualCircuitProviderStateEnum string

const (
	VIRTUAL_CIRCUIT_PROVIDER_STATE_ACTIVE   VirtualCircuitProviderStateEnum = "ACTIVE"
	VIRTUAL_CIRCUIT_PROVIDER_STATE_INACTIVE VirtualCircuitProviderStateEnum = "INACTIVE"
	VIRTUAL_CIRCUIT_PROVIDER_STATE_UNKNOWN  VirtualCircuitProviderStateEnum = "UNKNOWN"
)

var mapping_virtualcircuit_providerState = map[string]VirtualCircuitProviderStateEnum{
	"ACTIVE":   VIRTUAL_CIRCUIT_PROVIDER_STATE_ACTIVE,
	"INACTIVE": VIRTUAL_CIRCUIT_PROVIDER_STATE_INACTIVE,
	"UNKNOWN":  VIRTUAL_CIRCUIT_PROVIDER_STATE_UNKNOWN,
}

func GetVirtualCircuitProviderStateEnumValues() []VirtualCircuitProviderStateEnum {
	values := make([]VirtualCircuitProviderStateEnum, 0)
	for _, v := range mapping_virtualcircuit_providerState {
		if v != VIRTUAL_CIRCUIT_PROVIDER_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type VirtualCircuitServiceTypeEnum string

const (
	VIRTUAL_CIRCUIT_SERVICE_TYPE_COLOCATED VirtualCircuitServiceTypeEnum = "COLOCATED"
	VIRTUAL_CIRCUIT_SERVICE_TYPE_LAYER2    VirtualCircuitServiceTypeEnum = "LAYER2"
	VIRTUAL_CIRCUIT_SERVICE_TYPE_LAYER3    VirtualCircuitServiceTypeEnum = "LAYER3"
	VIRTUAL_CIRCUIT_SERVICE_TYPE_UNKNOWN   VirtualCircuitServiceTypeEnum = "UNKNOWN"
)

var mapping_virtualcircuit_serviceType = map[string]VirtualCircuitServiceTypeEnum{
	"COLOCATED": VIRTUAL_CIRCUIT_SERVICE_TYPE_COLOCATED,
	"LAYER2":    VIRTUAL_CIRCUIT_SERVICE_TYPE_LAYER2,
	"LAYER3":    VIRTUAL_CIRCUIT_SERVICE_TYPE_LAYER3,
	"UNKNOWN":   VIRTUAL_CIRCUIT_SERVICE_TYPE_UNKNOWN,
}

func GetVirtualCircuitServiceTypeEnumValues() []VirtualCircuitServiceTypeEnum {
	values := make([]VirtualCircuitServiceTypeEnum, 0)
	for _, v := range mapping_virtualcircuit_serviceType {
		if v != VIRTUAL_CIRCUIT_SERVICE_TYPE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type VirtualCircuitType_Enum string

const (
	VIRTUAL_CIRCUIT_TYPE__PUBLIC  VirtualCircuitType_Enum = "PUBLIC"
	VIRTUAL_CIRCUIT_TYPE__PRIVATE VirtualCircuitType_Enum = "PRIVATE"
	VIRTUAL_CIRCUIT_TYPE__UNKNOWN VirtualCircuitType_Enum = "UNKNOWN"
)

var mapping_virtualcircuit_type = map[string]VirtualCircuitType_Enum{
	"PUBLIC":  VIRTUAL_CIRCUIT_TYPE__PUBLIC,
	"PRIVATE": VIRTUAL_CIRCUIT_TYPE__PRIVATE,
	"UNKNOWN": VIRTUAL_CIRCUIT_TYPE__UNKNOWN,
}

func GetVirtualCircuitType_EnumValues() []VirtualCircuitType_Enum {
	values := make([]VirtualCircuitType_Enum, 0)
	for _, v := range mapping_virtualcircuit_type {
		if v != VIRTUAL_CIRCUIT_TYPE__UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
