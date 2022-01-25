// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration API
//
// Application Migration simplifies the migration of applications from Oracle Cloud Infrastructure Classic to Oracle Cloud Infrastructure.
// You can use Application Migration API to migrate applications, such as Oracle Java Cloud Service, SOA Cloud Service, and Integration Classic
// instances, to Oracle Cloud Infrastructure. For more information, see
// Overview of Application Migration (https://docs.cloud.oracle.com/iaas/application-migration/appmigrationoverview.htm).
//

package applicationmigration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ImportManifest Manifest describing details about an import source
type ImportManifest struct {

	// the version of the export tool that was used to generate the manifest
	Version *string `mandatory:"false" json:"version"`

	// the type of application that the export tool was executed against to generate this manifest
	ExportType *string `mandatory:"false" json:"exportType"`

	// application specific details as parsed from various sources of the application that was exported
	ExportDetails *interface{} `mandatory:"false" json:"exportDetails"`

	// when this manifest was generated
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`

	// the MD5 hash of the export artifact archive that was produced by the export tool and should be used with this manifest
	Md5 *string `mandatory:"false" json:"md5"`

	// a sha1 hash of all the fields of this manifest (excluding the signature)
	Signature *string `mandatory:"false" json:"signature"`
}

func (m ImportManifest) String() string {
	return common.PointerString(m)
}
