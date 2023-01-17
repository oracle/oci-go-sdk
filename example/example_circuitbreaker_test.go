// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package example

import (
	"context"
	"fmt"
	"time"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

// Example shows how to configure circuit breaker
func ExampleConfigureCircuitBreaker() {
	// If need to disable all service default circuit breaker, there are two ways: set circuit breaker environment variable to false
	// or use global variable.
	// common.GlobalCircuitBreakerSetting = common.NoCircuitBreaker()

	identityClient, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)
	// Add one more status code 404(compare with default circuit breaker setting) as the failure request in circuit breaker.
	successStatCodeMap := map[int]bool{
		429: false,
		404: false,
		500: false,
		502: false,
		503: false,
		504: false,
	}
	// Configure CircuitBreaker
	cbst := common.NewCircuitBreakerSettingWithOptions(
		common.WithName("myCircuitBreaker"),
		common.WithIsEnabled(true),
		common.WithMinimumRequests(5),
		common.WithCloseStateWindow(60*time.Second),
		common.WithFailureRateThreshold(0.70),
		common.WithSuccessStatCodeMap(successStatCodeMap),
		common.WithServiceName("Identity"))

	// if prefer to use default circuit breaker, no need to define successStatCodeMap and cb, but directly call:
	// cbst := common.DefaultCircuitBreakerSetting()

	identityClient.BaseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(cbst)

	// The OCID of the tenancy containing the compartment.
	tenancyID, err := common.DefaultConfigProvider().TenancyOCID()
	helpers.FatalIfError(err)

	// make the tenancyOCID incorrect on purpose - testing
	fakeTenancyID := tenancyID[1:len(tenancyID)-2] + "mm"

	request := identity.ListAvailabilityDomainsRequest{
		CompartmentId: &fakeTenancyID,
	}

	for i := 0; i < 5; i++ {
		identityClient.ListAvailabilityDomains(context.Background(), request)
		fmt.Println(i*10, "seconds CircuitBreaker state: "+identityClient.Configuration.CircuitBreaker.Cb.State().String())
		time.Sleep(10 * time.Second)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("After 55s, CircuitBreaker current state: " + identityClient.Configuration.CircuitBreaker.Cb.State().String())

	fmt.Println("Wait 30 sec...")
	time.Sleep(30 * time.Second)
	fmt.Println("Make a good API call")

	request = identity.ListAvailabilityDomainsRequest{
		CompartmentId: &tenancyID,
	}
	identityClient.ListAvailabilityDomains(context.Background(), request)
	time.Sleep(10 * time.Second)
	fmt.Println("check current CircuitBreaker state: " + identityClient.Configuration.CircuitBreaker.Cb.State().String())

	// Output:
	// 0 seconds CircuitBreaker state: closed
	// 10 seconds CircuitBreaker state: closed
	// 20 seconds CircuitBreaker state: closed
	// 30 seconds CircuitBreaker state: closed
	// 40 seconds CircuitBreaker state: open
	// After 55s, CircuitBreaker current state: open
	// Wait 30 sec...
	// Make a good API call
	// check current CircuitBreaker state: closed
}
