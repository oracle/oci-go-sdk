// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Audit API
//
// API for the Audit Service. You can use this API for queries, but not bulk-export operations.
//

package audit

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestResource. The resource entity.
type WorkRequestResource struct {

	// The identifier of the resource the work request affects.
	Identifier *string `mandatory:"true" json:"identifier,omitempty"`

	// The resource type the work request is affects.
	EntityType *string `mandatory:"true" json:"entityType,omitempty"`

	// The way in which this resource was affected by the work tracked by the work request.
	ActionType WorkRequestResourceActionTypeEnum `mandatory:"true" json:"actionType,omitempty"`

	// The URI path that the user can do a GET on to access the resource metadata.
	EntityUri *string `mandatory:"false" json:"entityUri,omitempty"`
}

func (model WorkRequestResource) String() string {
	return common.PointerString(model)
}

type WorkRequestResourceActionTypeEnum string

const (
	WORK_REQUEST_RESOURCE_ACTION_TYPE_CREATED     WorkRequestResourceActionTypeEnum = "CREATED"
	WORK_REQUEST_RESOURCE_ACTION_TYPE_UPDATED     WorkRequestResourceActionTypeEnum = "UPDATED"
	WORK_REQUEST_RESOURCE_ACTION_TYPE_DELETED     WorkRequestResourceActionTypeEnum = "DELETED"
	WORK_REQUEST_RESOURCE_ACTION_TYPE_RELATED     WorkRequestResourceActionTypeEnum = "RELATED"
	WORK_REQUEST_RESOURCE_ACTION_TYPE_IN_PROGRESS WorkRequestResourceActionTypeEnum = "IN_PROGRESS"
	WORK_REQUEST_RESOURCE_ACTION_TYPE_UNKNOWN     WorkRequestResourceActionTypeEnum = "UNKNOWN"
)

var mapping_workrequestresource_actionType = map[string]WorkRequestResourceActionTypeEnum{
	"CREATED":     WORK_REQUEST_RESOURCE_ACTION_TYPE_CREATED,
	"UPDATED":     WORK_REQUEST_RESOURCE_ACTION_TYPE_UPDATED,
	"DELETED":     WORK_REQUEST_RESOURCE_ACTION_TYPE_DELETED,
	"RELATED":     WORK_REQUEST_RESOURCE_ACTION_TYPE_RELATED,
	"IN_PROGRESS": WORK_REQUEST_RESOURCE_ACTION_TYPE_IN_PROGRESS,
	"UNKNOWN":     WORK_REQUEST_RESOURCE_ACTION_TYPE_UNKNOWN,
}

func GetWorkRequestResourceActionTypeEnumValues() []WorkRequestResourceActionTypeEnum {
	values := make([]WorkRequestResourceActionTypeEnum, 0)
	for _, v := range mapping_workrequestresource_actionType {
		if v != WORK_REQUEST_RESOURCE_ACTION_TYPE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
