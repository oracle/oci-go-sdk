// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"github.com/oracle/oci-go-sdk/v47/common"
)

// RoverCluster Description of RoverCluster.
type RoverCluster struct {

	// The OCID of RoverCluster.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the RoverCluster.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Size of the cluster.
	ClusterSize *int `mandatory:"true" json:"clusterSize"`

	// The current state of the RoverCluster.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the the RoverCluster was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// A property that can contain details on the lifecycle.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	CustomerShippingAddress *ShippingAddress `mandatory:"false" json:"customerShippingAddress"`

	// The summary of nodes that are part of this cluster.
	Nodes []RoverNodeSummary `mandatory:"false" json:"nodes"`

	// The type of enclosure rover nodes in this cluster are shipped in.
	EnclosureType EnclosureTypeEnum `mandatory:"false" json:"enclosureType,omitempty"`

	// Time when customer received the cluster.
	TimeCustomerReceived *common.SDKTime `mandatory:"false" json:"timeCustomerReceived"`

	// Time when customer returned the cluster.
	TimeCustomerReturned *common.SDKTime `mandatory:"false" json:"timeCustomerReturned"`

	// Tracking information for device shipping.
	DeliveryTrackingInfo *string `mandatory:"false" json:"deliveryTrackingInfo"`

	// List of existing workloads that should be provisioned on the nodes.
	ClusterWorkloads []RoverWorkload `mandatory:"false" json:"clusterWorkloads"`

	// Root password for the rover cluster.
	SuperUserPassword *string `mandatory:"false" json:"superUserPassword"`

	// Password to unlock the rover cluster.
	UnlockPassphrase *string `mandatory:"false" json:"unlockPassphrase"`

	// Name of point of contact for this order if customer is picking up.
	PointOfContact *string `mandatory:"false" json:"pointOfContact"`

	// Phone number of point of contact for this order if customer is picking up.
	PointOfContactPhoneNumber *string `mandatory:"false" json:"pointOfContactPhoneNumber"`

	// Preference for device delivery.
	ShippingPreference RoverClusterShippingPreferenceEnum `mandatory:"false" json:"shippingPreference,omitempty"`

	// Tracking Url for the shipped Rover Cluster.
	OracleShippingTrackingUrl *string `mandatory:"false" json:"oracleShippingTrackingUrl"`

	// Shipping vendor of choice for orace to customer shipping.
	ShippingVendor *string `mandatory:"false" json:"shippingVendor"`

	// Expected date when customer wants to pickup the device if they chose customer pickup.
	TimePickupExpected *common.SDKTime `mandatory:"false" json:"timePickupExpected"`

	// Start time for the window to pickup the device from customer.
	TimeReturnWindowStarts *common.SDKTime `mandatory:"false" json:"timeReturnWindowStarts"`

	// End time for the window to pickup the device from customer.
	TimeReturnWindowEnds *common.SDKTime `mandatory:"false" json:"timeReturnWindowEnds"`

	// Uri to download return shipping label.
	ReturnShippingLabelUri *string `mandatory:"false" json:"returnShippingLabelUri"`

	// The link to pre-authenticated request for a bucket where image workloads are moved.
	ImageExportPar *string `mandatory:"false" json:"imageExportPar"`

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

func (m RoverCluster) String() string {
	return common.PointerString(m)
}

// RoverClusterShippingPreferenceEnum Enum with underlying type: string
type RoverClusterShippingPreferenceEnum string

// Set of constants representing the allowable values for RoverClusterShippingPreferenceEnum
const (
	RoverClusterShippingPreferenceOracleShipped  RoverClusterShippingPreferenceEnum = "ORACLE_SHIPPED"
	RoverClusterShippingPreferenceCustomerPickup RoverClusterShippingPreferenceEnum = "CUSTOMER_PICKUP"
)

var mappingRoverClusterShippingPreference = map[string]RoverClusterShippingPreferenceEnum{
	"ORACLE_SHIPPED":  RoverClusterShippingPreferenceOracleShipped,
	"CUSTOMER_PICKUP": RoverClusterShippingPreferenceCustomerPickup,
}

// GetRoverClusterShippingPreferenceEnumValues Enumerates the set of values for RoverClusterShippingPreferenceEnum
func GetRoverClusterShippingPreferenceEnumValues() []RoverClusterShippingPreferenceEnum {
	values := make([]RoverClusterShippingPreferenceEnum, 0)
	for _, v := range mappingRoverClusterShippingPreference {
		values = append(values, v)
	}
	return values
}
