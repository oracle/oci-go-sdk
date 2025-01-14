// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataSourceDetails Details specific to the data source type.
type DataSourceDetails interface {
}

type datasourcedetails struct {
	JsonData               []byte
	DataSourceFeedProvider string `json:"dataSourceFeedProvider"`
}

// UnmarshalJSON unmarshals json
func (m *datasourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatasourcedetails datasourcedetails
	s := struct {
		Model Unmarshalerdatasourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DataSourceFeedProvider = s.Model.DataSourceFeedProvider

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *datasourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DataSourceFeedProvider {
	case "SCHEDULEDQUERY":
		mm := ScheduledQueryDataSourceObjDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOGGINGQUERY":
		mm := LoggingQueryDataSourceDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DataSourceDetails: %s.", m.DataSourceFeedProvider)
		return *m, nil
	}
}

func (m datasourcedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m datasourcedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
