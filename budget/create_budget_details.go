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

// CreateBudgetDetails The create budget details.
type CreateBudgetDetails struct {

	// The OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the compartment on which budget is applied
	TargetCompartmentId *string `mandatory:"true" json:"targetCompartmentId"`

	// The amount of the budget expressed as a whole number in the currency of the customer's rate card.
	Amount *float32 `mandatory:"true" json:"amount"`

	// The reset period for the budget. We will start with MONTHLY and look into QUARTERLY and maybe ANNUAL post-MVP.
	ResetPeriod CreateBudgetDetailsResetPeriodEnum `mandatory:"true" json:"resetPeriod"`

	// The displayName of the budget.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the budget.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateBudgetDetails) String() string {
	return common.PointerString(m)
}

// CreateBudgetDetailsResetPeriodEnum Enum with underlying type: string
type CreateBudgetDetailsResetPeriodEnum string

// Set of constants representing the allowable values for CreateBudgetDetailsResetPeriodEnum
const (
	CreateBudgetDetailsResetPeriodMonthly CreateBudgetDetailsResetPeriodEnum = "MONTHLY"
)

var mappingCreateBudgetDetailsResetPeriod = map[string]CreateBudgetDetailsResetPeriodEnum{
	"MONTHLY": CreateBudgetDetailsResetPeriodMonthly,
}

// GetCreateBudgetDetailsResetPeriodEnumValues Enumerates the set of values for CreateBudgetDetailsResetPeriodEnum
func GetCreateBudgetDetailsResetPeriodEnumValues() []CreateBudgetDetailsResetPeriodEnum {
	values := make([]CreateBudgetDetailsResetPeriodEnum, 0)
	for _, v := range mappingCreateBudgetDetailsResetPeriod {
		values = append(values, v)
	}
	return values
}
