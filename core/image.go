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

// Image. A boot disk image for launching an instance. For more information, see
// [Overview of the Compute Service]({{DOC_SERVER_URL}}/Content/Compute/Concepts/computeoverview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type Image struct {

	// The OCID of the compartment containing the instance you want to use as the basis for the image.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// Whether instances launched with this image can be used to create new images.
	// For example, you cannot create an image of an Oracle Database instance.
	// Example: `true`
	CreateImageAllowed *bool `mandatory:"true" json:"createImageAllowed,omitempty"`

	// The OCID of the image.
	ID *string `mandatory:"true" json:"id,omitempty"`

	LifecycleState ImageLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The image's operating system.
	// Example: `Oracle Linux`
	OperatingSystem *string `mandatory:"true" json:"operatingSystem,omitempty"`

	// The image's operating system version.
	// Example: `7.2`
	OperatingSystemVersion *string `mandatory:"true" json:"operatingSystemVersion,omitempty"`

	// The date and time the image was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated,omitempty"`

	// The OCID of the image originally used to launch the instance.
	BaseImageID *string `mandatory:"false" json:"baseImageId,omitempty"`

	// A user-friendly name for the image. It does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// You cannot use an Oracle-provided image name as a custom image name.
	// Example: `My custom Oracle Linux image`
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
}

func (model Image) String() string {
	return common.PointerString(model)
}

type ImageLifecycleStateEnum string

const (
	IMAGE_LIFECYCLE_STATE_PROVISIONING ImageLifecycleStateEnum = "PROVISIONING"
	IMAGE_LIFECYCLE_STATE_IMPORTING    ImageLifecycleStateEnum = "IMPORTING"
	IMAGE_LIFECYCLE_STATE_AVAILABLE    ImageLifecycleStateEnum = "AVAILABLE"
	IMAGE_LIFECYCLE_STATE_EXPORTING    ImageLifecycleStateEnum = "EXPORTING"
	IMAGE_LIFECYCLE_STATE_DISABLED     ImageLifecycleStateEnum = "DISABLED"
	IMAGE_LIFECYCLE_STATE_DELETED      ImageLifecycleStateEnum = "DELETED"
	IMAGE_LIFECYCLE_STATE_UNKNOWN      ImageLifecycleStateEnum = "UNKNOWN"
)

var mapping_image_lifecycleState = map[string]ImageLifecycleStateEnum{
	"PROVISIONING": IMAGE_LIFECYCLE_STATE_PROVISIONING,
	"IMPORTING":    IMAGE_LIFECYCLE_STATE_IMPORTING,
	"AVAILABLE":    IMAGE_LIFECYCLE_STATE_AVAILABLE,
	"EXPORTING":    IMAGE_LIFECYCLE_STATE_EXPORTING,
	"DISABLED":     IMAGE_LIFECYCLE_STATE_DISABLED,
	"DELETED":      IMAGE_LIFECYCLE_STATE_DELETED,
	"UNKNOWN":      IMAGE_LIFECYCLE_STATE_UNKNOWN,
}

func GetImageLifecycleStateEnumValues() []ImageLifecycleStateEnum {
	values := make([]ImageLifecycleStateEnum, 0)
	for _, v := range mapping_image_lifecycleState {
		if v != IMAGE_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
