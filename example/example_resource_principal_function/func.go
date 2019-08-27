package main

import (
	"context"
	"encoding/json"
	"io"

	fdk "github.com/fnproject/fdk-go"

	"github.com/oracle/oci-go-sdk/common/auth"
	"github.com/oracle/oci-go-sdk/identity"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))
}

type Message struct {
	Region       string   `json:"region"`
	Tenancy      string   `json:"tenant"`
	Compartment  string   `json:"compartment"`
	Compartments []string `json:"compartments"`
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {
	var err error
	var msg Message
	var compartments []string

	rp, err := auth.ResourcePrincipalConfigurationProvider()
	if err != nil {
		panic(err)
	}

	tenancy, _ := rp.TenancyOCID()
	compartment, _ := rp.GetClaim(auth.CompartmentOCIDClaimKey)
	region, _ := rp.Region()

	iam, err := identity.NewIdentityClientWithConfigurationProvider(rp)
	if err != nil {
		panic(err)
	}

	// As an example, use the resource principal to construct an identity client and exercise that.
	yes := true
	resp, err := iam.ListCompartments(context.Background(),
		identity.ListCompartmentsRequest{CompartmentId: &tenancy, CompartmentIdInSubtree: &yes})
	if err != nil {
		panic(err)
	}

	for _, cc := range resp.Items {
		compartments = append(compartments, *cc.Name)
	}

	msg = Message{
		Region:       region,
		Tenancy:      tenancy,
		Compartment:  compartment.(string),
		Compartments: compartments,
	}

	json.NewEncoder(out).Encode(&msg)
}
