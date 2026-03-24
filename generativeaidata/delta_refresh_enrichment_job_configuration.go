// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Service Nl2sql API
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

// DeltaRefreshEnrichmentJobConfiguration A DeltaRefreshEnrichmentJobConfiguration is an EnrichmentJobConfiguration [indicated by the first item under `allOf`, which is a reference to EnrichmentJobConfiguration]
// that describes the database schema that will be the scope of the enrichment job and the schedule on which the job will run. As distinguished by enrichmentJobType [with specific characteristics defined by the second item under `allOf`].
type DeltaRefreshEnrichmentJobConfiguration struct {

	// Name of the DB Schema to be enriched
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// Schedule for occurrences of Delta Refresh jobs
	DeltaRefreshSchedule *interface{} `mandatory:"true" json:"deltaRefreshSchedule"`
}

func (m DeltaRefreshEnrichmentJobConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeltaRefreshEnrichmentJobConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DeltaRefreshEnrichmentJobConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDeltaRefreshEnrichmentJobConfiguration DeltaRefreshEnrichmentJobConfiguration
	s := struct {
		DiscriminatorParam string `json:"enrichmentJobType"`
		MarshalTypeDeltaRefreshEnrichmentJobConfiguration
	}{
		"DELTA_REFRESH",
		(MarshalTypeDeltaRefreshEnrichmentJobConfiguration)(m),
	}

	return json.Marshal(&s)
}
