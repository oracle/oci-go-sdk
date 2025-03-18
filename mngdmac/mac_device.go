// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Managed Services for Mac API
//
// Use the OCI Managed Services for Mac API to create and manage orders for dedicated, partially-managed Mac servers hosted in an OCI Data Center. For more information, see the user guide documentation for the [OCI Managed Services for Mac]()
//

package mngdmac

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MacDevice Represents a MacDevice resource.
type MacDevice struct {

	// The unique ID of the MacDevice.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the compartment to which the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the resource.
	MacOrderId *string `mandatory:"true" json:"macOrderId"`

	// The serial number of the MacDevice.
	SerialNumber *string `mandatory:"true" json:"serialNumber"`

	// The IP address assigned to the MacDevice.
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// The current status of the MacDevice.
	LifecycleState MacDeviceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The shape of the Mac.
	Shape MacOrderShapeEnum `mandatory:"true" json:"shape"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A flag that indicates if this MacDevice is decommissioned.
	IsMarkedDecom *bool `mandatory:"true" json:"isMarkedDecom"`

	// An RFC3339-formatted datetime string containing the time this MacDevice was decommissioned.
	TimeDecom *common.SDKTime `mandatory:"true" json:"timeDecom"`
}

func (m MacDevice) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MacDevice) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMacDeviceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMacDeviceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMacOrderShapeEnum(string(m.Shape)); !ok && m.Shape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shape: %s. Supported values are: %s.", m.Shape, strings.Join(GetMacOrderShapeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MacDeviceLifecycleStateEnum Enum with underlying type: string
type MacDeviceLifecycleStateEnum string

// Set of constants representing the allowable values for MacDeviceLifecycleStateEnum
const (
	MacDeviceLifecycleStateCreating       MacDeviceLifecycleStateEnum = "CREATING"
	MacDeviceLifecycleStateActive         MacDeviceLifecycleStateEnum = "ACTIVE"
	MacDeviceLifecycleStateNeedsAttention MacDeviceLifecycleStateEnum = "NEEDS_ATTENTION"
	MacDeviceLifecycleStateDeleting       MacDeviceLifecycleStateEnum = "DELETING"
	MacDeviceLifecycleStateDeleted        MacDeviceLifecycleStateEnum = "DELETED"
)

var mappingMacDeviceLifecycleStateEnum = map[string]MacDeviceLifecycleStateEnum{
	"CREATING":        MacDeviceLifecycleStateCreating,
	"ACTIVE":          MacDeviceLifecycleStateActive,
	"NEEDS_ATTENTION": MacDeviceLifecycleStateNeedsAttention,
	"DELETING":        MacDeviceLifecycleStateDeleting,
	"DELETED":         MacDeviceLifecycleStateDeleted,
}

var mappingMacDeviceLifecycleStateEnumLowerCase = map[string]MacDeviceLifecycleStateEnum{
	"creating":        MacDeviceLifecycleStateCreating,
	"active":          MacDeviceLifecycleStateActive,
	"needs_attention": MacDeviceLifecycleStateNeedsAttention,
	"deleting":        MacDeviceLifecycleStateDeleting,
	"deleted":         MacDeviceLifecycleStateDeleted,
}

// GetMacDeviceLifecycleStateEnumValues Enumerates the set of values for MacDeviceLifecycleStateEnum
func GetMacDeviceLifecycleStateEnumValues() []MacDeviceLifecycleStateEnum {
	values := make([]MacDeviceLifecycleStateEnum, 0)
	for _, v := range mappingMacDeviceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMacDeviceLifecycleStateEnumStringValues Enumerates the set of values in String for MacDeviceLifecycleStateEnum
func GetMacDeviceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
	}
}

// GetMappingMacDeviceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMacDeviceLifecycleStateEnum(val string) (MacDeviceLifecycleStateEnum, bool) {
	enum, ok := mappingMacDeviceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
