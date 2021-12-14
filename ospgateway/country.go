// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Billing Center Gateway API
//
// This site describes all the Rest endpoints of Billing Center Gateway.
//

package ospgateway

import (
	"github.com/oracle/oci-go-sdk/v54/common"
)

// Country Country details model
type Country struct {

	// Indentifier of the country. This is a DB side unique id which was generated when the entity was created in the table
	CountryId *float32 `mandatory:"false" json:"countryId"`

	// Country code in ISO-3166-1 2-letter format
	CountryCode *string `mandatory:"false" json:"countryCode"`

	// Name of the country
	CountryName *string `mandatory:"false" json:"countryName"`

	// Language identifier
	LanguageId *float32 `mandatory:"false" json:"languageId"`

	// Country code in ISO-3166-1 3-letter format
	Ascii3CountryCode *string `mandatory:"false" json:"ascii3CountryCode"`
}

func (m Country) String() string {
	return common.PointerString(m)
}
