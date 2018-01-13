// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateIdentityProviderDetails is an interface representing the polymorphic json shape of this model
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
	JsonData      []byte
	CompartmentID *string                                      `mandatory:"true" json:"compartmentId,omitempty"`
	Name          *string                                      `mandatory:"true" json:"name,omitempty"`
	Description   *string                                      `mandatory:"true" json:"description,omitempty"`
	ProductType   CreateIdentityProviderDetailsProductTypeEnum `mandatory:"true" json:"productType,omitempty"`
	Protocol      string                                       `json:"protocol"`
}

// UnmarshalJSON unmarshals json
func (m *createidentityproviderdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateidentityproviderdetails createidentityproviderdetails
	s := struct {
		Model Unmarshalercreateidentityproviderdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentID = s.Model.CompartmentID
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.ProductType = s.Model.ProductType
	m.Protocol = s.Model.Protocol

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createidentityproviderdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	var err error
	switch m.Protocol {
	case "SAML2":
		mm := CreateSaml2IdentityProviderDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

//GetCompartmentID returns CompartmentID
func (m createidentityproviderdetails) GetCompartmentID() *string {
	return m.CompartmentID
}

//GetName returns Name
func (m createidentityproviderdetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m createidentityproviderdetails) GetDescription() *string {
	return m.Description
}

//GetProductType returns ProductType
func (m createidentityproviderdetails) GetProductType() CreateIdentityProviderDetailsProductTypeEnum {
	return m.ProductType
}

func (m createidentityproviderdetails) String() string {
	return common.PointerString(m)
}

// CreateIdentityProviderDetailsProductTypeEnum Enum with underlying type: string
type CreateIdentityProviderDetailsProductTypeEnum string

// Set of constants representing the allowable values for CreateIdentityProviderDetailsProductType
const (
	CreateIdentityProviderDetailsProductTypeIdcs    CreateIdentityProviderDetailsProductTypeEnum = "IDCS"
	CreateIdentityProviderDetailsProductTypeAdfs    CreateIdentityProviderDetailsProductTypeEnum = "ADFS"
	CreateIdentityProviderDetailsProductTypeUnknown CreateIdentityProviderDetailsProductTypeEnum = "UNKNOWN"
)

var mappingCreateIdentityProviderDetailsProductType = map[string]CreateIdentityProviderDetailsProductTypeEnum{
	"IDCS":    CreateIdentityProviderDetailsProductTypeIdcs,
	"ADFS":    CreateIdentityProviderDetailsProductTypeAdfs,
	"UNKNOWN": CreateIdentityProviderDetailsProductTypeUnknown,
}

// GetCreateIdentityProviderDetailsProductTypeEnumValues Enumerates the set of values for CreateIdentityProviderDetailsProductType
func GetCreateIdentityProviderDetailsProductTypeEnumValues() []CreateIdentityProviderDetailsProductTypeEnum {
	values := make([]CreateIdentityProviderDetailsProductTypeEnum, 0)
	for _, v := range mappingCreateIdentityProviderDetailsProductType {
		if v != CreateIdentityProviderDetailsProductTypeUnknown {
			values = append(values, v)
		}
	}
	return values
}
