// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming Service API
//
// The API for the Streaming Service.
//

package streaming

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateGroupDetails Request body for operationally managing a group.
type UpdateGroupDetails struct {

	// The type of the cursor.
	Type UpdateGroupDetailsTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The time to consume from if type is AT_TIME.
	Time *common.SDKTime `mandatory:"false" json:"time"`
}

func (m UpdateGroupDetails) String() string {
	return common.PointerString(m)
}

// UpdateGroupDetailsTypeEnum Enum with underlying type: string
type UpdateGroupDetailsTypeEnum string

// Set of constants representing the allowable values for UpdateGroupDetailsTypeEnum
const (
	UpdateGroupDetailsTypeAtTime      UpdateGroupDetailsTypeEnum = "AT_TIME"
	UpdateGroupDetailsTypeLatest      UpdateGroupDetailsTypeEnum = "LATEST"
	UpdateGroupDetailsTypeTrimHorizon UpdateGroupDetailsTypeEnum = "TRIM_HORIZON"
)

var mappingUpdateGroupDetailsType = map[string]UpdateGroupDetailsTypeEnum{
	"AT_TIME":      UpdateGroupDetailsTypeAtTime,
	"LATEST":       UpdateGroupDetailsTypeLatest,
	"TRIM_HORIZON": UpdateGroupDetailsTypeTrimHorizon,
}

// GetUpdateGroupDetailsTypeEnumValues Enumerates the set of values for UpdateGroupDetailsTypeEnum
func GetUpdateGroupDetailsTypeEnumValues() []UpdateGroupDetailsTypeEnum {
	values := make([]UpdateGroupDetailsTypeEnum, 0)
	for _, v := range mappingUpdateGroupDetailsType {
		values = append(values, v)
	}
	return values
}
