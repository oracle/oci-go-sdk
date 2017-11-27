// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oci-go-sdk/common"
	"encoding/json"
)

// CreateSaml2IdentityProviderDetails.
type CreateSaml2IdentityProviderDetails struct {

	// The OCID of your tenancy.
	CompartmentID *string `mandatory:"true" json:"compartmentId,omitempty"`

	// The name you assign to the `IdentityProvider` during creation.
	// The name must be unique across all `IdentityProvider` objects in the
	// tenancy and cannot be changed.
	Name *string `mandatory:"true" json:"name,omitempty"`

	// The description you assign to the `IdentityProvider` during creation.
	// Does not have to be unique, and it's changeable.
	Description *string `mandatory:"true" json:"description,omitempty"`

	// The URL for retrieving the identity provider's metadata,
	// which contains information required for federating.
	MetadataURL *string `mandatory:"true" json:"metadataUrl,omitempty"`

	// The XML that contains the information required for federating.
	Metadata *string `mandatory:"true" json:"metadata,omitempty"`

	// The identity provider service or product.
	// Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft
	// Active Directory Federation Services (ADFS).
	// Example: `IDCS`
	ProductType CreateIdentityProviderDetailsProductTypeEnum `mandatory:"true" json:"producttype,omitempty"`
}

func (model CreateSaml2IdentityProviderDetails) GetCompartmentID() *string {
	return model.CompartmentID
}
func (model CreateSaml2IdentityProviderDetails) GetName() *string {
	return model.Name
}
func (model CreateSaml2IdentityProviderDetails) GetDescription() *string {
	return model.Description
}
func (model CreateSaml2IdentityProviderDetails) GetProductType() CreateIdentityProviderDetailsProductTypeEnum {
	return model.ProductType
}

func (m CreateSaml2IdentityProviderDetails) String() string {
	return common.PointerString(m)
}

func (m CreateSaml2IdentityProviderDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateSaml2IdentityProviderDetails CreateSaml2IdentityProviderDetails
	s := struct {
		DiscriminatorParam string `json:"protocol"`
		MarshalTypeCreateSaml2IdentityProviderDetails
	}{
		"SAML2",
		(MarshalTypeCreateSaml2IdentityProviderDetails)(m),
	}

	return json.Marshal(&s)
}
