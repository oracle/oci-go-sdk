// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DbSystem The Database Service supports several types of DB Systems, ranging in size, price, and performance. For details about each type of system, see:
// - [Exadata DB Systems]({{DOC_SERVER_URL}}/Content/Database/Concepts/exaoverview.htm)
// - [Bare Metal or VM DB Systems]({{DOC_SERVER_URL}}/Content/Database/Concepts/overview.htm)
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
//
// For information about access control and compartments, see
// [Overview of the Identity Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm).
// For information about Availability Domains, see
// [Regions and Availability Domains]({{DOC_SERVER_URL}}/Content/General/Concepts/regions.htm).
// To get a list of Availability Domains, use the `ListAvailabilityDomains` operation
// in the Identity Service API.
type DbSystem struct {

	// The name of the Availability Domain that the DB System is located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain,omitempty"`

	// The OCID of the compartment.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The number of CPU cores enabled on the DB System.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount,omitempty"`

	// The Oracle Database Edition that applies to all the databases on the DB System.
	DatabaseEdition DbSystemDatabaseEditionEnum `mandatory:"true" json:"databaseEdition,omitempty"`

	// The user-friendly name for the DB System. It does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName,omitempty"`

	// The domain name for the DB System.
	Domain *string `mandatory:"true" json:"domain,omitempty"`

	// The host name for the DB Node.
	Hostname *string `mandatory:"true" json:"hostname,omitempty"`

	// The OCID of the DB System.
	ID *string `mandatory:"true" json:"id,omitempty"`

	// The current state of the DB System.
	LifecycleState DbSystemLifecycleStateEnum `mandatory:"true" json:"lifecycleState,omitempty"`

	// The shape of the DB System. The shape determines resources to allocate to the DB system - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes.
	Shape *string `mandatory:"true" json:"shape,omitempty"`

	// The public key portion of one or more key pairs used for SSH access to the DB System.
	SshPublicKeys []string `mandatory:"true" json:"sshPublicKeys,omitempty"`

	// The OCID of the subnet the DB System is associated with.
	// **Subnet Restrictions:**
	// - For single node and 2-node (RAC) DB Systems, do not use a subnet that overlaps with 192.168.16.16/28
	// - For Exadata and VM-based RAC DB Systems, do not use a subnet that overlaps with 192.168.128.0/20
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and backup subnet.
	SubnetID *string `mandatory:"true" json:"subnetId,omitempty"`

	// The OCID of the backup network subnet the DB System is associated with. Applicable only to Exadata.
	// **Subnet Restriction:** See above subnetId's 'Subnet Restriction'.
	// to malfunction.
	BackupSubnetID *string `mandatory:"false" json:"backupSubnetId,omitempty"`

	// Cluster name for Exadata and 2-node RAC DB Systems. The cluster name must begin with an an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive.
	ClusterName *string `mandatory:"false" json:"clusterName,omitempty"`

	// The percentage assigned to DATA storage (user data and database files).
	// The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups). Accepted values are 40 and 80.
	DataStoragePercentage *int `mandatory:"false" json:"dataStoragePercentage,omitempty"`

	// Data storage size, in GBs, that is currently available to the DB system. This is applicable only for VM-based DBs.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs,omitempty"`

	// The type of redundancy configured for the DB System.
	// Normal is 2-way redundancy.
	// High is 3-way redundancy.
	DiskRedundancy DbSystemDiskRedundancyEnum `mandatory:"false" json:"diskRedundancy,omitempty"`

	// The OCID of the last patch history. This is updated as soon as a patch operation is started.
	LastPatchHistoryEntryID *string `mandatory:"false" json:"lastPatchHistoryEntryId,omitempty"`

	// The Oracle license model that applies to all the databases on the DB System. The default is LICENSE_INCLUDED.
	LicenseModel DbSystemLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// The port number configured for the listener on the DB System.
	ListenerPort *int `mandatory:"false" json:"listenerPort,omitempty"`

	// Number of nodes in this DB system. For RAC DBs, this will be greater than 1.
	NodeCount *int `mandatory:"false" json:"nodeCount,omitempty"`

	// RECO/REDO storage size, in GBs, that is currently allocated to the DB system. This is applicable only for VM-based DBs.
	RecoStorageSizeInGB *int `mandatory:"false" json:"recoStorageSizeInGB,omitempty"`

	// The OCID of the DNS record for the SCAN IP addresses that are associated with the DB System.
	ScanDnsRecordID *string `mandatory:"false" json:"scanDnsRecordId,omitempty"`

	// The OCID of the Single Client Access Name (SCAN) IP addresses associated with the DB System.
	// SCAN IP addresses are typically used for load balancing and are not assigned to any interface.
	// Clusterware directs the requests to the appropriate nodes in the cluster.
	// - For a single-node DB System, this list is empty.
	ScanIpIds []string `mandatory:"false" json:"scanIpIds,omitempty"`

	// The date and time the DB System was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`

	// The version of the DB System.
	Version *string `mandatory:"false" json:"version,omitempty"`

	// The OCID of the virtual IP (VIP) addresses associated with the DB System.
	// The Cluster Ready Services (CRS) creates and maintains one VIP address for each node in the DB System to
	// enable failover. If one node fails, the VIP is reassigned to another active node in the cluster.
	// - For a single-node DB System, this list is empty.
	VipIds []string `mandatory:"false" json:"vipIds,omitempty"`
}

func (model DbSystem) String() string {
	return common.PointerString(model)
}

// DbSystemDatabaseEditionEnum Enum with underlying type: string
type DbSystemDatabaseEditionEnum string

// Set of constants representing the allowable values for DbSystemDatabaseEdition
const (
	DbSystemDatabaseEditionStandardEdition                     DbSystemDatabaseEditionEnum = "STANDARD_EDITION"
	DbSystemDatabaseEditionEnterpriseEdition                   DbSystemDatabaseEditionEnum = "ENTERPRISE_EDITION"
	DbSystemDatabaseEditionEnterpriseEditionExtremePerformance DbSystemDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
	DbSystemDatabaseEditionEnterpriseEditionHighPerformance    DbSystemDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	DbSystemDatabaseEditionUnknown                             DbSystemDatabaseEditionEnum = "UNKNOWN"
)

var mappingDbSystemDatabaseEdition = map[string]DbSystemDatabaseEditionEnum{
	"STANDARD_EDITION":                       DbSystemDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     DbSystemDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": DbSystemDatabaseEditionEnterpriseEditionExtremePerformance,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    DbSystemDatabaseEditionEnterpriseEditionHighPerformance,
	"UNKNOWN": DbSystemDatabaseEditionUnknown,
}

// GetDbSystemDatabaseEditionEnumValues Enumerates the set of values for DbSystemDatabaseEdition
func GetDbSystemDatabaseEditionEnumValues() []DbSystemDatabaseEditionEnum {
	values := make([]DbSystemDatabaseEditionEnum, 0)
	for _, v := range mappingDbSystemDatabaseEdition {
		if v != DbSystemDatabaseEditionUnknown {
			values = append(values, v)
		}
	}
	return values
}

// DbSystemDiskRedundancyEnum Enum with underlying type: string
type DbSystemDiskRedundancyEnum string

// Set of constants representing the allowable values for DbSystemDiskRedundancy
const (
	DbSystemDiskRedundancyHigh    DbSystemDiskRedundancyEnum = "HIGH"
	DbSystemDiskRedundancyNormal  DbSystemDiskRedundancyEnum = "NORMAL"
	DbSystemDiskRedundancyUnknown DbSystemDiskRedundancyEnum = "UNKNOWN"
)

var mappingDbSystemDiskRedundancy = map[string]DbSystemDiskRedundancyEnum{
	"HIGH":    DbSystemDiskRedundancyHigh,
	"NORMAL":  DbSystemDiskRedundancyNormal,
	"UNKNOWN": DbSystemDiskRedundancyUnknown,
}

// GetDbSystemDiskRedundancyEnumValues Enumerates the set of values for DbSystemDiskRedundancy
func GetDbSystemDiskRedundancyEnumValues() []DbSystemDiskRedundancyEnum {
	values := make([]DbSystemDiskRedundancyEnum, 0)
	for _, v := range mappingDbSystemDiskRedundancy {
		if v != DbSystemDiskRedundancyUnknown {
			values = append(values, v)
		}
	}
	return values
}

// DbSystemLicenseModelEnum Enum with underlying type: string
type DbSystemLicenseModelEnum string

// Set of constants representing the allowable values for DbSystemLicenseModel
const (
	DbSystemLicenseModelLicenseIncluded     DbSystemLicenseModelEnum = "LICENSE_INCLUDED"
	DbSystemLicenseModelBringYourOwnLicense DbSystemLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
	DbSystemLicenseModelUnknown             DbSystemLicenseModelEnum = "UNKNOWN"
)

var mappingDbSystemLicenseModel = map[string]DbSystemLicenseModelEnum{
	"LICENSE_INCLUDED":       DbSystemLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": DbSystemLicenseModelBringYourOwnLicense,
	"UNKNOWN":                DbSystemLicenseModelUnknown,
}

// GetDbSystemLicenseModelEnumValues Enumerates the set of values for DbSystemLicenseModel
func GetDbSystemLicenseModelEnumValues() []DbSystemLicenseModelEnum {
	values := make([]DbSystemLicenseModelEnum, 0)
	for _, v := range mappingDbSystemLicenseModel {
		if v != DbSystemLicenseModelUnknown {
			values = append(values, v)
		}
	}
	return values
}

// DbSystemLifecycleStateEnum Enum with underlying type: string
type DbSystemLifecycleStateEnum string

// Set of constants representing the allowable values for DbSystemLifecycleState
const (
	DbSystemLifecycleStateProvisioning DbSystemLifecycleStateEnum = "PROVISIONING"
	DbSystemLifecycleStateAvailable    DbSystemLifecycleStateEnum = "AVAILABLE"
	DbSystemLifecycleStateUpdating     DbSystemLifecycleStateEnum = "UPDATING"
	DbSystemLifecycleStateTerminating  DbSystemLifecycleStateEnum = "TERMINATING"
	DbSystemLifecycleStateTerminated   DbSystemLifecycleStateEnum = "TERMINATED"
	DbSystemLifecycleStateFailed       DbSystemLifecycleStateEnum = "FAILED"
	DbSystemLifecycleStateUnknown      DbSystemLifecycleStateEnum = "UNKNOWN"
)

var mappingDbSystemLifecycleState = map[string]DbSystemLifecycleStateEnum{
	"PROVISIONING": DbSystemLifecycleStateProvisioning,
	"AVAILABLE":    DbSystemLifecycleStateAvailable,
	"UPDATING":     DbSystemLifecycleStateUpdating,
	"TERMINATING":  DbSystemLifecycleStateTerminating,
	"TERMINATED":   DbSystemLifecycleStateTerminated,
	"FAILED":       DbSystemLifecycleStateFailed,
	"UNKNOWN":      DbSystemLifecycleStateUnknown,
}

// GetDbSystemLifecycleStateEnumValues Enumerates the set of values for DbSystemLifecycleState
func GetDbSystemLifecycleStateEnumValues() []DbSystemLifecycleStateEnum {
	values := make([]DbSystemLifecycleStateEnum, 0)
	for _, v := range mappingDbSystemLifecycleState {
		if v != DbSystemLifecycleStateUnknown {
			values = append(values, v)
		}
	}
	return values
}
