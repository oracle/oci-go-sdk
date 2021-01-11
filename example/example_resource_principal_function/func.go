// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for resource principle auth

package main

import (
	"context"
	"encoding/json"
	"io"

	fdk "github.com/fnproject/fdk-go"

	"github.com/oracle/oci-go-sdk/v32/common/auth"
	"github.com/oracle/oci-go-sdk/v32/identity"
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
