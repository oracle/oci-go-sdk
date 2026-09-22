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

// OpnMembership OPN membership information
type OpnMembership struct {

	// OPN membership start date. An RFC3339 formatted datetime string
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// OPN membership end date. An RFC3339 formatted datetime string
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// OPN status
	OpnStatus AdminOpnPartnerSummaryMembershipStatusEnum `mandatory:"false" json:"opnStatus,omitempty"`

	// OPN Number number
	OpnNumber *string `mandatory:"false" json:"opnNumber"`

	// OPN membership type
	OpnMembershipType *string `mandatory:"false" json:"opnMembershipType"`
}

func (m OpnMembership) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OpnMembership) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAdminOpnPartnerSummaryMembershipStatusEnum(string(m.OpnStatus)); !ok && m.OpnStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OpnStatus: %s. Supported values are: %s.", m.OpnStatus, strings.Join(GetAdminOpnPartnerSummaryMembershipStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
