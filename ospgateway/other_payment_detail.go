// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Billing Center Gateway API
//
// This site describes all the Rest endpoints of Billing Center Gateway.
//

package ospgateway

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v54/common"
)

// OtherPaymentDetail Other Payment related details
type OtherPaymentDetail struct {

	// Paid the invoice on this day
	TimePaidOn *common.SDKTime `mandatory:"false" json:"timePaidOn"`

	// example
	PaidBy *string `mandatory:"false" json:"paidBy"`

	// Amount that paid
	AmountPaid *float32 `mandatory:"false" json:"amountPaid"`
}

//GetTimePaidOn returns TimePaidOn
func (m OtherPaymentDetail) GetTimePaidOn() *common.SDKTime {
	return m.TimePaidOn
}

//GetPaidBy returns PaidBy
func (m OtherPaymentDetail) GetPaidBy() *string {
	return m.PaidBy
}

//GetAmountPaid returns AmountPaid
func (m OtherPaymentDetail) GetAmountPaid() *float32 {
	return m.AmountPaid
}

func (m OtherPaymentDetail) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OtherPaymentDetail) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOtherPaymentDetail OtherPaymentDetail
	s := struct {
		DiscriminatorParam string `json:"paymentMethod"`
		MarshalTypeOtherPaymentDetail
	}{
		"OTHER",
		(MarshalTypeOtherPaymentDetail)(m),
	}

	return json.Marshal(&s)
}
