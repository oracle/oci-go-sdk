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

// SoftwarePackageSearchSummary Summary information for a software package
type SoftwarePackageSearchSummary struct {

	// Package name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique identifier for the package. NOTE - This is not an OCID
	Name *string `mandatory:"true" json:"name"`

	// Type of the package
	Type *string `mandatory:"true" json:"type"`

	// Version of the package
	Version *string `mandatory:"true" json:"version"`

	// the architecture for which this software was built
	Architecture *string `mandatory:"false" json:"architecture"`

	// a summary description of the software package
	Summary *string `mandatory:"false" json:"summary"`

	// Type of the erratum.
	AdvisoryType UpdateTypesEnum `mandatory:"false" json:"advisoryType,omitempty"`

	// List of errata containing this software package
	Errata []Id `mandatory:"false" json:"errata"`

	// list of software sources that provide the software package
	SoftwareSources []SoftwareSourceId `mandatory:"false" json:"softwareSources"`
}

func (m SoftwarePackageSearchSummary) String() string {
	return common.PointerString(m)
}
