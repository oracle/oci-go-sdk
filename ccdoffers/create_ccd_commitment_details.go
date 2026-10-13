// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Continuous Consumption Discounts API
//
// Use the Continuous Consumption Discounts API to manage continuous consumption discount operations.
//

package ccdoffers

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateCcdCommitmentDetails CCD commitment creation details.
type CreateCcdCommitmentDetails struct {

	// The hourly quantity that the customer is committing to a CCD offer's main SKU.
	// For FLEX offers, this quantity must equal the CCD offer's minQuantity and maxQuantity when both are provided.
	// For example, if the quantity is 1000, then you are committing to a 1000 usage quantity value per hour.
	Quantity *int64 `mandatory:"true" json:"quantity"`

	// The CCD commitment name. Ensure the name is unique across subscription IDs. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The subscription OCID value the CCD offer is applicable to.
	SubscriptionId *string `mandatory:"true" json:"subscriptionId"`

	// The compartment OCID where you want to create the CCD commitment.
	// Must be the root compartment OCID of a parent tenancy (if part of an organization), or a standalone tenancy to which the subscription is associated.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The inclusive date and time when the CCD commitment will be effective.
	// The date and time must be rounded up to the exact hour, and must be in the future. The format is defined by RFC 3339.
	// For example, if you want to commit to a CCD on 2025-06-01 12:00:00, you should create a request any time before
	// 2025-06-01 12:00:00. You can only create CCD for a Scheduled or Current CCD offer.
	TimeCcdCommitmentStarts *common.SDKTime `mandatory:"true" json:"timeCcdCommitmentStarts"`

	// The tenancy OCID you want to apply filler usage to.
	FillerTenancyId *string `mandatory:"false" json:"fillerTenancyId"`

	UsageEligibility *UsageEligibility `mandatory:"false" json:"usageEligibility"`

	// The bundled SKUs selected for this commitment. Supply `quantity` only for components whose offer quantity rule is `CUSTOMER_SELECTED`.
	SecondarySkus []SecondarySkus `mandatory:"false" json:"secondarySkus"`

	// The exclusive time the CCD commitment ends. The time gets computed using `timeCcdCommitmentStarts` from the `ccdCommitment`,
	// and the minimum commitment term associated with `CcdOffer`. The default value is set as null at the time of CCD commitment creation.
	TimeCcdCommitmentEnds *common.SDKTime `mandatory:"false" json:"timeCcdCommitmentEnds"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateCcdCommitmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCcdCommitmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
