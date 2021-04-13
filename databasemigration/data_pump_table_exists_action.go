// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// REST API for Zero Downtime Migration (Oracle Database Migration Service --ODMS-- as customer-facing service name)
//
// Provides users the ability to perform Zero Downtime migration operations
//

package databasemigration

// DataPumpTableExistsActionEnum Enum with underlying type: string
type DataPumpTableExistsActionEnum string

// Set of constants representing the allowable values for DataPumpTableExistsActionEnum
const (
	DataPumpTableExistsActionTruncate DataPumpTableExistsActionEnum = "TRUNCATE"
	DataPumpTableExistsActionReplace  DataPumpTableExistsActionEnum = "REPLACE"
	DataPumpTableExistsActionAppend   DataPumpTableExistsActionEnum = "APPEND"
	DataPumpTableExistsActionSkip     DataPumpTableExistsActionEnum = "SKIP"
)

var mappingDataPumpTableExistsAction = map[string]DataPumpTableExistsActionEnum{
	"TRUNCATE": DataPumpTableExistsActionTruncate,
	"REPLACE":  DataPumpTableExistsActionReplace,
	"APPEND":   DataPumpTableExistsActionAppend,
	"SKIP":     DataPumpTableExistsActionSkip,
}

// GetDataPumpTableExistsActionEnumValues Enumerates the set of values for DataPumpTableExistsActionEnum
func GetDataPumpTableExistsActionEnumValues() []DataPumpTableExistsActionEnum {
	values := make([]DataPumpTableExistsActionEnum, 0)
	for _, v := range mappingDataPumpTableExistsAction {
		values = append(values, v)
	}
	return values
}
