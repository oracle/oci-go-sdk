// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PublicLoggingDataplane API
//
// PublicLoggingDataplane API specification
//

package loggingingestion

import (
	"github.com/oracle/oci-go-sdk/common"
)

// LogEntry Contains the content of the log with associated timestamp and id. Each
// entry should be less than 1 MB size.
type LogEntry struct {

	// The content of the log entry.
	Data *string `mandatory:"true" json:"data"`

	// UUID uniquely representing this logEntry. This is not an OCID related
	// to any oracle resource.
	Id *string `mandatory:"true" json:"id"`

	// Optional. The timestamp associated with the log entry. Defaults to
	// PutLogsDetails.defaultlogentrytime if unspecified. An RFC3339 formatted
	// datetime string.
	Time *common.SDKTime `mandatory:"false" json:"time"`
}

func (m LogEntry) String() string {
	return common.PointerString(m)
}
