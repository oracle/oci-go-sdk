// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SourceDetailsResponse An object that represents the source of the flow defined by the connector.
// An example source is the VCNFlow logs within the NetworkLogs group.
// For more information about flows defined by connectors, see
// Overview of Connector Hub (https://docs.oracle.com/iaas/Content/connector-hub/overview.htm).
// For configuration instructions, see
// Creating a Connector (https://docs.oracle.com/iaas/Content/connector-hub/create-service-connector.htm).
type SourceDetailsResponse interface {
	GetPrivateEndpointMetadata() *PrivateEndpointMetadata
}

type sourcedetailsresponse struct {
	JsonData                []byte
	PrivateEndpointMetadata *PrivateEndpointMetadata `mandatory:"false" json:"privateEndpointMetadata"`
	Kind                    string                   `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *sourcedetailsresponse) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersourcedetailsresponse sourcedetailsresponse
	s := struct {
		Model Unmarshalersourcedetailsresponse
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PrivateEndpointMetadata = s.Model.PrivateEndpointMetadata
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sourcedetailsresponse) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "logging":
		mm := LoggingSourceDetailsResponse{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "monitoring":
		mm := MonitoringSourceDetailsResponse{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "streaming":
		mm := StreamingSourceDetailsResponse{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "plugin":
		mm := PluginSourceDetailsResponse{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SourceDetailsResponse: %s.", m.Kind)
		return *m, nil
	}
}

// GetPrivateEndpointMetadata returns PrivateEndpointMetadata
func (m sourcedetailsresponse) GetPrivateEndpointMetadata() *PrivateEndpointMetadata {
	return m.PrivateEndpointMetadata
}

func (m sourcedetailsresponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sourcedetailsresponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SourceDetailsResponseKindEnum Enum with underlying type: string
type SourceDetailsResponseKindEnum string

// Set of constants representing the allowable values for SourceDetailsResponseKindEnum
const (
	SourceDetailsResponseKindLogging    SourceDetailsResponseKindEnum = "logging"
	SourceDetailsResponseKindMonitoring SourceDetailsResponseKindEnum = "monitoring"
	SourceDetailsResponseKindStreaming  SourceDetailsResponseKindEnum = "streaming"
	SourceDetailsResponseKindPlugin     SourceDetailsResponseKindEnum = "plugin"
)

var mappingSourceDetailsResponseKindEnum = map[string]SourceDetailsResponseKindEnum{
	"logging":    SourceDetailsResponseKindLogging,
	"monitoring": SourceDetailsResponseKindMonitoring,
	"streaming":  SourceDetailsResponseKindStreaming,
	"plugin":     SourceDetailsResponseKindPlugin,
}

var mappingSourceDetailsResponseKindEnumLowerCase = map[string]SourceDetailsResponseKindEnum{
	"logging":    SourceDetailsResponseKindLogging,
	"monitoring": SourceDetailsResponseKindMonitoring,
	"streaming":  SourceDetailsResponseKindStreaming,
	"plugin":     SourceDetailsResponseKindPlugin,
}

// GetSourceDetailsResponseKindEnumValues Enumerates the set of values for SourceDetailsResponseKindEnum
func GetSourceDetailsResponseKindEnumValues() []SourceDetailsResponseKindEnum {
	values := make([]SourceDetailsResponseKindEnum, 0)
	for _, v := range mappingSourceDetailsResponseKindEnum {
		values = append(values, v)
	}
	return values
}

// GetSourceDetailsResponseKindEnumStringValues Enumerates the set of values in String for SourceDetailsResponseKindEnum
func GetSourceDetailsResponseKindEnumStringValues() []string {
	return []string{
		"logging",
		"monitoring",
		"streaming",
		"plugin",
	}
}

// GetMappingSourceDetailsResponseKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSourceDetailsResponseKindEnum(val string) (SourceDetailsResponseKindEnum, bool) {
	enum, ok := mappingSourceDetailsResponseKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
