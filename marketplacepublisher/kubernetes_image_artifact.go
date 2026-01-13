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

// KubernetesImageArtifact Kubernetes HelmChart Image artifact details.
type KubernetesImageArtifact struct {

	// Unique OCID identifier for the artifact.
	Id *string `mandatory:"true" json:"id"`

	// A display name for the artifact.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the artifact was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique identifier for the publisher.
	PublisherId *string `mandatory:"true" json:"publisherId"`

	// The date and time the artifact was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2022-09-15T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	HelmChart *HelmChartImageDetails `mandatory:"true" json:"helmChart"`

	// Status notes for the Artifact.
	StatusNotes *string `mandatory:"false" json:"statusNotes"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// List of container image artifact unique identifiers included in the helm chart.
	ContainerImageArtifactIds []string `mandatory:"false" json:"containerImageArtifactIds"`

	// The current status for the Artifact.
	Status ArtifactStatusEnum `mandatory:"true" json:"status"`

	// The current state for the Artifact.
	LifecycleState ArtifactLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m KubernetesImageArtifact) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m KubernetesImageArtifact) GetDisplayName() *string {
	return m.DisplayName
}

// GetStatus returns Status
func (m KubernetesImageArtifact) GetStatus() ArtifactStatusEnum {
	return m.Status
}

// GetStatusNotes returns StatusNotes
func (m KubernetesImageArtifact) GetStatusNotes() *string {
	return m.StatusNotes
}

// GetLifecycleState returns LifecycleState
func (m KubernetesImageArtifact) GetLifecycleState() ArtifactLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m KubernetesImageArtifact) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetCompartmentId returns CompartmentId
func (m KubernetesImageArtifact) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetPublisherId returns PublisherId
func (m KubernetesImageArtifact) GetPublisherId() *string {
	return m.PublisherId
}

// GetTimeUpdated returns TimeUpdated
func (m KubernetesImageArtifact) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetFreeformTags returns FreeformTags
func (m KubernetesImageArtifact) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m KubernetesImageArtifact) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m KubernetesImageArtifact) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m KubernetesImageArtifact) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KubernetesImageArtifact) ValidateEnumValue() (bool, error) {
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

// MarshalJSON marshals to json representation
func (m KubernetesImageArtifact) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeKubernetesImageArtifact KubernetesImageArtifact
	s := struct {
		DiscriminatorParam string `json:"artifactType"`
		MarshalTypeKubernetesImageArtifact
	}{
		"HELM_CHART",
		(MarshalTypeKubernetesImageArtifact)(m),
	}

	return json.Marshal(&s)
}
