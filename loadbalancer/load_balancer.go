// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LoadBalancer. The properties that define a load balancer. For more information, see
// [Managing a Load Balancer]({{DOC_SERVER_URL}}/Content/Balance/Tasks/managingloadbalancer.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
// For information about endpoints and signing API requests, see
// [About the API]({{DOC_SERVER_URL}}/Content/API/Concepts/usingapi.htm). For information about available SDKs and tools, see
// [SDKS and Other Tools]({{DOC_SERVER_URL}}/Content/API/Concepts/sdks.htm).
type LoadBalancer struct {

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancer.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Example: `My load balancer`
	DisplayName *string `mandatory:"true" json:"displayName,omitempty"`

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the load balancer.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The current state of the load balancer.
	LifecycleState LoadBalancerLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// A template that determines the total pre-provisioned bandwidth (ingress plus egress).
	// To get a list of available shapes, use the ListShapes
	// operation.
	// Example: `100Mbps`
	ShapeName *string `mandatory:"true" json:"shapeName,omitempty"`

	// The date and time the load balancer was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	BackendSets map[string]BackendSet `mandatory:"false" json:"backendSets,omitempty"`

	Certificates map[string]Certificate `mandatory:"false" json:"certificates,omitempty"`

	// An array of IP addresses.
	IpAddresses []IpAddress `mandatory:"false" json:"ipAddresses,omitempty"`

	// Whether the load balancer has a VCN-local (private) IP address.
	// If "true", the service assigns a private IP address to the load balancer. The load balancer requires only one subnet
	// to host both the primary and secondary load balancers. The private IP address is local to the subnet. The load balancer
	// is accessible only from within the VCN that contains the associated subnet, or as further restricted by your security
	// list rules. The load balancer can route traffic to any backend server that is reachable from the VCN.
	// For a private load balancer, both the primary and secondary load balancer hosts are within the same Availability Domain.
	// If "false", the service assigns a public IP address to the load balancer. A load balancer with a public IP address
	// requires two subnets, each in a different Availability Domain. One subnet hosts the primary load balancer and the other
	// hosts the secondary (standby) load balancer. A public load balancer is accessible from the internet, depending on your
	// VCN's [security list rules]({{DOC_SERVER_URL}}/Content/Network/Concepts/securitylists.htm).
	IsPrivate *bool `mandatory:"false" json:"isPrivate,omitempty"`

	Listeners map[string]Listener `mandatory:"false" json:"listeners,omitempty"`

	PathRouteSets map[string]PathRouteSet `mandatory:"false" json:"pathRouteSets,omitempty"`

	// An array of subnet [OCIDs]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
	SubnetIds []string `mandatory:"false" json:"subnetIds,omitempty"`
}

func (model LoadBalancer) String() string {
	return common.PointerString(model)
}

type LoadBalancerLifecycleStateEnum string

const (
	LOAD_BALANCER_LIFECYCLE_STATE_CREATING LoadBalancerLifecycleStateEnum = "CREATING"
	LOAD_BALANCER_LIFECYCLE_STATE_FAILED   LoadBalancerLifecycleStateEnum = "FAILED"
	LOAD_BALANCER_LIFECYCLE_STATE_ACTIVE   LoadBalancerLifecycleStateEnum = "ACTIVE"
	LOAD_BALANCER_LIFECYCLE_STATE_DELETING LoadBalancerLifecycleStateEnum = "DELETING"
	LOAD_BALANCER_LIFECYCLE_STATE_DELETED  LoadBalancerLifecycleStateEnum = "DELETED"
	LOAD_BALANCER_LIFECYCLE_STATE_UNKNOWN  LoadBalancerLifecycleStateEnum = "UNKNOWN"
)

var mapping_loadbalancer_lifecycleState = map[string]LoadBalancerLifecycleStateEnum{
	"CREATING": LOAD_BALANCER_LIFECYCLE_STATE_CREATING,
	"FAILED":   LOAD_BALANCER_LIFECYCLE_STATE_FAILED,
	"ACTIVE":   LOAD_BALANCER_LIFECYCLE_STATE_ACTIVE,
	"DELETING": LOAD_BALANCER_LIFECYCLE_STATE_DELETING,
	"DELETED":  LOAD_BALANCER_LIFECYCLE_STATE_DELETED,
	"UNKNOWN":  LOAD_BALANCER_LIFECYCLE_STATE_UNKNOWN,
}

func GetLoadBalancerLifecycleStateEnumValues() []LoadBalancerLifecycleStateEnum {
	values := make([]LoadBalancerLifecycleStateEnum, 0)
	for _, v := range mapping_loadbalancer_lifecycleState {
		if v != LOAD_BALANCER_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
