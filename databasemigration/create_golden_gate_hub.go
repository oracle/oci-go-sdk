// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// REST API for Zero Downtime Migration (Oracle Database Migration Service --ODMS-- as customer-facing service name)
//
// Provides users the ability to perform Zero Downtime migration operations
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v40/common"
)

// CreateGoldenGateHub Details about Oracle GoldenGate Microservices. Required for online logical migration.
type CreateGoldenGateHub struct {
	RestAdminCredentials *CreateAdminCredentials `mandatory:"true" json:"restAdminCredentials"`

	SourceDbAdminCredentials *CreateAdminCredentials `mandatory:"true" json:"sourceDbAdminCredentials"`

	TargetDbAdminCredentials *CreateAdminCredentials `mandatory:"true" json:"targetDbAdminCredentials"`

	// Oracle GoldenGate hub's REST endpoint.
	// Refer to https://docs.oracle.com/en/middleware/goldengate/core/19.1/securing/network.html#GUID-A709DA55-111D-455E-8942-C9BDD1E38CAA
	Url *string `mandatory:"true" json:"url"`

	// Name of Microservices deployment to operate on source DB
	SourceMicroservicesDeploymentName *string `mandatory:"true" json:"sourceMicroservicesDeploymentName"`

	// Name of Microservices deployment to operate on target DB
	TargetMicroservicesDeploymentName *string `mandatory:"true" json:"targetMicroservicesDeploymentName"`

	SourceContainerDbAdminCredentials *CreateAdminCredentials `mandatory:"false" json:"sourceContainerDbAdminCredentials"`

	// OCID of Golden Gate compute instance.
	ComputeId *string `mandatory:"false" json:"computeId"`
}

func (m CreateGoldenGateHub) String() string {
	return common.PointerString(m)
}
