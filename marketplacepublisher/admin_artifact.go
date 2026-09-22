// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MarketplacePublisherService API
//
// Use the Marketplace Publisher API to manage the publishing of applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplacepublisher

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AdminArtifact Base model object for the artifacts.
type AdminArtifact interface {

	// Unique OCID identifier for the artifact.
	GetId() *string

	// A display name for the artifact.
	GetDisplayName() *string

	// The current status for the Artifact.
	GetStatus() ArtifactStatusEnum

	// The current state for the Artifact.
	GetLifecycleState() ArtifactLifecycleStateEnum

	// The date and time the artifact was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	GetTimeCreated() *common.SDKTime

	// The unique identifier for the compartment.
	GetCompartmentId() *string

	// The unique identifier for the publisher.
	GetPublisherId() *string

	// The date and time the artifact was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	GetTimeUpdated() *common.SDKTime

	// The unique legacy identifier for the listing.
	GetLegacyId() *string

	// Status notes for the Artifact.
	GetStatusNotes() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type adminartifact struct {
	JsonData       []byte
	LegacyId       *string                           `mandatory:"false" json:"legacyId"`
	StatusNotes    *string                           `mandatory:"false" json:"statusNotes"`
	FreeformTags   map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags    map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags     map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id             *string                           `mandatory:"true" json:"id"`
	DisplayName    *string                           `mandatory:"true" json:"displayName"`
	Status         ArtifactStatusEnum                `mandatory:"true" json:"status"`
	LifecycleState ArtifactLifecycleStateEnum        `mandatory:"true" json:"lifecycleState"`
	TimeCreated    *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	CompartmentId  *string                           `mandatory:"true" json:"compartmentId"`
	PublisherId    *string                           `mandatory:"true" json:"publisherId"`
	TimeUpdated    *common.SDKTime                   `mandatory:"true" json:"timeUpdated"`
	ArtifactType   string                            `json:"artifactType"`
}

// UnmarshalJSON unmarshals json
func (m *adminartifact) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleradminartifact adminartifact
	s := struct {
		Model Unmarshaleradminartifact
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.Status = s.Model.Status
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.CompartmentId = s.Model.CompartmentId
	m.PublisherId = s.Model.PublisherId
	m.TimeUpdated = s.Model.TimeUpdated
	m.LegacyId = s.Model.LegacyId
	m.StatusNotes = s.Model.StatusNotes
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.ArtifactType = s.Model.ArtifactType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *adminartifact) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ArtifactType {
	case "STACK":
		mm := AdminStackArtifact{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MACHINE_IMAGE":
		mm := AdminMachineImageArtifact{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HELM_CHART":
		mm := AdminKubernetesImageArtifact{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CONTAINER_IMAGE":
		mm := AdminContainerImageArtifact{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for AdminArtifact: %s.", m.ArtifactType)
		return *m, nil
	}
}

// GetLegacyId returns LegacyId
func (m adminartifact) GetLegacyId() *string {
	return m.LegacyId
}

// GetStatusNotes returns StatusNotes
func (m adminartifact) GetStatusNotes() *string {
	return m.StatusNotes
}

// GetFreeformTags returns FreeformTags
func (m adminartifact) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m adminartifact) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m adminartifact) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m adminartifact) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m adminartifact) GetDisplayName() *string {
	return m.DisplayName
}

// GetStatus returns Status
func (m adminartifact) GetStatus() ArtifactStatusEnum {
	return m.Status
}

// GetLifecycleState returns LifecycleState
func (m adminartifact) GetLifecycleState() ArtifactLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m adminartifact) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetCompartmentId returns CompartmentId
func (m adminartifact) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetPublisherId returns PublisherId
func (m adminartifact) GetPublisherId() *string {
	return m.PublisherId
}

// GetTimeUpdated returns TimeUpdated
func (m adminartifact) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m adminartifact) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m adminartifact) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingArtifactStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetArtifactStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArtifactLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetArtifactLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
