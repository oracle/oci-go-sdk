// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
//     pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
//     This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

// NerModelVersionEnum Enum with underlying type: string
type NerModelVersionEnum string

// Set of constants representing the allowable values for NerModelVersionEnum
const (
	NerModelVersionV21 NerModelVersionEnum = "V2.1"
	NerModelVersionV11 NerModelVersionEnum = "V1.1"
)

var mappingNerModelVersion = map[string]NerModelVersionEnum{
	"V2.1": NerModelVersionV21,
	"V1.1": NerModelVersionV11,
}

// GetNerModelVersionEnumValues Enumerates the set of values for NerModelVersionEnum
func GetNerModelVersionEnumValues() []NerModelVersionEnum {
	values := make([]NerModelVersionEnum, 0)
	for _, v := range mappingNerModelVersion {
		values = append(values, v)
	}
	return values
}
