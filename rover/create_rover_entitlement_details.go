// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"github.com/oracle/oci-go-sdk/v46/common"
)

// CreateRoverEntitlementDetails Information required to create a RoverEntitlement.
type CreateRoverEntitlementDetails struct {

	// The OCID of the compartment containing the RoverEntitlement.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Requestor name for the entitlement.
	RequestorName *string `mandatory:"true" json:"requestorName"`

	// Requestor email for the entitlement.
	RequestorEmail *string `mandatory:"true" json:"requestorEmail"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Details about the entitlement.
	EntitlementDetails *string `mandatory:"false" json:"entitlementDetails"`

	// The current state of the RoverNode.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A property that can contain details on the lifecycle.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// tenant Id.
	TenantId *string `mandatory:"false" json:"tenantId"`

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

func (m CreateRoverEntitlementDetails) String() string {
	return common.PointerString(m)
}
