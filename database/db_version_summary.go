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

// DbVersionSummary The Oracle Database software version.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type DbVersionSummary struct {

	// A valid Oracle Database version.
	Version *string `mandatory:"true" json:"version"`

	// True if this version of the Oracle Database software is the latest version for a release.
	IsLatestForMajorVersion *bool `mandatory:"false" json:"isLatestForMajorVersion"`

	// True if this version of the Oracle Database software supports pluggable databases.
	SupportsPdb *bool `mandatory:"false" json:"supportsPdb"`

	// True if this version of the Oracle Database software is the preview version.
	IsPreviewDbVersion *bool `mandatory:"false" json:"isPreviewDbVersion"`
}

func (m DbVersionSummary) String() string {
	return common.PointerString(m)
}
