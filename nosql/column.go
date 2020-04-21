// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ndcs-control-plane API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Column A column of a table.
type Column struct {

	// The column name.
	Name *string `mandatory:"false" json:"name"`

	// The column type.
	Type *string `mandatory:"false" json:"type"`

	// The column nullable flag.
	IsNullable *bool `mandatory:"false" json:"isNullable"`

	// The column default value.
	DefaultValue *string `mandatory:"false" json:"defaultValue"`
}

func (m Column) String() string {
	return common.PointerString(m)
}
