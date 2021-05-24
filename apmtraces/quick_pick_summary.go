// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Traces API
//
// API for APM Trace service. Use this API to query the Traces and associated Spans.
//

package apmtraces

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// QuickPickSummary Summary of quick pick query objects that contains the quick pick queries.
type QuickPickSummary struct {

	// Quick pick name for the query.
	QuickPickName *string `mandatory:"true" json:"quickPickName"`

	// Query for the quick pick.
	QuickPickQuery *string `mandatory:"true" json:"quickPickQuery"`
}

func (m QuickPickSummary) String() string {
	return common.PointerString(m)
}
