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

// LocalPeeringGateway. A local peering gateway (LPG) is an object on a VCN that lets that VCN peer
// with another VCN in the same region. *Peering* means that the two VCNs can
// communicate using private IP addresses, but without the traffic traversing the
// internet or routing through your on-premises network. For more information,
// see [VCN Peering]({{DOC_SERVER_URL}}/Content/Network/Tasks/VCNpeering.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type LocalPeeringGateway struct {

	// The OCID of the compartment containing the LPG.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName,omitempty"`

	// The LPG's Oracle ID (OCID).
	ID *string `mandatory:"true" json:"id,omitempty"`

	// Whether the VCN at the other end of the peering is in a different tenancy.
	// Example: `false`
	IsCrossTenancyPeering *bool `mandatory:"true" json:"isCrossTenancyPeering,omitempty"`

	// The LPG's current lifecycle state.
	LifecycleState LocalPeeringGatewayLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// Whether the LPG is peered with another LPG. `NEW` means the LPG has not yet been
	// peered. `PENDING` means the peering is being established. `REVOKED` means the
	// LPG at the other end of the peering has been deleted.
	PeeringStatus LocalPeeringGatewayPeeringStatusEnum `mandatory:"true" json:"peeringStatus,omitempty"`

	// The date and time the LPG was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The OCID of the VCN the LPG belongs to.
	VcnID *string `mandatory:"true" json:"vcnId,omitempty"`

	// The range of IP addresses available on the VCN at the other
	// end of the peering from this LPG. The value is `null` if the LPG is not peered.
	// You can use this as the destination CIDR for a route rule to route a subnet's
	// traffic to this LPG.
	// Example: `192.168.0.0/16`
	PeerAdvertisedCidr *string `mandatory:"false" json:"peerAdvertisedCidr,omitempty"`

	// Additional information regarding the peering status, if applicable.
	PeeringStatusDetails *string `mandatory:"false" json:"peeringStatusDetails,omitempty"`
}

func (model LocalPeeringGateway) String() string {
	return common.PointerString(model)
}

type LocalPeeringGatewayLifecycleStateEnum string

const (
	LOCAL_PEERING_GATEWAY_LIFECYCLE_STATE_PROVISIONING LocalPeeringGatewayLifecycleStateEnum = "PROVISIONING"
	LOCAL_PEERING_GATEWAY_LIFECYCLE_STATE_AVAILABLE    LocalPeeringGatewayLifecycleStateEnum = "AVAILABLE"
	LOCAL_PEERING_GATEWAY_LIFECYCLE_STATE_TERMINATING  LocalPeeringGatewayLifecycleStateEnum = "TERMINATING"
	LOCAL_PEERING_GATEWAY_LIFECYCLE_STATE_TERMINATED   LocalPeeringGatewayLifecycleStateEnum = "TERMINATED"
	LOCAL_PEERING_GATEWAY_LIFECYCLE_STATE_UNKNOWN      LocalPeeringGatewayLifecycleStateEnum = "UNKNOWN"
)

var mapping_localpeeringgateway_lifecycleState = map[string]LocalPeeringGatewayLifecycleStateEnum{
	"PROVISIONING": LOCAL_PEERING_GATEWAY_LIFECYCLE_STATE_PROVISIONING,
	"AVAILABLE":    LOCAL_PEERING_GATEWAY_LIFECYCLE_STATE_AVAILABLE,
	"TERMINATING":  LOCAL_PEERING_GATEWAY_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":   LOCAL_PEERING_GATEWAY_LIFECYCLE_STATE_TERMINATED,
	"UNKNOWN":      LOCAL_PEERING_GATEWAY_LIFECYCLE_STATE_UNKNOWN,
}

func GetLocalPeeringGatewayLifecycleStateEnumValues() []LocalPeeringGatewayLifecycleStateEnum {
	values := make([]LocalPeeringGatewayLifecycleStateEnum, 0)
	for _, v := range mapping_localpeeringgateway_lifecycleState {
		if v != LOCAL_PEERING_GATEWAY_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type LocalPeeringGatewayPeeringStatusEnum string

const (
	LOCAL_PEERING_GATEWAY_PEERING_STATUS_INVALID LocalPeeringGatewayPeeringStatusEnum = "INVALID"
	LOCAL_PEERING_GATEWAY_PEERING_STATUS_NEW     LocalPeeringGatewayPeeringStatusEnum = "NEW"
	LOCAL_PEERING_GATEWAY_PEERING_STATUS_PEERED  LocalPeeringGatewayPeeringStatusEnum = "PEERED"
	LOCAL_PEERING_GATEWAY_PEERING_STATUS_PENDING LocalPeeringGatewayPeeringStatusEnum = "PENDING"
	LOCAL_PEERING_GATEWAY_PEERING_STATUS_REVOKED LocalPeeringGatewayPeeringStatusEnum = "REVOKED"
	LOCAL_PEERING_GATEWAY_PEERING_STATUS_UNKNOWN LocalPeeringGatewayPeeringStatusEnum = "UNKNOWN"
)

var mapping_localpeeringgateway_peeringStatus = map[string]LocalPeeringGatewayPeeringStatusEnum{
	"INVALID": LOCAL_PEERING_GATEWAY_PEERING_STATUS_INVALID,
	"NEW":     LOCAL_PEERING_GATEWAY_PEERING_STATUS_NEW,
	"PEERED":  LOCAL_PEERING_GATEWAY_PEERING_STATUS_PEERED,
	"PENDING": LOCAL_PEERING_GATEWAY_PEERING_STATUS_PENDING,
	"REVOKED": LOCAL_PEERING_GATEWAY_PEERING_STATUS_REVOKED,
	"UNKNOWN": LOCAL_PEERING_GATEWAY_PEERING_STATUS_UNKNOWN,
}

func GetLocalPeeringGatewayPeeringStatusEnumValues() []LocalPeeringGatewayPeeringStatusEnum {
	values := make([]LocalPeeringGatewayPeeringStatusEnum, 0)
	for _, v := range mapping_localpeeringgateway_peeringStatus {
		if v != LOCAL_PEERING_GATEWAY_PEERING_STATUS_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
