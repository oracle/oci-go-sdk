// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Retrieval API
//
// API for retrieving certificates.
//

package certificates

import (
	"github.com/oracle/oci-go-sdk/v53/common"
)

// CertificateAuthorityBundleVersionCollection The results of a certificate authority (CA) version search. Results contain CA version summary objects and other data.
type CertificateAuthorityBundleVersionCollection struct {

	// A list of CA version summary objects.
	Items []CertificateAuthorityBundleVersionSummary `mandatory:"true" json:"items"`
}

func (m CertificateAuthorityBundleVersionCollection) String() string {
	return common.PointerString(m)
}
