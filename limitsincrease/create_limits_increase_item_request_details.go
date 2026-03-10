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

// CreateLimitsIncreaseItemRequestDetails The configuration details for creating a limit increase item within a limit increase request.
// For more information, see
// Creating a Limit Increase Request (https://docs.oracle.com/iaas/Content/General/service-limits/create-requests.htm).
type CreateLimitsIncreaseItemRequestDetails struct {

	// The name of the service that owns the limit.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// The name of the limit to be increased.
	LimitName *string `mandatory:"true" json:"limitName"`

	// The region for the limit increase.
	Region *string `mandatory:"true" json:"region"`

	// The requested value of the increase.
	Value *int64 `mandatory:"true" json:"value"`

	// The scope of the limit increase item.
	// For OCI, use the logical availability domain (AD) of the request.
	// For Multicloud, use the availability zone (AZ) of the request
	Scope *string `mandatory:"false" json:"scope"`

	// List of questionnaire responses.
	QuestionnaireResponse []LimitsIncreaseItemQuestionRequest `mandatory:"false" json:"questionnaireResponse"`
}

func (m CreateLimitsIncreaseItemRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateLimitsIncreaseItemRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
