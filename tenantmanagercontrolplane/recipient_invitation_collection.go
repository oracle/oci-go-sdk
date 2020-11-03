// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// TenantManager API
//
// A description of the TenantManager API
//

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v27/common"
)

// RecipientInvitationCollection Result of a query request for a list of recipient invitations. Contains RecipientInvitationSummary items.
type RecipientInvitationCollection struct {

	// Array containing RecipientInvitationSummary items.
	Items []RecipientInvitationSummary `mandatory:"true" json:"items"`
}

func (m RecipientInvitationCollection) String() string {
	return common.PointerString(m)
}
