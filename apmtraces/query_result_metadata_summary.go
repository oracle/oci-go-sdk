// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Traces API
//
// API for APM Trace service. Use this API to query the Traces and associated Spans.
//

package apmtraces

import (
	"github.com/oracle/oci-go-sdk/v38/common"
)

// QueryResultMetadataSummary Summary containing the metadata about the query result set.
type QueryResultMetadataSummary struct {

	// A collection of QueryResultRowTypeSummary objects that describe the type and properties of the individual row elements of the query rows
	// being returned.  The ith element in this list contains the QueryResultRowTypeSummary of the ith key value pair in the QueryResultRowData map.
	QueryResultRowTypeSummaries []QueryResultRowTypeSummary `mandatory:"false" json:"queryResultRowTypeSummaries"`

	// Source of the query result set (traces, spans, etc).
	SourceName *string `mandatory:"false" json:"sourceName"`

	// Columns or attributes of the query rows  which are group by values.  This is a list of ResultsGroupedBy summary objects,
	// and the list will contain as many elements as the attributes and aggregate functions in the group by clause in the select query.
	QueryResultsGroupedBy []QueryResultsGroupedBySummary `mandatory:"false" json:"queryResultsGroupedBy"`

	// Order by which the query results are organized.  This is a list of queryResultsOrderedBy summary objects, and the list
	// will contain more than one OrderedBy summary object, if the sort was multidimensional.
	QueryResultsOrderedBy []QueryResultsOrderedBySummary `mandatory:"false" json:"queryResultsOrderedBy"`

	// Interval for the time series function in minutes.
	TimeSeriesIntervalInMins *int `mandatory:"false" json:"timeSeriesIntervalInMins"`
}

func (m QueryResultMetadataSummary) String() string {
	return common.PointerString(m)
}
