// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VerticalScalingScheduleDetails Details of a vertical scaling schedule.
type VerticalScalingScheduleDetails interface {
}

type verticalscalingscheduledetails struct {
	JsonData     []byte
	ScheduleType string `json:"scheduleType"`
}

// UnmarshalJSON unmarshals json
func (m *verticalscalingscheduledetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerverticalscalingscheduledetails verticalscalingscheduledetails
	s := struct {
		Model Unmarshalerverticalscalingscheduledetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ScheduleType = s.Model.ScheduleType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *verticalscalingscheduledetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ScheduleType {
	case "DAY_BASED":
		mm := DayBasedVerticalScalingScheduleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for VerticalScalingScheduleDetails: %s.", m.ScheduleType)
		return *m, nil
	}
}

func (m verticalscalingscheduledetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m verticalscalingscheduledetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
