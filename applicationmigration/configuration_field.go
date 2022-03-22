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
	"fmt"
	"github.com/oracle/oci-go-sdk/v63/common"
	"strings"
)

// ConfigurationField Provide configuration information about the application in the target environment. Application Migration migrates the application to
// the target environment only after you provide this information. The information that you must provide varies depending on the type of
// application that you are migrating.
type ConfigurationField struct {

	// The name of the configuration field.
	Name *string `mandatory:"false" json:"name"`

	// The name of the group to which this field belongs, if any.
	Group *string `mandatory:"false" json:"group"`

	// The type of the configuration field.
	Type *string `mandatory:"false" json:"type"`

	// The value of the field.
	Value *string `mandatory:"false" json:"value"`

	// Help text to guide the user in setting the configuration value.
	Description *string `mandatory:"false" json:"description"`

	// A list of resources associated with a specific configuration object.
	ResourceList []ResourceField `mandatory:"false" json:"resourceList"`

	// Indicates whether or not the field is required (defaults to `true`).
	IsRequired *bool `mandatory:"false" json:"isRequired"`

	// Indicates whether or not the field may be modified (defaults to `true`).
	IsMutable *bool `mandatory:"false" json:"isMutable"`
}

func (m ConfigurationField) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConfigurationField) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
