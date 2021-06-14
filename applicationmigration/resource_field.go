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
	"github.com/oracle/oci-go-sdk/v42/common"
)

// ResourceField Resource object that can be used to pass details about any list of resources associated with Migrations. The List of resources are added to ConfigurationField to add the capability to pass lists of resources of any type and group.
type ResourceField struct {

	// The type of the resource field.
	Type *string `mandatory:"true" json:"type"`

	// The value of the field.
	Value *string `mandatory:"true" json:"value"`

	// The display name of the resource field.
	Name *string `mandatory:"false" json:"name"`

	// The name of the group to which this field belongs to.
	Group *string `mandatory:"false" json:"group"`
}

func (m ResourceField) String() string {
	return common.PointerString(m)
}
