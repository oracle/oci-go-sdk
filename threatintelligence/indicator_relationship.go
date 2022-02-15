// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Threat Intelligence API
//
// Use the Threat Intelligence API to view indicators of compromise and related items. For more information, see Overview of Threat Intelligence (https://docs.cloud.oracle.com/Content/ThreatIntelligence/Concepts/threatintelligenceoverview.htm).
//

package threatintelligence

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// IndicatorRelationship An relationship name and list of releated entities.
type IndicatorRelationship struct {

	// The name of the attribute
	Name *string `mandatory:"true" json:"name"`

	RelatedEntity EntityReference `mandatory:"true" json:"relatedEntity"`

	// The array of attribution data that support this SourcedRelationship
	Attribution []DataAttribution `mandatory:"true" json:"attribution"`
}

func (m IndicatorRelationship) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IndicatorRelationship) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *IndicatorRelationship) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name          *string           `json:"name"`
		RelatedEntity entityreference   `json:"relatedEntity"`
		Attribution   []DataAttribution `json:"attribution"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Name = model.Name

	nn, e = model.RelatedEntity.UnmarshalPolymorphicJSON(model.RelatedEntity.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RelatedEntity = nn.(EntityReference)
	} else {
		m.RelatedEntity = nil
	}

	m.Attribution = make([]DataAttribution, len(model.Attribution))
	for i, n := range model.Attribution {
		m.Attribution[i] = n
	}

	return
}
