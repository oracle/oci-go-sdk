// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateBackupDestinationDetails For a RECOVERY_APPLIANCE backup destination, used to update the connection string and/or the list of VPC users.
// For an NFS backup destination, used to update the NFS location.
type UpdateBackupDestinationDetails struct {

	// For a RECOVERY_APPLIANCE backup destination, the Virtual Private Catalog (VPC) users that are used to access the Recovery Appliance.
	VpcUsers []string `mandatory:"false" json:"vpcUsers"`

	// For a RECOVERY_APPLIANCE backup destination, the connection string for connecting to the Recovery Appliance.
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// The local directory path on each VM cluster node where the NFS server location is mounted. The local directory path and the NFS server location must each be the same across all of the VM cluster nodes. Ensure that the NFS mount is maintained continuously on all of the VM cluster nodes.
	LocalMountPointPath *string `mandatory:"false" json:"localMountPointPath"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateBackupDestinationDetails) String() string {
	return common.PointerString(m)
}
