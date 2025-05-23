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

// AddDistributedAutonomousDatabaseGdsControlNodeDetails Details required to create a new Global database services control(GDS CTL) compute node.
type AddDistributedAutonomousDatabaseGdsControlNodeDetails struct {

	// The public sshKey for Global database services control(GDS CTL) node.
	PublicSshKey *string `mandatory:"true" json:"publicSshKey"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet on which the Global database services control(GDS CTL) node should be created.
	SubnetId *string `mandatory:"false" json:"subnetId"`
}

func (m AddDistributedAutonomousDatabaseGdsControlNodeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddDistributedAutonomousDatabaseGdsControlNodeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
