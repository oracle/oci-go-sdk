// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for Budget API

package example

import (
	"context"
	"fmt"
	"strings"

	"github.com/oracle/oci-go-sdk/budget"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/example/helpers"
)

// ExampleBudget shows the sample for budget operations: create, update, get, list etc...
func ExampleBudget() {
	client, err := budget.NewBudgetClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	ctx := context.Background()

	// create a COMPARTMENT budget
	// TargetType is CreateBudgetDetailsTargetTypeCompartment
	// Target is the target compartment OCID
	createBudgetResp, err := createBudget(
		context.Background(),
		client,
		budget.CreateBudgetDetailsTargetTypeCompartment,
		helpers.CompartmentID())
	helpers.FatalIfError(err)
	fmt.Println("Compartment budget created")

	budgetId := createBudgetResp.Budget.Id

	// create a TAG budget
	// TargetType is CreateBudgetDetailsTargetTypeTag
	// Target is the target tag in string format {TagNamespace}.{TagDefinition}.{TagValue}
	tagNamespace := "tag_namespace_for_go_sdk_example"
	tagDefinition := "tag_definition_for_go_sdk_example"
	tagValue := "tag_value_for_go_sdk_example"
	createBudgetResp, err = createBudget(
		context.Background(),
		client,
		budget.CreateBudgetDetailsTargetTypeTag,
		common.String(strings.Join([]string{tagNamespace, tagDefinition, tagValue}, ".")))
	helpers.FatalIfError(err)
	fmt.Println("Tag budget created")

	tagBudgetId := createBudgetResp.Budget.Id

	// list budgets, list operations are paginated and take a "page" parameter
	// to allow you to get the next batch of items from the server
	// for pagination sample, please refer to 'example_core_pagination_test.go'
	listBudgetsReq := budget.ListBudgetsRequest{
		CompartmentId: helpers.CompartmentID(),
	}
	_, err = client.ListBudgets(ctx, listBudgetsReq)
	helpers.FatalIfError(err)
	fmt.Println("List budgets")

	// get budget
	getBudgetReq := budget.GetBudgetRequest{
		BudgetId: budgetId,
	}
	_, err = client.GetBudget(ctx, getBudgetReq)
	helpers.FatalIfError(err)
	fmt.Println("Get budget")

	// update budget
	updateBudgetReq := budget.UpdateBudgetRequest{
		BudgetId: budgetId,
		UpdateBudgetDetails: budget.UpdateBudgetDetails{
			DisplayName: common.String("GOSDKSampleBudgetUpdated"),
			Description: common.String("GOSDK Sample Budget Description - Updated"),
			Amount:      common.Float32(20000),
		},
	}

	_, err = client.UpdateBudget(ctx, updateBudgetReq)
	helpers.FatalIfError(err)
	fmt.Println("Budget updated")

	// Create alert rule
	createAlertRuleDetails := budget.CreateAlertRuleDetails{
		Type:          budget.CreateAlertRuleDetailsTypeActual,
		Threshold:     common.Float32(8000),
		ThresholdType: budget.CreateAlertRuleDetailsThresholdTypeAbsolute,
		Recipients:    common.String("oci_go_sdk_example@oracle.com"),
	}

	createAlertRuleReq := budget.CreateAlertRuleRequest{
		BudgetId:               budgetId,
		CreateAlertRuleDetails: createAlertRuleDetails,
	}

	createAlertRuleResp, err := client.CreateAlertRule(ctx, createAlertRuleReq)
	helpers.FatalIfError(err)
	fmt.Println("AlertRule created")

	alertRuleId := createAlertRuleResp.AlertRule.Id

	// List alert rules
	listAlertRulesReq := budget.ListAlertRulesRequest{
		BudgetId: budgetId,
	}

	_, err = client.ListAlertRules(ctx, listAlertRulesReq)
	helpers.FatalIfError(err)
	fmt.Println("List alert rules")

	// Get alert rule
	getAlertRuleReq := budget.GetAlertRuleRequest{
		BudgetId:    budgetId,
		AlertRuleId: alertRuleId,
	}

	_, err = client.GetAlertRule(ctx, getAlertRuleReq)
	helpers.FatalIfError(err)
	fmt.Println("Get alert rule")

	// Update alert rule
	updateAlertRuleDetails := budget.UpdateAlertRuleDetails{
		DisplayName: common.String("GOSDKSampleBudgetUpdated"),
		Description: common.String("GOSDK Sample AlertRule Description - Updated"),
		Threshold:   common.Float32(16000),
	}

	updateAlertRuleReq := budget.UpdateAlertRuleRequest{
		BudgetId:               budgetId,
		AlertRuleId:            alertRuleId,
		UpdateAlertRuleDetails: updateAlertRuleDetails,
	}

	_, err = client.UpdateAlertRule(ctx, updateAlertRuleReq)
	helpers.FatalIfError(err)
	fmt.Println("AlertRule updated")

	// Delete alert rule
	deleteAlertRuleReq := budget.DeleteAlertRuleRequest{
		BudgetId:    budgetId,
		AlertRuleId: alertRuleId,
	}

	_, err = client.DeleteAlertRule(ctx, deleteAlertRuleReq)
	helpers.FatalIfError(err)
	fmt.Println("AlertRule deleted")

	// Delete Compartment budget
	deleteBudgetReq := budget.DeleteBudgetRequest{
		BudgetId: budgetId,
	}

	_, err = client.DeleteBudget(ctx, deleteBudgetReq)
	helpers.FatalIfError(err)
	fmt.Println("Compartment budget deleted")

	// Delete Tag budget
	deleteBudgetReq = budget.DeleteBudgetRequest{
		BudgetId: tagBudgetId,
	}

	_, err = client.DeleteBudget(ctx, deleteBudgetReq)
	helpers.FatalIfError(err)
	fmt.Println("Tag budget deleted")

	// Output:
	// Compartment budget created
	// Tag budget created
	// List budgets
	// Get budget
	// Budget updated
	// AlertRule created
	// List alert rules
	// Get alert rule
	// AlertRule updated
	// AlertRule deleted
	// Compartment budget deleted
	// Tag budget deleted
}

func createBudget(ctx context.Context, client budget.BudgetClient, targetType budget.CreateBudgetDetailsTargetTypeEnum, target *string) (budget.CreateBudgetResponse, error) {
	createBudgetDetails := budget.CreateBudgetDetails{
		CompartmentId: helpers.RootCompartmentID(),
		TargetType:    targetType,
		Targets:       []string{*target},
		DisplayName:   common.String(strings.Join([]string{"GOSDKSampleBudget", helpers.GetRandomString(8)}, "_")),
		Description:   common.String(strings.Join([]string{"GOSDK Sample Budget Description", helpers.GetRandomString(8)}, "_")),
		Amount:        common.Float32(10000),
		ResetPeriod:   budget.CreateBudgetDetailsResetPeriodMonthly,
	}

	createBudgetReq := budget.CreateBudgetRequest{
		CreateBudgetDetails: createBudgetDetails,
	}

	createBudgetResp, err := client.CreateBudget(context.Background(), createBudgetReq)
	return createBudgetResp, err
}
