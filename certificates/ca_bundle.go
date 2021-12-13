// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Retrieval API
//
// API for retrieving certificates.
//

package certificates

import (
	"github.com/oracle/oci-go-sdk/v54/common"
)

// CaBundle The contents of the CA bundle (root and intermediate certificates), properties of the CA bundle, and user-provided contextual metadata for the CA bundle.
type CaBundle struct {

	// The OCID of the CA bundle.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name for the CA bundle. Names are unique within a compartment. Valid characters include uppercase or lowercase letters, numbers, hyphens, underscores, and periods.
	Name *string `mandatory:"true" json:"name"`

	// Certificates (in PEM format) in the CA bundle. Can be of arbitrary length.
	CaBundlePem *string `mandatory:"true" json:"caBundlePem"`
}

func (m CaBundle) String() string {
	return common.PointerString(m)
}
