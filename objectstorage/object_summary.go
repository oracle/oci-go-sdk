// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// APIs for managing buckets and objects.
//

package objectstorage

import (
	"github.com/oci-go-sdk/common"
)

// ObjectSummary. To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type ObjectSummary struct {

	// The name of the object.
	Name *string `mandatory:"true" json:"name,omitempty"`

	// Size of the object in bytes.
	Size *int64 `mandatory:"false" json:"size,omitempty"`

	// Base64-encoded MD5 hash of the object data.
	Md5 *string `mandatory:"false" json:"md5,omitempty"`

	// Date and time of object creation.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`
}

func (model ObjectSummary) String() string {
	return common.PointerString(model)
}
