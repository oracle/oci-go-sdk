// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Labeling API
//
// Use Data Labeling API to create Annotations on Images, Texts & Documents, and generate snapshots.
//

package datalabelingservicedataplane

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// InitialRecordGenerationConfiguration The initial generate records configuration It generates records from the dataset's source.
type InitialRecordGenerationConfiguration struct {

	// The maximum number of records to generate.
	Limit *float32 `mandatory:"false" json:"limit"`
}

func (m InitialRecordGenerationConfiguration) String() string {
	return common.PointerString(m)
}
