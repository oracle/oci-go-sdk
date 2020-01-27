// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// SoacsDiscoveryDetails Specifies the credentials to access the source SOACS instance
type SoacsDiscoveryDetails struct {

	// The SOACS instance weblogic admin user
	WeblogicUser *string `mandatory:"true" json:"weblogicUser"`

	// The SOACS instance weblogic admin password
	WeblogicPassword *string `mandatory:"true" json:"weblogicPassword"`
}

func (m SoacsDiscoveryDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m SoacsDiscoveryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSoacsDiscoveryDetails SoacsDiscoveryDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeSoacsDiscoveryDetails
	}{
		"SOACS",
		(MarshalTypeSoacsDiscoveryDetails)(m),
	}

	return json.Marshal(&s)
}
