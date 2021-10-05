// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Ingestion API
//
// Use the Logging Ingestion API to ingest your application logs.
//

package loggingingestion

import (
	"github.com/oracle/oci-go-sdk/v49/common"
)

// LogEntryBatch A single batch of Log Entries.
type LogEntryBatch struct {

	// List of data entries.
	Entries []LogEntry `mandatory:"true" json:"entries"`

	// Source of the logs that generated the message. This could be the
	// instance name, hostname, or the source used to read the event. For example, "ServerA".
	Source *string `mandatory:"true" json:"source"`

	// This field signifies the type of logs being ingested.
	// For example: ServerA.requestLogs.
	Type *string `mandatory:"true" json:"type"`

	// The timestamp for all log entries in this batch. This can be
	// considered as the default timestamp for each entry, unless it is
	// overwritten by the entry time. An RFC3339-formatted date-time string
	// with milliseconds precision.
	Defaultlogentrytime *common.SDKTime `mandatory:"true" json:"defaultlogentrytime"`

	// This optional field is useful for specifying the specific sub-resource
	// or input file used to read the event.
	// For example: "/var/log/application.log".
	Subject *string `mandatory:"false" json:"subject"`
}

func (m LogEntryBatch) String() string {
	return common.PointerString(m)
}
