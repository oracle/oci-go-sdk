// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"strings"
)

// CertSignatureAlgorithmEnum Enum with underlying type: string
type CertSignatureAlgorithmEnum string

// Set of constants representing the allowable values for CertSignatureAlgorithmEnum
const (
	CertSignatureAlgorithmSha256WithRsa   CertSignatureAlgorithmEnum = "SHA256_WITH_RSA"
	CertSignatureAlgorithmSha384WithRsa   CertSignatureAlgorithmEnum = "SHA384_WITH_RSA"
	CertSignatureAlgorithmSha512WithRsa   CertSignatureAlgorithmEnum = "SHA512_WITH_RSA"
	CertSignatureAlgorithmSha256WithEcdsa CertSignatureAlgorithmEnum = "SHA256_WITH_ECDSA"
	CertSignatureAlgorithmSha384WithEcdsa CertSignatureAlgorithmEnum = "SHA384_WITH_ECDSA"
	CertSignatureAlgorithmSha512WithEcdsa CertSignatureAlgorithmEnum = "SHA512_WITH_ECDSA"
)

var mappingCertSignatureAlgorithmEnum = map[string]CertSignatureAlgorithmEnum{
	"SHA256_WITH_RSA":   CertSignatureAlgorithmSha256WithRsa,
	"SHA384_WITH_RSA":   CertSignatureAlgorithmSha384WithRsa,
	"SHA512_WITH_RSA":   CertSignatureAlgorithmSha512WithRsa,
	"SHA256_WITH_ECDSA": CertSignatureAlgorithmSha256WithEcdsa,
	"SHA384_WITH_ECDSA": CertSignatureAlgorithmSha384WithEcdsa,
	"SHA512_WITH_ECDSA": CertSignatureAlgorithmSha512WithEcdsa,
}

var mappingCertSignatureAlgorithmEnumLowerCase = map[string]CertSignatureAlgorithmEnum{
	"sha256_with_rsa":   CertSignatureAlgorithmSha256WithRsa,
	"sha384_with_rsa":   CertSignatureAlgorithmSha384WithRsa,
	"sha512_with_rsa":   CertSignatureAlgorithmSha512WithRsa,
	"sha256_with_ecdsa": CertSignatureAlgorithmSha256WithEcdsa,
	"sha384_with_ecdsa": CertSignatureAlgorithmSha384WithEcdsa,
	"sha512_with_ecdsa": CertSignatureAlgorithmSha512WithEcdsa,
}

// GetCertSignatureAlgorithmEnumValues Enumerates the set of values for CertSignatureAlgorithmEnum
func GetCertSignatureAlgorithmEnumValues() []CertSignatureAlgorithmEnum {
	values := make([]CertSignatureAlgorithmEnum, 0)
	for _, v := range mappingCertSignatureAlgorithmEnum {
		values = append(values, v)
	}
	return values
}

// GetCertSignatureAlgorithmEnumStringValues Enumerates the set of values in String for CertSignatureAlgorithmEnum
func GetCertSignatureAlgorithmEnumStringValues() []string {
	return []string{
		"SHA256_WITH_RSA",
		"SHA384_WITH_RSA",
		"SHA512_WITH_RSA",
		"SHA256_WITH_ECDSA",
		"SHA384_WITH_ECDSA",
		"SHA512_WITH_ECDSA",
	}
}

// GetMappingCertSignatureAlgorithmEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCertSignatureAlgorithmEnum(val string) (CertSignatureAlgorithmEnum, bool) {
	enum, ok := mappingCertSignatureAlgorithmEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
