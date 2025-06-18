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

// OpsActionUpdateMacDeviceDetails The data to update a MacDevice in MAC_DEVICE bucket.
type OpsActionUpdateMacDeviceDetails struct {

	// The current status of the MacDevice.
	LifecycleState OpsActionUpdateMacDeviceDetailsLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The IP address assigned to the MacDevice.
	IpAddress *string `mandatory:"false" json:"ipAddress"`
}

func (m OpsActionUpdateMacDeviceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OpsActionUpdateMacDeviceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOpsActionUpdateMacDeviceDetailsLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOpsActionUpdateMacDeviceDetailsLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OpsActionUpdateMacDeviceDetailsLifecycleStateEnum Enum with underlying type: string
type OpsActionUpdateMacDeviceDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for OpsActionUpdateMacDeviceDetailsLifecycleStateEnum
const (
	OpsActionUpdateMacDeviceDetailsLifecycleStateCreating       OpsActionUpdateMacDeviceDetailsLifecycleStateEnum = "CREATING"
	OpsActionUpdateMacDeviceDetailsLifecycleStateActive         OpsActionUpdateMacDeviceDetailsLifecycleStateEnum = "ACTIVE"
	OpsActionUpdateMacDeviceDetailsLifecycleStateNeedsAttention OpsActionUpdateMacDeviceDetailsLifecycleStateEnum = "NEEDS_ATTENTION"
	OpsActionUpdateMacDeviceDetailsLifecycleStateDeleting       OpsActionUpdateMacDeviceDetailsLifecycleStateEnum = "DELETING"
	OpsActionUpdateMacDeviceDetailsLifecycleStateDeleted        OpsActionUpdateMacDeviceDetailsLifecycleStateEnum = "DELETED"
)

var mappingOpsActionUpdateMacDeviceDetailsLifecycleStateEnum = map[string]OpsActionUpdateMacDeviceDetailsLifecycleStateEnum{
	"CREATING":        OpsActionUpdateMacDeviceDetailsLifecycleStateCreating,
	"ACTIVE":          OpsActionUpdateMacDeviceDetailsLifecycleStateActive,
	"NEEDS_ATTENTION": OpsActionUpdateMacDeviceDetailsLifecycleStateNeedsAttention,
	"DELETING":        OpsActionUpdateMacDeviceDetailsLifecycleStateDeleting,
	"DELETED":         OpsActionUpdateMacDeviceDetailsLifecycleStateDeleted,
}

var mappingOpsActionUpdateMacDeviceDetailsLifecycleStateEnumLowerCase = map[string]OpsActionUpdateMacDeviceDetailsLifecycleStateEnum{
	"creating":        OpsActionUpdateMacDeviceDetailsLifecycleStateCreating,
	"active":          OpsActionUpdateMacDeviceDetailsLifecycleStateActive,
	"needs_attention": OpsActionUpdateMacDeviceDetailsLifecycleStateNeedsAttention,
	"deleting":        OpsActionUpdateMacDeviceDetailsLifecycleStateDeleting,
	"deleted":         OpsActionUpdateMacDeviceDetailsLifecycleStateDeleted,
}

// GetOpsActionUpdateMacDeviceDetailsLifecycleStateEnumValues Enumerates the set of values for OpsActionUpdateMacDeviceDetailsLifecycleStateEnum
func GetOpsActionUpdateMacDeviceDetailsLifecycleStateEnumValues() []OpsActionUpdateMacDeviceDetailsLifecycleStateEnum {
	values := make([]OpsActionUpdateMacDeviceDetailsLifecycleStateEnum, 0)
	for _, v := range mappingOpsActionUpdateMacDeviceDetailsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOpsActionUpdateMacDeviceDetailsLifecycleStateEnumStringValues Enumerates the set of values in String for OpsActionUpdateMacDeviceDetailsLifecycleStateEnum
func GetOpsActionUpdateMacDeviceDetailsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
	}
}

// GetMappingOpsActionUpdateMacDeviceDetailsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOpsActionUpdateMacDeviceDetailsLifecycleStateEnum(val string) (OpsActionUpdateMacDeviceDetailsLifecycleStateEnum, bool) {
	enum, ok := mappingOpsActionUpdateMacDeviceDetailsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
