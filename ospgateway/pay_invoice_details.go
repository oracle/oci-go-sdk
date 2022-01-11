// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Billing Center Gateway API
//
// This site describes all the Rest endpoints of Billing Center Gateway.
//

package ospgateway

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// PayInvoiceDetails Request object for invoice payment
type PayInvoiceDetails struct {

	// User email
	Email *string `mandatory:"true" json:"email"`

	// Language code
	LanguageCode *string `mandatory:"false" json:"languageCode"`

	// Callback URL
	ReturnUrl *string `mandatory:"false" json:"returnUrl"`
}

func (m PayInvoiceDetails) String() string {
	return common.PointerString(m)
}
