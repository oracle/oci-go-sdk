// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cims

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetCsiNumberRequest wrapper for the GetCsiNumber operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cims/GetCsiNumber.go.html to see an example of how to use GetCsiNumberRequest.
type GetCsiNumberRequest struct {

	// Tenancy Ocid in oracle cloud Infrastructure
	TenantId *string `mandatory:"true" contributesTo:"query" name:"tenantId"`

	// Home region of the customer which is part of oracle cloud infrastructure regions
	Region *string `mandatory:"true" contributesTo:"query" name:"region"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// User OCID for Oracle Identity Cloud Service (IDCS) users who also have a federated Oracle Cloud Infrastructure account.
	Ocid *string `mandatory:"false" contributesTo:"header" name:"ocid"`

	// The region of the tenancy.
	Homeregion *string `mandatory:"false" contributesTo:"header" name:"homeregion"`

	// Token type that determine which cloud provider the request come from.
	Bearertokentype *string `mandatory:"false" contributesTo:"header" name:"bearertokentype"`

	// Token that provided by multi cloud provider, which help to validate the email.
	Bearertoken *string `mandatory:"false" contributesTo:"header" name:"bearertoken"`

	// IdToken that provided by multi cloud provider, which help to validate the email.
	Idtoken *string `mandatory:"false" contributesTo:"header" name:"idtoken"`

	// The OCID of identity domain.
	Domainid *string `mandatory:"false" contributesTo:"header" name:"domainid"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetCsiNumberRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetCsiNumberRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetCsiNumberRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetCsiNumberRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetCsiNumberRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetCsiNumberResponse wrapper for the GetCsiNumber operation
type GetCsiNumberResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The string instance
	Value *string `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetCsiNumberResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetCsiNumberResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
