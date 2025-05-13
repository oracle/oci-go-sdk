// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Ingestion API
//
// Use the Logging Ingestion API to ingest your application logs.
//

package loggingingestion

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogEntry Contains the log content with the associated timestamp and ID. Each
// entry should be less than 1 MB size.
type LogEntry struct {

	// The log entry content.
	Data *string `mandatory:"true" json:"data"`

	// UUID uniquely representing this logEntry. This is not an OCID related
	// to any oracle resource.
	Id *string `mandatory:"true" json:"id"`

	// Optional. The timestamp associated with the log entry. An RFC3339-formatted date-time string with milliseconds precision.
	// If unspecified, defaults to PutLogsDetails.defaultlogentrytime.
	Time *common.SDKTime `mandatory:"false" json:"time"`
}

func (m LogEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
