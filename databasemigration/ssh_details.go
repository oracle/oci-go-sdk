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

// SshDetails Details of the ssh key that will be used.
type SshDetails struct {

	// Name of the host the sshkey is valid for.
	Host *string `mandatory:"true" json:"host"`

	// SSH user
	User *string `mandatory:"true" json:"user"`

	// Sudo location
	SudoLocation *string `mandatory:"true" json:"sudoLocation"`
}

func (m SshDetails) String() string {
	return common.PointerString(m)
}
