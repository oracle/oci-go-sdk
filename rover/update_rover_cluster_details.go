// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// UpdateRoverClusterDetails The information required to update a RoverCluster.
type UpdateRoverClusterDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Number of nodes desired in the cluster, between 5 and 15.
	ClusterSize *int `mandatory:"false" json:"clusterSize"`

	CustomerShippingAddress *ShippingAddress `mandatory:"false" json:"customerShippingAddress"`

	// List of existing workloads that should be provisioned on the nodes.
	ClusterWorkloads []RoverWorkload `mandatory:"false" json:"clusterWorkloads"`

	// Root password for the rover cluster.
	SuperUserPassword *string `mandatory:"false" json:"superUserPassword"`

	// The current state of the RoverCluster.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A property that can contain details on the lifecycle.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Password to unlock the rover cluster.
	UnlockPassphrase *string `mandatory:"false" json:"unlockPassphrase"`

	// The type of enclosure rover nodes in this cluster are shipped in.
	EnclosureType EnclosureTypeEnum `mandatory:"false" json:"enclosureType,omitempty"`

	// Name of point of contact for this order if customer is picking up.
	PointOfContact *string `mandatory:"false" json:"pointOfContact"`

	// Phone number of point of contact for this order if customer is picking up.
	PointOfContactPhoneNumber *string `mandatory:"false" json:"pointOfContactPhoneNumber"`

	// Preference for device delivery.
	ShippingPreference UpdateRoverClusterDetailsShippingPreferenceEnum `mandatory:"false" json:"shippingPreference,omitempty"`

	// Tracking Url for the shipped Rover Cluster.
	OracleShippingTrackingUrl *string `mandatory:"false" json:"oracleShippingTrackingUrl"`

	// Shipping vendor of choice for orace to customer shipping.
	ShippingVendor *string `mandatory:"false" json:"shippingVendor"`

	// Expected date when customer wants to pickup the device if they chose customer pickup.
	TimePickupExpected *common.SDKTime `mandatory:"false" json:"timePickupExpected"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m UpdateRoverClusterDetails) String() string {
	return common.PointerString(m)
}

// UpdateRoverClusterDetailsShippingPreferenceEnum Enum with underlying type: string
type UpdateRoverClusterDetailsShippingPreferenceEnum string

// Set of constants representing the allowable values for UpdateRoverClusterDetailsShippingPreferenceEnum
const (
	UpdateRoverClusterDetailsShippingPreferenceOracleShipped  UpdateRoverClusterDetailsShippingPreferenceEnum = "ORACLE_SHIPPED"
	UpdateRoverClusterDetailsShippingPreferenceCustomerPickup UpdateRoverClusterDetailsShippingPreferenceEnum = "CUSTOMER_PICKUP"
)

var mappingUpdateRoverClusterDetailsShippingPreference = map[string]UpdateRoverClusterDetailsShippingPreferenceEnum{
	"ORACLE_SHIPPED":  UpdateRoverClusterDetailsShippingPreferenceOracleShipped,
	"CUSTOMER_PICKUP": UpdateRoverClusterDetailsShippingPreferenceCustomerPickup,
}

// GetUpdateRoverClusterDetailsShippingPreferenceEnumValues Enumerates the set of values for UpdateRoverClusterDetailsShippingPreferenceEnum
func GetUpdateRoverClusterDetailsShippingPreferenceEnumValues() []UpdateRoverClusterDetailsShippingPreferenceEnum {
	values := make([]UpdateRoverClusterDetailsShippingPreferenceEnum, 0)
	for _, v := range mappingUpdateRoverClusterDetailsShippingPreference {
		values = append(values, v)
	}
	return values
}
