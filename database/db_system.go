// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DbSystem The Database Service supports several types of DB systems, ranging in size, price, and performance. For details about each type of system, see:
// - Exadata DB Systems (https://docs.us-phoenix-1.oraclecloud.com/Content/Database/Concepts/exaoverview.htm)
// - Bare Metal and Virtual Machine DB Systems (https://docs.us-phoenix-1.oraclecloud.com/Content/Database/Concepts/overview.htm)
// To use any of the API operations, you must be authorized in an IAM policy. If you are not authorized, talk to an administrator. If you are an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).
//
// For information about access control and compartments, see
// Overview of the Identity Service (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
// For information about availability domains, see
// Regions and Availability Domains (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/regions.htm).
// To get a list of availability domains, use the `ListAvailabilityDomains` operation
// in the Identity Service API.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type DbSystem struct {

	// The name of the availability domain that the DB system is located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The number of CPU cores enabled on the DB system.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The Oracle Database edition that applies to all the databases on the DB system.
	DatabaseEdition DbSystemDatabaseEditionEnum `mandatory:"true" json:"databaseEdition"`

	// The user-friendly name for the DB system. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The domain name for the DB system.
	Domain *string `mandatory:"true" json:"domain"`

	// The hostname for the DB system.
	Hostname *string `mandatory:"true" json:"hostname"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the DB system.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the DB system.
	LifecycleState DbSystemLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The shape of the DB system. The shape determines resources to allocate to the DB system.
	// - For virtual machine shapes, the number of CPU cores and memory
	// - For bare metal and Exadata shapes, the number of CPU cores, storage, and memory
	Shape *string `mandatory:"true" json:"shape"`

	// The public key portion of one or more key pairs used for SSH access to the DB system.
	SshPublicKeys []string `mandatory:"true" json:"sshPublicKeys"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the subnet the DB system is associated with.
	// **Subnet Restrictions:**
	// - For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	// - For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and backup subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the backup network subnet the DB system is associated with. Applicable only to Exadata DB systems.
	// **Subnet Restriction:** See the subnet restrictions information for **subnetId**.
	BackupSubnetId *string `mandatory:"false" json:"backupSubnetId"`

	// The cluster name for Exadata and 2-node RAC virtual machine DB systems. The cluster name must begin with an an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive.
	ClusterName *string `mandatory:"false" json:"clusterName"`

	// The percentage assigned to DATA storage (user data and database files).
	// The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups). Accepted values are 40 and 80. The default is 80 percent assigned to DATA storage. Not applicable for virtual machine DB systems.
	DataStoragePercentage *int `mandatory:"false" json:"dataStoragePercentage"`

	// The data storage size, in gigabytes, that is currently available to the DB system. Applies only for virtual machine DB systems.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The type of redundancy configured for the DB system.
	// NORMAL is 2-way redundancy.
	// HIGH is 3-way redundancy.
	DiskRedundancy DbSystemDiskRedundancyEnum `mandatory:"false" json:"diskRedundancy,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the last patch history. This value is updated as soon as a patch operation starts.
	LastPatchHistoryEntryId *string `mandatory:"false" json:"lastPatchHistoryEntryId"`

	// The Oracle license model that applies to all the databases on the DB system. The default is LICENSE_INCLUDED.
	LicenseModel DbSystemLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The port number configured for the listener on the DB system.
	ListenerPort *int `mandatory:"false" json:"listenerPort"`

	// The number of nodes in the DB system. For RAC DB systems, the value is greater than 1.
	NodeCount *int `mandatory:"false" json:"nodeCount"`

	// The RECO/REDO storage size, in gigabytes, that is currently allocated to the DB system. Applies only for virtual machine DB systems.
	RecoStorageSizeInGB *int `mandatory:"false" json:"recoStorageSizeInGB"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the DNS record for the SCAN IP addresses that are associated with the DB system.
	ScanDnsRecordId *string `mandatory:"false" json:"scanDnsRecordId"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the Single Client Access Name (SCAN) IP addresses associated with the DB system.
	// SCAN IP addresses are typically used for load balancing and are not assigned to any interface.
	// Oracle Clusterware directs the requests to the appropriate nodes in the cluster.
	// **Note:** For a single-node DB system, this list is empty.
	ScanIpIds []string `mandatory:"false" json:"scanIpIds"`

	// The date and time the DB system was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The Oracle Database version of the DB system.
	Version *string `mandatory:"false" json:"version"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the virtual IP (VIP) addresses associated with the DB system.
	// The Cluster Ready Services (CRS) creates and maintains one VIP address for each node in the DB system to
	// enable failover. If one node fails, the VIP is reassigned to another active node in the cluster.
	// **Note:** For a single-node DB system, this list is empty.
	VipIds []string `mandatory:"false" json:"vipIds"`
}

func (m DbSystem) String() string {
	return common.PointerString(m)
}

// DbSystemDatabaseEditionEnum Enum with underlying type: string
type DbSystemDatabaseEditionEnum string

// Set of constants representing the allowable values for DbSystemDatabaseEdition
const (
	DbSystemDatabaseEditionStandardEdition                     DbSystemDatabaseEditionEnum = "STANDARD_EDITION"
	DbSystemDatabaseEditionEnterpriseEdition                   DbSystemDatabaseEditionEnum = "ENTERPRISE_EDITION"
	DbSystemDatabaseEditionEnterpriseEditionHighPerformance    DbSystemDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	DbSystemDatabaseEditionEnterpriseEditionExtremePerformance DbSystemDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
)

var mappingDbSystemDatabaseEdition = map[string]DbSystemDatabaseEditionEnum{
	"STANDARD_EDITION":                       DbSystemDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     DbSystemDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    DbSystemDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": DbSystemDatabaseEditionEnterpriseEditionExtremePerformance,
}

// GetDbSystemDatabaseEditionEnumValues Enumerates the set of values for DbSystemDatabaseEdition
func GetDbSystemDatabaseEditionEnumValues() []DbSystemDatabaseEditionEnum {
	values := make([]DbSystemDatabaseEditionEnum, 0)
	for _, v := range mappingDbSystemDatabaseEdition {
		values = append(values, v)
	}
	return values
}

// DbSystemDiskRedundancyEnum Enum with underlying type: string
type DbSystemDiskRedundancyEnum string

// Set of constants representing the allowable values for DbSystemDiskRedundancy
const (
	DbSystemDiskRedundancyHigh   DbSystemDiskRedundancyEnum = "HIGH"
	DbSystemDiskRedundancyNormal DbSystemDiskRedundancyEnum = "NORMAL"
)

var mappingDbSystemDiskRedundancy = map[string]DbSystemDiskRedundancyEnum{
	"HIGH":   DbSystemDiskRedundancyHigh,
	"NORMAL": DbSystemDiskRedundancyNormal,
}

// GetDbSystemDiskRedundancyEnumValues Enumerates the set of values for DbSystemDiskRedundancy
func GetDbSystemDiskRedundancyEnumValues() []DbSystemDiskRedundancyEnum {
	values := make([]DbSystemDiskRedundancyEnum, 0)
	for _, v := range mappingDbSystemDiskRedundancy {
		values = append(values, v)
	}
	return values
}

// DbSystemLicenseModelEnum Enum with underlying type: string
type DbSystemLicenseModelEnum string

// Set of constants representing the allowable values for DbSystemLicenseModel
const (
	DbSystemLicenseModelLicenseIncluded     DbSystemLicenseModelEnum = "LICENSE_INCLUDED"
	DbSystemLicenseModelBringYourOwnLicense DbSystemLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingDbSystemLicenseModel = map[string]DbSystemLicenseModelEnum{
	"LICENSE_INCLUDED":       DbSystemLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": DbSystemLicenseModelBringYourOwnLicense,
}

// GetDbSystemLicenseModelEnumValues Enumerates the set of values for DbSystemLicenseModel
func GetDbSystemLicenseModelEnumValues() []DbSystemLicenseModelEnum {
	values := make([]DbSystemLicenseModelEnum, 0)
	for _, v := range mappingDbSystemLicenseModel {
		values = append(values, v)
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
)

var mappingDbSystemLifecycleState = map[string]DbSystemLifecycleStateEnum{
	"PROVISIONING": DbSystemLifecycleStateProvisioning,
	"AVAILABLE":    DbSystemLifecycleStateAvailable,
	"UPDATING":     DbSystemLifecycleStateUpdating,
	"TERMINATING":  DbSystemLifecycleStateTerminating,
	"TERMINATED":   DbSystemLifecycleStateTerminated,
	"FAILED":       DbSystemLifecycleStateFailed,
}

// GetDbSystemLifecycleStateEnumValues Enumerates the set of values for DbSystemLifecycleState
func GetDbSystemLifecycleStateEnumValues() []DbSystemLifecycleStateEnum {
	values := make([]DbSystemLifecycleStateEnum, 0)
	for _, v := range mappingDbSystemLifecycleState {
		values = append(values, v)
	}
	return values
}
