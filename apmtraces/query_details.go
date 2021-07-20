// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Traces API
//
// API for APM Trace service. Use this API to query the Traces and associated Spans.
//

package apmtraces

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// QueryDetails Request object containing the query to be run against our repository.
type QueryDetails struct {

	// APM defined query string to run against our repository and return results for.
	QueryText *string `mandatory:"false" json:"queryText"`
}

func (m QueryDetails) String() string {
	return common.PointerString(m)
}
