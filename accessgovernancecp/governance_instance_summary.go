// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// GovernanceInstanceSummary The summary of an GovernanceInstance.
type GovernanceInstanceSummary struct {

	// The name for the GovernanceInstance.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment where the GovernanceInstance resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the the GovernanceInstance was created in an RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the GovernanceInstance.
	LifecycleState InstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The unique OCID of the GovernanceInstance.
	Id *string `mandatory:"false" json:"id"`

	// The description of the GovernanceInstance.
	Description *string `mandatory:"false" json:"description"`

	// The time the GovernanceInstance was updated in an RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The licenseType being used.
	LicenseType LicenseTypeEnum `mandatory:"false" json:"licenseType,omitempty"`

	// The access URL of the GovernanceInstance.
	InstanceUrl *string `mandatory:"false" json:"instanceUrl"`

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

func (m GovernanceInstanceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GovernanceInstanceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInstanceLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLicenseTypeEnum(string(m.LicenseType)); !ok && m.LicenseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseType: %s. Supported values are: %s.", m.LicenseType, strings.Join(GetLicenseTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
