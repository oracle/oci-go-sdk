// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// API for the File Storage service. Use this API to manage file systems, mount targets, and snapshots. For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateMountTargetDetails Details for creating the mount target.
type CreateMountTargetDetails struct {

	// The availability domain in which to create the mount target.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the compartment in which to create the mount target.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the subnet in which to create the mount target.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My mount target`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The hostname for the mount target's IP address, used for
	// DNS resolution. The value is the hostname portion of the private IP
	// address's fully qualified domain name (FQDN). For example,
	// `files-1` in the FQDN `files-1.subnet123.vcn1.oraclevcn.com`.
	// Must be unique across all VNICs in the subnet and comply
	// with RFC 952 (https://tools.ietf.org/html/rfc952)
	// and RFC 1123 (https://tools.ietf.org/html/rfc1123).
	// For more information, see
	// DNS in Your Virtual Cloud Network (https://docs.cloud.oracle.com/Content/Network/Concepts/dns.htm).
	// Example: `files-1`
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// A private IP address of your choice. Must be an available IP address within
	// the subnet's CIDR. If you don't specify a value, Oracle automatically
	// assigns a private IP address from the subnet.
	// Example: `10.0.3.3`
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// A list of Network Security Group OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) associated with this mount target.
	// A maximum of 5 is allowed.
	// Setting this to an empty array after the list is created removes the mount target from all NSGs.
	// For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	//  with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateMountTargetDetails) String() string {
	return common.PointerString(m)
}
