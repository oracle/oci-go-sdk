// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// BudgetsControlPlane API
//
// Use the BudgetsControlPlane API to manage budgets and budget alerts.
//

package budget

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateAlertRuleDetails The update alert rule details.
type UpdateAlertRuleDetails struct {

	// The name of the alert rule.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Type of alert. Valid values are ACTUAL (the alert will trigger based on actual usage) or
	// FORECAST (the alert will trigger based on predicted usage).
	Type UpdateAlertRuleDetailsTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The threshold for triggering the alert expressed as a whole number or decimal value.
	// If thresholdType is ABSOLUTE, threshold can have at most 12 digits before the decimal point and up to 2 digits after the decimal point.
	// If thresholdType is PERCENTAGE, the maximum value is 10000 and can have up to 2 digits after the decimal point.
	Threshold *float32 `mandatory:"false" json:"threshold"`

	// The type of threshold.
	ThresholdType UpdateAlertRuleDetailsThresholdTypeEnum `mandatory:"false" json:"thresholdType,omitempty"`

	// The audience that will received the alert when it triggers.
	Recipients *string `mandatory:"false" json:"recipients"`

	// The description of the alert rule
	Description *string `mandatory:"false" json:"description"`

	// The message to be delivered to the recipients when alert is triggered
	Message *string `mandatory:"false" json:"message"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateAlertRuleDetails) String() string {
	return common.PointerString(m)
}

// UpdateAlertRuleDetailsTypeEnum Enum with underlying type: string
type UpdateAlertRuleDetailsTypeEnum string

// Set of constants representing the allowable values for UpdateAlertRuleDetailsTypeEnum
const (
	UpdateAlertRuleDetailsTypeActual   UpdateAlertRuleDetailsTypeEnum = "ACTUAL"
	UpdateAlertRuleDetailsTypeForecast UpdateAlertRuleDetailsTypeEnum = "FORECAST"
)

var mappingUpdateAlertRuleDetailsType = map[string]UpdateAlertRuleDetailsTypeEnum{
	"ACTUAL":   UpdateAlertRuleDetailsTypeActual,
	"FORECAST": UpdateAlertRuleDetailsTypeForecast,
}

// GetUpdateAlertRuleDetailsTypeEnumValues Enumerates the set of values for UpdateAlertRuleDetailsTypeEnum
func GetUpdateAlertRuleDetailsTypeEnumValues() []UpdateAlertRuleDetailsTypeEnum {
	values := make([]UpdateAlertRuleDetailsTypeEnum, 0)
	for _, v := range mappingUpdateAlertRuleDetailsType {
		values = append(values, v)
	}
	return values
}

// UpdateAlertRuleDetailsThresholdTypeEnum Enum with underlying type: string
type UpdateAlertRuleDetailsThresholdTypeEnum string

// Set of constants representing the allowable values for UpdateAlertRuleDetailsThresholdTypeEnum
const (
	UpdateAlertRuleDetailsThresholdTypePercentage UpdateAlertRuleDetailsThresholdTypeEnum = "PERCENTAGE"
	UpdateAlertRuleDetailsThresholdTypeAbsolute   UpdateAlertRuleDetailsThresholdTypeEnum = "ABSOLUTE"
)

var mappingUpdateAlertRuleDetailsThresholdType = map[string]UpdateAlertRuleDetailsThresholdTypeEnum{
	"PERCENTAGE": UpdateAlertRuleDetailsThresholdTypePercentage,
	"ABSOLUTE":   UpdateAlertRuleDetailsThresholdTypeAbsolute,
}

// GetUpdateAlertRuleDetailsThresholdTypeEnumValues Enumerates the set of values for UpdateAlertRuleDetailsThresholdTypeEnum
func GetUpdateAlertRuleDetailsThresholdTypeEnumValues() []UpdateAlertRuleDetailsThresholdTypeEnum {
	values := make([]UpdateAlertRuleDetailsThresholdTypeEnum, 0)
	for _, v := range mappingUpdateAlertRuleDetailsThresholdType {
		values = append(values, v)
	}
	return values
}
