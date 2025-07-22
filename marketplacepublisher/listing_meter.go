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

// ListingMeter A listing meter information attached by marketplace admin.
type ListingMeter struct {

	// The meter name, ex - MP_BOBO_AP.
	Name *string `mandatory:"true" json:"name"`

	// rate allocation, these are calculated based on rate information at parent part/sku and listing revision.
	RateAllocation *float32 `mandatory:"true" json:"rateAllocation"`

	// Additional metadata key/value pairs for the listing meter.
	// For example:
	// `{"pausedOnInstanceStop": "True","coreSuffixMeter": "_VMWARE", "minimumBillingPeriodInHours": "755", "weight": "1.0" }`
	ExtendedMetadata map[string]string `mandatory:"true" json:"extendedMetadata"`
}

func (m ListingMeter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ListingMeter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
