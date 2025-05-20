// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DistributedAutonomousDatabaseAssociatedWithPrivateEndpoint The details of distributed autonomous database associated with the private endpoint.
type DistributedAutonomousDatabaseAssociatedWithPrivateEndpoint struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the distributed database.
	Id *string `mandatory:"true" json:"id"`

	// The dbDeploymentType associated with the distributed autonomous database.
	DbDeploymentType DistributedAutonomousDatabaseDbDeploymentTypeEnum `mandatory:"true" json:"dbDeploymentType"`
}

func (m DistributedAutonomousDatabaseAssociatedWithPrivateEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedAutonomousDatabaseAssociatedWithPrivateEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedAutonomousDatabaseDbDeploymentTypeEnum(string(m.DbDeploymentType)); !ok && m.DbDeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbDeploymentType: %s. Supported values are: %s.", m.DbDeploymentType, strings.Join(GetDistributedAutonomousDatabaseDbDeploymentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
