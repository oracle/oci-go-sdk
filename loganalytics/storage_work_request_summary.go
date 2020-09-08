// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/common"
)

// StorageWorkRequestSummary Storage work request summary for list operation.
type StorageWorkRequestSummary struct {

	// Unique OCID identifier to reference this storage work Request with.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Work request status.
	Status WorkRequestStatusEnum `mandatory:"true" json:"status"`

	// the end of the time interval
	TimeDataEnded *common.SDKTime `mandatory:"true" json:"timeDataEnded"`

	// the type of the log data to be purged
	DataType StorageDataTypeEnum `mandatory:"true" json:"dataType"`

	// Asynchronous storage request name.
	OperationType StorageOperationTypeEnum `mandatory:"true" json:"operationType"`

	// When the work request started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// When the work request was accepted. Should match timeStarted in all cases.
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// When the work request finished execution.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// When the work request will expire.
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`

	// Percentage progress completion of the work request.
	PercentComplete *int `mandatory:"false" json:"percentComplete"`

	// the start of the time interval
	TimeDataStarted *common.SDKTime `mandatory:"false" json:"timeDataStarted"`

	// the solr data filter query, '*' means all
	PurgeQueryString *string `mandatory:"false" json:"purgeQueryString"`

	// more detailed status if applicable
	StatusDetails *string `mandatory:"false" json:"statusDetails"`

	// more detailed info about this operation if applicable
	OperationDetails *string `mandatory:"false" json:"operationDetails"`

	// policy name if applicable (e.g. purge policy)
	PolicyName *string `mandatory:"false" json:"policyName"`

	// purge policy ID
	PolicyId *string `mandatory:"false" json:"policyId"`

	// storage usage in bytes if applicable
	StorageUsageInBytes *int64 `mandatory:"false" json:"storageUsageInBytes"`

	// if true, purge child compartments data, only applicable to purge request
	CompartmentIdInSubtree *bool `mandatory:"false" json:"compartmentIdInSubtree"`
}

func (m StorageWorkRequestSummary) String() string {
	return common.PointerString(m)
}
