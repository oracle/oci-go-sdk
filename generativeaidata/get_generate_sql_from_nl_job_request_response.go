// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeaidata

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetGenerateSqlFromNlJobRequest wrapper for the GetGenerateSqlFromNlJob operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeaidata/GetGenerateSqlFromNlJob.go.html to see an example of how to use GetGenerateSqlFromNlJobRequest.
type GetGenerateSqlFromNlJobRequest struct {

	// The OCID of the semantic store
	SemanticStoreId *string `mandatory:"true" contributesTo:"path" name:"semanticStoreId"`

	// The OCID of the Semantic Store job for a GenerateSqlFromNl job.
	GenerateSqlFromNlJobId *string `mandatory:"true" contributesTo:"path" name:"generateSqlFromNlJobId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetGenerateSqlFromNlJobRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetGenerateSqlFromNlJobRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetGenerateSqlFromNlJobRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetGenerateSqlFromNlJobRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetGenerateSqlFromNlJobRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetGenerateSqlFromNlJobResponse wrapper for the GetGenerateSqlFromNlJob operation
type GetGenerateSqlFromNlJobResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The GenerateSqlFromNlJob instance
	GenerateSqlFromNlJob `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetGenerateSqlFromNlJobResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetGenerateSqlFromNlJobResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
