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

type CreatePreauthenticatedRequestDetails struct {

	// user specified name for pre-authenticated request. Helpful for management purposes.
	Name *string `mandatory:"true" json:"name,omitempty"`

	// the operation that can be performed on this resource e.g PUT or GET.
	AccessType CreatePreauthenticatedRequestDetailsAccessTypeEnum `mandatory:"true" json:"accessType,omitempty"`

	// The expiration date after which the pre-authenticated request will no longer be valid per spec
	// [RFC 3339](https://tools.ietf.org/rfc/rfc3339)
	TimeExpires *common.SDKTime `mandatory:"true" json:"timeExpires,omitempty"`

	// Name of object that is being granted access to by the pre-authenticated request. This can be null and that would mean that the pre-authenticated request is granting access to the entire bucket
	ObjectName *string `mandatory:"false" json:"objectName,omitempty"`
}

func (model CreatePreauthenticatedRequestDetails) String() string {
	return common.PointerString(model)
}

type CreatePreauthenticatedRequestDetailsAccessTypeEnum string
type CreatePreauthenticatedRequestDetailsAccessType struct{}

const (
	CREATE_PREAUTHENTICATED_REQUEST_DETAILS_ACCESS_TYPE_OBJECTREAD      CreatePreauthenticatedRequestDetailsAccessTypeEnum = "ObjectRead"
	CREATE_PREAUTHENTICATED_REQUEST_DETAILS_ACCESS_TYPE_OBJECTWRITE     CreatePreauthenticatedRequestDetailsAccessTypeEnum = "ObjectWrite"
	CREATE_PREAUTHENTICATED_REQUEST_DETAILS_ACCESS_TYPE_OBJECTREADWRITE CreatePreauthenticatedRequestDetailsAccessTypeEnum = "ObjectReadWrite"
	CREATE_PREAUTHENTICATED_REQUEST_DETAILS_ACCESS_TYPE_ANYOBJECTWRITE  CreatePreauthenticatedRequestDetailsAccessTypeEnum = "AnyObjectWrite"
	CREATE_PREAUTHENTICATED_REQUEST_DETAILS_ACCESS_TYPE_UNKNOWN         CreatePreauthenticatedRequestDetailsAccessTypeEnum = "UNKNOWN"
)

var mapping_createpreauthenticatedrequestdetails_accessType = map[string]CreatePreauthenticatedRequestDetailsAccessTypeEnum{
	"ObjectRead":      CREATE_PREAUTHENTICATED_REQUEST_DETAILS_ACCESS_TYPE_OBJECTREAD,
	"ObjectWrite":     CREATE_PREAUTHENTICATED_REQUEST_DETAILS_ACCESS_TYPE_OBJECTWRITE,
	"ObjectReadWrite": CREATE_PREAUTHENTICATED_REQUEST_DETAILS_ACCESS_TYPE_OBJECTREADWRITE,
	"AnyObjectWrite":  CREATE_PREAUTHENTICATED_REQUEST_DETAILS_ACCESS_TYPE_ANYOBJECTWRITE,
	"UNKNOWN":         CREATE_PREAUTHENTICATED_REQUEST_DETAILS_ACCESS_TYPE_UNKNOWN,
}

func (receiver CreatePreauthenticatedRequestDetailsAccessType) Values() []CreatePreauthenticatedRequestDetailsAccessTypeEnum {
	values := make([]CreatePreauthenticatedRequestDetailsAccessTypeEnum, 0)
	for _, v := range mapping_createpreauthenticatedrequestdetails_accessType {
		if v != CREATE_PREAUTHENTICATED_REQUEST_DETAILS_ACCESS_TYPE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

func (receiver CreatePreauthenticatedRequestDetailsAccessType) IsValid(toBeChecked string) bool {
	for _, v := range receiver.Values() {
		if CreatePreauthenticatedRequestDetailsAccessTypeEnum(toBeChecked) == v {
			return true
		}
	}
	return false
}

func (receiver CreatePreauthenticatedRequestDetailsAccessType) From(toBeConverted string) CreatePreauthenticatedRequestDetailsAccessTypeEnum {
	if val, ok := mapping_createpreauthenticatedrequestdetails_accessType[toBeConverted]; ok {
		return val
	}
	return CREATE_PREAUTHENTICATED_REQUEST_DETAILS_ACCESS_TYPE_UNKNOWN
}
