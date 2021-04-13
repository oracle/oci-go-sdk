// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Traces API
//
// API for APM Trace service. Use this API to query the Traces and associated Spans.
//

package apmtraces

import (
	"github.com/oracle/oci-go-sdk/v39/common"
)

// QueryResultResponse A response containing a collection of query rows (selected attributes and aggregations) filtered, grouped and
// sorted by the specified criteria from the query that is run, and the associated summary describing the corresponding
// query result metadata.
type QueryResultResponse struct {
	QueryResultMetadataSummary *QueryResultMetadataSummary `mandatory:"true" json:"queryResultMetadataSummary"`

	// A collection of objects with each object representing an individual row of the query result set.  The total number of objects
	// returned in this collection correspond to the total number of rows returned by the actual query that is run against
	// the queried entity.
	QueryResultRows []QueryResultRow `mandatory:"true" json:"queryResultRows"`
}

func (m QueryResultResponse) String() string {
	return common.PointerString(m)
}
