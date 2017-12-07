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

// CreateBucketDetails. To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type CreateBucketDetails struct {

	// The name of the bucket. Valid characters are uppercase or lowercase letters,
	// numbers, and dashes. Bucket names must be unique within the namespace.
	Name *string `mandatory:"true" json:"name,omitempty"`

	// The ID of the compartment in which to create the bucket.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// Arbitrary string, up to 4KB, of keys and values for user-defined metadata.
	Metadata map[string]string `mandatory:"false" json:"metadata,omitempty"`

	// The type of public access available on this bucket. Allows authenticated caller to access the bucket or
	// contents of this bucket. By default a bucket is set to NoPublicAccess. It is treated as NoPublicAccess
	// when this value is not specified. When the type is NoPublicAccess the bucket does not allow any public access.
	// When the type is ObjectRead the bucket allows public access to the GetObject, HeadObject, ListObjects.
	PublicAccessType CreateBucketDetailsPublicAccessTypeEnum `mandatory:"false" json:"publicAccessType,omitempty"`
}

func (model CreateBucketDetails) String() string {
	return common.PointerString(model)
}

type CreateBucketDetailsPublicAccessTypeEnum string

const (
	CREATE_BUCKET_DETAILS_PUBLIC_ACCESS_TYPE_NOPUBLICACCESS CreateBucketDetailsPublicAccessTypeEnum = "NoPublicAccess"
	CREATE_BUCKET_DETAILS_PUBLIC_ACCESS_TYPE_OBJECTREAD     CreateBucketDetailsPublicAccessTypeEnum = "ObjectRead"
	CREATE_BUCKET_DETAILS_PUBLIC_ACCESS_TYPE_UNKNOWN        CreateBucketDetailsPublicAccessTypeEnum = "UNKNOWN"
)

var mapping_createbucketdetails_publicAccessType = map[string]CreateBucketDetailsPublicAccessTypeEnum{
	"NoPublicAccess": CREATE_BUCKET_DETAILS_PUBLIC_ACCESS_TYPE_NOPUBLICACCESS,
	"ObjectRead":     CREATE_BUCKET_DETAILS_PUBLIC_ACCESS_TYPE_OBJECTREAD,
	"UNKNOWN":        CREATE_BUCKET_DETAILS_PUBLIC_ACCESS_TYPE_UNKNOWN,
}

func GetCreateBucketDetailsPublicAccessTypeEnumValues() []CreateBucketDetailsPublicAccessTypeEnum {
	values := make([]CreateBucketDetailsPublicAccessTypeEnum, 0)
	for _, v := range mapping_createbucketdetails_publicAccessType {
		if v != CREATE_BUCKET_DETAILS_PUBLIC_ACCESS_TYPE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
