// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TagSummary A tag definition that belongs to a specific tag namespace.
type TagSummary struct {

	// The OCID of the compartment that contains the tag definition.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID of the tag definition.
	Id *string `mandatory:"false" json:"id"`

	// The name assigned to the tag during creation. This is the tag key definition.
	// The name must be unique within the tag namespace and cannot be changed.
	Name *string `mandatory:"false" json:"name"`

	// The description you assign to the tag.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Whether the tag is retired.
	// See Retiring Key Definitions and Namespace Definitions (https://docs.cloud.oracle.com/Content/Identity/Concepts/taggingoverview.htm#Retiring).
	IsRetired *bool `mandatory:"false" json:"isRetired"`

	// The tag's current state. After creating a tag, make sure its `lifecycleState` is ACTIVE before using it. After retiring a tag, make sure its `lifecycleState` is INACTIVE before using it. If you delete a tag, you cannot delete another tag until the deleted tag's `lifecycleState` changes from DELETING to DELETED.
	LifecycleState TagLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Date and time the tag was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Indicates whether the tag is enabled for cost tracking.
	IsCostTracking *bool `mandatory:"false" json:"isCostTracking"`
}

func (m TagSummary) String() string {
	return common.PointerString(m)
}
