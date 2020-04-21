// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AnalyticsInstanceSummary Analytics Instance metadata (summary view).
type AnalyticsInstanceSummary struct {

	// The resource OCID.
	Id *string `mandatory:"true" json:"id"`

	// The name of the Analytics instance. This name must be unique in the tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of an instance.
	LifecycleState AnalyticsInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Analytics feature set.
	FeatureSet FeatureSetEnum `mandatory:"true" json:"featureSet"`

	Capacity *Capacity `mandatory:"true" json:"capacity"`

	// The date and time the instance was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Optional description.
	Description *string `mandatory:"false" json:"description"`

	// The license used for the service.
	LicenseType LicenseTypeEnum `mandatory:"false" json:"licenseType,omitempty"`

	// Email address receiving notifications.
	EmailNotification *string `mandatory:"false" json:"emailNotification"`

	// URL of the Analytics service.
	ServiceUrl *string `mandatory:"false" json:"serviceUrl"`

	// The date and time the instance was last updated (in the format defined by RFC3339).
	// This timestamp represents updates made through this API. External events do not
	// influence it.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m AnalyticsInstanceSummary) String() string {
	return common.PointerString(m)
}
