// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
// Use this pattern to construct endpoints for identity domains: `https://<domainURL>/admin/v1/`. See Finding an Identity Domain URL (https://docs.oracle.com/en-us/iaas/Content/Identity/api-getstarted/locate-identity-domain-url.htm) to locate the domain URL you need.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExtensionDbUserUser DB User extension
type ExtensionDbUserUser struct {

	// If true, indicates this is a database user.
	// **Added In:** 18.2.2
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: boolean
	//  - uniqueness: none
	IsDbUser *bool `mandatory:"false" json:"isDbUser"`

	// Password Verifiers for DB User.
	// **Added In:** 18.2.2
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [type]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	PasswordVerifiers []UserExtPasswordVerifiers `mandatory:"false" json:"passwordVerifiers"`

	// DB domain level schema to which the user is granted access.
	// **Added In:** 18.2.2
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - idcsSensitive: none
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	DomainLevelSchema *string `mandatory:"false" json:"domainLevelSchema"`

	// DB instance level schema to which the user is granted access.
	// **Added In:** 18.2.2
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - idcsSensitive: none
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	InstanceLevelSchema *string `mandatory:"false" json:"instanceLevelSchema"`

	// DB global roles to which the user is granted access.
	// **Added In:** 18.2.2
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - idcsSensitive: none
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	DbGlobalRoles []string `mandatory:"false" json:"dbGlobalRoles"`
}

func (m ExtensionDbUserUser) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExtensionDbUserUser) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
