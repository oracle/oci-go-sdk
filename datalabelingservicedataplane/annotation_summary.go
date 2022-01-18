// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling API
//
// Use Data Labeling API to create Annotations on Images, Texts & Documents, and generate snapshots.
//

package datalabelingservicedataplane

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// AnnotationSummary An annotation summary is the representation returned in list views.  It is usually a subset of the full annotation entity and should not contain any potentially sensitive information.
type AnnotationSummary struct {

	// The OCID of the annotation.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the annotation was created, in the timestamp format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the resource was updated, in the timestamp format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID of the record annotated.
	RecordId *string `mandatory:"true" json:"recordId"`

	// The OCID of the compartment for the annotation.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Describes the lifecycle state.
	LifecycleState AnnotationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only.
	// For example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m AnnotationSummary) String() string {
	return common.PointerString(m)
}
