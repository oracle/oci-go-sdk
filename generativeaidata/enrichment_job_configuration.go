// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service NL2SQL API
//
// A description of the ReferenceService API. in progress
//

package generativeaidata

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EnrichmentJobConfiguration An EnrichmentJobConfiguration describes the database objects that will be targeted by an EnrichmentJob, as described by the job's type.
type EnrichmentJobConfiguration interface {
}

type enrichmentjobconfiguration struct {
	JsonData          []byte
	EnrichmentJobType string `json:"enrichmentJobType"`
}

// UnmarshalJSON unmarshals json
func (m *enrichmentjobconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerenrichmentjobconfiguration enrichmentjobconfiguration
	s := struct {
		Model Unmarshalerenrichmentjobconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.EnrichmentJobType = s.Model.EnrichmentJobType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *enrichmentjobconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EnrichmentJobType {
	case "DELTA_REFRESH":
		mm := DeltaRefreshEnrichmentJobConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PARTIAL_BUILD":
		mm := PartialBuildEnrichmentJobConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FULL_BUILD":
		mm := FullBuildEnrichmentJobConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for EnrichmentJobConfiguration: %s.", m.EnrichmentJobType)
		return *m, nil
	}
}

func (m enrichmentjobconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m enrichmentjobconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
