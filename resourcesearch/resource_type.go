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

// ResourceType Defines a type of resource that you can find with a search or query.
type ResourceType struct {

	// The unique name of the resource type, which matches the value returned as part of the ResourceSummary object.
	Name *string `mandatory:"true" json:"name"`

	// List of all the fields and their value type that are indexed for querying.
	Fields []QueryableFieldDescription `mandatory:"true" json:"fields"`
}

func (m ResourceType) String() string {
	return common.PointerString(m)
}
