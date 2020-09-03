// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Attribute Attribute
type Attribute struct {

	// default value
	DefaultValue *interface{} `mandatory:"false" json:"defaultValue"`

	// dynamic value range reference attribute
	DynamicValueRangeRefAttr *string `mandatory:"false" json:"dynamicValueRangeRefAttr"`

	// maximum length
	MaximumLen AttributeMaximumLenEnum `mandatory:"false" json:"maximumLen,omitempty"`

	// name
	Name *string `mandatory:"false" json:"name"`

	// populated by
	PopulatedBy AttributePopulatedByEnum `mandatory:"false" json:"populatedBy,omitempty"`

	// required in JSON
	RequiredInJSON AttributeRequiredInJSONEnum `mandatory:"false" json:"requiredInJSON,omitempty"`

	// schema column
	SchemaColumn *string `mandatory:"false" json:"schemaColumn"`

	// is string exceed maximum length
	IsStringExceedMaximumLength *bool `mandatory:"false" json:"isStringExceedMaximumLength"`

	// usage senario
	UsageSenario AttributeUsageSenarioEnum `mandatory:"false" json:"usageSenario,omitempty"`

	// value data type
	ValueDataType AttributeValueDataTypeEnum `mandatory:"false" json:"valueDataType,omitempty"`

	// value population priority
	ValuePopulationPriority AttributeValuePopulationPriorityEnum `mandatory:"false" json:"valuePopulationPriority,omitempty"`
}

func (m Attribute) String() string {
	return common.PointerString(m)
}

// AttributeMaximumLenEnum Enum with underlying type: string
type AttributeMaximumLenEnum string

// Set of constants representing the allowable values for AttributeMaximumLenEnum
const (
	AttributeMaximumLenFive         AttributeMaximumLenEnum = "LENGTH_FIVE"
	AttributeMaximumLenSixteen      AttributeMaximumLenEnum = "LENGTH_SIXTEEN"
	AttributeMaximumLenThirtytwo    AttributeMaximumLenEnum = "LENGTH_THIRTYTWO"
	AttributeMaximumLenSixtyfour    AttributeMaximumLenEnum = "LENGTH_SIXTYFOUR"
	AttributeMaximumLenOnetwoeight  AttributeMaximumLenEnum = "LENGTH_ONETWOEIGHT"
	AttributeMaximumLenTwofiftysix  AttributeMaximumLenEnum = "LENGTH_TWOFIFTYSIX"
	AttributeMaximumLenFivetwelve   AttributeMaximumLenEnum = "LENGTH_FIVETWELVE"
	AttributeMaximumLenSevenfifty   AttributeMaximumLenEnum = "LENGTH_SEVENFIFTY"
	AttributeMaximumLenOneThousand  AttributeMaximumLenEnum = "LENGTH_ONE_THOUSAND"
	AttributeMaximumLenTwoThousand  AttributeMaximumLenEnum = "LENGTH_TWO_THOUSAND"
	AttributeMaximumLenFourThousand AttributeMaximumLenEnum = "LENGTH_FOUR_THOUSAND"
)

var mappingAttributeMaximumLen = map[string]AttributeMaximumLenEnum{
	"LENGTH_FIVE":          AttributeMaximumLenFive,
	"LENGTH_SIXTEEN":       AttributeMaximumLenSixteen,
	"LENGTH_THIRTYTWO":     AttributeMaximumLenThirtytwo,
	"LENGTH_SIXTYFOUR":     AttributeMaximumLenSixtyfour,
	"LENGTH_ONETWOEIGHT":   AttributeMaximumLenOnetwoeight,
	"LENGTH_TWOFIFTYSIX":   AttributeMaximumLenTwofiftysix,
	"LENGTH_FIVETWELVE":    AttributeMaximumLenFivetwelve,
	"LENGTH_SEVENFIFTY":    AttributeMaximumLenSevenfifty,
	"LENGTH_ONE_THOUSAND":  AttributeMaximumLenOneThousand,
	"LENGTH_TWO_THOUSAND":  AttributeMaximumLenTwoThousand,
	"LENGTH_FOUR_THOUSAND": AttributeMaximumLenFourThousand,
}

// GetAttributeMaximumLenEnumValues Enumerates the set of values for AttributeMaximumLenEnum
func GetAttributeMaximumLenEnumValues() []AttributeMaximumLenEnum {
	values := make([]AttributeMaximumLenEnum, 0)
	for _, v := range mappingAttributeMaximumLen {
		values = append(values, v)
	}
	return values
}

// AttributePopulatedByEnum Enum with underlying type: string
type AttributePopulatedByEnum string

// Set of constants representing the allowable values for AttributePopulatedByEnum
const (
	AttributePopulatedByBackendGen AttributePopulatedByEnum = "BACKEND_GEN"
	AttributePopulatedByCallerGen  AttributePopulatedByEnum = "CALLER_GEN"
)

var mappingAttributePopulatedBy = map[string]AttributePopulatedByEnum{
	"BACKEND_GEN": AttributePopulatedByBackendGen,
	"CALLER_GEN":  AttributePopulatedByCallerGen,
}

// GetAttributePopulatedByEnumValues Enumerates the set of values for AttributePopulatedByEnum
func GetAttributePopulatedByEnumValues() []AttributePopulatedByEnum {
	values := make([]AttributePopulatedByEnum, 0)
	for _, v := range mappingAttributePopulatedBy {
		values = append(values, v)
	}
	return values
}

// AttributeRequiredInJSONEnum Enum with underlying type: string
type AttributeRequiredInJSONEnum string

// Set of constants representing the allowable values for AttributeRequiredInJSONEnum
const (
	AttributeRequiredInJSONMandatory AttributeRequiredInJSONEnum = "MANDATORY"
	AttributeRequiredInJSONOptional  AttributeRequiredInJSONEnum = "OPTIONAL"
)

var mappingAttributeRequiredInJSON = map[string]AttributeRequiredInJSONEnum{
	"MANDATORY": AttributeRequiredInJSONMandatory,
	"OPTIONAL":  AttributeRequiredInJSONOptional,
}

// GetAttributeRequiredInJSONEnumValues Enumerates the set of values for AttributeRequiredInJSONEnum
func GetAttributeRequiredInJSONEnumValues() []AttributeRequiredInJSONEnum {
	values := make([]AttributeRequiredInJSONEnum, 0)
	for _, v := range mappingAttributeRequiredInJSON {
		values = append(values, v)
	}
	return values
}

// AttributeUsageSenarioEnum Enum with underlying type: string
type AttributeUsageSenarioEnum string

// Set of constants representing the allowable values for AttributeUsageSenarioEnum
const (
	AttributeUsageSenarioCreate             AttributeUsageSenarioEnum = "CREATE"
	AttributeUsageSenarioUpdate             AttributeUsageSenarioEnum = "UPDATE"
	AttributeUsageSenarioCreateAndUpdate    AttributeUsageSenarioEnum = "CREATE_AND_UPDATE"
	AttributeUsageSenarioDelete             AttributeUsageSenarioEnum = "DELETE"
	AttributeUsageSenarioReCreate           AttributeUsageSenarioEnum = "RE_CREATE"
	AttributeUsageSenarioDetail             AttributeUsageSenarioEnum = "DETAIL"
	AttributeUsageSenarioList               AttributeUsageSenarioEnum = "LIST"
	AttributeUsageSenarioFunctionWithLookup AttributeUsageSenarioEnum = "FUNCTION_WITH_LOOKUP"
	AttributeUsageSenarioDbPattern          AttributeUsageSenarioEnum = "DB_PATTERN"
	AttributeUsageSenarioCreateFirsttimeT1  AttributeUsageSenarioEnum = "CREATE_FIRSTTIME_T1"
	AttributeUsageSenarioUpdateOobMetric    AttributeUsageSenarioEnum = "UPDATE_OOB_METRIC"
)

var mappingAttributeUsageSenario = map[string]AttributeUsageSenarioEnum{
	"CREATE":               AttributeUsageSenarioCreate,
	"UPDATE":               AttributeUsageSenarioUpdate,
	"CREATE_AND_UPDATE":    AttributeUsageSenarioCreateAndUpdate,
	"DELETE":               AttributeUsageSenarioDelete,
	"RE_CREATE":            AttributeUsageSenarioReCreate,
	"DETAIL":               AttributeUsageSenarioDetail,
	"LIST":                 AttributeUsageSenarioList,
	"FUNCTION_WITH_LOOKUP": AttributeUsageSenarioFunctionWithLookup,
	"DB_PATTERN":           AttributeUsageSenarioDbPattern,
	"CREATE_FIRSTTIME_T1":  AttributeUsageSenarioCreateFirsttimeT1,
	"UPDATE_OOB_METRIC":    AttributeUsageSenarioUpdateOobMetric,
}

// GetAttributeUsageSenarioEnumValues Enumerates the set of values for AttributeUsageSenarioEnum
func GetAttributeUsageSenarioEnumValues() []AttributeUsageSenarioEnum {
	values := make([]AttributeUsageSenarioEnum, 0)
	for _, v := range mappingAttributeUsageSenario {
		values = append(values, v)
	}
	return values
}

// AttributeValueDataTypeEnum Enum with underlying type: string
type AttributeValueDataTypeEnum string

// Set of constants representing the allowable values for AttributeValueDataTypeEnum
const (
	AttributeValueDataTypeInteger         AttributeValueDataTypeEnum = "INTEGER"
	AttributeValueDataTypeLong            AttributeValueDataTypeEnum = "LONG"
	AttributeValueDataTypeFloat           AttributeValueDataTypeEnum = "FLOAT"
	AttributeValueDataTypeString          AttributeValueDataTypeEnum = "STRING"
	AttributeValueDataTypeTimestamp       AttributeValueDataTypeEnum = "TIMESTAMP"
	AttributeValueDataTypeDate            AttributeValueDataTypeEnum = "DATE"
	AttributeValueDataTypeClob            AttributeValueDataTypeEnum = "CLOB"
	AttributeValueDataTypeTagRef          AttributeValueDataTypeEnum = "TAG_REF"
	AttributeValueDataTypeParserRef       AttributeValueDataTypeEnum = "PARSER_REF"
	AttributeValueDataTypeSttRef          AttributeValueDataTypeEnum = "STT_REF"
	AttributeValueDataTypeLookupRef       AttributeValueDataTypeEnum = "LOOKUP_REF"
	AttributeValueDataTypeMetaFunctionRef AttributeValueDataTypeEnum = "META_FUNCTION_REF"
	AttributeValueDataTypeCommonFieldRef  AttributeValueDataTypeEnum = "COMMON_FIELD_REF"
)

var mappingAttributeValueDataType = map[string]AttributeValueDataTypeEnum{
	"INTEGER":           AttributeValueDataTypeInteger,
	"LONG":              AttributeValueDataTypeLong,
	"FLOAT":             AttributeValueDataTypeFloat,
	"STRING":            AttributeValueDataTypeString,
	"TIMESTAMP":         AttributeValueDataTypeTimestamp,
	"DATE":              AttributeValueDataTypeDate,
	"CLOB":              AttributeValueDataTypeClob,
	"TAG_REF":           AttributeValueDataTypeTagRef,
	"PARSER_REF":        AttributeValueDataTypeParserRef,
	"STT_REF":           AttributeValueDataTypeSttRef,
	"LOOKUP_REF":        AttributeValueDataTypeLookupRef,
	"META_FUNCTION_REF": AttributeValueDataTypeMetaFunctionRef,
	"COMMON_FIELD_REF":  AttributeValueDataTypeCommonFieldRef,
}

// GetAttributeValueDataTypeEnumValues Enumerates the set of values for AttributeValueDataTypeEnum
func GetAttributeValueDataTypeEnumValues() []AttributeValueDataTypeEnum {
	values := make([]AttributeValueDataTypeEnum, 0)
	for _, v := range mappingAttributeValueDataType {
		values = append(values, v)
	}
	return values
}

// AttributeValuePopulationPriorityEnum Enum with underlying type: string
type AttributeValuePopulationPriorityEnum string

// Set of constants representing the allowable values for AttributeValuePopulationPriorityEnum
const (
	AttributeValuePopulationPriorityNone AttributeValuePopulationPriorityEnum = "NONE"
	AttributeValuePopulationPriorityLow  AttributeValuePopulationPriorityEnum = "LOW"
	AttributeValuePopulationPriorityHigh AttributeValuePopulationPriorityEnum = "HIGH"
)

var mappingAttributeValuePopulationPriority = map[string]AttributeValuePopulationPriorityEnum{
	"NONE": AttributeValuePopulationPriorityNone,
	"LOW":  AttributeValuePopulationPriorityLow,
	"HIGH": AttributeValuePopulationPriorityHigh,
}

// GetAttributeValuePopulationPriorityEnumValues Enumerates the set of values for AttributeValuePopulationPriorityEnum
func GetAttributeValuePopulationPriorityEnumValues() []AttributeValuePopulationPriorityEnum {
	values := make([]AttributeValuePopulationPriorityEnum, 0)
	for _, v := range mappingAttributeValuePopulationPriority {
		values = append(values, v)
	}
	return values
}
