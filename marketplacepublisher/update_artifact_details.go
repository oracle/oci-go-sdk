// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateArtifactDetails The model for an Oracle Cloud Infrastructure Marketplace artifact.
type UpdateArtifactDetails interface {

	// The unique identifier for the compartment.
	GetCompartmentId() *string

	// The display name for the artifact.
	GetDisplayName() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type updateartifactdetails struct {
	JsonData      []byte
	CompartmentId *string                           `mandatory:"false" json:"compartmentId"`
	DisplayName   *string                           `mandatory:"false" json:"displayName"`
	FreeformTags  map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags   map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	ArtifactType  string                            `json:"artifactType"`
}

// UnmarshalJSON unmarshals json
func (m *updateartifactdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateartifactdetails updateartifactdetails
	s := struct {
		Model Unmarshalerupdateartifactdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.ArtifactType = s.Model.ArtifactType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateartifactdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ArtifactType {
	case "HELM_CHART":
		mm := UpdateKubernetesImageArtifactDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CONTAINER_IMAGE":
		mm := UpdateContainerImageArtifactDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateArtifactDetails: %s.", m.ArtifactType)
		return *m, nil
	}
}

// GetCompartmentId returns CompartmentId
func (m updateartifactdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m updateartifactdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m updateartifactdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m updateartifactdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m updateartifactdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateartifactdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
