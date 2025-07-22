// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CustomerInstanceReportRecord The model for instance report item
type CustomerInstanceReportRecord struct {

	// publisher OCID
	PublisherId *string `mandatory:"false" json:"publisherId"`

	// Name of the listing
	ListingName *string `mandatory:"false" json:"listingName"`

	// Listing ID
	ListingId *string `mandatory:"false" json:"listingId"`

	// Listing version ID
	ListingVersionId *string `mandatory:"false" json:"listingVersionId"`

	// The version for the package
	PackageVersion *string `mandatory:"false" json:"packageVersion"`

	// Package version ID
	PackageVersionId *string `mandatory:"false" json:"packageVersionId"`

	// Instance status
	Status *string `mandatory:"false" json:"status"`

	// Instance OCID
	InstanceId *string `mandatory:"false" json:"instanceId"`

	// The date and time that instance was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeInstanceCreation *common.SDKTime `mandatory:"false" json:"timeInstanceCreation"`

	// The date and time that instance was terminated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeInstanceTermination *common.SDKTime `mandatory:"false" json:"timeInstanceTermination"`

	// The region of the instance
	Region *string `mandatory:"false" json:"region"`

	// The realm of the instance
	Realm *string `mandatory:"false" json:"realm"`

	// The shape of the instance
	Shape *string `mandatory:"false" json:"shape"`

	// The tenancy of the instance
	OciTenancy *string `mandatory:"false" json:"ociTenancy"`

	// The name of the tenant adminstrator
	TenantAdminName *string `mandatory:"false" json:"tenantAdminName"`

	// The email of the tenant adminstrator
	TenantAdminEmail *string `mandatory:"false" json:"tenantAdminEmail"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CustomerInstanceReportRecord) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CustomerInstanceReportRecord) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
