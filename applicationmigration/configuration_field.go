// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ConfigurationField Information required to migrate an application. Populated by the service as the source application is introspected
type ConfigurationField struct {

	// The name of the configuration field
	Name *string `mandatory:"false" json:"name"`

	// The name of the group to which this field belongs, if any.
	Group *string `mandatory:"false" json:"group"`

	// The configuration field type
	Type *string `mandatory:"false" json:"type"`

	// The value of the field
	Value *string `mandatory:"false" json:"value"`

	// Help text to guide the customer in setting the configuration value
	Description *string `mandatory:"false" json:"description"`

	// Indicates whether or not the field is required (defaults to true)
	IsRequired *bool `mandatory:"false" json:"isRequired"`

	// Indicates whether or not the field may be modified (defaults to true)
	IsMutable *bool `mandatory:"false" json:"isMutable"`
}

func (m ConfigurationField) String() string {
	return common.PointerString(m)
}
