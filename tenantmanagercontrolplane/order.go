// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v47/common"
)

// Order Order Details for Console plugin display
type Order struct {

	// Immutable and unique order number holding customer subscription information
	OrderNumber *string `mandatory:"true" json:"orderNumber"`

	// Admin e-mail owning subscription.
	AdminEmail *string `mandatory:"true" json:"adminEmail"`

	// State of order.
	OrderState *string `mandatory:"true" json:"orderState"`

	// Array of subscriptions associated with the order.
	Subscriptions []SubscriptionInfo `mandatory:"true" json:"subscriptions"`

	// Order's data center region.
	DataCenterRegion *string `mandatory:"false" json:"dataCenterRegion"`
}

func (m Order) String() string {
	return common.PointerString(m)
}
