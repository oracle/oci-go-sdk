// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v53/common"
)

// UpdateOrganizationDetails UpdateOrganizationDetails contains subscriptionId information.
type UpdateOrganizationDetails struct {

	// OCID of the default Annual Universal Credits subscription. Any tenancy joining the organization will automatically get assigned this subscription if a subscription is not explictly assigned.
	DefaultUcmSubscriptionId *string `mandatory:"true" json:"defaultUcmSubscriptionId"`
}

func (m UpdateOrganizationDetails) String() string {
	return common.PointerString(m)
}
