// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListObjectsRequest wrapper for the ListObjects operation
type ListObjectsRequest struct {

	// The top-level namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The name of the bucket.
	// Example: `my-new-bucket1`
	BucketName *string `mandatory:"true" contributesTo:"path" name:"bucketName"`

	// The string to use for matching against the start of object names in a list query.
	Prefix *string `mandatory:"false" contributesTo:"query" name:"prefix"`

	// Object names returned by a list query must be greater or equal to this parameter.
	Start *string `mandatory:"false" contributesTo:"query" name:"start"`

	// Object names returned by a list query must be strictly less than this parameter.
	End *string `mandatory:"false" contributesTo:"query" name:"end"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// When this parameter is set, only objects whose names do not contain the delimiter character
	// (after an optionally specified prefix) are returned. Scanned objects whose names contain the
	// delimiter have part of their name up to the last occurrence of the delimiter (after the optional
	// prefix) returned as a set of prefixes. Note that only '/' is a supported delimiter character at
	// this time.
	Delimiter *string `mandatory:"false" contributesTo:"query" name:"delimiter"`

	// Object summary in list of objects includes the 'name' field. This parameter can also include 'size'
	// (object size in bytes), 'md5', and 'timeCreated' (object creation date and time) fields.
	// Value of this parameter should be a comma-separated, case-insensitive list of those field names.
	// For example 'name,timeCreated,md5'.
	Fields ListObjectsFieldsEnum `mandatory:"false" contributesTo:"query" name:"fields"`

	// The client request ID for tracing.
	OpcClientRequestID *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`
}

func (request ListObjectsRequest) String() string {
	return common.PointerString(request)
}

// ListObjectsResponse wrapper for the ListObjects operation
type ListObjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ListObjects instance
	ListObjects `presentIn:"body"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestID *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, please provide this request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListObjectsResponse) String() string {
	return common.PointerString(response)
}

type ListObjectsFieldsEnum string

const (
	LIST_OBJECTS_FIELDS_NAME        ListObjectsFieldsEnum = "name"
	LIST_OBJECTS_FIELDS_SIZE        ListObjectsFieldsEnum = "size"
	LIST_OBJECTS_FIELDS_TIMECREATED ListObjectsFieldsEnum = "timeCreated"
	LIST_OBJECTS_FIELDS_MD5         ListObjectsFieldsEnum = "md5"
	LIST_OBJECTS_FIELDS_UNKNOWN     ListObjectsFieldsEnum = "UNKNOWN"
)

var mapping_listobjectsfields = map[string]ListObjectsFieldsEnum{
	"name":        LIST_OBJECTS_FIELDS_NAME,
	"size":        LIST_OBJECTS_FIELDS_SIZE,
	"timeCreated": LIST_OBJECTS_FIELDS_TIMECREATED,
	"md5":         LIST_OBJECTS_FIELDS_MD5,
	"UNKNOWN":     LIST_OBJECTS_FIELDS_UNKNOWN,
}

func GetListObjectsFieldsEnumValues() []ListObjectsFieldsEnum {
	values := make([]ListObjectsFieldsEnum, 0)
	for _, v := range mapping_listobjectsfields {
		if v != LIST_OBJECTS_FIELDS_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
