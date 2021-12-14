// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v54/common"
)

// SubscriptionInfo A single subscription's details.
type SubscriptionInfo struct {

	// Subscription ID.
	SpmSubscriptionId *string `mandatory:"true" json:"spmSubscriptionId"`

	// Service name for subscription.
	Service *string `mandatory:"true" json:"service"`

	// Subscription start date. An RFC 3339-formatted date and time string.
	StartDate *common.SDKTime `mandatory:"true" json:"startDate"`

	// Subscription end date. An RFC 3339-formatted date and time string.
	EndDate *common.SDKTime `mandatory:"true" json:"endDate"`

	// List of SKUs the subscription contains.
	Skus []Sku `mandatory:"true" json:"skus"`
}

func (m SubscriptionInfo) String() string {
	return common.PointerString(m)
}
