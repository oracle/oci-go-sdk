// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration API
//
// Application Migration simplifies the migration of applications from Oracle Cloud Infrastructure Classic to Oracle Cloud Infrastructure.
// You can use Application Migration API to migrate applications, such as Oracle Java Cloud Service, SOA Cloud Service, and Integration Classic
// instances, to Oracle Cloud Infrastructure. For more information, see
// Overview of Application Migration (https://docs.cloud.oracle.com/iaas/application-migration/appmigrationoverview.htm).
//

package applicationmigration

// SourceTypesEnum Enum with underlying type: string
type SourceTypesEnum string

// Set of constants representing the allowable values for SourceTypesEnum
const (
	SourceTypesOcic            SourceTypesEnum = "OCIC"
	SourceTypesInternalCompute SourceTypesEnum = "INTERNAL_COMPUTE"
)

var mappingSourceTypes = map[string]SourceTypesEnum{
	"OCIC":             SourceTypesOcic,
	"INTERNAL_COMPUTE": SourceTypesInternalCompute,
}

// GetSourceTypesEnumValues Enumerates the set of values for SourceTypesEnum
func GetSourceTypesEnumValues() []SourceTypesEnum {
	values := make([]SourceTypesEnum, 0)
	for _, v := range mappingSourceTypes {
		values = append(values, v)
	}
	return values
}
