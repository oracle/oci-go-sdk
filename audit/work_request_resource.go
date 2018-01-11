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

// WorkRequestResource The resource entity.
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

func (m WorkRequestResource) String() string {
	return common.PointerString(m)
}

// WorkRequestResourceActionTypeEnum Enum with underlying type: string
type WorkRequestResourceActionTypeEnum string

// Set of constants representing the allowable values for WorkRequestResourceActionType
const (
	WorkRequestResourceActionTypeCreated    WorkRequestResourceActionTypeEnum = "CREATED"
	WorkRequestResourceActionTypeUpdated    WorkRequestResourceActionTypeEnum = "UPDATED"
	WorkRequestResourceActionTypeDeleted    WorkRequestResourceActionTypeEnum = "DELETED"
	WorkRequestResourceActionTypeRelated    WorkRequestResourceActionTypeEnum = "RELATED"
	WorkRequestResourceActionTypeInProgress WorkRequestResourceActionTypeEnum = "IN_PROGRESS"
	WorkRequestResourceActionTypeUnknown    WorkRequestResourceActionTypeEnum = "UNKNOWN"
)

var mappingWorkRequestResourceActionType = map[string]WorkRequestResourceActionTypeEnum{
	"CREATED":     WorkRequestResourceActionTypeCreated,
	"UPDATED":     WorkRequestResourceActionTypeUpdated,
	"DELETED":     WorkRequestResourceActionTypeDeleted,
	"RELATED":     WorkRequestResourceActionTypeRelated,
	"IN_PROGRESS": WorkRequestResourceActionTypeInProgress,
	"UNKNOWN":     WorkRequestResourceActionTypeUnknown,
}

// GetWorkRequestResourceActionTypeEnumValues Enumerates the set of values for WorkRequestResourceActionType
func GetWorkRequestResourceActionTypeEnumValues() []WorkRequestResourceActionTypeEnum {
	values := make([]WorkRequestResourceActionTypeEnum, 0)
	for _, v := range mappingWorkRequestResourceActionType {
		if v != WorkRequestResourceActionTypeUnknown {
			values = append(values, v)
		}
	}
	return values
}
