// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// InternalAuthorizationDetails Specifies the credentials to access the source Oracle Cloud Infrastructure - Classic environment.
type InternalAuthorizationDetails struct {

	// User with Compute Operations role in Oracle Cloud Infrastructure - Classic.
	Username *string `mandatory:"true" json:"username"`

	// Password for this user.
	Password *string `mandatory:"true" json:"password"`
}

func (m InternalAuthorizationDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InternalAuthorizationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInternalAuthorizationDetails InternalAuthorizationDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeInternalAuthorizationDetails
	}{
		"INTERNAL_COMPUTE",
		(MarshalTypeInternalAuthorizationDetails)(m),
	}

	return json.Marshal(&s)
}
