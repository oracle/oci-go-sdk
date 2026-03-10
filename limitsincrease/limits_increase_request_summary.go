// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Limits Increase API
//
// Use the Limits Increase API to work with limit increase requests.
// For more information, see
// Working with Limit Increase Requests (https://docs.oracle.com/iaas/Content/General/service-limits/requests.htm).
//

package limitsincrease

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LimitsIncreaseRequestSummary A summary of properties for a limit increase request.
type LimitsIncreaseRequestSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the limit increase request.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the limit increase request.
	// Note: The tenancy is the root compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the limit increase request.
	LifecycleState LimitsIncreaseRequestLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The user-friendly name of the limit increase request. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Number of limit increase items in the limit increase request.
	ItemsInRequest *int64 `mandatory:"true" json:"itemsInRequest"`

	// The time that the limit increase request was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy subscription for the limit increase request.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// Reason for the limit increase request.
	Justification *string `mandatory:"false" json:"justification"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m LimitsIncreaseRequestSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LimitsIncreaseRequestSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLimitsIncreaseRequestLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLimitsIncreaseRequestLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
