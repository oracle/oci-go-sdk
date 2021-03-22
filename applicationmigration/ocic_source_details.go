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
	"github.com/oracle/oci-go-sdk/v37/common"
)

// OcicSourceDetails Details about the Oracle Cloud Infrastructure Classic account, the source environment from which you want to migrate the application.
type OcicSourceDetails struct {

	// The Oracle Cloud Infrastructure - Classic region from which you want to migrate your applications. For example, uscom-east-1 or uscom-central-1.
	Region *string `mandatory:"true" json:"region"`

	// If you are using a Oracle Cloud Infrastructure - Classic account with Identity Cloud Service (IDCS), enter the service instance ID.
	// For example, if Compute-567890123 is the account name of your Oracle Cloud Infrastructure Classic Compute service entitlement,
	// then enter 567890123.
	// If you are using a traditional Oracle Cloud Infrastructure - Classic account, enter your identity domain ID.
	ComputeAccount *string `mandatory:"true" json:"computeAccount"`
}

func (m OcicSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OcicSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOcicSourceDetails OcicSourceDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeOcicSourceDetails
	}{
		"OCIC",
		(MarshalTypeOcicSourceDetails)(m),
	}

	return json.Marshal(&s)
}
