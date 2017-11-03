// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"net/http"
)

// InstanceActionRequest wrapper for the InstanceAction operation
type InstanceActionRequest struct {

	// The OCID of the instance.
	InstanceID *string `mandatory:"true" contributesTo:"path" name:"instanceId"`

	// The action to perform on the instance.
	Action InstanceActionActionEnum `mandatory:"true" contributesTo:"query" name:"action"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations (for example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

func (request InstanceActionRequest) String() string {
	return common.PointerString(request)
}

// InstanceActionResponse wrapper for the InstanceAction operation
type InstanceActionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Instance instance
	Instance `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response InstanceActionResponse) String() string {
	return common.PointerString(response)
}

type InstanceActionActionEnum string
type InstanceActionAction struct{}

const (
	INSTANCE_ACTION_ACTION_STOP      InstanceActionActionEnum = "STOP"
	INSTANCE_ACTION_ACTION_START     InstanceActionActionEnum = "START"
	INSTANCE_ACTION_ACTION_SOFTRESET InstanceActionActionEnum = "SOFTRESET"
	INSTANCE_ACTION_ACTION_RESET     InstanceActionActionEnum = "RESET"
	INSTANCE_ACTION_ACTION_UNKNOWN   InstanceActionActionEnum = "UNKNOWN"
)

var mapping_Compute_action = map[string]InstanceActionActionEnum{
	"STOP":      INSTANCE_ACTION_ACTION_STOP,
	"START":     INSTANCE_ACTION_ACTION_START,
	"SOFTRESET": INSTANCE_ACTION_ACTION_SOFTRESET,
	"RESET":     INSTANCE_ACTION_ACTION_RESET,
	"UNKNOWN":   INSTANCE_ACTION_ACTION_UNKNOWN,
}

func (receiver InstanceActionAction) Values() []InstanceActionActionEnum {
	values := make([]InstanceActionActionEnum, 0)
	for _, v := range mapping_Compute_action {
		if v != INSTANCE_ACTION_ACTION_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

func (receiver InstanceActionAction) IsValid(toBeChecked string) bool {
	for _, v := range receiver.Values() {
		if InstanceActionActionEnum(toBeChecked) == v {
			return true
		}
	}
	return false
}

func (receiver InstanceActionAction) From(toBeConverted string) InstanceActionActionEnum {
	if val, ok := mapping_Compute_action[toBeConverted]; ok {
		return val
	}
	return INSTANCE_ACTION_ACTION_UNKNOWN
}
