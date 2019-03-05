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

// UpdateBudgetDetails The update budget details.
type UpdateBudgetDetails struct {

	// The displayName of the budget.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the budget.
	Description *string `mandatory:"false" json:"description"`

	// The amount of the budget expressed as a whole number in the currency of the customer's rate card.
	Amount *float32 `mandatory:"false" json:"amount"`

	// The reset period for the budget. We will start with MONTHLY and look into QUARTERLY and maybe ANNUAL post-MVP.
	ResetPeriod UpdateBudgetDetailsResetPeriodEnum `mandatory:"false" json:"resetPeriod,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateBudgetDetails) String() string {
	return common.PointerString(m)
}

// UpdateBudgetDetailsResetPeriodEnum Enum with underlying type: string
type UpdateBudgetDetailsResetPeriodEnum string

// Set of constants representing the allowable values for UpdateBudgetDetailsResetPeriodEnum
const (
	UpdateBudgetDetailsResetPeriodMonthly UpdateBudgetDetailsResetPeriodEnum = "MONTHLY"
)

var mappingUpdateBudgetDetailsResetPeriod = map[string]UpdateBudgetDetailsResetPeriodEnum{
	"MONTHLY": UpdateBudgetDetailsResetPeriodMonthly,
}

// GetUpdateBudgetDetailsResetPeriodEnumValues Enumerates the set of values for UpdateBudgetDetailsResetPeriodEnum
func GetUpdateBudgetDetailsResetPeriodEnumValues() []UpdateBudgetDetailsResetPeriodEnum {
	values := make([]UpdateBudgetDetailsResetPeriodEnum, 0)
	for _, v := range mappingUpdateBudgetDetailsResetPeriod {
		values = append(values, v)
	}
	return values
}
