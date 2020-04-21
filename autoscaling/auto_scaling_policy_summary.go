// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// APIs for dynamically scaling Compute resources to meet application requirements.
// For information about the Compute service, see Overview of the Compute Service (https://docs.cloud.oracle.com/Content/Compute/Concepts/computeoverview.htm).
//

package autoscaling

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutoScalingPolicySummary Summary information for an autoscaling policy.
type AutoScalingPolicySummary struct {

	// The ID of the autoscaling policy that is assigned after creation.
	Id *string `mandatory:"true" json:"id"`

	// The type of autoscaling policy.
	PolicyType *string `mandatory:"true" json:"policyType"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m AutoScalingPolicySummary) String() string {
	return common.PointerString(m)
}
