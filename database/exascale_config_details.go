// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExascaleConfigDetails The exascale config response details for the Exadata Cloud@Customer infrastructure or cloud Exadata infrastructure . Applies to both Exadata Cloud@Customer instances and Exadata Cloud Service instances.
type ExascaleConfigDetails struct {

	// Storage size needed for Exascale in GBs.
	TotalStorageInGBs *int `mandatory:"true" json:"totalStorageInGBs"`

	// Available storage size for Exascale in GBs.
	AvailableStorageInGBs *int `mandatory:"false" json:"availableStorageInGBs"`
}

func (m ExascaleConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExascaleConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
