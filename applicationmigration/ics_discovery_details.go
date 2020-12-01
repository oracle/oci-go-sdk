// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v30/common"
)

// IcsDiscoveryDetails Credentials to access the Oracle Integration Cloud Service application in the source environment. Application Migration connects
// to the application in the source environment with the supplied credentials.
type IcsDiscoveryDetails struct {

	// Application administrator username to access the Oracle Integration Cloud Service application in the source environment.
	ServiceInstanceUser *string `mandatory:"true" json:"serviceInstanceUser"`

	// Password for this user.
	ServiceInstancePassword *string `mandatory:"true" json:"serviceInstancePassword"`
}

func (m IcsDiscoveryDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m IcsDiscoveryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeIcsDiscoveryDetails IcsDiscoveryDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeIcsDiscoveryDetails
	}{
		"ICS",
		(MarshalTypeIcsDiscoveryDetails)(m),
	}

	return json.Marshal(&s)
}
