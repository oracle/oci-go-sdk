// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ApplianceExportJob The representation of ApplianceExportJob
type ApplianceExportJob struct {
	Id *string `mandatory:"true" json:"id"`

	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	BucketName *string `mandatory:"false" json:"bucketName"`

	DisplayName *string `mandatory:"false" json:"displayName"`

	CreationTime *common.SDKTime `mandatory:"false" json:"creationTime"`

	LifecycleState ApplianceExportJobLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A property that can contain details on the lifecycle.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Serial Number of the Appliance associated with this Export Job.
	ApplianceSerialNumber *string `mandatory:"false" json:"applianceSerialNumber"`

	// Passphrase associated with the Appliance.
	ApplianceDecryptionPassphrase *string `mandatory:"false" json:"applianceDecryptionPassphrase"`

	// Shipping Vendor selected to ship the Appliance associated with this job.
	ApplianceDeliveryVendor *string `mandatory:"false" json:"applianceDeliveryVendor"`

	// Tracking number associated with the shipment while shipping the Appliance to Customer.
	ApplianceDeliveryTrackingNumber *string `mandatory:"false" json:"applianceDeliveryTrackingNumber"`

	// Tracking number associated with the shipment while shipping the Appliance back to Oracle.
	ApplianceReturnDeliveryTrackingNumber *string `mandatory:"false" json:"applianceReturnDeliveryTrackingNumber"`

	// Unique number associated with the security tie used to seal the Appliance case.
	SendingSecurityTie *string `mandatory:"false" json:"sendingSecurityTie"`

	// Unique number associated with the return security tie used to seal the Appliance case.
	ReceivingSecurityTie *string `mandatory:"false" json:"receivingSecurityTie"`

	// List of objects with names matching this prefix would be part of this export job.
	Prefix *string `mandatory:"false" json:"prefix"`

	// The name of the first object in the range of objects that are expected to be part of this export job.
	RangeStart *string `mandatory:"false" json:"rangeStart"`

	// The name of the last object in the range of objects that are expected to be part of this export job.
	RangeEnd *string `mandatory:"false" json:"rangeEnd"`

	// Total number of objects that are exported in this job.
	NumberOfObjects *string `mandatory:"false" json:"numberOfObjects"`

	// Total size of objects in Bytes that are exported in this job.
	TotalSizeInBytes *string `mandatory:"false" json:"totalSizeInBytes"`

	// First object in the list of objects that are exported in this job.
	FirstObject *string `mandatory:"false" json:"firstObject"`

	// Last object in the list of objects that are exported in this job.
	LastObject *string `mandatory:"false" json:"lastObject"`

	// First object from which the next potential export job could start.
	NextObject *string `mandatory:"false" json:"nextObject"`

	// Url of the Manifest File associated with this export job.
	ManifestFile *string `mandatory:"false" json:"manifestFile"`

	// md5 digest of the manifest file.
	ManifestMd5 *string `mandatory:"false" json:"manifestMd5"`

	// Polices to grant Data Transfer Service to access objects in the Bucket
	BucketAccessPolicies []string `mandatory:"false" json:"bucketAccessPolicies"`

	CustomerShippingAddress *ShippingAddress `mandatory:"false" json:"customerShippingAddress"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ApplianceExportJob) String() string {
	return common.PointerString(m)
}

// ApplianceExportJobLifecycleStateEnum Enum with underlying type: string
type ApplianceExportJobLifecycleStateEnum string

// Set of constants representing the allowable values for ApplianceExportJobLifecycleStateEnum
const (
	ApplianceExportJobLifecycleStateCreating   ApplianceExportJobLifecycleStateEnum = "CREATING"
	ApplianceExportJobLifecycleStateActive     ApplianceExportJobLifecycleStateEnum = "ACTIVE"
	ApplianceExportJobLifecycleStateInprogress ApplianceExportJobLifecycleStateEnum = "INPROGRESS"
	ApplianceExportJobLifecycleStateSucceeded  ApplianceExportJobLifecycleStateEnum = "SUCCEEDED"
	ApplianceExportJobLifecycleStateFailed     ApplianceExportJobLifecycleStateEnum = "FAILED"
	ApplianceExportJobLifecycleStateCancelled  ApplianceExportJobLifecycleStateEnum = "CANCELLED"
	ApplianceExportJobLifecycleStateDeleted    ApplianceExportJobLifecycleStateEnum = "DELETED"
)

var mappingApplianceExportJobLifecycleState = map[string]ApplianceExportJobLifecycleStateEnum{
	"CREATING":   ApplianceExportJobLifecycleStateCreating,
	"ACTIVE":     ApplianceExportJobLifecycleStateActive,
	"INPROGRESS": ApplianceExportJobLifecycleStateInprogress,
	"SUCCEEDED":  ApplianceExportJobLifecycleStateSucceeded,
	"FAILED":     ApplianceExportJobLifecycleStateFailed,
	"CANCELLED":  ApplianceExportJobLifecycleStateCancelled,
	"DELETED":    ApplianceExportJobLifecycleStateDeleted,
}

// GetApplianceExportJobLifecycleStateEnumValues Enumerates the set of values for ApplianceExportJobLifecycleStateEnum
func GetApplianceExportJobLifecycleStateEnumValues() []ApplianceExportJobLifecycleStateEnum {
	values := make([]ApplianceExportJobLifecycleStateEnum, 0)
	for _, v := range mappingApplianceExportJobLifecycleState {
		values = append(values, v)
	}
	return values
}
