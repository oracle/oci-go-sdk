// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// FailoverDataGuardAssociationDetails. The Data Guard association failover parameters.
type FailoverDataGuardAssociationDetails struct {

	// The DB System administrator password.
	DatabaseAdminPassword *string `mandatory:"true" json:"databaseAdminPassword,omitempty"`
}

func (model FailoverDataGuardAssociationDetails) String() string {
	return common.PointerString(model)
}
