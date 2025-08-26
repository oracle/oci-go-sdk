// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RoverBundleStatus The status of the rover bundle status by a specified work request id.
type RoverBundleStatus struct {

	// The progress of the workflow.
	Status RoverBundleStatusStatusEnum `mandatory:"true" json:"status"`

	// Percentage of the work request completed.
	PercentComplete *float32 `mandatory:"false" json:"percentComplete"`

	// The date and time the work request was created. An RFC3339 formatted datetime string.
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// The date and time the work request was started. An RFC3339 formatted datetime string.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the work request was finished. An RFC3339 formatted datetime string.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// The full name of the bundle.
	BundleName *string `mandatory:"false" json:"bundleName"`

	// The error message if work request fails.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`
}

func (m RoverBundleStatus) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RoverBundleStatus) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRoverBundleStatusStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetRoverBundleStatusStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RoverBundleStatusStatusEnum Enum with underlying type: string
type RoverBundleStatusStatusEnum string

// Set of constants representing the allowable values for RoverBundleStatusStatusEnum
const (
	RoverBundleStatusStatusAccepted   RoverBundleStatusStatusEnum = "ACCEPTED"
	RoverBundleStatusStatusInProgress RoverBundleStatusStatusEnum = "IN_PROGRESS"
	RoverBundleStatusStatusFailed     RoverBundleStatusStatusEnum = "FAILED"
	RoverBundleStatusStatusCompleted  RoverBundleStatusStatusEnum = "COMPLETED"
	RoverBundleStatusStatusCanceling  RoverBundleStatusStatusEnum = "CANCELING"
	RoverBundleStatusStatusCanceled   RoverBundleStatusStatusEnum = "CANCELED"
)

var mappingRoverBundleStatusStatusEnum = map[string]RoverBundleStatusStatusEnum{
	"ACCEPTED":    RoverBundleStatusStatusAccepted,
	"IN_PROGRESS": RoverBundleStatusStatusInProgress,
	"FAILED":      RoverBundleStatusStatusFailed,
	"COMPLETED":   RoverBundleStatusStatusCompleted,
	"CANCELING":   RoverBundleStatusStatusCanceling,
	"CANCELED":    RoverBundleStatusStatusCanceled,
}

var mappingRoverBundleStatusStatusEnumLowerCase = map[string]RoverBundleStatusStatusEnum{
	"accepted":    RoverBundleStatusStatusAccepted,
	"in_progress": RoverBundleStatusStatusInProgress,
	"failed":      RoverBundleStatusStatusFailed,
	"completed":   RoverBundleStatusStatusCompleted,
	"canceling":   RoverBundleStatusStatusCanceling,
	"canceled":    RoverBundleStatusStatusCanceled,
}

// GetRoverBundleStatusStatusEnumValues Enumerates the set of values for RoverBundleStatusStatusEnum
func GetRoverBundleStatusStatusEnumValues() []RoverBundleStatusStatusEnum {
	values := make([]RoverBundleStatusStatusEnum, 0)
	for _, v := range mappingRoverBundleStatusStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetRoverBundleStatusStatusEnumStringValues Enumerates the set of values in String for RoverBundleStatusStatusEnum
func GetRoverBundleStatusStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"COMPLETED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingRoverBundleStatusStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRoverBundleStatusStatusEnum(val string) (RoverBundleStatusStatusEnum, bool) {
	enum, ok := mappingRoverBundleStatusStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
