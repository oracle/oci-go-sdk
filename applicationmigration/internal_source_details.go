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

// InternalSourceDetails Specifies configuration specific to the source environment.
type InternalSourceDetails struct {

	// The tradition cloud account name
	AccountName *string `mandatory:"true" json:"accountName"`
}

func (m InternalSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InternalSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInternalSourceDetails InternalSourceDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeInternalSourceDetails
	}{
		"INTERNAL_COMPUTE",
		(MarshalTypeInternalSourceDetails)(m),
	}

	return json.Marshal(&s)
}
