// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v47/common"
)

// SenderInvitation The invitation model that the sender owns.
type SenderInvitation struct {

	// OCID of the sender invitation.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the sender tenancy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCID of the recipient tenancy.
	RecipientTenancyId *string `mandatory:"true" json:"recipientTenancyId"`

	// Lifecycle state of the sender invitation.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Status of the sender invitation.
	Status SenderInvitationStatusEnum `mandatory:"true" json:"status"`

	// Date-time when this sender invitation was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// OCID of the corresponding recipient invitation.
	RecipientInvitationId *string `mandatory:"false" json:"recipientInvitationId"`

	// A user-created name to describe the invitation. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Date-time when this sender invitation was last updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Email address of the recipient.
	RecipientEmailAddress *string `mandatory:"false" json:"recipientEmailAddress"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m SenderInvitation) String() string {
	return common.PointerString(m)
}
