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

// EstimatePurgeDataSizeDetails Parameters used to estimate purge data size
type EstimatePurgeDataSizeDetails struct {

	// the compartment OCID under which the data will be purged
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// the time before which data will be purged
	TimeDataEnded *common.SDKTime `mandatory:"true" json:"timeDataEnded"`

	// if true, purge child compartments data
	CompartmentIdInSubtree *bool `mandatory:"false" json:"compartmentIdInSubtree"`

	// the solr data filter query, '*' means all
	PurgeQueryString *string `mandatory:"false" json:"purgeQueryString"`

	// the type of the log data to be purged
	DataType StorageDataTypeEnum `mandatory:"false" json:"dataType,omitempty"`
}

func (m EstimatePurgeDataSizeDetails) String() string {
	return common.PointerString(m)
}
