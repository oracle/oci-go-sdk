// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// File Storage Service API
//
// APIs for OCI file storage service.
//

package ffsw

import (
    "github.com/oracle/oci-go-sdk/common"
)


    
 // CreateMountTargetDetails The representation of CreateMountTargetDetails
type CreateMountTargetDetails struct {
    
 // The Availability Domain to create the mount target in.
 // Example: `Uocm:PHX-AD-1`
    AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`
    
 // The OCID of the compartment to create the mount target in.
    CompartmentId *string `mandatory:"true" json:"compartmentId"`
    
 // The OCID of the subnet to create the mount target in.
    SubnetId *string `mandatory:"true" json:"subnetId"`
    
 // A user-friendly name. Does not have to be unique, and it is changeable.
 // Avoid entering confidential information.
 // Example: `My mount target`
    DisplayName *string `mandatory:"false" json:"displayName"`
    
 // The hostname for the mount target's IP address. Used for
 // DNS. The value is the hostname portion of the private IP's
 // fully qualified domain name (FQDN) (for example,
 // `-1` in FQDN
 // `files-1.subnet123.vcn1.oraclevcn.com`).  Must be
 // unique across all VNICs in the subnet and comply with RFC
 // 952 (https://tools.ietf.org/html/rfc952) and RFC
 // 1123 (https://tools.ietf.org/html/rfc1123).
 // For more information, see
 // DNS in Your Virtual Cloud Network ({{DOC_SERVER_URL}}/Content/Network/Concepts/dns.htm).
 // Example: `files-1`
    HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`
    
 // A private IP address of your choice. Must be an available IP address within
 // the subnet's CIDR. If you don't specify a value, Oracle automatically
 // assigns a private IP address from the subnet.
 // Example: `10.0.3.3`
    IpAddress *string `mandatory:"false" json:"ipAddress"`
}

func (m CreateMountTargetDetails) String() string {
    return common.PointerString(m)
}





