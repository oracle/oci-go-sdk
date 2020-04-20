// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SourceSummary The Source object. Sources represent external locations from which
// applications may be imported into an OCI tenancy.
type SourceSummary struct {

	// Unique identifier (OCID) for the source
	Id *string `mandatory:"false" json:"id"`

	// The type of source environment
	Type SourceTypesEnum `mandatory:"false" json:"type,omitempty"`

	// Unique idenfifier (OCID) for the compartment where the Source is located.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Human-readable name of the source.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the source.
	Description *string `mandatory:"false" json:"description"`

	// The date and time at which the source was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The current state of the Source
	LifecycleState SourceLifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Details about the current lifecycle state
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m SourceSummary) String() string {
	return common.PointerString(m)
}
