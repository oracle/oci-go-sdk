// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DlsDataPlane API
//
// A description of the DlsDataPlane API.
//

package datalabelingservicedataplane

import (
	"github.com/oracle/oci-go-sdk/v49/common"
)

// AnnotationSummary An annotation summary is the representation returned in list views.  It is usually a subset of the full annotation entity and should not contain any potentially sensitive information.
type AnnotationSummary struct {

	// The OCID of the annotation
	Id *string `mandatory:"true" json:"id"`

	// The date and time the annotation was created, in the timestamp format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the resource was updated, in the timestamp format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID of the record annotated
	RecordId *string `mandatory:"true" json:"recordId"`

	// The OCID of the compartment for the annotation
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Describes the lifecycle state.
	LifecycleState AnnotationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m AnnotationSummary) String() string {
	return common.PointerString(m)
}
