// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for sending raw request to  Service API

package example

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
)

// ExampleRawRequest compose a request, sign it and send to server
func ExampleListUsers_RawRequest() {
	// build the url
	url := "https://identity.us-phoenix-1.oraclecloud.com/20160918/users/?compartmentId=" + *helpers.RootCompartmentID()

	// create request
	request, err := http.NewRequest("GET", url, nil)
	helpers.FatalIfError(err)

	// Set the Date header
	request.Header.Set("Date", time.Now().UTC().Format(http.TimeFormat))

	// And a provider of cryptographic keys
	provider := common.DefaultConfigProvider()

	// Build the signer
	signer := common.DefaultRequestSigner(provider)

	// Sign the request
	signer.Sign(request)

	client := http.Client{}

	fmt.Println("send request")

	// Execute the request
	resp, err := client.Do(request)
	helpers.FatalIfError(err)

	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	fmt.Println("receive response")

	// Output:
	// send request
	// receive response
}
