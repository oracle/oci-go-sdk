// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

type CreateIdentityProviderDetails struct {

	// The OCID of your tenancy.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The name you assign to the `IdentityProvider` during creation.
	// The name must be unique across all `IdentityProvider` objects in the
	// tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name,omitempty"`

	// The description you assign to the `IdentityProvider` during creation.
	// Does not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description,omitempty"`

	// The identity provider service or product.
	// Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft
	// Active Directory Federation Services (ADFS).
	// Example: `IDCS`
	ProductType CreateIdentityProviderDetailsProductTypeEnum `mandatory:"true" json:"productType,omitempty"`

	// The protocol used for federation.
	// Example: `SAML2`
	Protocol CreateIdentityProviderDetailsProtocolEnum `mandatory:"true" json:"protocol,omitempty"`
}

func (model CreateIdentityProviderDetails) String() string {
	return common.PointerString(model)
}

type CreateIdentityProviderDetailsProductTypeEnum string
type CreateIdentityProviderDetailsProductType struct{}

const (
	CREATE_IDENTITY_PROVIDER_DETAILS_PRODUCT_TYPE_IDCS    CreateIdentityProviderDetailsProductTypeEnum = "IDCS"
	CREATE_IDENTITY_PROVIDER_DETAILS_PRODUCT_TYPE_ADFS    CreateIdentityProviderDetailsProductTypeEnum = "ADFS"
	CREATE_IDENTITY_PROVIDER_DETAILS_PRODUCT_TYPE_UNKNOWN CreateIdentityProviderDetailsProductTypeEnum = "UNKNOWN"
)

var mapping_createidentityproviderdetails_productType = map[string]CreateIdentityProviderDetailsProductTypeEnum{
	"IDCS":    CREATE_IDENTITY_PROVIDER_DETAILS_PRODUCT_TYPE_IDCS,
	"ADFS":    CREATE_IDENTITY_PROVIDER_DETAILS_PRODUCT_TYPE_ADFS,
	"UNKNOWN": CREATE_IDENTITY_PROVIDER_DETAILS_PRODUCT_TYPE_UNKNOWN,
}

func (receiver CreateIdentityProviderDetailsProductType) Values() []CreateIdentityProviderDetailsProductTypeEnum {
	values := make([]CreateIdentityProviderDetailsProductTypeEnum, 0)
	for _, v := range mapping_createidentityproviderdetails_productType {
		if v != CREATE_IDENTITY_PROVIDER_DETAILS_PRODUCT_TYPE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

func (receiver CreateIdentityProviderDetailsProductType) IsValid(toBeChecked string) bool {
	for _, v := range receiver.Values() {
		if CreateIdentityProviderDetailsProductTypeEnum(toBeChecked) == v {
			return true
		}
	}
	return false
}

func (receiver CreateIdentityProviderDetailsProductType) From(toBeConverted string) CreateIdentityProviderDetailsProductTypeEnum {
	if val, ok := mapping_createidentityproviderdetails_productType[toBeConverted]; ok {
		return val
	}
	return CREATE_IDENTITY_PROVIDER_DETAILS_PRODUCT_TYPE_UNKNOWN
}

type CreateIdentityProviderDetailsProtocolEnum string
type CreateIdentityProviderDetailsProtocol struct{}

const (
	CREATE_IDENTITY_PROVIDER_DETAILS_PROTOCOL_SAML2   CreateIdentityProviderDetailsProtocolEnum = "SAML2"
	CREATE_IDENTITY_PROVIDER_DETAILS_PROTOCOL_UNKNOWN CreateIdentityProviderDetailsProtocolEnum = "UNKNOWN"
)

var mapping_createidentityproviderdetails_protocol = map[string]CreateIdentityProviderDetailsProtocolEnum{
	"SAML2":   CREATE_IDENTITY_PROVIDER_DETAILS_PROTOCOL_SAML2,
	"UNKNOWN": CREATE_IDENTITY_PROVIDER_DETAILS_PROTOCOL_UNKNOWN,
}

func (receiver CreateIdentityProviderDetailsProtocol) Values() []CreateIdentityProviderDetailsProtocolEnum {
	values := make([]CreateIdentityProviderDetailsProtocolEnum, 0)
	for _, v := range mapping_createidentityproviderdetails_protocol {
		if v != CREATE_IDENTITY_PROVIDER_DETAILS_PROTOCOL_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

func (receiver CreateIdentityProviderDetailsProtocol) IsValid(toBeChecked string) bool {
	for _, v := range receiver.Values() {
		if CreateIdentityProviderDetailsProtocolEnum(toBeChecked) == v {
			return true
		}
	}
	return false
}

func (receiver CreateIdentityProviderDetailsProtocol) From(toBeConverted string) CreateIdentityProviderDetailsProtocolEnum {
	if val, ok := mapping_createidentityproviderdetails_protocol[toBeConverted]; ok {
		return val
	}
	return CREATE_IDENTITY_PROVIDER_DETAILS_PROTOCOL_UNKNOWN
}
