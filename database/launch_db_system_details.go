// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oci-go-sdk/common"
)

type LaunchDbSystemDetails struct {

	// The Availability Domain where the DB System is located.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain,omitempty"`

	// The Oracle Cloud ID (OCID) of the compartment the DB System  belongs in.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The number of CPU cores to enable. The valid values depend on the specified shape:
	// - BM.DenseIO1.36 and BM.HighIO1.36 - Specify a multiple of 2, from 2 to 36.
	// - BM.RACLocalStorage1.72 - Specify a multiple of 4, from 4 to 72.
	// - Exadata.Quarter1.84 - Specify a multiple of 2, from 22 to 84.
	// - Exadata.Half1.168 - Specify a multiple of 4, from 44 to 168.
	// - Exadata.Full1.336 - Specify a multiple of 8, from 88 to 336.
	// For VM DB systems, the core count is inferred from the specific VM shape chosen, so this parameter is not used.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount,omitempty"`

	// The Oracle Database Edition that applies to all the databases on the DB System.
	// Exadata DB Systems and 2-node RAC DB Systems require ENTERPRISE_EDITION_EXTREME_PERFORMANCE.
	DatabaseEdition LaunchDbSystemDetailsDatabaseEditionEnum `mandatory:"true" json:"databaseEdition,omitempty"`

	DbHome *CreateDbHomeDetails `mandatory:"true" json:"dbHome,omitempty"`

	// The host name for the DB System. The host name must begin with an alphabetic character and
	// can contain a maximum of 30 alphanumeric characters, including hyphens (-).
	// The maximum length of the combined hostname and domain is 63 characters.
	// **Note:** The hostname must be unique within the subnet. If it is not unique,
	// the DB System will fail to provision.
	Hostname *string `mandatory:"true" json:"hostname,omitempty"`

	// The shape of the DB System. The shape determines resources allocated to the DB System - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes. To get a list of shapes, use the ListDbSystemShapes operation.
	Shape *string `mandatory:"true" json:"shape,omitempty"`

	// The public key portion of the key pair to use for SSH access to the DB System. Multiple public keys can be provided. The length of the combined keys cannot exceed 10,000 characters.
	SshPublicKeys *[]string `mandatory:"true" json:"sshPublicKeys,omitempty"`

	// The OCID of the subnet the DB System is associated with.
	// **Subnet Restrictions:**
	// - For single node and 2-node (RAC) DB Systems, do not use a subnet that overlaps with 192.168.16.16/28
	// - For Exadata and VM-based RAC DB Systems, do not use a subnet that overlaps with 192.168.128.0/20
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and backup subnet.
	SubnetID *string `mandatory:"true" json:"subnetId,omitempty"`

	// The OCID of the backup network subnet the DB System is associated with. Applicable only to Exadata.
	// **Subnet Restrictions:** See above subnetId's **Subnet Restriction**.
	BackupSubnetID *string `mandatory:"false" json:"backupSubnetId,omitempty"`

	// Cluster name for Exadata and 2-node RAC DB Systems. The cluster name must begin with an an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive.
	ClusterName *string `mandatory:"false" json:"clusterName,omitempty"`

	// The percentage assigned to DATA storage (user data and database files).
	// The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups).
	// Specify 80 or 40. The default is 80 percent assigned to DATA storage. This is not applicable for VM based DB systems.
	DataStoragePercentage *int `mandatory:"false" json:"dataStoragePercentage,omitempty"`

	// The type of redundancy configured for the DB System.
	// Normal is 2-way redundancy, recommended for test and development systems.
	// High is 3-way redundancy, recommended for production systems.
	DiskRedundancy LaunchDbSystemDetailsDiskRedundancyEnum `mandatory:"false" json:"diskRedundancy,omitempty"`

	// The user-friendly name for the DB System. It does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// A domain name used for the DB System. If the Oracle-provided Internet and VCN
	// Resolver is enabled for the specified subnet, the domain name for the subnet is used
	// (don't provide one). Otherwise, provide a valid DNS domain name. Hyphens (-) are not permitted.
	Domain *string `mandatory:"false" json:"domain,omitempty"`

	// Size, in GBs, of the initial data volume that will be created and attached to VM-shape based DB system. This storage can later be scaled up if needed. Note that the total storage size attached will be more than what is requested, to account for REDO/RECO space and software volume.
	InitialDataStorageSizeInGB *int `mandatory:"false" json:"initialDataStorageSizeInGB,omitempty"`

	// The Oracle license model that applies to all the databases on the DB System. The default is LICENSE_INCLUDED.
	LicenseModel LaunchDbSystemDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Number of nodes to launch for a VM-shape based RAC DB system.
	NodeCount *int `mandatory:"false" json:"nodeCount,omitempty"`
}

func (model LaunchDbSystemDetails) String() string {
	return common.PointerString(model)
}

type LaunchDbSystemDetailsDatabaseEditionEnum string

const (
	LAUNCH_DB_SYSTEM_DETAILS_DATABASE_EDITION_STANDARD_EDITION                       LaunchDbSystemDetailsDatabaseEditionEnum = "STANDARD_EDITION"
	LAUNCH_DB_SYSTEM_DETAILS_DATABASE_EDITION_ENTERPRISE_EDITION                     LaunchDbSystemDetailsDatabaseEditionEnum = "ENTERPRISE_EDITION"
	LAUNCH_DB_SYSTEM_DETAILS_DATABASE_EDITION_ENTERPRISE_EDITION_EXTREME_PERFORMANCE LaunchDbSystemDetailsDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
	LAUNCH_DB_SYSTEM_DETAILS_DATABASE_EDITION_ENTERPRISE_EDITION_HIGH_PERFORMANCE    LaunchDbSystemDetailsDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	LAUNCH_DB_SYSTEM_DETAILS_DATABASE_EDITION_UNKNOWN                                LaunchDbSystemDetailsDatabaseEditionEnum = "UNKNOWN"
)

var mapping_launchdbsystemdetails_databaseEdition = map[string]LaunchDbSystemDetailsDatabaseEditionEnum{
	"STANDARD_EDITION":                       LAUNCH_DB_SYSTEM_DETAILS_DATABASE_EDITION_STANDARD_EDITION,
	"ENTERPRISE_EDITION":                     LAUNCH_DB_SYSTEM_DETAILS_DATABASE_EDITION_ENTERPRISE_EDITION,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": LAUNCH_DB_SYSTEM_DETAILS_DATABASE_EDITION_ENTERPRISE_EDITION_EXTREME_PERFORMANCE,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    LAUNCH_DB_SYSTEM_DETAILS_DATABASE_EDITION_ENTERPRISE_EDITION_HIGH_PERFORMANCE,
	"UNKNOWN": LAUNCH_DB_SYSTEM_DETAILS_DATABASE_EDITION_UNKNOWN,
}

func GetLaunchDbSystemDetailsDatabaseEditionEnumValues() []LaunchDbSystemDetailsDatabaseEditionEnum {
	values := make([]LaunchDbSystemDetailsDatabaseEditionEnum, 0)
	for _, v := range mapping_launchdbsystemdetails_databaseEdition {
		if v != LAUNCH_DB_SYSTEM_DETAILS_DATABASE_EDITION_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type LaunchDbSystemDetailsDiskRedundancyEnum string

const (
	LAUNCH_DB_SYSTEM_DETAILS_DISK_REDUNDANCY_HIGH    LaunchDbSystemDetailsDiskRedundancyEnum = "HIGH"
	LAUNCH_DB_SYSTEM_DETAILS_DISK_REDUNDANCY_NORMAL  LaunchDbSystemDetailsDiskRedundancyEnum = "NORMAL"
	LAUNCH_DB_SYSTEM_DETAILS_DISK_REDUNDANCY_UNKNOWN LaunchDbSystemDetailsDiskRedundancyEnum = "UNKNOWN"
)

var mapping_launchdbsystemdetails_diskRedundancy = map[string]LaunchDbSystemDetailsDiskRedundancyEnum{
	"HIGH":    LAUNCH_DB_SYSTEM_DETAILS_DISK_REDUNDANCY_HIGH,
	"NORMAL":  LAUNCH_DB_SYSTEM_DETAILS_DISK_REDUNDANCY_NORMAL,
	"UNKNOWN": LAUNCH_DB_SYSTEM_DETAILS_DISK_REDUNDANCY_UNKNOWN,
}

func GetLaunchDbSystemDetailsDiskRedundancyEnumValues() []LaunchDbSystemDetailsDiskRedundancyEnum {
	values := make([]LaunchDbSystemDetailsDiskRedundancyEnum, 0)
	for _, v := range mapping_launchdbsystemdetails_diskRedundancy {
		if v != LAUNCH_DB_SYSTEM_DETAILS_DISK_REDUNDANCY_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

type LaunchDbSystemDetailsLicenseModelEnum string

const (
	LAUNCH_DB_SYSTEM_DETAILS_LICENSE_MODEL_LICENSE_INCLUDED       LaunchDbSystemDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	LAUNCH_DB_SYSTEM_DETAILS_LICENSE_MODEL_BRING_YOUR_OWN_LICENSE LaunchDbSystemDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
	LAUNCH_DB_SYSTEM_DETAILS_LICENSE_MODEL_UNKNOWN                LaunchDbSystemDetailsLicenseModelEnum = "UNKNOWN"
)

var mapping_launchdbsystemdetails_licenseModel = map[string]LaunchDbSystemDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       LAUNCH_DB_SYSTEM_DETAILS_LICENSE_MODEL_LICENSE_INCLUDED,
	"BRING_YOUR_OWN_LICENSE": LAUNCH_DB_SYSTEM_DETAILS_LICENSE_MODEL_BRING_YOUR_OWN_LICENSE,
	"UNKNOWN":                LAUNCH_DB_SYSTEM_DETAILS_LICENSE_MODEL_UNKNOWN,
}

func GetLaunchDbSystemDetailsLicenseModelEnumValues() []LaunchDbSystemDetailsLicenseModelEnum {
	values := make([]LaunchDbSystemDetailsLicenseModelEnum, 0)
	for _, v := range mapping_launchdbsystemdetails_licenseModel {
		if v != LAUNCH_DB_SYSTEM_DETAILS_LICENSE_MODEL_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
