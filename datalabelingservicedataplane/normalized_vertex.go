// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DlsDataPlane API
//
// A description of the DlsDataPlane API.
//

package datalabelingservicedataplane

import (
	"github.com/oracle/oci-go-sdk/v49/common"
)

// NormalizedVertex A NormalizedVertex is a cartesian coordinate that represents a corner between two segments of a polygon.
type NormalizedVertex struct {

	// X axis coordinate
	X *float32 `mandatory:"true" json:"x"`

	// Y axis coordinate
	Y *float32 `mandatory:"true" json:"y"`
}

func (m NormalizedVertex) String() string {
	return common.PointerString(m)
}
