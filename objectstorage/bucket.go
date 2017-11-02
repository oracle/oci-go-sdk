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

// Bucket. To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type Bucket struct {

	// The namespace in which the bucket lives.
	Namespace *string `mandatory:"true" json:"namespace,omitempty"`

	// The name of the bucket.
	Name *string `mandatory:"true" json:"name,omitempty"`

	// The compartment ID in which the bucket is authorized.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// Arbitrary string keys and values for user-defined metadata.
	Metadata *map[string]string `mandatory:"true" json:"metadata,omitempty"`

	// The OCID of the user who created the bucket.
	CreatedBy *string `mandatory:"true" json:"createdBy,omitempty"`

	// The date and time at which the bucket was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The entity tag for the bucket.
	Etag *string `mandatory:"true" json:"etag,omitempty"`

	// The type of public access available on this bucket. Allows authenticated caller to access the bucket or
	// contents of this bucket. By default a bucket is set to NoPublicAccess. It is treated as NoPublicAccess
	// when this value is not specified. When the type is NoPublicAccess the bucket does not allow any public access.
	// When the type is ObjectRead the bucket allows public access to the GetObject, HeadObject, ListObjects.
	PublicAccessType BucketPublicAccessTypeEnum `mandatory:"false" json:"publicAccessType,omitempty"`
}

func (model Bucket) String() string {
	return common.PointerString(model)
}

type BucketPublicAccessTypeEnum string
type BucketPublicAccessType struct{}

const (
	BUCKET_PUBLIC_ACCESS_TYPE_NOPUBLICACCESS BucketPublicAccessTypeEnum = "NoPublicAccess"
	BUCKET_PUBLIC_ACCESS_TYPE_OBJECTREAD     BucketPublicAccessTypeEnum = "ObjectRead"
	BUCKET_PUBLIC_ACCESS_TYPE_UNKNOWN        BucketPublicAccessTypeEnum = "UNKNOWN"
)

var mapping_bucket_publicAccessType = map[string]BucketPublicAccessTypeEnum{
	"NoPublicAccess": BUCKET_PUBLIC_ACCESS_TYPE_NOPUBLICACCESS,
	"ObjectRead":     BUCKET_PUBLIC_ACCESS_TYPE_OBJECTREAD,
	"UNKNOWN":        BUCKET_PUBLIC_ACCESS_TYPE_UNKNOWN,
}

func (receiver BucketPublicAccessType) Values() []BucketPublicAccessTypeEnum {
	values := make([]BucketPublicAccessTypeEnum, 0)
	for _, v := range mapping_bucket_publicAccessType {
		if v != BUCKET_PUBLIC_ACCESS_TYPE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

func (receiver BucketPublicAccessType) IsValid(toBeChecked string) bool {
	for _, v := range receiver.Values() {
		if BucketPublicAccessTypeEnum(toBeChecked) == v {
			return true
		}
	}
	return false
}

func (receiver BucketPublicAccessType) From(toBeConverted string) BucketPublicAccessTypeEnum {
	if val, ok := mapping_bucket_publicAccessType[toBeConverted]; ok {
		return val
	}
	return BUCKET_PUBLIC_ACCESS_TYPE_UNKNOWN
}
