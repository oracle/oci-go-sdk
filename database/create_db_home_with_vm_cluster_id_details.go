// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateDbHomeWithVmClusterIdDetails Note that a valid `vmClusterId` value must be supplied for the `CreateDbHomeWithVmClusterId` API operation to successfully complete.
type CreateDbHomeWithVmClusterIdDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VM cluster.
	VmClusterId *string `mandatory:"true" json:"vmClusterId"`

	// A valid Oracle Database version. To get a list of supported versions, use the ListDbVersions operation.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// The user-provided name of the Database Home.
	DisplayName *string `mandatory:"false" json:"displayName"`

	Database *CreateDatabaseDetails `mandatory:"false" json:"database"`
}

//GetDisplayName returns DisplayName
func (m CreateDbHomeWithVmClusterIdDetails) GetDisplayName() *string {
	return m.DisplayName
}

func (m CreateDbHomeWithVmClusterIdDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDbHomeWithVmClusterIdDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDbHomeWithVmClusterIdDetails CreateDbHomeWithVmClusterIdDetails
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeCreateDbHomeWithVmClusterIdDetails
	}{
		"VM_CLUSTER_NEW",
		(MarshalTypeCreateDbHomeWithVmClusterIdDetails)(m),
	}

	return json.Marshal(&s)
}
