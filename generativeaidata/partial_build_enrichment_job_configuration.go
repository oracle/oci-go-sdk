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

// PartialBuildEnrichmentJobConfiguration A PartialBuildEnrichmentJobConfiguration is an EnrichmentJobConfiguration [indicated by the first item under `allOf`, which is a reference to EnrichmentJobConfiguration]
// that describes the database schema that will be the scope of the enrichment job. As distinguished by enrichmentJobType [with specific characteristics defined by the second item under `allOf`].
type PartialBuildEnrichmentJobConfiguration struct {

	// Name of the DB Schema to be enriched
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// Collection of the DatabaseObjects to be enriched for the given schema.
	DatabaseObjects []DatabaseObject `mandatory:"true" json:"databaseObjects"`
}

func (m PartialBuildEnrichmentJobConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PartialBuildEnrichmentJobConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PartialBuildEnrichmentJobConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePartialBuildEnrichmentJobConfiguration PartialBuildEnrichmentJobConfiguration
	s := struct {
		DiscriminatorParam string `json:"enrichmentJobType"`
		MarshalTypePartialBuildEnrichmentJobConfiguration
	}{
		"PARTIAL_BUILD",
		(MarshalTypePartialBuildEnrichmentJobConfiguration)(m),
	}

	return json.Marshal(&s)
}
