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

// SearchResponse Search response object.
type SearchResponse struct {
	Summary *SearchResultSummary `mandatory:"true" json:"summary"`

	// List of search results
	Results []SearchResult `mandatory:"false" json:"results"`

	// List of log field schema information.
	Fields []FieldInfo `mandatory:"false" json:"fields"`
}

func (m SearchResponse) String() string {
	return common.PointerString(m)
}
