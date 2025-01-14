// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// Artifact Base model object for the artifacts.
type Artifact interface {

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

type artifact struct {
	JsonData       []byte
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
func (m *artifact) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerartifact artifact
	s := struct {
		Model Unmarshalerartifact
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
	m.StatusNotes = s.Model.StatusNotes
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.ArtifactType = s.Model.ArtifactType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *artifact) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ArtifactType {
	case "CONTAINER_IMAGE":
		mm := ContainerImageArtifact{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HELM_CHART":
		mm := KubernetesImageArtifact{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Artifact: %s.", m.ArtifactType)
		return *m, nil
	}
}

// GetStatusNotes returns StatusNotes
func (m artifact) GetStatusNotes() *string {
	return m.StatusNotes
}

// GetFreeformTags returns FreeformTags
func (m artifact) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m artifact) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m artifact) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m artifact) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m artifact) GetDisplayName() *string {
	return m.DisplayName
}

// GetStatus returns Status
func (m artifact) GetStatus() ArtifactStatusEnum {
	return m.Status
}

// GetLifecycleState returns LifecycleState
func (m artifact) GetLifecycleState() ArtifactLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m artifact) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetCompartmentId returns CompartmentId
func (m artifact) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetPublisherId returns PublisherId
func (m artifact) GetPublisherId() *string {
	return m.PublisherId
}

// GetTimeUpdated returns TimeUpdated
func (m artifact) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m artifact) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m artifact) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingArtifactStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetArtifactStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingArtifactLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetArtifactLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ArtifactStatusEnum Enum with underlying type: string
type ArtifactStatusEnum string

// Set of constants representing the allowable values for ArtifactStatusEnum
const (
	ArtifactStatusInProgress  ArtifactStatusEnum = "IN_PROGRESS"
	ArtifactStatusAvailable   ArtifactStatusEnum = "AVAILABLE"
	ArtifactStatusUnavailable ArtifactStatusEnum = "UNAVAILABLE"
)

var mappingArtifactStatusEnum = map[string]ArtifactStatusEnum{
	"IN_PROGRESS": ArtifactStatusInProgress,
	"AVAILABLE":   ArtifactStatusAvailable,
	"UNAVAILABLE": ArtifactStatusUnavailable,
}

var mappingArtifactStatusEnumLowerCase = map[string]ArtifactStatusEnum{
	"in_progress": ArtifactStatusInProgress,
	"available":   ArtifactStatusAvailable,
	"unavailable": ArtifactStatusUnavailable,
}

// GetArtifactStatusEnumValues Enumerates the set of values for ArtifactStatusEnum
func GetArtifactStatusEnumValues() []ArtifactStatusEnum {
	values := make([]ArtifactStatusEnum, 0)
	for _, v := range mappingArtifactStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetArtifactStatusEnumStringValues Enumerates the set of values in String for ArtifactStatusEnum
func GetArtifactStatusEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"AVAILABLE",
		"UNAVAILABLE",
	}
}

// GetMappingArtifactStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArtifactStatusEnum(val string) (ArtifactStatusEnum, bool) {
	enum, ok := mappingArtifactStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ArtifactLifecycleStateEnum Enum with underlying type: string
type ArtifactLifecycleStateEnum string

// Set of constants representing the allowable values for ArtifactLifecycleStateEnum
const (
	ArtifactLifecycleStateCreating ArtifactLifecycleStateEnum = "CREATING"
	ArtifactLifecycleStateUpdating ArtifactLifecycleStateEnum = "UPDATING"
	ArtifactLifecycleStateActive   ArtifactLifecycleStateEnum = "ACTIVE"
	ArtifactLifecycleStateDeleting ArtifactLifecycleStateEnum = "DELETING"
	ArtifactLifecycleStateDeleted  ArtifactLifecycleStateEnum = "DELETED"
	ArtifactLifecycleStateFailed   ArtifactLifecycleStateEnum = "FAILED"
)

var mappingArtifactLifecycleStateEnum = map[string]ArtifactLifecycleStateEnum{
	"CREATING": ArtifactLifecycleStateCreating,
	"UPDATING": ArtifactLifecycleStateUpdating,
	"ACTIVE":   ArtifactLifecycleStateActive,
	"DELETING": ArtifactLifecycleStateDeleting,
	"DELETED":  ArtifactLifecycleStateDeleted,
	"FAILED":   ArtifactLifecycleStateFailed,
}

var mappingArtifactLifecycleStateEnumLowerCase = map[string]ArtifactLifecycleStateEnum{
	"creating": ArtifactLifecycleStateCreating,
	"updating": ArtifactLifecycleStateUpdating,
	"active":   ArtifactLifecycleStateActive,
	"deleting": ArtifactLifecycleStateDeleting,
	"deleted":  ArtifactLifecycleStateDeleted,
	"failed":   ArtifactLifecycleStateFailed,
}

// GetArtifactLifecycleStateEnumValues Enumerates the set of values for ArtifactLifecycleStateEnum
func GetArtifactLifecycleStateEnumValues() []ArtifactLifecycleStateEnum {
	values := make([]ArtifactLifecycleStateEnum, 0)
	for _, v := range mappingArtifactLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetArtifactLifecycleStateEnumStringValues Enumerates the set of values in String for ArtifactLifecycleStateEnum
func GetArtifactLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingArtifactLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingArtifactLifecycleStateEnum(val string) (ArtifactLifecycleStateEnum, bool) {
	enum, ok := mappingArtifactLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
