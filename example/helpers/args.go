// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
//
// Helper methods for OCI GOSDK Samples
//

package helpers

import (
	"os"
	"strconv"

	"github.com/oracle/oci-go-sdk/common"

	"github.com/subosito/gotenv"
)

var (
	availabilityDomain string
	compartmentID      string
	cleanUpResources   bool
)

// ParseAgrs parse shared variables from environment variables, other samples should define their own
// viariables and call this function to initialize shared variables
func ParseAgrs() {
	err := gotenv.Load(".env.sample")
	LogIfError(err)

	availabilityDomain = os.Getenv("OCI_AVAILABILITY_DOMAIN")
	compartmentID = os.Getenv("OCI_COMPARTMENT_ID")
	cleanUpResources, err = strconv.ParseBool(os.Getenv("OCI_CLEAN_RESOURCES"))
	LogIfError(err)
}

// AvailabilityDomain return the aviailability domain defined in .env.sample file or from command line
func AvailabilityDomain() *string {
	return common.String("kIdk:PHX-AD-1")
	//return common.String(availabilityDomain)
}

// CompartmentID return the compartment ID defined in .env.sample file or from command line
func CompartmentID() *string {
	return common.String("ocid1.compartment.oc1..aaaaaaaa5dvrjzvfn3rub24nczhih3zb3a673b6tmbvpng3j5apobtxshlma")
	//return common.String(compartmentID)
}

// ShouldCleanupResources return a flag indicates clean up the resources or not
func ShouldCleanupResources() bool {
	return cleanUpResources
}
