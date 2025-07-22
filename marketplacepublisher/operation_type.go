// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MarketplacePublisherService API
//
// Use the Marketplace Publisher API to manage the publishing of applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplacepublisher

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateOffer                         OperationTypeEnum = "CREATE_OFFER"
	OperationTypeUpdateOffer                         OperationTypeEnum = "UPDATE_OFFER"
	OperationTypeDeleteOffer                         OperationTypeEnum = "DELETE_OFFER"
	OperationTypeMoveOffer                           OperationTypeEnum = "MOVE_OFFER"
	OperationTypeDeleteOfferAttachment               OperationTypeEnum = "DELETE_OFFER_ATTACHMENT"
	OperationTypePublishListingRevisionPackage       OperationTypeEnum = "PUBLISH_LISTING_REVISION_PACKAGE"
	OperationTypeUnpublishListingRevisionPackage     OperationTypeEnum = "UNPUBLISH_LISTING_REVISION_PACKAGE"
	OperationTypePublishListingRevision              OperationTypeEnum = "PUBLISH_LISTING_REVISION"
	OperationTypePublishListingRevisionAsPrivate     OperationTypeEnum = "PUBLISH_LISTING_REVISION_AS_PRIVATE"
	OperationTypeWithdrawListingRevision             OperationTypeEnum = "WITHDRAW_LISTING_REVISION"
	OperationTypeCloneListingRevision                OperationTypeEnum = "CLONE_LISTING_REVISION"
	OperationTypeCascadingDeleteListing              OperationTypeEnum = "CASCADING_DELETE_LISTING"
	OperationTypeCascadingDeleteListingRevision      OperationTypeEnum = "CASCADING_DELETE_LISTING_REVISION"
	OperationTypeMarkAsDefaultListingRevisionPackage OperationTypeEnum = "MARK_AS_DEFAULT_LISTING_REVISION_PACKAGE"
	OperationTypeChangeListingCompartment            OperationTypeEnum = "CHANGE_LISTING_COMPARTMENT"
	OperationTypeCreateArtifact                      OperationTypeEnum = "CREATE_ARTIFACT"
	OperationTypeValidateAndPublishArtifact          OperationTypeEnum = "VALIDATE_AND_PUBLISH_ARTIFACT"
	OperationTypeChangeArtifactCompartment           OperationTypeEnum = "CHANGE_ARTIFACT_COMPARTMENT"
	OperationTypeChangeTermCompartment               OperationTypeEnum = "CHANGE_TERM_COMPARTMENT"
	OperationTypeDeleteArtifact                      OperationTypeEnum = "DELETE_ARTIFACT"
	OperationTypeUpdateArtifact                      OperationTypeEnum = "UPDATE_ARTIFACT"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_OFFER":                             OperationTypeCreateOffer,
	"UPDATE_OFFER":                             OperationTypeUpdateOffer,
	"DELETE_OFFER":                             OperationTypeDeleteOffer,
	"MOVE_OFFER":                               OperationTypeMoveOffer,
	"DELETE_OFFER_ATTACHMENT":                  OperationTypeDeleteOfferAttachment,
	"PUBLISH_LISTING_REVISION_PACKAGE":         OperationTypePublishListingRevisionPackage,
	"UNPUBLISH_LISTING_REVISION_PACKAGE":       OperationTypeUnpublishListingRevisionPackage,
	"PUBLISH_LISTING_REVISION":                 OperationTypePublishListingRevision,
	"PUBLISH_LISTING_REVISION_AS_PRIVATE":      OperationTypePublishListingRevisionAsPrivate,
	"WITHDRAW_LISTING_REVISION":                OperationTypeWithdrawListingRevision,
	"CLONE_LISTING_REVISION":                   OperationTypeCloneListingRevision,
	"CASCADING_DELETE_LISTING":                 OperationTypeCascadingDeleteListing,
	"CASCADING_DELETE_LISTING_REVISION":        OperationTypeCascadingDeleteListingRevision,
	"MARK_AS_DEFAULT_LISTING_REVISION_PACKAGE": OperationTypeMarkAsDefaultListingRevisionPackage,
	"CHANGE_LISTING_COMPARTMENT":               OperationTypeChangeListingCompartment,
	"CREATE_ARTIFACT":                          OperationTypeCreateArtifact,
	"VALIDATE_AND_PUBLISH_ARTIFACT":            OperationTypeValidateAndPublishArtifact,
	"CHANGE_ARTIFACT_COMPARTMENT":              OperationTypeChangeArtifactCompartment,
	"CHANGE_TERM_COMPARTMENT":                  OperationTypeChangeTermCompartment,
	"DELETE_ARTIFACT":                          OperationTypeDeleteArtifact,
	"UPDATE_ARTIFACT":                          OperationTypeUpdateArtifact,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_offer":                             OperationTypeCreateOffer,
	"update_offer":                             OperationTypeUpdateOffer,
	"delete_offer":                             OperationTypeDeleteOffer,
	"move_offer":                               OperationTypeMoveOffer,
	"delete_offer_attachment":                  OperationTypeDeleteOfferAttachment,
	"publish_listing_revision_package":         OperationTypePublishListingRevisionPackage,
	"unpublish_listing_revision_package":       OperationTypeUnpublishListingRevisionPackage,
	"publish_listing_revision":                 OperationTypePublishListingRevision,
	"publish_listing_revision_as_private":      OperationTypePublishListingRevisionAsPrivate,
	"withdraw_listing_revision":                OperationTypeWithdrawListingRevision,
	"clone_listing_revision":                   OperationTypeCloneListingRevision,
	"cascading_delete_listing":                 OperationTypeCascadingDeleteListing,
	"cascading_delete_listing_revision":        OperationTypeCascadingDeleteListingRevision,
	"mark_as_default_listing_revision_package": OperationTypeMarkAsDefaultListingRevisionPackage,
	"change_listing_compartment":               OperationTypeChangeListingCompartment,
	"create_artifact":                          OperationTypeCreateArtifact,
	"validate_and_publish_artifact":            OperationTypeValidateAndPublishArtifact,
	"change_artifact_compartment":              OperationTypeChangeArtifactCompartment,
	"change_term_compartment":                  OperationTypeChangeTermCompartment,
	"delete_artifact":                          OperationTypeDeleteArtifact,
	"update_artifact":                          OperationTypeUpdateArtifact,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_OFFER",
		"UPDATE_OFFER",
		"DELETE_OFFER",
		"MOVE_OFFER",
		"DELETE_OFFER_ATTACHMENT",
		"PUBLISH_LISTING_REVISION_PACKAGE",
		"UNPUBLISH_LISTING_REVISION_PACKAGE",
		"PUBLISH_LISTING_REVISION",
		"PUBLISH_LISTING_REVISION_AS_PRIVATE",
		"WITHDRAW_LISTING_REVISION",
		"CLONE_LISTING_REVISION",
		"CASCADING_DELETE_LISTING",
		"CASCADING_DELETE_LISTING_REVISION",
		"MARK_AS_DEFAULT_LISTING_REVISION_PACKAGE",
		"CHANGE_LISTING_COMPARTMENT",
		"CREATE_ARTIFACT",
		"VALIDATE_AND_PUBLISH_ARTIFACT",
		"CHANGE_ARTIFACT_COMPARTMENT",
		"CHANGE_TERM_COMPARTMENT",
		"DELETE_ARTIFACT",
		"UPDATE_ARTIFACT",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
