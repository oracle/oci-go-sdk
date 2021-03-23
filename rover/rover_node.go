// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"github.com/oracle/oci-go-sdk/v37/common"
)

// RoverNode Information about a RoverNode.
type RoverNode struct {

	// The OCID of RoverNode.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the RoverNode.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the RoverNode.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The cluster ID if the node is part of a cluster.
	ClusterId *string `mandatory:"false" json:"clusterId"`

	// The type of node indicating if it belongs to a cluster
	NodeType NodeTypeEnum `mandatory:"false" json:"nodeType,omitempty"`

	// The type of enclosure rover node is shipped in.
	EnclosureType EnclosureTypeEnum `mandatory:"false" json:"enclosureType,omitempty"`

	// Serial number of the node.
	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	// The time the the RoverNode was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// A property that can contain details on the lifecycle.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	CustomerShippingAddress *ShippingAddress `mandatory:"false" json:"customerShippingAddress"`

	// List of existing workloads that should be provisioned on the node.
	NodeWorkloads []RoverWorkload `mandatory:"false" json:"nodeWorkloads"`

	// Date and time when customer received tne node.
	TimeCustomerReceieved *common.SDKTime `mandatory:"false" json:"timeCustomerReceieved"`

	// Date and time when customer returned the node.
	TimeCustomerReturned *common.SDKTime `mandatory:"false" json:"timeCustomerReturned"`

	// Tracking information for device shipping.
	DeliveryTrackingInfo *string `mandatory:"false" json:"deliveryTrackingInfo"`

	// Root password for the rover node.
	SuperUserPassword *string `mandatory:"false" json:"superUserPassword"`

	// Password to unlock the rover node.
	UnlockPassphrase *string `mandatory:"false" json:"unlockPassphrase"`

	// Name of point of contact for this order if customer is picking up.
	PointOfContact *string `mandatory:"false" json:"pointOfContact"`

	// Phone number of point of contact for this order if customer is picking up.
	PointOfContactPhoneNumber *string `mandatory:"false" json:"pointOfContactPhoneNumber"`

	// Preference for device delivery.
	ShippingPreference RoverNodeShippingPreferenceEnum `mandatory:"false" json:"shippingPreference,omitempty"`

	// Shipping vendor of choice for orace to customer shipping.
	ShippingVendor *string `mandatory:"false" json:"shippingVendor"`

	// Expected date when customer wants to pickup the device if they chose customer pickup.
	TimePickupExpected *common.SDKTime `mandatory:"false" json:"timePickupExpected"`

	// Start time for the window to pickup the device from customer.
	TimeReturnWindowStarts *common.SDKTime `mandatory:"false" json:"timeReturnWindowStarts"`

	// Tracking Url for the shipped RoverNode.
	OracleShippingTrackingUrl *string `mandatory:"false" json:"oracleShippingTrackingUrl"`

	// End time for the window to pickup the device from customer.
	TimeReturnWindowEnds *common.SDKTime `mandatory:"false" json:"timeReturnWindowEnds"`

	// Uri to download return shipping label.
	ReturnShippingLabelUri *string `mandatory:"false" json:"returnShippingLabelUri"`

	// The public key of the resource principal
	PublicKey *string `mandatory:"false" json:"publicKey"`

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

func (m RoverNode) String() string {
	return common.PointerString(m)
}

// RoverNodeShippingPreferenceEnum Enum with underlying type: string
type RoverNodeShippingPreferenceEnum string

// Set of constants representing the allowable values for RoverNodeShippingPreferenceEnum
const (
	RoverNodeShippingPreferenceOracleShipped  RoverNodeShippingPreferenceEnum = "ORACLE_SHIPPED"
	RoverNodeShippingPreferenceCustomerPickup RoverNodeShippingPreferenceEnum = "CUSTOMER_PICKUP"
)

var mappingRoverNodeShippingPreference = map[string]RoverNodeShippingPreferenceEnum{
	"ORACLE_SHIPPED":  RoverNodeShippingPreferenceOracleShipped,
	"CUSTOMER_PICKUP": RoverNodeShippingPreferenceCustomerPickup,
}

// GetRoverNodeShippingPreferenceEnumValues Enumerates the set of values for RoverNodeShippingPreferenceEnum
func GetRoverNodeShippingPreferenceEnumValues() []RoverNodeShippingPreferenceEnum {
	values := make([]RoverNodeShippingPreferenceEnum, 0)
	for _, v := range mappingRoverNodeShippingPreference {
		values = append(values, v)
	}
	return values
}
