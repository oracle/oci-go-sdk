// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ShapeSummary A shape of a node on a Rover device.
type ShapeSummary struct {

	// The name of the shape.
	Shape *string `mandatory:"true" json:"shape"`

	// A short description of the graphics processing unit (GPU) available for this shape.
	GpuDescription *string `mandatory:"false" json:"gpuDescription"`

	// The number of GPUs available for this shape.
	Gpus *int `mandatory:"false" json:"gpus"`

	// The default amount of memory available for this shape, in gigabytes.
	MemoryInGBs *float32 `mandatory:"false" json:"memoryInGBs"`

	// The networking bandwidth available for this shape, in gigabits per second.
	NetworkingBandwidthInGbps *float32 `mandatory:"false" json:"networkingBandwidthInGbps"`

	// The default number of OCPUs available for this shape.
	Ocpus *int `mandatory:"false" json:"ocpus"`

	// A short description of the shape's processor (CPU).
	ProcessorDescription *string `mandatory:"false" json:"processorDescription"`

	// A short description of the USB controller available for this shape.
	UsbControllerDescription *string `mandatory:"false" json:"usbControllerDescription"`

	// The number of USB controllers available for this shape.
	NumberOfUsbControllers *int `mandatory:"false" json:"numberOfUsbControllers"`

	// The tags associated with tagSlug.
	Tags *string `mandatory:"false" json:"tags"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ShapeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShapeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
