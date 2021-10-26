// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Retrieval API
//
// API for retrieving certificates.
//

package certificates

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// Validity An object that describes a period of time during which an entity is valid.
type Validity struct {

	// The date on which the certificate validity period begins, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfValidityNotBefore *common.SDKTime `mandatory:"true" json:"timeOfValidityNotBefore"`

	// The date on which the certificate validity period ends, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfValidityNotAfter *common.SDKTime `mandatory:"true" json:"timeOfValidityNotAfter"`
}

func (m Validity) String() string {
	return common.PointerString(m)
}
