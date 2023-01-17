// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JcsDiscoveryDetails Credentials to access the Oracle Java Cloud Service application in the source environment. When you create and update a migration,
// Application Migration connects to the application in the source environment with the supplied credentials and exports the domain
// configuration.
type JcsDiscoveryDetails struct {

	// WebLogic administrator username for the Oracle Java Cloud Service application in the source environment.
	WeblogicUser *string `mandatory:"true" json:"weblogicUser"`

	// The password of the WebLogic administrator for the Oracle Java Cloud Service application in the source environment.
	WeblogicPassword *string `mandatory:"true" json:"weblogicPassword"`
}

func (m JcsDiscoveryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JcsDiscoveryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m JcsDiscoveryDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeJcsDiscoveryDetails JcsDiscoveryDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeJcsDiscoveryDetails
	}{
		"JCS",
		(MarshalTypeJcsDiscoveryDetails)(m),
	}

	return json.Marshal(&s)
}
