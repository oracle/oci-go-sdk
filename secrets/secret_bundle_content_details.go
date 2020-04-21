// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secrets
//
// API for retrieving secrets from vaults.
//

package secrets

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// SecretBundleContentDetails The contents of the secret.
type SecretBundleContentDetails interface {
}

type secretbundlecontentdetails struct {
	JsonData    []byte
	ContentType string `json:"contentType"`
}

// UnmarshalJSON unmarshals json
func (m *secretbundlecontentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersecretbundlecontentdetails secretbundlecontentdetails
	s := struct {
		Model Unmarshalersecretbundlecontentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ContentType = s.Model.ContentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *secretbundlecontentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ContentType {
	case "BASE64":
		mm := Base64SecretBundleContentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m secretbundlecontentdetails) String() string {
	return common.PointerString(m)
}

// SecretBundleContentDetailsContentTypeEnum Enum with underlying type: string
type SecretBundleContentDetailsContentTypeEnum string

// Set of constants representing the allowable values for SecretBundleContentDetailsContentTypeEnum
const (
	SecretBundleContentDetailsContentTypeBase64 SecretBundleContentDetailsContentTypeEnum = "BASE64"
)

var mappingSecretBundleContentDetailsContentType = map[string]SecretBundleContentDetailsContentTypeEnum{
	"BASE64": SecretBundleContentDetailsContentTypeBase64,
}

// GetSecretBundleContentDetailsContentTypeEnumValues Enumerates the set of values for SecretBundleContentDetailsContentTypeEnum
func GetSecretBundleContentDetailsContentTypeEnumValues() []SecretBundleContentDetailsContentTypeEnum {
	values := make([]SecretBundleContentDetailsContentTypeEnum, 0)
	for _, v := range mappingSecretBundleContentDetailsContentType {
		values = append(values, v)
	}
	return values
}
