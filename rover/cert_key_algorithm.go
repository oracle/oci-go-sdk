// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CertKeyAlgorithmEnum Enum with underlying type: string
type CertKeyAlgorithmEnum string

// Set of constants representing the allowable values for CertKeyAlgorithmEnum
const (
	CertKeyAlgorithmRsa2048   CertKeyAlgorithmEnum = "RSA2048"
	CertKeyAlgorithmRsa4096   CertKeyAlgorithmEnum = "RSA4096"
	CertKeyAlgorithmEcdsaP256 CertKeyAlgorithmEnum = "ECDSA_P256"
	CertKeyAlgorithmEcdsaP384 CertKeyAlgorithmEnum = "ECDSA_P384"
)

var mappingCertKeyAlgorithmEnum = map[string]CertKeyAlgorithmEnum{
	"RSA2048":    CertKeyAlgorithmRsa2048,
	"RSA4096":    CertKeyAlgorithmRsa4096,
	"ECDSA_P256": CertKeyAlgorithmEcdsaP256,
	"ECDSA_P384": CertKeyAlgorithmEcdsaP384,
}

var mappingCertKeyAlgorithmEnumLowerCase = map[string]CertKeyAlgorithmEnum{
	"rsa2048":    CertKeyAlgorithmRsa2048,
	"rsa4096":    CertKeyAlgorithmRsa4096,
	"ecdsa_p256": CertKeyAlgorithmEcdsaP256,
	"ecdsa_p384": CertKeyAlgorithmEcdsaP384,
}

// GetCertKeyAlgorithmEnumValues Enumerates the set of values for CertKeyAlgorithmEnum
func GetCertKeyAlgorithmEnumValues() []CertKeyAlgorithmEnum {
	values := make([]CertKeyAlgorithmEnum, 0)
	for _, v := range mappingCertKeyAlgorithmEnum {
		values = append(values, v)
	}
	return values
}

// GetCertKeyAlgorithmEnumStringValues Enumerates the set of values in String for CertKeyAlgorithmEnum
func GetCertKeyAlgorithmEnumStringValues() []string {
	return []string{
		"RSA2048",
		"RSA4096",
		"ECDSA_P256",
		"ECDSA_P384",
	}
}

// GetMappingCertKeyAlgorithmEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCertKeyAlgorithmEnum(val string) (CertKeyAlgorithmEnum, bool) {
	enum, ok := mappingCertKeyAlgorithmEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
