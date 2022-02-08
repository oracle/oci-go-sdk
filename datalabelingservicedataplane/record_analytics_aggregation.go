// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling API
//
// Use Data Labeling API to create Annotations on Images, Texts & Documents, and generate snapshots.
//

package datalabelingservicedataplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"strings"
)

// RecordAnalyticsAggregation Aggregation entities are required by the API consistency guidelines for API Consistency Guidelines#AnalyticsAPIs.  These are used to summarize record information for a given dataset and are used to populate UI elements.  Aggregations need to have the fields that identify the exact scope that they're summarizing.  Any filters applied to the list API, have to show up in the aggregation.
type RecordAnalyticsAggregation struct {

	// the count of the matching results
	Count *float32 `mandatory:"true" json:"count"`

	// ocid of the dataset the annotation belongs to
	DatasetId *string `mandatory:"true" json:"datasetId"`

	// ocid of the compartment the records
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	Dimensions *RecordAggregationDimensions `mandatory:"false" json:"dimensions"`

	// Describes the lifecycle state.
	LifecycleState RecordLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m RecordAnalyticsAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecordAnalyticsAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := mappingRecordLifecycleStateEnum[string(m.LifecycleState)]; !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRecordLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
