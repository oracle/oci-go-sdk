package identity

import (
	"context"
	"log"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/example/helper"
	"github.com/oracle/oci-go-sdk/identity"
)

// ListAvailabilityDomains Lists the Availability Domains in your tenancy.
// Specify the OCID of either the tenancy or another of your compartments as
// the value for the compartment ID (remember that the tenancy is simply the root compartment).
func ListAvailabilityDomains() []identity.AvailabilityDomain {
	c, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	helper.LogIfError(err)

	// The OCID of the tenancy containing the compartment.
	tenancyID, err := common.DefaultConfigProvider().TenancyOCID()
	helper.LogIfError(err)

	request := identity.ListAvailabilityDomainsRequest{
		CompartmentId: &tenancyID,
	}

	r, err := c.ListAvailabilityDomains(context.Background(), request)
	helper.LogIfError(err)

	if r.Items == nil || len(r.Items) == 0 {
		log.Fatalln("Invalid response from ListAvailabilityDomains!")
	}

	return r.Items
}
