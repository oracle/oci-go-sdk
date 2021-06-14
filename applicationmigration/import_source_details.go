// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
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
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v42/common"
)

// ImportSourceDetails / Basic details about the source, import manifest and object storage bucket as well as object name of the archive that should be used during import
type ImportSourceDetails struct {
	Manifest *ImportManifest `mandatory:"true" json:"manifest"`

	// the object storage namespace where the bucket and uploaded object resides
	Namespace *string `mandatory:"true" json:"namespace"`

	// the bucket wherein the export archive exists in object storage
	Bucket *string `mandatory:"true" json:"bucket"`

	// the name of the archive as it exists in object storage
	ObjectName *string `mandatory:"true" json:"objectName"`
}

func (m ImportSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ImportSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeImportSourceDetails ImportSourceDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeImportSourceDetails
	}{
		"IMPORT",
		(MarshalTypeImportSourceDetails)(m),
	}

	return json.Marshal(&s)
}
