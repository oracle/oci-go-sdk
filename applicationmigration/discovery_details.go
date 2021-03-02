// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v36/common"
)

// DiscoveryDetails Base model for different application discovery requirements.
type DiscoveryDetails interface {
}

type discoverydetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *discoverydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdiscoverydetails discoverydetails
	s := struct {
		Model Unmarshalerdiscoverydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *discoverydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OIC":
		mm := OicDiscoveryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PCS":
		mm := PcsDiscoveryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ICS":
		mm := IcsDiscoveryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OAC":
		mm := OacDiscoveryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JCS":
		mm := JcsDiscoveryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SOACS":
		mm := SoacsDiscoveryDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m discoverydetails) String() string {
	return common.PointerString(m)
}
