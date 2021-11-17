// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// ChangeRoverEntitlementCompartmentDetails Object for moving an entitlement to a different compartment.
type ChangeRoverEntitlementCompartmentDetails struct {

	// The OCID  (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment into which the resources should be moved.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeRoverEntitlementCompartmentDetails) String() string {
	return common.PointerString(m)
}
