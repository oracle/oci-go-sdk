// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MarketplacePublisherService API
//
// Use the Marketplace Publisher API to manage the publishing of applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplacepublisher

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PartnerPortalOptions Partner Portal Options.
type PartnerPortalOptions struct {

	// Enable Oracle Cloud Marketplace Publisher Agreement validity check for partner before listing is published.
	IsEnablePublisherAgreementCheckBeforeListingPublish *bool `mandatory:"false" json:"isEnablePublisherAgreementCheckBeforeListingPublish"`

	// Enable OPN membership status check for partner before listing is published.
	IsEnableOpnMembershipStatusCheckBeforeListingPublish *bool `mandatory:"false" json:"isEnableOpnMembershipStatusCheckBeforeListingPublish"`

	// Enable the Reports feature for this Partner.
	IsEnableReports *bool `mandatory:"false" json:"isEnableReports"`

	// Enable the private listing feature for this Partner.
	IsEnablePrivateListing *bool `mandatory:"false" json:"isEnablePrivateListing"`

	// Enable Auto-publish on listing submit for this partner for Lead generation listings.
	IsEnableAutoPublishOnSubmitForLeadGenListings *bool `mandatory:"false" json:"isEnableAutoPublishOnSubmitForLeadGenListings"`

	// Disable Auto-publish option visibility for listing submit flow. Defaults to false when unset.
	IsDisableAutoPublishOnSubmit *bool `mandatory:"false" json:"isDisableAutoPublishOnSubmit"`

	// Enable Auto approval options for new listings - Include OCI Application listings with Image artifact.
	IsEnableAutoApprovalForImageArtifactNewAppListings *bool `mandatory:"false" json:"isEnableAutoApprovalForImageArtifactNewAppListings"`

	// Enable Auto approval options for new listings - Include OCI Application listings with Terraform artifact.
	IsEnableAutoApprovalForTerraformArtifactNewAppListings *bool `mandatory:"false" json:"isEnableAutoApprovalForTerraformArtifactNewAppListings"`

	// Enable Auto approval options for new listings - Include OCI Paid Application listings.
	IsEnableAutoApprovalForNewPaidAppListings *bool `mandatory:"false" json:"isEnableAutoApprovalForNewPaidAppListings"`

	// Enable Auto approval options for new listings - Include Lead generation listings.
	IsEnableAutoApprovalForNewLeadGenListings *bool `mandatory:"false" json:"isEnableAutoApprovalForNewLeadGenListings"`

	// Disable Auto approval options for versioned listings - Include OCI Application listings with Image artifact.
	IsDisableAutoApprovalForImageArtifactVersionedAppListings *bool `mandatory:"false" json:"isDisableAutoApprovalForImageArtifactVersionedAppListings"`

	// Disable Auto approval options for versioned listings - Include OCI Application listings with Terraform artifact.
	IsDisableAutoApprovalForTerraformArtifactVersionedAppListings *bool `mandatory:"false" json:"isDisableAutoApprovalForTerraformArtifactVersionedAppListings"`

	// Disable Auto approval options for versioned listings - Include OCI Application listings with Container artifact.
	IsDisableAutoApprovalForContainerArtifactVersionedAppListings *bool `mandatory:"false" json:"isDisableAutoApprovalForContainerArtifactVersionedAppListings"`

	// Disable Auto approval options for versioned listings - Include OCI Application listings with Helm chart artifact.
	IsDisableAutoApprovalForHelmChartVersionedAppListings *bool `mandatory:"false" json:"isDisableAutoApprovalForHelmChartVersionedAppListings"`

	// Enable Auto approval options for versioned listings - Include Service listings.
	IsEnableAutoApprovalForServiceVersionedListings *bool `mandatory:"false" json:"isEnableAutoApprovalForServiceVersionedListings"`

	// Enable Auto approval options for versioned listings - Include Fusion AiAgent listings.
	IsEnableAutoApprovalForFusionAiAgentVersionedListings *bool `mandatory:"false" json:"isEnableAutoApprovalForFusionAiAgentVersionedListings"`

	// Enable Auto approval options for versioned listings - Include Lead generation listings.
	IsEnableAutoApprovalForLeadGenVersionedListings *bool `mandatory:"false" json:"isEnableAutoApprovalForLeadGenVersionedListings"`

	// Disable Auto approval options for versioned listings - Include SaaS listings.
	IsDisableAutoApprovalForSaaSVersionedListings *bool `mandatory:"false" json:"isDisableAutoApprovalForSaaSVersionedListings"`

	// Disable Auto approval options for versioned listings - Include Paid App listings.
	IsDisableAutoApprovalForPaidAppVersionedListings *bool `mandatory:"false" json:"isDisableAutoApprovalForPaidAppVersionedListings"`
}

func (m PartnerPortalOptions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PartnerPortalOptions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
