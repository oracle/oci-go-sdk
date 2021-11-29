// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Traces API
//
// API for APM Trace service. Use this API to query the Traces and associated Spans.
//

package apmtraces

import (
	"github.com/oracle/oci-go-sdk/v53/common"
)

// QueryResultsGroupedBySummary Summary of the attribute based on which the query results are grouped by.
type QueryResultsGroupedBySummary struct {

	// Column or attribute in the query result which is a group by value.
	QueryResultsGroupedByColumn *string `mandatory:"false" json:"queryResultsGroupedByColumn"`
}

func (m QueryResultsGroupedBySummary) String() string {
	return common.PointerString(m)
}
