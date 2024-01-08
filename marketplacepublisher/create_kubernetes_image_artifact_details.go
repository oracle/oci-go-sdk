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

// CreateKubernetesImageArtifactDetails Details to create a new helm chart image artifact.
type CreateKubernetesImageArtifactDetails struct {

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	HelmChart *CreateHelmChartImageDetails `mandatory:"true" json:"helmChart"`

	// The display name for the artifact.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// List of container image artifact uniquie identifiers included in the helm chart.
	ContainerImageArtifactIds []string `mandatory:"false" json:"containerImageArtifactIds"`
}

// GetCompartmentId returns CompartmentId
func (m CreateKubernetesImageArtifactDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m CreateKubernetesImageArtifactDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetFreeformTags returns FreeformTags
func (m CreateKubernetesImageArtifactDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateKubernetesImageArtifactDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateKubernetesImageArtifactDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateKubernetesImageArtifactDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateKubernetesImageArtifactDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateKubernetesImageArtifactDetails CreateKubernetesImageArtifactDetails
	s := struct {
		DiscriminatorParam string `json:"artifactType"`
		MarshalTypeCreateKubernetesImageArtifactDetails
	}{
		"HELM_CHART",
		(MarshalTypeCreateKubernetesImageArtifactDetails)(m),
	}

	return json.Marshal(&s)
}
