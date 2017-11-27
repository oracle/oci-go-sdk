// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"encoding/json"
)

type CreateIdentityProviderDetails interface {

	// The OCID of your tenancy.
	GetCompartmentID() *string

	// The name you assign to the `IdentityProvider` during creation.
	// The name must be unique across all `IdentityProvider` objects in the
	// tenancy and cannot be changed.
	GetName() *string

	// The description you assign to the `IdentityProvider` during creation.
	// Does not have to be unique, and it's changeable.
	GetDescription() *string

	// The identity provider service or product.
	// Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft
	// Active Directory Federation Services (ADFS).
	// Example: `IDCS`
	GetProductType() CreateIdentityProviderDetailsProductTypeEnum
}

type createidentityproviderdetails struct {
	CompartmentID *string                                      `mandatory:"true" json:"compartmentId,omitempty"`
	Name          *string                                      `mandatory:"true" json:"name,omitempty"`
	Description   *string                                      `mandatory:"true" json:"description,omitempty"`
	ProductType   CreateIdentityProviderDetailsProductTypeEnum `mandatory:"true" json:"productType,omitempty"`
	Protocol      string                                       `json:"protocol"`
}

func (m *createidentityproviderdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	err := json.Unmarshal(data, m)
	if err != nil {
		return nil, err
	}

	switch m.Protocol {
	case "SAML2":
		mm := CreateSaml2IdentityProviderDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

func (m createidentityproviderdetails) GetCompartmentID() *string {
	return m.CompartmentID
}
func (m createidentityproviderdetails) GetName() *string {
	return m.Name
}
func (m createidentityproviderdetails) GetDescription() *string {
	return m.Description
}
func (m createidentityproviderdetails) GetProductType() CreateIdentityProviderDetailsProductTypeEnum {
	return m.ProductType
}

func (model createidentityproviderdetails) String() string {
	return common.PointerString(model)
}

type CreateIdentityProviderDetailsProductTypeEnum string

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

func GetCreateIdentityProviderDetailsProductTypeEnumValues() []CreateIdentityProviderDetailsProductTypeEnum {
	values := make([]CreateIdentityProviderDetailsProductTypeEnum, 0)
	for _, v := range mapping_createidentityproviderdetails_productType {
		if v != CREATE_IDENTITY_PROVIDER_DETAILS_PRODUCT_TYPE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}
