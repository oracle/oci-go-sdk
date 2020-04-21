// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Search Service API
//
// Search for resources in your cloud network.
//

package resourcesearch

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// SearchDetails A base request type that contains common criteria for searching for resources.
type SearchDetails interface {

	// The type of matching context returned in the response. If you specify `HIGHLIGHTS`, then the service will highlight fragments in its response. (For more information, see ResourceSummary.searchContext and SearchContext.) The default setting is `NONE`.
	GetMatchingContextType() SearchDetailsMatchingContextTypeEnum
}

type searchdetails struct {
	JsonData            []byte
	MatchingContextType SearchDetailsMatchingContextTypeEnum `mandatory:"false" json:"matchingContextType,omitempty"`
	Type                string                               `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *searchdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersearchdetails searchdetails
	s := struct {
		Model Unmarshalersearchdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.MatchingContextType = s.Model.MatchingContextType
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *searchdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "Structured":
		mm := StructuredSearchDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FreeText":
		mm := FreeTextSearchDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetMatchingContextType returns MatchingContextType
func (m searchdetails) GetMatchingContextType() SearchDetailsMatchingContextTypeEnum {
	return m.MatchingContextType
}

func (m searchdetails) String() string {
	return common.PointerString(m)
}

// SearchDetailsMatchingContextTypeEnum Enum with underlying type: string
type SearchDetailsMatchingContextTypeEnum string

// Set of constants representing the allowable values for SearchDetailsMatchingContextTypeEnum
const (
	SearchDetailsMatchingContextTypeNone       SearchDetailsMatchingContextTypeEnum = "NONE"
	SearchDetailsMatchingContextTypeHighlights SearchDetailsMatchingContextTypeEnum = "HIGHLIGHTS"
)

var mappingSearchDetailsMatchingContextType = map[string]SearchDetailsMatchingContextTypeEnum{
	"NONE":       SearchDetailsMatchingContextTypeNone,
	"HIGHLIGHTS": SearchDetailsMatchingContextTypeHighlights,
}

// GetSearchDetailsMatchingContextTypeEnumValues Enumerates the set of values for SearchDetailsMatchingContextTypeEnum
func GetSearchDetailsMatchingContextTypeEnumValues() []SearchDetailsMatchingContextTypeEnum {
	values := make([]SearchDetailsMatchingContextTypeEnum, 0)
	for _, v := range mappingSearchDetailsMatchingContextType {
		values = append(values, v)
	}
	return values
}
