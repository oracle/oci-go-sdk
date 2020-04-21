// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Search Service API
//
// Search for resources in your cloud network.
//

package resourcesearch

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ResourceSummaryCollection A summary representation of resources that matched the search criteria.
type ResourceSummaryCollection struct {

	// A list of resources.
	Items []ResourceSummary `mandatory:"false" json:"items"`
}

func (m ResourceSummaryCollection) String() string {
	return common.PointerString(m)
}
