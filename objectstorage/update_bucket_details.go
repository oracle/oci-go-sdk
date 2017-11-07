// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// APIs for managing buckets and objects.
//

package objectstorage

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// UpdateBucketDetails. To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type UpdateBucketDetails struct {

	// The namespace in which the bucket lives.
	Namespace *string `mandatory:"false" json:"namespace,omitempty"`

	// The name of the bucket.
	Name *string `mandatory:"false" json:"name,omitempty"`

	// Arbitrary string, up to 4KB, of keys and values for user-defined metadata.
	Metadata *map[string]string `mandatory:"false" json:"metadata,omitempty"`

	// The type of public access available on this bucket. Allows authenticated caller to access the bucket or
	// contents of this bucket. By default a bucket is set to NoPublicAccess. It is treated as NoPublicAccess
	// when this value is not specified. When the type is NoPublicAccess the bucket does not allow any public access.
	// When the type is ObjectRead the bucket allows public access to the GetObject, HeadObject, ListObjects.
	PublicAccessType UpdateBucketDetailsPublicAccessTypeEnum `mandatory:"false" json:"publicAccessType,omitempty"`
}

func (model UpdateBucketDetails) String() string {
	return common.PointerString(model)
}

type UpdateBucketDetailsPublicAccessTypeEnum string

const (
	UPDATE_BUCKET_DETAILS_PUBLIC_ACCESS_TYPE_NOPUBLICACCESS UpdateBucketDetailsPublicAccessTypeEnum = "NoPublicAccess"
	UPDATE_BUCKET_DETAILS_PUBLIC_ACCESS_TYPE_OBJECTREAD     UpdateBucketDetailsPublicAccessTypeEnum = "ObjectRead"
	UPDATE_BUCKET_DETAILS_PUBLIC_ACCESS_TYPE_UNKNOWN        UpdateBucketDetailsPublicAccessTypeEnum = "UNKNOWN"
)

var mapping_updatebucketdetails_publicAccessType = map[string]UpdateBucketDetailsPublicAccessTypeEnum{
	"NoPublicAccess": UPDATE_BUCKET_DETAILS_PUBLIC_ACCESS_TYPE_NOPUBLICACCESS,
	"ObjectRead":     UPDATE_BUCKET_DETAILS_PUBLIC_ACCESS_TYPE_OBJECTREAD,
	"UNKNOWN":        UPDATE_BUCKET_DETAILS_PUBLIC_ACCESS_TYPE_UNKNOWN,
}

func GetUpdateBucketDetailsPublicAccessTypeEnumValues() []UpdateBucketDetailsPublicAccessTypeEnum {
	values := make([]UpdateBucketDetailsPublicAccessTypeEnum, 0)
	for _, v := range mapping_updatebucketdetails_publicAccessType {
		if v != UPDATE_BUCKET_DETAILS_PUBLIC_ACCESS_TYPE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
