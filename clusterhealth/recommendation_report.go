// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ITV Control Plane - Diagnosis Store API
//
// Use the ITV Control Plane Diagnosis Store API to manage diagnosis.
//

package clusterhealth

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RecommendationReport A recommendation report
type RecommendationReport struct {

	// summary
	Summary *string `mandatory:"true" json:"summary"`

	// total issue count
	TotalIssues *int `mandatory:"true" json:"totalIssues"`

	// critical issue count
	CriticalIssues *int `mandatory:"true" json:"criticalIssues"`

	// warning issue count
	WarningIssues *int `mandatory:"true" json:"warningIssues"`

	// info issue count
	InfoIssues *int `mandatory:"true" json:"infoIssues"`

	// recommendations
	Recommendations []Recommendation `mandatory:"true" json:"recommendations"`

	// time generated
	TimeGenerated *common.SDKTime `mandatory:"true" json:"timeGenerated"`

	Metadata *Metadata `mandatory:"false" json:"metadata"`
}

func (m RecommendationReport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecommendationReport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
