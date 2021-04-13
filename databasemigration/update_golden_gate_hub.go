// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// REST API for Zero Downtime Migration (Oracle Database Migration Service --ODMS-- as customer-facing service name)
//
// Provides users the ability to perform Zero Downtime migration operations
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v39/common"
)

// UpdateGoldenGateHub Details about Oracle GoldenGate Microservices.
type UpdateGoldenGateHub struct {
	RestAdminCredentials *UpdateAdminCredentials `mandatory:"false" json:"restAdminCredentials"`

	SourceDbAdminCredentials *UpdateAdminCredentials `mandatory:"false" json:"sourceDbAdminCredentials"`

	SourceContainerDbAdminCredentials *UpdateAdminCredentials `mandatory:"false" json:"sourceContainerDbAdminCredentials"`

	TargetDbAdminCredentials *UpdateAdminCredentials `mandatory:"false" json:"targetDbAdminCredentials"`

	// Oracle GoldenGate hub's REST endpoint.
	// Refer to https://docs.oracle.com/en/middleware/goldengate/core/19.1/securing/network.html#GUID-A709DA55-111D-455E-8942-C9BDD1E38CAA
	Url *string `mandatory:"false" json:"url"`

	// Name of Microservices deployment to operate on source DB
	SourceMicroservicesDeploymentName *string `mandatory:"false" json:"sourceMicroservicesDeploymentName"`

	// Name of Microservices deployment to operate on target DB
	TargetMicroservicesDeploymentName *string `mandatory:"false" json:"targetMicroservicesDeploymentName"`

	// OCID of Golden Gate compute instance. An empty value will remove the stored computeId.
	ComputeId *string `mandatory:"false" json:"computeId"`
}

func (m UpdateGoldenGateHub) String() string {
	return common.PointerString(m)
}
