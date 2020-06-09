// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for HealthChecks API

package example

import (
	"context"
	"fmt"
	"testing"

	"github.com/oracle/oci-go-sdk/common"

	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/oracle/oci-go-sdk/healthchecks"
)

// Shows listing vantage points
func exampleListVantagePoints(ctx context.Context, client healthchecks.HealthChecksClient) []healthchecks.HealthChecksVantagePointSummary {
	vantagePoints := []healthchecks.HealthChecksVantagePointSummary{}

	request := healthchecks.ListHealthChecksVantagePointsRequest{}

	for r, err := client.ListHealthChecksVantagePoints(ctx, request); ; r, err = client.ListHealthChecksVantagePoints(ctx, request) {
		helpers.FatalIfError(err)

		vantagePoints = append(vantagePoints, r.Items...)

		if r.OpcNextPage != nil {
			// if there are more items in next page, fetch items from next page
			request.Page = r.OpcNextPage
		} else {
			// no more result, break the loop
			break
		}
	}

	return vantagePoints
}

// Shows creating a new Http Monitor:
func exampleCreateHttpMonitor(ctx context.Context, client healthchecks.HealthChecksClient, compartmentId *string) healthchecks.HttpMonitor {

	createDetails := healthchecks.CreateHttpMonitorDetails{
		CompartmentId:     compartmentId,
		DisplayName:       common.String("Monitor Name"),
		Targets:           []string{"example.com"},
		Protocol:          healthchecks.HttpProbeProtocolHttps,
		Port:              common.Int(443),
		Path:              common.String("/"),
		IsEnabled:         common.Bool(false),
		IntervalInSeconds: common.Int(30),
		TimeoutInSeconds:  common.Int(30),
	}

	resp, crerr := client.CreateHttpMonitor(ctx, healthchecks.CreateHttpMonitorRequest{
		CreateHttpMonitorDetails: createDetails,
	})

	helpers.FatalIfError(crerr)

	return resp.HttpMonitor
}

// Shows updating an existing Http Monitor:
func exampleUpdateHttpMonitor(ctx context.Context, client healthchecks.HealthChecksClient, monitorId *string) healthchecks.HttpMonitor {

	updateDetails := healthchecks.UpdateHttpMonitorDetails{
		IsEnabled: common.Bool(true),
		Targets:   []string{"example.com", "other.example.com"},
	}

	resp, crerr := client.UpdateHttpMonitor(ctx, healthchecks.UpdateHttpMonitorRequest{
		MonitorId:                monitorId,
		UpdateHttpMonitorDetails: updateDetails,
	})

	helpers.FatalIfError(crerr)

	return resp.HttpMonitor
}

// Shows getting monitor results.
func exampleListHttpMonitorResults(ctx context.Context, client healthchecks.HealthChecksClient, monitorId *string) []healthchecks.HttpProbeResultSummary {
	monitorResponses := []healthchecks.HttpProbeResultSummary{}

	request := healthchecks.ListHttpProbeResultsRequest{
		ProbeConfigurationId: monitorId,
	}

	for r, err := client.ListHttpProbeResults(ctx, request); ; r, err = client.ListHttpProbeResults(ctx, request) {
		helpers.FatalIfError(err)

		monitorResponses = append(monitorResponses, r.Items...)

		if r.OpcNextPage != nil {
			// if there are more items in next page, fetch items from next page
			request.Page = r.OpcNextPage
		} else {
			// no more result, break the loop
			break
		}
	}

	return monitorResponses
}

// Shows moving between compartments
func exampleMoveCompartmentHttpMonitor(ctx context.Context, client healthchecks.HealthChecksClient, monitorId *string, newCompartment *string) {
	request := healthchecks.ChangeHttpMonitorCompartmentRequest{
		MonitorId: monitorId,
		ChangeHttpMonitorCompartmentDetails: healthchecks.ChangeHttpMonitorCompartmentDetails{
			CompartmentId: newCompartment,
		},
	}

	_, err := client.ChangeHttpMonitorCompartment(ctx, request)
	helpers.FatalIfError(err)
}

// Shows deleting an existing monitor
func exampleDeleteHttpMonitor(ctx context.Context, client healthchecks.HealthChecksClient, monitorId *string) {
	deleteRequest := healthchecks.DeleteHttpMonitorRequest{
		MonitorId: monitorId,
	}

	_, err := client.DeleteHttpMonitor(ctx, deleteRequest)
	// resp will be blank if it succeeded

	helpers.FatalIfError(err)
}

func ExampleHealthChecksHttpSamples() {
	ctx := context.Background()

	// Initialize default config provider
	configProvider := common.DefaultConfigProvider()
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		panic(err)
	}

	client, clerr := healthchecks.NewHealthChecksClientWithConfigurationProvider(configProvider)
	helpers.FatalIfError(clerr)

	compartmentId, _ := configProvider.TenancyOCID()

	_ = exampleListVantagePoints(ctx, client)
	//fmt.Println(vantagePoints)
	fmt.Println("List Vantage Points")

	httpMonitor := exampleCreateHttpMonitor(ctx, client, &compartmentId)
	fmt.Println("Create Monitor")
	httpMonitor = exampleUpdateHttpMonitor(ctx, client, httpMonitor.Id)
	fmt.Println("Update Monitor")
	//fmt.Println(httpMonitor)

	_ = exampleListHttpMonitorResults(ctx, client, httpMonitor.Id)
	fmt.Println("Retrieved Results")

	// We need a different compartment to run this.
	// exampleMoveCompartmentHttpMonitor(ctx, client, httpMonitor.Id, &compartmentId)

	exampleDeleteHttpMonitor(ctx, client, httpMonitor.Id)
	fmt.Println("Deleted Monitor")

	// Output:
	// List Vantage Points
	// Create Monitor
	// Update Monitor
	// Retrieved Results
	// Deleted Monitor
}

// Shows creating a new Ping Monitor:
func exampleCreatePingMonitor(ctx context.Context, client healthchecks.HealthChecksClient, compartmentId *string) healthchecks.PingMonitor {

	createDetails := healthchecks.CreatePingMonitorDetails{
		CompartmentId:     compartmentId,
		DisplayName:       common.String("Monitor Name"),
		Targets:           []string{"example.com"},
		Protocol:          healthchecks.PingProbeProtocolIcmp,
		IsEnabled:         common.Bool(false),
		IntervalInSeconds: common.Int(30),
		TimeoutInSeconds:  common.Int(30),
	}

	resp, crerr := client.CreatePingMonitor(ctx, healthchecks.CreatePingMonitorRequest{
		CreatePingMonitorDetails: createDetails,
	})

	helpers.FatalIfError(crerr)

	return resp.PingMonitor
}

// Shows updating an existing Ping Monitor:
func exampleUpdatePingMonitor(ctx context.Context, client healthchecks.HealthChecksClient, monitorId *string) healthchecks.PingMonitor {

	updateDetails := healthchecks.UpdatePingMonitorDetails{
		IsEnabled: common.Bool(true),
		Targets:   []string{"example.com", "other.example.com"},
	}

	resp, crerr := client.UpdatePingMonitor(ctx, healthchecks.UpdatePingMonitorRequest{
		MonitorId:                monitorId,
		UpdatePingMonitorDetails: updateDetails,
	})

	helpers.FatalIfError(crerr)

	return resp.PingMonitor
}

// Shows getting monitor results.
func exampleListPingMonitorResults(ctx context.Context, client healthchecks.HealthChecksClient, monitorId *string) []healthchecks.PingProbeResultSummary {
	monitorResponses := []healthchecks.PingProbeResultSummary{}

	request := healthchecks.ListPingProbeResultsRequest{
		ProbeConfigurationId: monitorId,
	}

	for r, err := client.ListPingProbeResults(ctx, request); ; r, err = client.ListPingProbeResults(ctx, request) {
		helpers.FatalIfError(err)

		monitorResponses = append(monitorResponses, r.Items...)

		if r.OpcNextPage != nil {
			// if there are more items in next page, fetch items from next page
			request.Page = r.OpcNextPage
		} else {
			// no more result, break the loop
			break
		}
	}

	return monitorResponses
}

// Shows moving between compartments
func exampleMoveCompartmentPingMonitor(ctx context.Context, client healthchecks.HealthChecksClient, monitorId *string, newCompartment *string) {
	request := healthchecks.ChangePingMonitorCompartmentRequest{
		MonitorId: monitorId,
		ChangePingMonitorCompartmentDetails: healthchecks.ChangePingMonitorCompartmentDetails{
			CompartmentId: newCompartment,
		},
	}

	_, err := client.ChangePingMonitorCompartment(ctx, request)
	helpers.FatalIfError(err)
}

// Shows deleting an existing monitor
func exampleDeletePingMonitor(ctx context.Context, client healthchecks.HealthChecksClient, monitorId *string) {
	deleteRequest := healthchecks.DeletePingMonitorRequest{
		MonitorId: monitorId,
	}

	_, err := client.DeletePingMonitor(ctx, deleteRequest)
	// resp will be blank if it succeeded

	helpers.FatalIfError(err)
}

func ExampleHealthChecksPingSamples() {
	ctx := context.Background()

	// Initialize default config provider
	configProvider := common.DefaultConfigProvider()
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		panic(err)
	}

	client, clerr := healthchecks.NewHealthChecksClientWithConfigurationProvider(configProvider)
	helpers.FatalIfError(clerr)

	compartmentId, _ := configProvider.TenancyOCID()

	_ = exampleListVantagePoints(ctx, client)
	//fmt.Println(vantagePoints)
	fmt.Println("List Vantage Points")

	pingMonitor := exampleCreatePingMonitor(ctx, client, &compartmentId)
	fmt.Println("Create Monitor")
	pingMonitor = exampleUpdatePingMonitor(ctx, client, pingMonitor.Id)
	fmt.Println("Update Monitor")
	//fmt.Println(pingMonitor)

	_ = exampleListPingMonitorResults(ctx, client, pingMonitor.Id)
	fmt.Println("Retrieved Results")

	// We need a different compartment to run this.
	// exampleMoveCompartmentPingMonitor(ctx, client, pingMonitor.Id, &compartmentId)

	exampleDeletePingMonitor(ctx, client, pingMonitor.Id)
	fmt.Println("Deleted Monitor")

	// Output:
	// List Vantage Points
	// Create Monitor
	// Update Monitor
	// Retrieved Results
	// Deleted Monitor
}

func TestHealthChecks(t *testing.T) {
	ExampleHealthChecksHttpSamples()

	ExampleHealthChecksPingSamples()

	// Output:
	// List Vantage Points
	// Create Monitor
	// Update Monitor
	// Retrieved Results
	// Deleted Monitor
	// List Vantage Points
	// Create Monitor
	// Update Monitor
	// Retrieved Results
	// Deleted Monitor

}
