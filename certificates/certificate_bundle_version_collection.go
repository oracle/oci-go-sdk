// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Retrieval API
//
// API for retrieving certificates.
//

package certificates

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CertificateBundleVersionCollection The results of a certificate bundle versions search. Results contain certificate bundle version summary objects.
type CertificateBundleVersionCollection struct {

	// A list of certificate bundle version summary objects.
	Items []CertificateBundleVersionSummary `mandatory:"true" json:"items"`
}

func (m CertificateBundleVersionCollection) String() string {
	return common.PointerString(m)
}
