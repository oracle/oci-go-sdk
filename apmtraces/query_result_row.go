// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Traces API
//
// API for APM Trace service. Use this API to query the Traces and associated Spans.
//

package apmtraces

import (
	"github.com/oracle/oci-go-sdk/v46/common"
)

// QueryResultRow An object that represents a single row of the query result.  It contains the queryResultRowData object that holds the actual data
// represented by the elements of the query result row, and a queryResultRowMetadata object that holds the metadata about the data contained in
// the query result row.
type QueryResultRow struct {

	// A map containing the actual data represented by a single row of the query result.
	// The key is the column name or attribute specified in the show clause, or an aggregate function in the show clause.
	// The value is the actual value of that attribute or aggregate function of the corresponding single row of the query result set.
	// If an alias name is specified for an attribute or an aggregate function, then the key will be the alias name specified in the show
	// clause.  If an alias name is not specified for the group by aggregate function in the show clause, then the corresponding key
	// will be the appropriate aggregate_function_name_column_name (Eg: count(traces) will be keyed as count_traces).  For more details
	// on the supported aggregate functions, look at the APM confluence doc on High Level Query Aggregation.  The datatype of the value
	// is presented in the queryResultRowTypeSummaries list in the queryResultMetadata structure, where the ith queryResultRowTypeSummary object
	// represents the datatype of the ith value when this map is iterated in order.
	QueryResultRowData map[string]interface{} `mandatory:"true" json:"queryResultRowData"`

	// A map containing metadata or add-on data for the data presented in the queryResultRowData map.  Data required to present drill down
	// information from the queryResultRowData is presented as key value pairs.
	QueryResultRowMetadata map[string]interface{} `mandatory:"true" json:"queryResultRowMetadata"`
}

func (m QueryResultRow) String() string {
	return common.PointerString(m)
}
