// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

// ChecksumTypesEnum Enum with underlying type: string
type ChecksumTypesEnum string

// Set of constants representing the allowable values for ChecksumTypesEnum
const (
	ChecksumTypesSha1   ChecksumTypesEnum = "SHA1"
	ChecksumTypesSha256 ChecksumTypesEnum = "SHA256"
	ChecksumTypesSha384 ChecksumTypesEnum = "SHA384"
	ChecksumTypesSha512 ChecksumTypesEnum = "SHA512"
)

var mappingChecksumTypes = map[string]ChecksumTypesEnum{
	"SHA1":   ChecksumTypesSha1,
	"SHA256": ChecksumTypesSha256,
	"SHA384": ChecksumTypesSha384,
	"SHA512": ChecksumTypesSha512,
}

// GetChecksumTypesEnumValues Enumerates the set of values for ChecksumTypesEnum
func GetChecksumTypesEnumValues() []ChecksumTypesEnum {
	values := make([]ChecksumTypesEnum, 0)
	for _, v := range mappingChecksumTypes {
		values = append(values, v)
	}
	return values
}
