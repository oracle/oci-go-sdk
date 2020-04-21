// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science APIs to organize your data science work, access data and computing resources, and build, train, deploy, and manage models on Oracle Cloud.
//

package datascience

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ModelSummary Summary information for a model.
type ModelSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the model's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the project associated with the model.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the model.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly display name for the resource. Does not have to be unique, and can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the user who created the model.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The date and time the resource was created, in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: 2019-08-25T21:10:29.41Z
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The state of the model.
	LifecycleState ModelLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ModelSummary) String() string {
	return common.PointerString(m)
}
