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

// RevocationStatus The current revocation status of the certificate or certificate authority (CA).
type RevocationStatus struct {

	// The time when the certificate or CA was revoked.
	TimeRevoked *common.SDKTime `mandatory:"true" json:"timeRevoked"`

	// The reason that the certificate or CA was revoked.
	RevocationReason RevocationReasonEnum `mandatory:"true" json:"revocationReason"`
}

func (m RevocationStatus) String() string {
	return common.PointerString(m)
}
