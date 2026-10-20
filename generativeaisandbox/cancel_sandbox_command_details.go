// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GenAI Sandbox API
//
// API for managing sandbox resources and exchanging sandbox access tokens.
//

package generativeaisandbox

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CancelSandboxCommandDetails Request body used to cancel a sandbox command.
type CancelSandboxCommandDetails struct {

	// Signal to send to the running command.
	Signal CancelSandboxCommandDetailsSignalEnum `mandatory:"false" json:"signal,omitempty"`
}

func (m CancelSandboxCommandDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CancelSandboxCommandDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCancelSandboxCommandDetailsSignalEnum(string(m.Signal)); !ok && m.Signal != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Signal: %s. Supported values are: %s.", m.Signal, strings.Join(GetCancelSandboxCommandDetailsSignalEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CancelSandboxCommandDetailsSignalEnum Enum with underlying type: string
type CancelSandboxCommandDetailsSignalEnum string

// Set of constants representing the allowable values for CancelSandboxCommandDetailsSignalEnum
const (
	CancelSandboxCommandDetailsSignalSigterm CancelSandboxCommandDetailsSignalEnum = "SIGTERM"
	CancelSandboxCommandDetailsSignalSigkill CancelSandboxCommandDetailsSignalEnum = "SIGKILL"
)

var mappingCancelSandboxCommandDetailsSignalEnum = map[string]CancelSandboxCommandDetailsSignalEnum{
	"SIGTERM": CancelSandboxCommandDetailsSignalSigterm,
	"SIGKILL": CancelSandboxCommandDetailsSignalSigkill,
}

var mappingCancelSandboxCommandDetailsSignalEnumLowerCase = map[string]CancelSandboxCommandDetailsSignalEnum{
	"sigterm": CancelSandboxCommandDetailsSignalSigterm,
	"sigkill": CancelSandboxCommandDetailsSignalSigkill,
}

// GetCancelSandboxCommandDetailsSignalEnumValues Enumerates the set of values for CancelSandboxCommandDetailsSignalEnum
func GetCancelSandboxCommandDetailsSignalEnumValues() []CancelSandboxCommandDetailsSignalEnum {
	values := make([]CancelSandboxCommandDetailsSignalEnum, 0)
	for _, v := range mappingCancelSandboxCommandDetailsSignalEnum {
		values = append(values, v)
	}
	return values
}

// GetCancelSandboxCommandDetailsSignalEnumStringValues Enumerates the set of values in String for CancelSandboxCommandDetailsSignalEnum
func GetCancelSandboxCommandDetailsSignalEnumStringValues() []string {
	return []string{
		"SIGTERM",
		"SIGKILL",
	}
}

// GetMappingCancelSandboxCommandDetailsSignalEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCancelSandboxCommandDetailsSignalEnum(val string) (CancelSandboxCommandDetailsSignalEnum, bool) {
	enum, ok := mappingCancelSandboxCommandDetailsSignalEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
