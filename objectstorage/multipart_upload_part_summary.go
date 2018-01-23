// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// APIs for managing buckets and objects.
//

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
)

// MultipartUploadPartSummary To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm.
type MultipartUploadPartSummary struct {

	// the current entity tag for the part.
	Etag *string `mandatory:"true" json:"etag,omitempty"`

	// the MD5 hash of the bytes of the part.
	Md5 *string `mandatory:"true" json:"md5,omitempty"`

	// the size of the part in bytes.
	Size *int `mandatory:"true" json:"size,omitempty"`

	// the part number for this part.
	PartNumber *int `mandatory:"true" json:"partNumber,omitempty"`
}

func (m MultipartUploadPartSummary) String() string {
	return common.PointerString(m)
}
