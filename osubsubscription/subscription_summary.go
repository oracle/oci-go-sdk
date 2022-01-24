// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription API Subscription, Commitment and and Rate Card Details
//
// Set of APIs that return the Subscription Details, Commitment and Effective Rate Card Details
//

package osubsubscription

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// SubscriptionSummary Subscription summary
type SubscriptionSummary struct {

	// Status of the plan
	Status *string `mandatory:"false" json:"status"`

	// Represents the date when the first service of the subscription was activated
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Represents the date when the last service of the subscription ends
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	Currency *Currency `mandatory:"false" json:"currency"`

	// Customer friendly service name provided by PRG
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// List of Subscribed Services of the plan
	SubscribedServices []SubscribedServiceSummary `mandatory:"false" json:"subscribedServices"`
}

func (m SubscriptionSummary) String() string {
	return common.PointerString(m)
}
