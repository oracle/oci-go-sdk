// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GenAI Sandbox API
//
// API for managing sandbox resources and exchanging sandbox access tokens.
//

package generativeaisandbox

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateSandboxDetails Request body used to create a sandbox.
type CreateSandboxDetails struct {

	// Name of a MANAGED or CUSTOM warm pool in the owning project. Runtime and shape
	// are derived from the pool; explicit values must match. Deleting pools cannot be used.
	// Snapshot restore, credential brokerage, and nested resource principals are not
	// supported with named pools. An empty pool may cold-create using its same template,
	// including when desired capacity is zero. Zero capacity stops idle prewarming;
	// it does not prevent creation from an available applied template.
	// When warmpoolName is omitted, a built-in runtime can automatically select an eligible
	// named MANAGED pool in the same project matching its runtime and shape. Omitted runtime
	// defaults to python-3.11 and omitted shape defaults to SMALL for automatic matching.
	// A digest-pinned customer image can automatically select an eligible CUSTOM pool in the
	// same project matching its canonical registry host, repository, digest and shape.
	// Tag-based customer images are not automatically matched. Matching pools are considered
	// in name order; the first with a valid applied template is selected without reading ready counts.
	// Empty pools and zero desired or applied capacity still allow cold creation through
	// the selected template. If no eligible pool is found or automatic lookup is unavailable,
	// built-in runtimes use the existing managed runtime path, and customer images use the
	// ordinary direct customer-image path with its existing validation.
	// Automatic matching does not apply when source is supplied, including snapshot restore,
	// or when credential brokerage or nested resource principals are used. Explicit names
	// retain strict validation and do not fall back to automatic matching or another creation path.
	// Retries preserve the original persisted pool binding or the original managed or direct
	// customer-image route, even when matching pools change.
	// Sandbox and SandboxSummary report an automatically selected pool in warmpoolName.
	WarmpoolName *string `mandatory:"false" json:"warmpoolName"`

	Source SandboxSource `mandatory:"false" json:"source"`

	// User-provided display name for the sandbox.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Runtime identifier for the sandbox, either a built-in runtime such as python-3.11 or a same-region Oracle Cloud Infrastructure Registry customer image reference.
	Runtime *string `mandatory:"false" json:"runtime"`

	// Compute shape for the sandbox.
	Shape SandboxShapeEnum `mandatory:"false" json:"shape,omitempty"`

	// Authorization for all published-port HTTP requests and WebSocket handshakes.
	// OCI_IAM requires Sandbox use and project read permission on the owning project.
	// Omission defaults to OCI_IAM for new sandboxes with published ports in enabled
	// tenancies; elsewhere omission retains legacy behavior. Explicit OCI_IAM outside
	// the rollout, null, unsupported values, or a mode without published ports are rejected.
	// The effective mode is immutable and is preserved when restoring protected snapshots.
	PublishedPortAuthMode SandboxPublishedPortAuthModeEnum `mandatory:"false" json:"publishedPortAuthMode,omitempty"`

	// Published ports to expose for the sandbox.
	PublishPorts []string `mandatory:"false" json:"publishPorts"`

	// Number of minutes from creation until automatic sandbox expiration. Use `-1` to disable automatic expiration.
	ExpiresInMinutes *int `mandatory:"false" json:"expiresInMinutes"`

	// Deprecated. Use `expiresInMinutes` instead. Existing clients may pass `-1` to disable automatic expiration.
	ExpirationDuration *string `mandatory:"false" json:"expirationDuration"`

	// Environment variables to inject into the sandbox runtime.
	EnvironmentVariables []EnvironmentVariable `mandatory:"false" json:"environmentVariables"`

	// The Project credential profile selected for this Sandbox. Omit this field to select
	// no credential profile. The selection and its configuration snapshot are immutable for
	// the lifetime of the Sandbox; later Project profile updates apply only to subsequently
	// created Sandboxes after configuration propagation.
	CredentialProfileLabel *string `mandatory:"false" json:"credentialProfileLabel"`
}

func (m CreateSandboxDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSandboxDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSandboxShapeEnum(string(m.Shape)); !ok && m.Shape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shape: %s. Supported values are: %s.", m.Shape, strings.Join(GetSandboxShapeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSandboxPublishedPortAuthModeEnum(string(m.PublishedPortAuthMode)); !ok && m.PublishedPortAuthMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PublishedPortAuthMode: %s. Supported values are: %s.", m.PublishedPortAuthMode, strings.Join(GetSandboxPublishedPortAuthModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateSandboxDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		WarmpoolName           *string                          `json:"warmpoolName"`
		Source                 sandboxsource                    `json:"source"`
		DisplayName            *string                          `json:"displayName"`
		Runtime                *string                          `json:"runtime"`
		Shape                  SandboxShapeEnum                 `json:"shape"`
		PublishedPortAuthMode  SandboxPublishedPortAuthModeEnum `json:"publishedPortAuthMode"`
		PublishPorts           []string                         `json:"publishPorts"`
		ExpiresInMinutes       *int                             `json:"expiresInMinutes"`
		ExpirationDuration     *string                          `json:"expirationDuration"`
		EnvironmentVariables   []EnvironmentVariable            `json:"environmentVariables"`
		CredentialProfileLabel *string                          `json:"credentialProfileLabel"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.WarmpoolName = model.WarmpoolName

	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(SandboxSource)
	} else {
		m.Source = nil
	}

	m.DisplayName = model.DisplayName

	m.Runtime = model.Runtime

	m.Shape = model.Shape

	m.PublishedPortAuthMode = model.PublishedPortAuthMode

	m.PublishPorts = make([]string, len(model.PublishPorts))
	copy(m.PublishPorts, model.PublishPorts)
	m.ExpiresInMinutes = model.ExpiresInMinutes

	m.ExpirationDuration = model.ExpirationDuration

	m.EnvironmentVariables = make([]EnvironmentVariable, len(model.EnvironmentVariables))
	copy(m.EnvironmentVariables, model.EnvironmentVariables)
	m.CredentialProfileLabel = model.CredentialProfileLabel

	return
}
