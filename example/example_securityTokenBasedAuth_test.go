// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for using security token based auth

package example

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/oracle/oci-go-sdk/identity"
	"log"
	"os"
	"os/user"
	"path"
)

// This test requires that you have a specified profile setup for security-token based authentication, for the detail
// how to set up the configuration, please refer https://docs.cloud.oracle.com/en-us/iaas/Content/API/SDKDocs/clitoken.htm?Highlight=security_token_file
// In this example the [security_token_based_auth] is created and lists Lists the Availability Domains in current tenancy.
// The token will expire in 1 hour, user needs to check if the token is valid and then decide refresh steps via OCI CLI command.

const (
	profileName = "security_token_based_auth"
	cfgDirName  = ".oci"
	cfgFileName = "config"
)

func ExampleCreateAndUseSecurityTokenBasedConfiguration() {
	homeFolder := getHomeFolder()
	configFilePath := path.Join(homeFolder, cfgDirName, cfgFileName)
	securityTokenBasedAuthConfigProvider := common.CustomProfileConfigProvider(configFilePath, profileName)
	c, err := identity.NewIdentityClientWithConfigurationProvider(securityTokenBasedAuthConfigProvider)
	helpers.FatalIfError(err)

	// The OCID of the tenancy containing the compartment.
	tenancyID, err := securityTokenBasedAuthConfigProvider.TenancyOCID()
	helpers.FatalIfError(err)

	request := identity.ListAvailabilityDomainsRequest{
		CompartmentId: &tenancyID,
	}

	r, err := c.ListAvailabilityDomains(context.Background(), request)
	helpers.FatalIfError(err)

	log.Printf("list of available domains: %v", r.Items)
	fmt.Println("list available domains completed")

	// Output:
	// list available domains completed
}

func getHomeFolder() string {
	current, e := user.Current()
	if e != nil {
		//Give up and try to return something sensible
		home := os.Getenv("HOME")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return current.HomeDir
}
