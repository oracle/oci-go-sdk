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
	"github.com/oracle/oci-go-sdk/v52/common"
)

// OacDiscoveryDetails Details about the Oracle Analytics Cloud - Classic application in the source environment.
type OacDiscoveryDetails struct {

	// This field is currently not supported. You must enter a value, such as <code>unused</code>. However, the value that you enter is ignored.
	ServiceInstanceUser *string `mandatory:"true" json:"serviceInstanceUser"`

	// This field is currently not supported. You must enter a value, such as <code>unused</code>. However, the value that you enter is ignored.
	ServiceInstancePassword *string `mandatory:"true" json:"serviceInstancePassword"`
}

func (m OacDiscoveryDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OacDiscoveryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOacDiscoveryDetails OacDiscoveryDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeOacDiscoveryDetails
	}{
		"OAC",
		(MarshalTypeOacDiscoveryDetails)(m),
	}

	return json.Marshal(&s)
}
