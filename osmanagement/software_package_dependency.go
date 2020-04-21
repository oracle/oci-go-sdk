// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SoftwarePackageDependency A dependecy for a software package
type SoftwarePackageDependency struct {

	// the software package's dependency
	Dependency *string `mandatory:"false" json:"dependency"`

	// the type of the dependency
	DependencyType *string `mandatory:"false" json:"dependencyType"`

	// the modifier for the dependency
	DependencyModifier *string `mandatory:"false" json:"dependencyModifier"`
}

func (m SoftwarePackageDependency) String() string {
	return common.PointerString(m)
}
