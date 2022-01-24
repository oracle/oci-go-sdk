// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling API
//
// Use Data Labeling API to create Annotations on Images, Texts & Documents, and generate snapshots.
//

package datalabelingservicedataplane

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// AnnotationAnalyticsAggregation Aggregation entities are required by the API consistency guidelines for API Consistency Guidelines#AnalyticsAPIs.  These are used to summarize annotations for a given dataset and will be used to populate UI elements.  Aggregations need to have the fields that identify the exact scope that they're summarizing.  Any filters applied to the list API, have to show up in the aggregation.
type AnnotationAnalyticsAggregation struct {

	// The count of the matching results.
	Count *float32 `mandatory:"true" json:"count"`

	// The OCID of the dataset the annotations belong to.
	DatasetId *string `mandatory:"true" json:"datasetId"`

	// The OCID of the compartment containing the annotations.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	Dimensions *AnnotationAggregationDimensions `mandatory:"false" json:"dimensions"`

	// The OCID of the principal which updated the annotation.
	UpdatedBy *string `mandatory:"false" json:"updatedBy"`

	// Describes the lifecycle state.
	LifecycleState AnnotationLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m AnnotationAnalyticsAggregation) String() string {
	return common.PointerString(m)
}
