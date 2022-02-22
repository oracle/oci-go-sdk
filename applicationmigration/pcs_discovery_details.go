// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Migration API
//
// Application Migration simplifies the migration of applications from Oracle Cloud Infrastructure Classic to Oracle Cloud Infrastructure.
// You can use Application Migration API to migrate applications, such as Oracle Java Cloud Service, SOA Cloud Service, and Integration Classic
// instances, to Oracle Cloud Infrastructure. For more information, see
// Overview of Application Migration (https://docs.cloud.oracle.com/iaas/application-migration/appmigrationoverview.htm).
//

package applicationmigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v59/common"
	"strings"
)

// PcsDiscoveryDetails Credentials to access the Oracle Process Cloud Service application in the source environment. Application Migration connects to the
// application in the source environment with the supplied credentials.
type PcsDiscoveryDetails struct {

	// Application administrator username to access the Oracle Process Cloud Service application in the source environment.
	ServiceInstanceUser *string `mandatory:"true" json:"serviceInstanceUser"`

	// Password for this user.
	ServiceInstancePassword *string `mandatory:"true" json:"serviceInstancePassword"`
}

func (m PcsDiscoveryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PcsDiscoveryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PcsDiscoveryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePcsDiscoveryDetails PcsDiscoveryDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypePcsDiscoveryDetails
	}{
		"PCS",
		(MarshalTypePcsDiscoveryDetails)(m),
	}

	return json.Marshal(&s)
}
