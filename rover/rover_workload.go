// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v62/common"
	"strings"
)

// RoverWorkload Information about a RoverWorkload.
type RoverWorkload struct {

	// The OCID of the compartment containing the workload.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Unique Oracle ID (OCID) that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The type of workload
	WorkloadType *string `mandatory:"true" json:"workloadType"`

	// Name of the Rover Workload
	Name *string `mandatory:"false" json:"name"`

	// Size of the workload.
	Size *string `mandatory:"false" json:"size"`

	// Number of objects in a workload.
	ObjectCount *string `mandatory:"false" json:"objectCount"`

	// Prefix to filter objects in case it is a bucket.
	Prefix *string `mandatory:"false" json:"prefix"`

	// Start of the range in a bucket.
	RangeStart *string `mandatory:"false" json:"rangeStart"`

	// End of the range in a bucket.
	RangeEnd *string `mandatory:"false" json:"rangeEnd"`

	// The compute work request id to track progress of custom image exports.
	WorkRequestId *string `mandatory:"false" json:"workRequestId"`
}

func (m RoverWorkload) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RoverWorkload) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
