// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Traces API
//
// API for APM Trace service. Use this API to query the Traces and associated Spans.
//

package apmtraces

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"strings"
)

// QueryResultsOrderedBySummary Summary of the sort attribute and order by which the query results are organized.
type QueryResultsOrderedBySummary struct {

	// Attribute by which the query result is sorted by.
	QueryResultsOrderedBy *string `mandatory:"false" json:"queryResultsOrderedBy"`

	// The sort order for the attribute, either 'ASC' or 'DESC'.
	QueryResultsSortOrder *string `mandatory:"false" json:"queryResultsSortOrder"`
}

func (m QueryResultsOrderedBySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueryResultsOrderedBySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
