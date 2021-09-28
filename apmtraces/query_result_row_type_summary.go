// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Traces API
//
// API for APM Trace service. Use this API to query the Traces and associated Spans.
//

package apmtraces

import (
	"github.com/oracle/oci-go-sdk/v48/common"
)

// QueryResultRowTypeSummary A summary of the datatype, unit and related metadata of an individual row element of a query result row that is returned.
type QueryResultRowTypeSummary struct {

	// Datatype of the query result row element.
	DataType *string `mandatory:"false" json:"dataType"`

	// Granular unit in which the query result row element's data is represented.
	Unit *string `mandatory:"false" json:"unit"`

	// Alias name if an alias is used for the query result row element or an assigned display name from the query language
	// in some default cases.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Actual show expression in the user typed query that produced this column.
	Expression *string `mandatory:"false" json:"expression"`

	// A query result row type summary object that represents a nested table structure.
	QueryResultRowTypeSummaries []QueryResultRowTypeSummary `mandatory:"false" json:"queryResultRowTypeSummaries"`
}

func (m QueryResultRowTypeSummary) String() string {
	return common.PointerString(m)
}
