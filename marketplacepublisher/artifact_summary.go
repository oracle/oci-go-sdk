// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MarketplacePublisherService API
//
// Use the Marketplace Publisher API to manage the publishing of applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplacepublisher

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ArtifactSummary The model for the artifact summary.
type ArtifactSummary struct {

	// Unique identifier for the artifact.
	Id *string `mandatory:"true" json:"id"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name for the artifact.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Artifact Type for the artifact.
	ArtifactType ArtifactTypeEnumEnum `mandatory:"true" json:"artifactType"`

	// The current state for the Artifact.
	LifecycleState ArtifactLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The current status for the Artifact.
	Status ArtifactStatusEnum `mandatory:"true" json:"status"`

	// The date and time the artifact was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2023-03-27T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the artifact was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2023-03-27T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ArtifactSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ArtifactSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingArtifactTypeEnumEnum(string(m.ArtifactType)); !ok && m.ArtifactType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ArtifactType: %s. Supported values are: %s.", m.ArtifactType, strings.Join(GetArtifactTypeEnumEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArtifactLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetArtifactLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArtifactStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetArtifactStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
