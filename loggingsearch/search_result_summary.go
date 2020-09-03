// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Search API
//
// Search for logs in your compartements / log groups / log objects.
//

package loggingsearch

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SearchResultSummary Summary of results.
type SearchResultSummary struct {

	// Total number of search results.
	ResultCount *int `mandatory:"false" json:"resultCount"`

	// Total number of field schema information.
	FieldCount *int `mandatory:"false" json:"fieldCount"`
}

func (m SearchResultSummary) String() string {
	return common.PointerString(m)
}
