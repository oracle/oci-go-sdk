// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// DomainGovernanceCollection Result of a query request for a list of domain governance entities. Contains DomainGovernanceSummary items.
type DomainGovernanceCollection struct {

	// Array containing DomainGovernanceSummary items.
	Items []DomainGovernanceSummary `mandatory:"true" json:"items"`
}

func (m DomainGovernanceCollection) String() string {
	return common.PointerString(m)
}
