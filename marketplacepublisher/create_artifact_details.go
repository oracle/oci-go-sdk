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

// CreateArtifactDetails Common Details to create Marketplace Publisher artifact.
type CreateArtifactDetails interface {

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

type createartifactdetails struct {
	JsonData      []byte
	DisplayName   *string                           `mandatory:"false" json:"displayName"`
	FreeformTags  map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags   map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	CompartmentId *string                           `mandatory:"true" json:"compartmentId"`
	ArtifactType  string                            `json:"artifactType"`
}

// UnmarshalJSON unmarshals json
func (m *createartifactdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateartifactdetails createartifactdetails
	s := struct {
		Model Unmarshalercreateartifactdetails
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
func (m *createartifactdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ArtifactType {
	case "HELM_CHART":
		mm := CreateKubernetesImageArtifactDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STACK":
		mm := CreateStackArtifactDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CONTAINER_IMAGE":
		mm := CreateContainerImageArtifactDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MACHINE_IMAGE":
		mm := CreateMachineImageArtifactDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateArtifactDetails: %s.", m.ArtifactType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m createartifactdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m createartifactdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createartifactdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetCompartmentId returns CompartmentId
func (m createartifactdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m createartifactdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createartifactdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
