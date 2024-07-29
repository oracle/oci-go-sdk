// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Speech API
//
// The OCI Speech Service harnesses the power of spoken language by allowing developers to easily convert file-based data containing human speech into highly accurate text transcriptions.
//

package aispeech

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RealtimeMessageAuthenticationCredentials The payload for credential-based authentication.
type RealtimeMessageAuthenticationCredentials struct {

	// Compartment ID to be used for authentication/authorization.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The signed request header object.
	Headers map[string]string `mandatory:"true" json:"headers"`
}

// GetCompartmentId returns CompartmentId
func (m RealtimeMessageAuthenticationCredentials) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m RealtimeMessageAuthenticationCredentials) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RealtimeMessageAuthenticationCredentials) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RealtimeMessageAuthenticationCredentials) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRealtimeMessageAuthenticationCredentials RealtimeMessageAuthenticationCredentials
	s := struct {
		DiscriminatorParam string `json:"authenticationType"`
		MarshalTypeRealtimeMessageAuthenticationCredentials
	}{
		"CREDENTIALS",
		(MarshalTypeRealtimeMessageAuthenticationCredentials)(m),
	}

	return json.Marshal(&s)
}
