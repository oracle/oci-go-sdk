// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/common"
)

// GenerateAgentObjectNameDetails Generate agent upload name for the given properties
type GenerateAgentObjectNameDetails struct {

	// Log group OCID
	LogGroupId *string `mandatory:"true" json:"logGroupId"`

	// Internal identifier used to uniquely identify the agent upload request
	UniqueId *string `mandatory:"true" json:"uniqueId"`

	// Metadata associated with the upload used during processing
	MetaProperties *string `mandatory:"true" json:"metaProperties"`

	// The time when this upload is created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m GenerateAgentObjectNameDetails) String() string {
	return common.PointerString(m)
}
