// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// Manage Oracle Cloud Infrastructure console announcements.
//

package announcementsservice

import (
	"github.com/oracle/oci-go-sdk/v47/common"
)

// Property A generic property that might be associated with the affected resource.
type Property struct {

	// The name of the property.
	Name *string `mandatory:"true" json:"name"`

	// The value of the property.
	Value *string `mandatory:"true" json:"value"`
}

func (m Property) String() string {
	return common.PointerString(m)
}
