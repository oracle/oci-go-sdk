// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling API
//
// Use Data Labeling API to create Annotations on Images, Texts & Documents, and generate snapshots.
//

package datalabelingservicedataplane

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// BoundingPolygon A polygon used to describe the location of an object.
type BoundingPolygon struct {

	// The normalized vertices that make up the polygon.  They are in the order of the segments they connect.
	NormalizedVertices []NormalizedVertex `mandatory:"true" json:"normalizedVertices"`
}

func (m BoundingPolygon) String() string {
	return common.PointerString(m)
}
