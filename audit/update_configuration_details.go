// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Audit API
//
// API for the Audit Service. You can use this API for queries, but not bulk-export operations.
//

package audit

import (
	"github.com/oracle/oci-go-sdk/common"
)

type UpdateConfigurationDetails struct {

	// The retention period days
	RetentionPeriodDays *int `mandatory:"false" json:"retentionPeriodDays,omitempty"`
}

func (m UpdateConfigurationDetails) String() string {
	return common.PointerString(m)
}
