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

// DefaultTagDefinitionValidator Use this validator to clear any existing validator on the tag key definition with the UpdateTag
// operation. Using this `validatorType` is the same as not setting any value on the validator field.
// The resultant value for `validatorType` returned in the response body is `null`.
type DefaultTagDefinitionValidator struct {
}

func (m DefaultTagDefinitionValidator) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DefaultTagDefinitionValidator) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DefaultTagDefinitionValidator) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDefaultTagDefinitionValidator DefaultTagDefinitionValidator
	s := struct {
		DiscriminatorParam string `json:"validatorType"`
		MarshalTypeDefaultTagDefinitionValidator
	}{
		"DEFAULT",
		(MarshalTypeDefaultTagDefinitionValidator)(m),
	}

	return json.Marshal(&s)
}
