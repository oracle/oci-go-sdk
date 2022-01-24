// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// AppMgmt Control API
//
// AppMgmt Control API
//

package appmgmtcontrol

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// MonitoredInstanceCollection Results of a monitored instance search. Contains MonitoredInstanceSummary items.
type MonitoredInstanceCollection struct {

	// List of monitored instances.
	Items []MonitoredInstanceSummary `mandatory:"true" json:"items"`
}

func (m MonitoredInstanceCollection) String() string {
	return common.PointerString(m)
}
