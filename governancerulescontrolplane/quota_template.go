// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GovernanceRulesControlPlane API
//
// A description of the GovernanceRulesControlPlane API
//

package governancerulescontrolplane

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// QuotaTemplate Quota template for governance rule.
type QuotaTemplate struct {

	// Display name of the quota resource.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// List of quota statements.
	Statements []string `mandatory:"true" json:"statements"`

	// Description of the quota resource.
	Description *string `mandatory:"false" json:"description"`
}

func (m QuotaTemplate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QuotaTemplate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m QuotaTemplate) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeQuotaTemplate QuotaTemplate
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeQuotaTemplate
	}{
		"QUOTA",
		(MarshalTypeQuotaTemplate)(m),
	}

	return json.Marshal(&s)
}
