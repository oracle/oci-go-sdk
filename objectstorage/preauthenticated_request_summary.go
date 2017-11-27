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

type PreauthenticatedRequestSummary struct {

	// the unique identifier to use when directly addressing the pre-authenticated request
	ID *string `mandatory:"true" json:"id,omitempty"`

	// the user supplied name of the pre-authenticated request
	Name *string `mandatory:"true" json:"name,omitempty"`

	// the operation that can be performed on this resource e.g PUT or GET.
	AccessType PreauthenticatedRequestSummaryAccessTypeEnum `mandatory:"true" json:"accessType,omitempty"`

	// the expiration date after which the pre authenticated request will no longer be valid as per spec
	// [RFC 3339](https://tools.ietf.org/rfc/rfc3339)
	TimeExpires *common.SDKTime `mandatory:"true" json:"timeExpires,omitempty"`

	// the date when the pre-authenticated request was created as per spec
	// [RFC 3339](https://tools.ietf.org/rfc/rfc3339)
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// Name of object that is being granted access to by the pre-authenticated request. This can be null and that would mean that the pre-authenticated request is granting access to the entire bucket
	ObjectName *string `mandatory:"false" json:"objectName,omitempty"`
}

func (model PreauthenticatedRequestSummary) String() string {
	return common.PointerString(model)
}

type PreauthenticatedRequestSummaryAccessTypeEnum string

const (
	PREAUTHENTICATED_REQUEST_SUMMARY_ACCESS_TYPE_OBJECTREAD      PreauthenticatedRequestSummaryAccessTypeEnum = "ObjectRead"
	PREAUTHENTICATED_REQUEST_SUMMARY_ACCESS_TYPE_OBJECTWRITE     PreauthenticatedRequestSummaryAccessTypeEnum = "ObjectWrite"
	PREAUTHENTICATED_REQUEST_SUMMARY_ACCESS_TYPE_OBJECTREADWRITE PreauthenticatedRequestSummaryAccessTypeEnum = "ObjectReadWrite"
	PREAUTHENTICATED_REQUEST_SUMMARY_ACCESS_TYPE_ANYOBJECTWRITE  PreauthenticatedRequestSummaryAccessTypeEnum = "AnyObjectWrite"
	PREAUTHENTICATED_REQUEST_SUMMARY_ACCESS_TYPE_UNKNOWN         PreauthenticatedRequestSummaryAccessTypeEnum = "UNKNOWN"
)

var mapping_preauthenticatedrequestsummary_accessType = map[string]PreauthenticatedRequestSummaryAccessTypeEnum{
	"ObjectRead":      PREAUTHENTICATED_REQUEST_SUMMARY_ACCESS_TYPE_OBJECTREAD,
	"ObjectWrite":     PREAUTHENTICATED_REQUEST_SUMMARY_ACCESS_TYPE_OBJECTWRITE,
	"ObjectReadWrite": PREAUTHENTICATED_REQUEST_SUMMARY_ACCESS_TYPE_OBJECTREADWRITE,
	"AnyObjectWrite":  PREAUTHENTICATED_REQUEST_SUMMARY_ACCESS_TYPE_ANYOBJECTWRITE,
	"UNKNOWN":         PREAUTHENTICATED_REQUEST_SUMMARY_ACCESS_TYPE_UNKNOWN,
}

func GetPreauthenticatedRequestSummaryAccessTypeEnumValues() []PreauthenticatedRequestSummaryAccessTypeEnum {
	values := make([]PreauthenticatedRequestSummaryAccessTypeEnum, 0)
	for _, v := range mapping_preauthenticatedrequestsummary_accessType {
		if v != PREAUTHENTICATED_REQUEST_SUMMARY_ACCESS_TYPE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
