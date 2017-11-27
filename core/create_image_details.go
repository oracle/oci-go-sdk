// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateImageDetails. Either instanceId or imageSourceDetails must be provided in addition to other required parameters.
type CreateImageDetails struct {

	// The OCID of the compartment containing the instance you want to use as the basis for the image.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// A user-friendly name for the image. It does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// You cannot use an Oracle-provided image name as a custom image name.
	// Example: `My Oracle Linux image`
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// Details for creating an image through import
	ImageSourceDetails *ImageSourceDetails `mandatory:"false" json:"imageSourceDetails,omitempty"`

	// The OCID of the instance you want to use as the basis for the image.
	InstanceID *string `mandatory:"false" json:"instanceId,omitempty"`
}

func (model CreateImageDetails) String() string {
	return common.PointerString(model)
}
