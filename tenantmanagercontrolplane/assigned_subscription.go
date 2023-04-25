// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AssignedSubscription Assigned subscription information.
type AssignedSubscription struct {

	// OCID of the subscription.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the compartment. Always a tenancy OCID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Subscription ID.
	ClassicSubscriptionId *string `mandatory:"true" json:"classicSubscriptionId"`

	// The type of subscription, such as 'CLOUDCM', 'SAAS', 'ERP', or 'CRM'.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// Denotes if the subscription is legacy or not.
	IsClassicSubscription *bool `mandatory:"false" json:"isClassicSubscription"`

	// Region for the subscription.
	RegionAssignment *string `mandatory:"false" json:"regionAssignment"`

	// Lifecycle state of the subscription.
	LifecycleState SubscriptionLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// List of SKUs linked to the subscription.
	Skus []SubscriptionSku `mandatory:"false" json:"skus"`

	// List of subscription order OCIDs that contributed to this subscription.
	OrderIds []string `mandatory:"false" json:"orderIds"`

	// Denotes any program that is associated with the subscription.
	ProgramType *string `mandatory:"false" json:"programType"`

	// The country code for the customer associated with the subscription.
	CustomerCountryCode *string `mandatory:"false" json:"customerCountryCode"`

	// The currency code for the customer associated with the subscription.
	CloudAmountCurrency *string `mandatory:"false" json:"cloudAmountCurrency"`

	// Customer service identifier for the customer associated with the subscription.
	CsiNumber *string `mandatory:"false" json:"csiNumber"`

	// Tier for the subscription, whether it is a free promotion subscription or a paid subscription.
	SubscriptionTier *string `mandatory:"false" json:"subscriptionTier"`

	// Denotes whether or not the subscription is a government subscription.
	IsGovernmentSubscription *bool `mandatory:"false" json:"isGovernmentSubscription"`

	// List of promotions related to the subscription.
	Promotion []Promotion `mandatory:"false" json:"promotion"`

	// Purchase entitlement ID associated with the subscription.
	PurchaseEntitlementId *string `mandatory:"false" json:"purchaseEntitlementId"`

	// Subscription start time.
	StartDate *common.SDKTime `mandatory:"false" json:"startDate"`

	// Subscription end time.
	EndDate *common.SDKTime `mandatory:"false" json:"endDate"`

	// Date-time when subscription is updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Date-time when subscription is created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m AssignedSubscription) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssignedSubscription) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSubscriptionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSubscriptionLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
