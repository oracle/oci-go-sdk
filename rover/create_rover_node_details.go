// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateRoverNodeDetails The information required to create a RoverNode.
type CreateRoverNodeDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment containing the RoverNode.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The shape of the node.
	Shape *string `mandatory:"false" json:"shape"`

	CustomerShippingAddress *ShippingAddress `mandatory:"false" json:"customerShippingAddress"`

	// List of existing workloads that should be provisioned on the node.
	NodeWorkloads []RoverWorkload `mandatory:"false" json:"nodeWorkloads"`

	// Root password for the rover node.
	SuperUserPassword *string `mandatory:"false" json:"superUserPassword"`

	// Passphrase to unlock the rover node.
	UnlockPassphrase *string `mandatory:"false" json:"unlockPassphrase"`

	// Name of point of contact for this order if customer is picking up.
	PointOfContact *string `mandatory:"false" json:"pointOfContact"`

	// Phone number of point of contact for this order if customer is picking up.
	PointOfContactPhoneNumber *string `mandatory:"false" json:"pointOfContactPhoneNumber"`

	// Preference for device delivery.
	ShippingPreference CreateRoverNodeDetailsShippingPreferenceEnum `mandatory:"false" json:"shippingPreference,omitempty"`

	// Shipping vendor of choice for orace to customer shipping.
	ShippingVendor *string `mandatory:"false" json:"shippingVendor"`

	// Expected date when customer wants to pickup the device if they chose customer pickup.
	TimePickupExpected *common.SDKTime `mandatory:"false" json:"timePickupExpected"`

	// The public key of the resource principal
	PublicKey *string `mandatory:"false" json:"publicKey"`

	// Start time for the window to pickup the device from customer.
	TimeReturnWindowStarts *common.SDKTime `mandatory:"false" json:"timeReturnWindowStarts"`

	// End time for the window to pickup the device from customer.
	TimeReturnWindowEnds *common.SDKTime `mandatory:"false" json:"timeReturnWindowEnds"`

	// The current state of the RoverNode.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The type of enclosure rover nodes in this cluster are shipped in.
	EnclosureType EnclosureTypeEnum `mandatory:"false" json:"enclosureType,omitempty"`

	// A property that can contain details on the lifecycle.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// Serial number of the node.
	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	// Tracking Url for the shipped FmsRoverNode.
	OracleShippingTrackingUrl *string `mandatory:"false" json:"oracleShippingTrackingUrl"`

	// The flag indicating that customer requests data to be imported to OCI upon Rover node return.
	IsImportRequested *bool `mandatory:"false" json:"isImportRequested"`

	// An OCID of a compartment where data will be imported to upon Rover node return.
	ImportCompartmentId *string `mandatory:"false" json:"importCompartmentId"`

	// Name of a bucket where files from NFS share will be imported to upon Rover node return.
	ImportFileBucket *string `mandatory:"false" json:"importFileBucket"`

	// Validation code returned by data validation tool. Required for return shipping label generation if data import was requested.
	DataValidationCode *string `mandatory:"false" json:"dataValidationCode"`

	// Customer provided master key ID to encrypt secret information. If not provided, Rover's master key will be used for encryption.
	MasterKeyId *string `mandatory:"false" json:"masterKeyId"`

	// The certificateAuthorityId of subordinate/intermediate certificate authority.
	CertificateAuthorityId *string `mandatory:"false" json:"certificateAuthorityId"`

	// The time after which leaf certificate will invalid.
	TimeCertValidityEnd *common.SDKTime `mandatory:"false" json:"timeCertValidityEnd"`

	// The common name for the leaf certificate.
	CommonName *string `mandatory:"false" json:"commonName"`

	// The compartmentId of the leaf certificate.
	CertCompartmentId *string `mandatory:"false" json:"certCompartmentId"`

	// key algorithm for issuing leaf certificate.
	CertKeyAlgorithm CertKeyAlgorithmEnum `mandatory:"false" json:"certKeyAlgorithm,omitempty"`

	// signature algorithm for issuing leaf certificate.
	CertSignatureAlgorithm CertSignatureAlgorithmEnum `mandatory:"false" json:"certSignatureAlgorithm,omitempty"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CreateRoverNodeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateRoverNodeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateRoverNodeDetailsShippingPreferenceEnum(string(m.ShippingPreference)); !ok && m.ShippingPreference != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShippingPreference: %s. Supported values are: %s.", m.ShippingPreference, strings.Join(GetCreateRoverNodeDetailsShippingPreferenceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEnclosureTypeEnum(string(m.EnclosureType)); !ok && m.EnclosureType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnclosureType: %s. Supported values are: %s.", m.EnclosureType, strings.Join(GetEnclosureTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCertKeyAlgorithmEnum(string(m.CertKeyAlgorithm)); !ok && m.CertKeyAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertKeyAlgorithm: %s. Supported values are: %s.", m.CertKeyAlgorithm, strings.Join(GetCertKeyAlgorithmEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCertSignatureAlgorithmEnum(string(m.CertSignatureAlgorithm)); !ok && m.CertSignatureAlgorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertSignatureAlgorithm: %s. Supported values are: %s.", m.CertSignatureAlgorithm, strings.Join(GetCertSignatureAlgorithmEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateRoverNodeDetailsShippingPreferenceEnum Enum with underlying type: string
type CreateRoverNodeDetailsShippingPreferenceEnum string

// Set of constants representing the allowable values for CreateRoverNodeDetailsShippingPreferenceEnum
const (
	CreateRoverNodeDetailsShippingPreferenceOracleShipped  CreateRoverNodeDetailsShippingPreferenceEnum = "ORACLE_SHIPPED"
	CreateRoverNodeDetailsShippingPreferenceCustomerPickup CreateRoverNodeDetailsShippingPreferenceEnum = "CUSTOMER_PICKUP"
)

var mappingCreateRoverNodeDetailsShippingPreferenceEnum = map[string]CreateRoverNodeDetailsShippingPreferenceEnum{
	"ORACLE_SHIPPED":  CreateRoverNodeDetailsShippingPreferenceOracleShipped,
	"CUSTOMER_PICKUP": CreateRoverNodeDetailsShippingPreferenceCustomerPickup,
}

var mappingCreateRoverNodeDetailsShippingPreferenceEnumLowerCase = map[string]CreateRoverNodeDetailsShippingPreferenceEnum{
	"oracle_shipped":  CreateRoverNodeDetailsShippingPreferenceOracleShipped,
	"customer_pickup": CreateRoverNodeDetailsShippingPreferenceCustomerPickup,
}

// GetCreateRoverNodeDetailsShippingPreferenceEnumValues Enumerates the set of values for CreateRoverNodeDetailsShippingPreferenceEnum
func GetCreateRoverNodeDetailsShippingPreferenceEnumValues() []CreateRoverNodeDetailsShippingPreferenceEnum {
	values := make([]CreateRoverNodeDetailsShippingPreferenceEnum, 0)
	for _, v := range mappingCreateRoverNodeDetailsShippingPreferenceEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateRoverNodeDetailsShippingPreferenceEnumStringValues Enumerates the set of values in String for CreateRoverNodeDetailsShippingPreferenceEnum
func GetCreateRoverNodeDetailsShippingPreferenceEnumStringValues() []string {
	return []string{
		"ORACLE_SHIPPED",
		"CUSTOMER_PICKUP",
	}
}

// GetMappingCreateRoverNodeDetailsShippingPreferenceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateRoverNodeDetailsShippingPreferenceEnum(val string) (CreateRoverNodeDetailsShippingPreferenceEnum, bool) {
	enum, ok := mappingCreateRoverNodeDetailsShippingPreferenceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
