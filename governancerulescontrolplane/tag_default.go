// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GovernanceRulesControlPlane API
//
// A description of the GovernanceRulesControlPlane API
//

package governancerulescontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TagDefault Tag defaults let you specify a default tag (tagnamespace.tag="value") to apply to all resource types
// in a specified compartment. The tag default is applied at the time the resource is created. Resources
// that exist in the compartment before you create the tag default are not tagged.
type TagDefault struct {

	// The name of the tag. The tag default will always assign a default value for this tag name.
	TagName *string `mandatory:"true" json:"tagName"`

	// The default value for the tag name. This will be applied to all new resources created in the compartment.
	Value *string `mandatory:"true" json:"value"`

	// If you specify that a value is required, a value is set during resource creation (either by
	// the user creating the resource or another tag default). If no value is set, resource
	// creation is blocked.
	// * If the `isRequired` flag is set to "true", the value is set during resource creation.
	// * If the `isRequired` flag is set to "false", the value you enter is set during resource creation.
	// Example: `false`
	IsRequired *bool `mandatory:"true" json:"isRequired"`
}

func (m TagDefault) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TagDefault) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
