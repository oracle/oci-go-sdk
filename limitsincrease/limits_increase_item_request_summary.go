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

// LimitsIncreaseItemRequestSummary A summary of properties for a limit increase item within the indicated limit increase request.
type LimitsIncreaseItemRequestSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the limit increase item.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the limit increase item.
	// Note: The tenancy is the root compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The region code for the limit increase item.
	Region *string `mandatory:"true" json:"region"`

	// The name of the service that owns the limit.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// The name of the limit to be increased.
	LimitName *string `mandatory:"true" json:"limitName"`

	// The requested value for the limit.
	Value *int64 `mandatory:"true" json:"value"`

	// The value of the limit for the tenancy at the time of the request. Purely informative.
	CurrentValue *int64 `mandatory:"true" json:"currentValue"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the limit increase request that the limit increase item belongs to.
	LimitsIncreaseRequestId *string `mandatory:"true" json:"limitsIncreaseRequestId"`

	// The time that the limit increase request was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current status of the limit increase item.
	LifecycleState LimitsIncreaseItemRequestLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The scope of the limit increase item.
	// For OCI, use the logical availability domain (AD) of the request.
	// For Multicloud, use the availability zone (AZ) of the request
	Scope *string `mandatory:"false" json:"scope"`
}

func (m LimitsIncreaseItemRequestSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LimitsIncreaseItemRequestSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLimitsIncreaseItemRequestLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLimitsIncreaseItemRequestLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
