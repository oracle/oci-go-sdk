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
	"github.com/oracle/oci-go-sdk/v43/common"
)

// OicDiscoveryDetails Credentials to access the Oracle Integration application in the source environment. Application Migration connects to the
// application in the source environment with the supplied credentials.
type OicDiscoveryDetails struct {

	// Application administrator username to access the Oracle Integration Classic instance in the source environment.
	ServiceInstanceUser *string `mandatory:"true" json:"serviceInstanceUser"`

	// Password for this user.
	ServiceInstancePassword *string `mandatory:"true" json:"serviceInstancePassword"`
}

func (m OicDiscoveryDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OicDiscoveryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOicDiscoveryDetails OicDiscoveryDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeOicDiscoveryDetails
	}{
		"OIC",
		(MarshalTypeOicDiscoveryDetails)(m),
	}

	return json.Marshal(&s)
}
