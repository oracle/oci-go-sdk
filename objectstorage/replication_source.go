// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.oracle.com/iaas/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.oracle.com/iaas/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReplicationSource The details of a replication source bucket that replicates to a target destination bucket.
type ReplicationSource struct {

	// The name of the policy.
	PolicyName *string `mandatory:"true" json:"policyName"`

	// The source region replicating data from, for example "us-ashburn-1".
	SourceRegionName *string `mandatory:"true" json:"sourceRegionName"`

	// The source bucket replicating data from.
	SourceBucketName *string `mandatory:"true" json:"sourceBucketName"`
}

func (m ReplicationSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReplicationSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
