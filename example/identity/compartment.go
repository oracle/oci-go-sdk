package identity

import (
	"context"

	"github.com/oracle/oci-go-sdk/example/helper"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/example"
	"github.com/oracle/oci-go-sdk/identity"
)

// CreateCompartment creates a compartment and return it's ID
// if the compartment already exists just return the id of it
func CreateCompartment() (compartmentID *string, err error) {
	c, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	helper.LogIfError(err)

	// The OCID of the tenancy containing the compartment.
	tenancyID, err := common.DefaultConfigProvider().TenancyOCID()
	helper.LogIfError(err)

	request := identity.CreateCompartmentRequest{
		CreateCompartmentDetails: identity.CreateCompartmentDetails{
			Name:          common.String(example.Config.CompartmentName),
			Description:   common.String("OCI Go SDK Sample Comparment Test"),
			CompartmentId: &tenancyID,
		}}

	response, _ := c.CreateCompartment(context.Background(), request)
	compartmentID = response.CompartmentId

	// if compartment already exist, return the id of it
	if response.RawResponse.StatusCode == 409 {
		compartments := ListCompartments()

		// get the id of compartment which already been created
		for _, element := range compartments {
			if *element.Name == example.Config.CompartmentName {
				compartmentID = element.Id
				err = nil // ignore the error, just return the compartmentID
				break
			}
		}
	}

	return
}

// ListCompartments does list all compartments under your tenant
func ListCompartments() []identity.Compartment {
	c, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	helper.LogIfError(err)

	// The OCID of the tenancy containing the compartment.
	tenancyID, err := common.DefaultConfigProvider().TenancyOCID()
	helper.LogIfError(err)

	request := identity.ListCompartmentsRequest{
		CompartmentId: &tenancyID,
		Limit:         common.Int(100),
	}

	response, err := c.ListCompartments(context.Background(), request)
	helper.LogIfError(err)

	return response.Items
}

// GetCompartment does return a compartment details
func GetCompartment(compartmentID string) identity.Compartment {
	c, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	helper.LogIfError(err)

	// The OCID of the tenancy containing the compartment.
	request := identity.GetCompartmentRequest{CompartmentId: &compartmentID}

	response, err := c.GetCompartment(context.Background(), request)
	helper.LogIfError(err)

	return response.Compartment
}
