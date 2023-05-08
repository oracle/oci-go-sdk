// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Access Governance APIs
//
// Use the Oracle Access Governance API to create, view, and manage GovernanceInstances.
//

package accessgovernancecp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateGovernanceInstanceDetails The details about a new GovernanceInstance.
type CreateGovernanceInstanceDetails struct {

	// The name for the GovernanceInstance.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The licenseType being used.
	LicenseType LicenseTypeEnum `mandatory:"true" json:"licenseType"`

	// The namespace for tenancy object storage.
	TenancyNamespace *string `mandatory:"true" json:"tenancyNamespace"`

	// The OCID of the compartment where the GovernanceInstance resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// IDCS access token identifying a stripe and service administrator user.
	IdcsAccessToken *string `mandatory:"true" json:"idcsAccessToken"`

	// The description of the GovernanceInstance.
	Description *string `mandatory:"false" json:"description"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CreateGovernanceInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateGovernanceInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLicenseTypeEnum(string(m.LicenseType)); !ok && m.LicenseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseType: %s. Supported values are: %s.", m.LicenseType, strings.Join(GetLicenseTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
