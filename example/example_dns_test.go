// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for DNS API

package example

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/oracle/oci-go-sdk/dns"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/example/helpers"
)

// ExampleDnsZone creates, gets, lists, and deletes a DNS Zone.
// If optional TARGET_COMPARTMENT_ID env var is set, it will also move the DNS Zone to the compartment.
func ExampleDnsZone() {
	client, err := dns.NewDnsClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Optional parsed value from env var TARGET_COMPARTMENT_ID
	targetCompartmentId := os.Getenv("TARGET_COMPARTMENT_ID")
	log.Printf("TARGET_COMPARTMENT_ID: %s", targetCompartmentId)

	ctx := context.Background()

	// Create a new zone
	zoneName := common.String("testdomain." + helpers.GetRandomString(15))
	createReq := dns.CreateZoneRequest{
		CreateZoneDetails: dns.CreateZoneDetails{
			CompartmentId: helpers.CompartmentID(),
			Name:          zoneName,
			ZoneType:      dns.CreateZoneDetailsZoneTypePrimary,
		},
	}
	createResp, err := client.CreateZone(ctx, createReq)
	helpers.FatalIfError(err)
	fmt.Printf("created dns zone %s", *zoneName)

	// below logic is to wait until zone is in active state
	pollUntilAvailable := func(r common.OCIOperationResponse) bool {
		if converted, ok := r.Response.(dns.GetZoneResponse); ok {
			return converted.LifecycleState != dns.ZoneLifecycleStateActive
		}
		return true
	}
	getRequest := dns.GetZoneRequest{
		ZoneNameOrId:    createResp.Id,
		CompartmentId:   helpers.CompartmentID(),
		RequestMetadata: helpers.GetRequestMetadataWithCustomizedRetryPolicy(pollUntilAvailable),
	}
	getResp, err := client.GetZone(ctx, getRequest)
	helpers.FatalIfError(err)
	fmt.Printf("get dns zone %s", *zoneName)

	listResp, err := client.ListZones(ctx, dns.ListZonesRequest{
		CompartmentId: helpers.CompartmentID(),
		NameContains:  zoneName,
		Limit:         common.Int64(10),
		SortBy:        dns.ListZonesSortByTimecreated,
		SortOrder:     dns.ListZonesSortOrderAsc,
	})
	helpers.FatalIfError(err)
	fmt.Println("list dns zones")
	log.Printf("count of dns zones in compartment %s: %d", *helpers.CompartmentID(), listResp.OpcTotalItems)

	if targetCompartmentId != "" && targetCompartmentId != *helpers.CompartmentID() {
		changeRequest := dns.ChangeZoneCompartmentRequest{
			ZoneId: getResp.Id,
			ChangeZoneCompartmentDetails: dns.ChangeZoneCompartmentDetails{
				CompartmentId: &targetCompartmentId,
			},
		}
		_, err := client.ChangeZoneCompartment(ctx, changeRequest)
		helpers.FatalIfError(err)
		fmt.Printf("change dns zone compartment to %s", targetCompartmentId)
	}

	// Clean up
	defer func() {
		_, err = client.DeleteZone(ctx, dns.DeleteZoneRequest{
			ZoneNameOrId: getResp.Id,
		})
		helpers.FatalIfError(err)
	}()

	// Output:
	// created dns zone
	// get dns zone
	// list dns zone
	// deleted dns zone
}

// ExampleSteeringPolicy creates, gets, lists, and deletes a DNS Steering Policy.
// If optional TARGET_COMPARTMENT_ID env var is set, it will also move the DNS Steering Policy to the compartment.
func ExampleDnsSteeringPolicy() {
	client, err := dns.NewDnsClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Optional parsed value from env var TARGET_COMPARTMENT_ID
	targetCompartmentId := os.Getenv("TARGET_COMPARTMENT_ID")
	log.Printf("TARGET_COMPARTMENT_ID: %s", targetCompartmentId)

	ctx := context.Background()

	// Create a new steering policy
	displayName := common.String(helpers.GetRandomString(15))
	createReq := dns.CreateSteeringPolicyRequest{
		CreateSteeringPolicyDetails: dns.CreateSteeringPolicyDetails{
			CompartmentId: helpers.CompartmentID(),
			DisplayName:   displayName,
			Template:      dns.CreateSteeringPolicyDetailsTemplateLoadBalance,
			Ttl:           common.Int(60),
		},
	}
	createResp, err := client.CreateSteeringPolicy(ctx, createReq)
	helpers.FatalIfError(err)
	fmt.Printf("created dns steering policy %s", *displayName)

	// below logic is to wait until steering policy is in active state
	pollUntilAvailable := func(r common.OCIOperationResponse) bool {
		if converted, ok := r.Response.(dns.GetSteeringPolicyResponse); ok {
			return converted.LifecycleState != dns.SteeringPolicyLifecycleStateActive
		}
		return true
	}
	getRequest := dns.GetSteeringPolicyRequest{
		SteeringPolicyId: createResp.Id,
		RequestMetadata:  helpers.GetRequestMetadataWithCustomizedRetryPolicy(pollUntilAvailable),
	}
	getResp, err := client.GetSteeringPolicy(ctx, getRequest)
	helpers.FatalIfError(err)
	fmt.Printf("get dns steering policy %s", *displayName)

	listResp, err := client.ListSteeringPolicies(ctx, dns.ListSteeringPoliciesRequest{
		CompartmentId:       helpers.CompartmentID(),
		DisplayNameContains: displayName,
		Limit:               common.Int64(10),
		SortBy:              dns.ListSteeringPoliciesSortByTimecreated,
		SortOrder:           dns.ListSteeringPoliciesSortOrderAsc,
	})
	helpers.FatalIfError(err)
	fmt.Println("list dns steering policies")
	log.Printf("count of dns steering policies in compartment %s: %d", *helpers.CompartmentID(), listResp.OpcTotalItems)

	if targetCompartmentId != "" && targetCompartmentId != *helpers.CompartmentID() {
		changeRequest := dns.ChangeSteeringPolicyCompartmentRequest{
			SteeringPolicyId: getResp.Id,
			ChangeSteeringPolicyCompartmentDetails: dns.ChangeSteeringPolicyCompartmentDetails{
				CompartmentId: &targetCompartmentId,
			},
		}
		_, err := client.ChangeSteeringPolicyCompartment(ctx, changeRequest)
		helpers.FatalIfError(err)
		fmt.Printf("change dns steering policy compartment to %s", targetCompartmentId)
	}

	// Clean up
	defer func() {
		_, err = client.DeleteSteeringPolicy(ctx, dns.DeleteSteeringPolicyRequest{
			SteeringPolicyId: getResp.Id,
		})
		helpers.FatalIfError(err)
	}()

	// Output:
	// created dns steering policy
	// get dns steering policy
	// list dns steering policy
	// deleted dns steering policy
}
