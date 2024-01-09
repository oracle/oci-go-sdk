// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration API
//
// Application Migration simplifies the migration of applications from Oracle Cloud Infrastructure Classic to Oracle Cloud Infrastructure.
// You can use Application Migration API to migrate applications, such as Oracle Java Cloud Service, SOA Cloud Service, and Integration Classic
// instances, to Oracle Cloud Infrastructure. For more information, see
// Overview of Application Migration (https://docs.cloud.oracle.com/iaas/application-migration/appmigrationoverview.htm).
//

package applicationmigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AuthorizationDetails Details of the source environment from which you want to migrate applications to Oracle Cloud Infrastructure. It also contains access
// credentials.
type AuthorizationDetails interface {
}

type authorizationdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *authorizationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerauthorizationdetails authorizationdetails
	s := struct {
		Model Unmarshalerauthorizationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *authorizationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OCC":
		mm := OccAuthorizationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INTERNAL_COMPUTE":
		mm := InternalAuthorizationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCIC_IDCS":
		mm := OcicAuthorizationTokenDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCIC":
		mm := OcicAuthorizationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for AuthorizationDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m authorizationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m authorizationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
