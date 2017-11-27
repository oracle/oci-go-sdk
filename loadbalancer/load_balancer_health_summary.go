// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"github.com/oci-go-sdk/common"
)

// LoadBalancerHealthSummary. A health status summary for the specified load balancer.
type LoadBalancerHealthSummary struct {

	// The [OCID]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm) of the load balancer the health status is associated with.
	LoadBalancerID *string `mandatory:"true" json:"loadBalancerId,omitempty"`

	// The overall health status of the load balancer.
	// *  **OK:** All backend sets associated with the load balancer return a status of `OK`.
	// *  **WARNING:** At least one of the backend sets associated with the load balancer returns a status of `WARNING`,
	// no backend sets return a status of `CRITICAL`, and the load balancer life cycle state is `ACTIVE`.
	// *  **CRITICAL:** One or more of the backend sets associated with the load balancer return a status of `CRITICAL`.
	// *  **UNKNOWN:** If any one of the following conditions is true:
	//     *  The load balancer life cycle state is not `ACTIVE`.
	//     *  No backend sets are defined for the load balancer.
	//     *  More than half of the backend sets associated with the load balancer return a status of `UNKNOWN`, none of the backend
	//        sets return a status of `WARNING` or `CRITICAL`, and the load balancer life cycle state is `ACTIVE`.
	//     *  The system could not retrieve metrics for any reason.
	Status LoadBalancerHealthSummaryStatusEnum `mandatory:"true" json:"status,omitempty"`
}

func (model LoadBalancerHealthSummary) String() string {
	return common.PointerString(model)
}

type LoadBalancerHealthSummaryStatusEnum string

const (
	LOAD_BALANCER_HEALTH_SUMMARY_STATUS_OK       LoadBalancerHealthSummaryStatusEnum = "OK"
	LOAD_BALANCER_HEALTH_SUMMARY_STATUS_WARNING  LoadBalancerHealthSummaryStatusEnum = "WARNING"
	LOAD_BALANCER_HEALTH_SUMMARY_STATUS_CRITICAL LoadBalancerHealthSummaryStatusEnum = "CRITICAL"
	LOAD_BALANCER_HEALTH_SUMMARY_STATUS_UNKNOWN  LoadBalancerHealthSummaryStatusEnum = "UNKNOWN"
)

var mapping_loadbalancerhealthsummary_status = map[string]LoadBalancerHealthSummaryStatusEnum{
	"OK":       LOAD_BALANCER_HEALTH_SUMMARY_STATUS_OK,
	"WARNING":  LOAD_BALANCER_HEALTH_SUMMARY_STATUS_WARNING,
	"CRITICAL": LOAD_BALANCER_HEALTH_SUMMARY_STATUS_CRITICAL,
	"UNKNOWN":  LOAD_BALANCER_HEALTH_SUMMARY_STATUS_UNKNOWN,
}

func GetLoadBalancerHealthSummaryStatusEnumValues() []LoadBalancerHealthSummaryStatusEnum {
	values := make([]LoadBalancerHealthSummaryStatusEnum, 0)
	for _, v := range mapping_loadbalancerhealthsummary_status {
		values = append(values, v)
	}
	return values
}
