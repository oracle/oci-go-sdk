// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BackupDestinationProperties The properties of the backup destination associated with the Autonomous Container Database.
type BackupDestinationProperties struct {

	// The timestamps at which this backup destination is used as the preferred destination to host the Autonomous Container Database backups.
	BackupDestinationAttachHistory []common.SDKTime `mandatory:"false" json:"backupDestinationAttachHistory"`

	// The total space utilized (in GBs) by this Autonomous Container Database on this backup destination, rounded to the nearest integer.
	SpaceUtilizedInGBs *int `mandatory:"false" json:"spaceUtilizedInGBs"`

	// The latest timestamp when the backup destination details, such as 'spaceUtilized,' are updated.
	TimeAtWhichStorageDetailsAreUpdated *common.SDKTime `mandatory:"false" json:"timeAtWhichStorageDetailsAreUpdated"`
}

func (m BackupDestinationProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackupDestinationProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
