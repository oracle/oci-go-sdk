// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReplaceCaDetails Information about the detailed CA bundle replacement of the rover node.
type ReplaceCaDetails struct {

	// Plain text certificate chain in PEM format for the subordinate CA associated with given roverNode.
	CaBundlePem *string `mandatory:"false" json:"caBundlePem"`

	// Max validity of leaf certificates issued by the CA associated with given node, in days, in ISO 8601 format, example "P365D".
	CertificateMaxValidityDuration *string `mandatory:"false" json:"certificateMaxValidityDuration"`

	// key algorithm for issuing leaf certificate.
	CertKeyAlgorithm CertKeyAlgorithmEnum `mandatory:"false" json:"certKeyAlgorithm,omitempty"`

	// signature algorithm for issuing leaf certificate.
	CertSignatureAlgorithm CertSignatureAlgorithmEnum `mandatory:"false" json:"certSignatureAlgorithm,omitempty"`
}

func (m ReplaceCaDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReplaceCaDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCertKeyAlgorithmEnum(string(m.CertKeyAlgorithm)); !ok && m.CertKeyAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertKeyAlgorithm: %s. Supported values are: %s.", m.CertKeyAlgorithm, strings.Join(GetCertKeyAlgorithmEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCertSignatureAlgorithmEnum(string(m.CertSignatureAlgorithm)); !ok && m.CertSignatureAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertSignatureAlgorithm: %s. Supported values are: %s.", m.CertSignatureAlgorithm, strings.Join(GetCertSignatureAlgorithmEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
