// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// CaDetails Information about the detailed CA bundle content of the rover node.
type CaDetails struct {

	// Plain text certificate chain in PEM format for the subordinate CA associated with given roverNode.
	CaBundlePem *string `mandatory:"false" json:"caBundlePem"`

	// Max validity of leaf certificates issued by the CA associated with given node, in days, in ISO 8601 format, example "P365D".
	CertificateMaxValidityDuration *string `mandatory:"false" json:"certificateMaxValidityDuration"`
}

func (m CaDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CaDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
