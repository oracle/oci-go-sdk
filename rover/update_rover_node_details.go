// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// UpdateRoverNodeDetails The information required to update a RoverNode.
type UpdateRoverNodeDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Serial number of the node.
	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	CustomerShippingAddress *ShippingAddress `mandatory:"false" json:"customerShippingAddress"`

	// List of existing workloads that should be provisioned on the node.
	NodeWorkloads []RoverWorkload `mandatory:"false" json:"nodeWorkloads"`

	// Root password for the rover node.
	SuperUserPassword *string `mandatory:"false" json:"superUserPassword"`

	// Password to unlock the rover node.
	UnlockPassphrase *string `mandatory:"false" json:"unlockPassphrase"`

	// Name of point of contact for this order if customer is picking up.
	PointOfContact *string `mandatory:"false" json:"pointOfContact"`

	// Phone number of point of contact for this order if customer is picking up.
	PointOfContactPhoneNumber *string `mandatory:"false" json:"pointOfContactPhoneNumber"`

	// Tracking Url for the shipped FmsRoverNode.
	OracleShippingTrackingUrl *string `mandatory:"false" json:"oracleShippingTrackingUrl"`

	// Preference for device delivery.
	ShippingPreference UpdateRoverNodeDetailsShippingPreferenceEnum `mandatory:"false" json:"shippingPreference,omitempty"`

	// Shipping vendor of choice for orace to customer shipping.
	ShippingVendor *string `mandatory:"false" json:"shippingVendor"`

	// Expected date when customer wants to pickup the device if they chose customer pickup.
	TimePickupExpected *common.SDKTime `mandatory:"false" json:"timePickupExpected"`

	// The current state of the RoverNode.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The type of enclosure rover nodes in this cluster are shipped in.
	EnclosureType EnclosureTypeEnum `mandatory:"false" json:"enclosureType,omitempty"`

	// A property that can contain details on the lifecycle.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Start time for the window to pickup the device from customer.
	TimeReturnWindowStarts *common.SDKTime `mandatory:"false" json:"timeReturnWindowStarts"`

	// End time for the window to pickup the device from customer.
	TimeReturnWindowEnds *common.SDKTime `mandatory:"false" json:"timeReturnWindowEnds"`

	// The public key of the resource principal
	PublicKey *string `mandatory:"false" json:"publicKey"`

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

func (m UpdateRoverNodeDetails) String() string {
	return common.PointerString(m)
}

// UpdateRoverNodeDetailsShippingPreferenceEnum Enum with underlying type: string
type UpdateRoverNodeDetailsShippingPreferenceEnum string

// Set of constants representing the allowable values for UpdateRoverNodeDetailsShippingPreferenceEnum
const (
	UpdateRoverNodeDetailsShippingPreferenceOracleShipped  UpdateRoverNodeDetailsShippingPreferenceEnum = "ORACLE_SHIPPED"
	UpdateRoverNodeDetailsShippingPreferenceCustomerPickup UpdateRoverNodeDetailsShippingPreferenceEnum = "CUSTOMER_PICKUP"
)

var mappingUpdateRoverNodeDetailsShippingPreference = map[string]UpdateRoverNodeDetailsShippingPreferenceEnum{
	"ORACLE_SHIPPED":  UpdateRoverNodeDetailsShippingPreferenceOracleShipped,
	"CUSTOMER_PICKUP": UpdateRoverNodeDetailsShippingPreferenceCustomerPickup,
}

// GetUpdateRoverNodeDetailsShippingPreferenceEnumValues Enumerates the set of values for UpdateRoverNodeDetailsShippingPreferenceEnum
func GetUpdateRoverNodeDetailsShippingPreferenceEnumValues() []UpdateRoverNodeDetailsShippingPreferenceEnum {
	values := make([]UpdateRoverNodeDetailsShippingPreferenceEnum, 0)
	for _, v := range mappingUpdateRoverNodeDetailsShippingPreference {
		values = append(values, v)
	}
	return values
}
