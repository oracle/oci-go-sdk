// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListIdentityProvidersRequest wrapper for the ListIdentityProviders operation
type ListIdentityProvidersRequest struct {

	// The protocol used for federation.
	Protocol ListIdentityProvidersProtocolEnum `mandatory:"true" contributesTo:"query" name:"protocol" omitEmpty:"true"`

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIdentityProvidersRequest) String() string {
	return common.PointerString(request)
}

// GetHttpRequest implements the OciRequest interface
func (request ListIdentityProvidersRequest) GetHttpRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// GetRetryPolicy implements the OciRetryableRequest interface
// => assembles retry policy based on specified options and default behavior
func (request ListIdentityProvidersRequest) GetRetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListIdentityProvidersResponse wrapper for the ListIdentityProviders operation
type ListIdentityProvidersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []IdentityProvider instance
	Items []IdentityProvider `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListIdentityProvidersResponse) String() string {
	return common.PointerString(response)
}

// GetRawResponse implements the OciResponse interface
func (response ListIdentityProvidersResponse) GetRawResponse() *http.Response {
	return response.RawResponse
}

// ListIdentityProvidersProtocolEnum Enum with underlying type: string
type ListIdentityProvidersProtocolEnum string

// Set of constants representing the allowable values for ListIdentityProvidersProtocol
const (
	ListIdentityProvidersProtocolSaml2   ListIdentityProvidersProtocolEnum = "SAML2"
	ListIdentityProvidersProtocolUnknown ListIdentityProvidersProtocolEnum = "UNKNOWN"
)

var mappingListIdentityProvidersProtocol = map[string]ListIdentityProvidersProtocolEnum{
	"SAML2":   ListIdentityProvidersProtocolSaml2,
	"UNKNOWN": ListIdentityProvidersProtocolUnknown,
}

// GetListIdentityProvidersProtocolEnumValues Enumerates the set of values for ListIdentityProvidersProtocol
func GetListIdentityProvidersProtocolEnumValues() []ListIdentityProvidersProtocolEnum {
	values := make([]ListIdentityProvidersProtocolEnum, 0)
	for _, v := range mappingListIdentityProvidersProtocol {
		if v != ListIdentityProvidersProtocolUnknown {
			values = append(values, v)
		}
	}
	return values
}
