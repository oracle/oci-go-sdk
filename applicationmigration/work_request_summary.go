// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestSummary A summary of the status of a work request.
type WorkRequestSummary struct {

	// The asynchronous operation tracked by this work request.
	OperationType OperationTypesEnum `mandatory:"true" json:"operationType"`

	// The status of the work request.
	Status OperationStatusEnum `mandatory:"true" json:"status"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing
	// this work request.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The percentage complete of the operation tracked by this work request.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time the work request was created, in the format defined by RFC3339.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The date and time the work request transitioned from `ACCEPTED` to `IN_PROGRESS`, in
	// the format defined by RFC3339.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the work request reached a terminal state, either `FAILED` or `SUCCEEDED`.
	// Format is defined by RFC3339.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequestSummary) String() string {
	return common.PointerString(m)
}
