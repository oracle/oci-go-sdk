// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ScaleAnalyticsInstanceDetails Input payload to scale an Analytics instance up or down.
type ScaleAnalyticsInstanceDetails struct {
	Capacity *Capacity `mandatory:"true" json:"capacity"`
}

func (m ScaleAnalyticsInstanceDetails) String() string {
	return common.PointerString(m)
}
