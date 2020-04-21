// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

// ScheduleTypesEnum Enum with underlying type: string
type ScheduleTypesEnum string

// Set of constants representing the allowable values for ScheduleTypesEnum
const (
	ScheduleTypesOnetime   ScheduleTypesEnum = "ONETIME"
	ScheduleTypesRecurring ScheduleTypesEnum = "RECURRING"
)

var mappingScheduleTypes = map[string]ScheduleTypesEnum{
	"ONETIME":   ScheduleTypesOnetime,
	"RECURRING": ScheduleTypesRecurring,
}

// GetScheduleTypesEnumValues Enumerates the set of values for ScheduleTypesEnum
func GetScheduleTypesEnumValues() []ScheduleTypesEnum {
	values := make([]ScheduleTypesEnum, 0)
	for _, v := range mappingScheduleTypes {
		values = append(values, v)
	}
	return values
}
