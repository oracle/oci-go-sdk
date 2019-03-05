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

// Budget A budget.
type Budget struct {

	// The OCID of the budget
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the compartment on which budget is applied
	TargetCompartmentId *string `mandatory:"true" json:"targetCompartmentId"`

	// The display name of the budget.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The amount of the budget expressed in the currency of the customer's rate card.
	Amount *float32 `mandatory:"true" json:"amount"`

	// The reset period for the budget. We will start with MONTHLY and look into QUARTERLY and maybe ANNUAL post-MVP.
	ResetPeriod BudgetResetPeriodEnum `mandatory:"true" json:"resetPeriod"`

	// The current state of the budget.
	LifecycleState BudgetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Total number of alert rules in the budget
	AlertRuleCount *int `mandatory:"true" json:"alertRuleCount"`

	// Time that budget was created
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time that budget was updated
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The description of the budget.
	Description *string `mandatory:"false" json:"description"`

	// Version of the budget. Starts from 1 and increments by 1.
	Version *int `mandatory:"false" json:"version"`

	// The actual spend in currency for the current budget cycle
	ActualSpend *float32 `mandatory:"false" json:"actualSpend"`

	// The forecasted spend in currency by the end of the current budget cycle
	ForecastedSpend *float32 `mandatory:"false" json:"forecastedSpend"`

	// The time that the budget spend was last computed
	TimeSpendComputed *common.SDKTime `mandatory:"false" json:"timeSpendComputed"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Budget) String() string {
	return common.PointerString(m)
}

// BudgetResetPeriodEnum Enum with underlying type: string
type BudgetResetPeriodEnum string

// Set of constants representing the allowable values for BudgetResetPeriodEnum
const (
	BudgetResetPeriodMonthly BudgetResetPeriodEnum = "MONTHLY"
)

var mappingBudgetResetPeriod = map[string]BudgetResetPeriodEnum{
	"MONTHLY": BudgetResetPeriodMonthly,
}

// GetBudgetResetPeriodEnumValues Enumerates the set of values for BudgetResetPeriodEnum
func GetBudgetResetPeriodEnumValues() []BudgetResetPeriodEnum {
	values := make([]BudgetResetPeriodEnum, 0)
	for _, v := range mappingBudgetResetPeriod {
		values = append(values, v)
	}
	return values
}

// BudgetLifecycleStateEnum Enum with underlying type: string
type BudgetLifecycleStateEnum string

// Set of constants representing the allowable values for BudgetLifecycleStateEnum
const (
	BudgetLifecycleStateActive   BudgetLifecycleStateEnum = "ACTIVE"
	BudgetLifecycleStateInactive BudgetLifecycleStateEnum = "INACTIVE"
)

var mappingBudgetLifecycleState = map[string]BudgetLifecycleStateEnum{
	"ACTIVE":   BudgetLifecycleStateActive,
	"INACTIVE": BudgetLifecycleStateInactive,
}

// GetBudgetLifecycleStateEnumValues Enumerates the set of values for BudgetLifecycleStateEnum
func GetBudgetLifecycleStateEnumValues() []BudgetLifecycleStateEnum {
	values := make([]BudgetLifecycleStateEnum, 0)
	for _, v := range mappingBudgetLifecycleState {
		values = append(values, v)
	}
	return values
}
