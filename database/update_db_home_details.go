// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateDbHomeDetails Describes the modification parameters for the DB Home.
type UpdateDbHomeDetails struct {
	DbVersion *PatchDetails `mandatory:"false" json:"dbVersion,omitempty"`
}

func (model UpdateDbHomeDetails) String() string {
	return common.PointerString(model)
}
