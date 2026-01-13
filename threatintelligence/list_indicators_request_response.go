// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package threatintelligence

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListIndicatorsRequest wrapper for the ListIndicators operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/threatintelligence/ListIndicators.go.html to see an example of how to use ListIndicatorsRequest.
type ListIndicatorsRequest struct {

	// The OCID of the tenancy (root compartment) that is used to filter results.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The threat type of entites to be returned. To filter for multiple threat types, repeat this parameter.
	ThreatTypeName []string `contributesTo:"query" name:"threatTypeName" collectionFormat:"multi"`

	// The indicator type of entities to be returned.
	Type ListIndicatorsTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// The indicator value of entities to be returned.
	Value *string `mandatory:"false" contributesTo:"query" name:"value"`

	// The minimum confidence score of entities to be returned.
	ConfidenceGreaterThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"confidenceGreaterThanOrEqualTo"`

	// The oldest update time of entities to be returned.
	TimeUpdatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdatedGreaterThanOrEqualTo"`

	// Return indicators updated before the provided time.
	TimeUpdatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdatedLessThan"`

	// The oldest last seen time of entities to be returned.
	TimeLastSeenGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLastSeenGreaterThanOrEqualTo"`

	// Return indicators last seen before the provided time.
	TimeLastSeenLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLastSeenLessThan"`

	// The oldest created/first seen time of entities to be returned.
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// Return indicators created/first seen before the provided time.
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThan"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListIndicatorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one field to sort by may be provided.
	SortBy ListIndicatorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIndicatorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIndicatorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListIndicatorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIndicatorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListIndicatorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListIndicatorsTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetListIndicatorsTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIndicatorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListIndicatorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListIndicatorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListIndicatorsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListIndicatorsResponse wrapper for the ListIndicators operation
type ListIndicatorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of IndicatorSummaryCollection instances
	IndicatorSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListIndicatorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIndicatorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIndicatorsTypeEnum Enum with underlying type: string
type ListIndicatorsTypeEnum string

// Set of constants representing the allowable values for ListIndicatorsTypeEnum
const (
	ListIndicatorsTypeDomainName ListIndicatorsTypeEnum = "DOMAIN_NAME"
	ListIndicatorsTypeFileName   ListIndicatorsTypeEnum = "FILE_NAME"
	ListIndicatorsTypeMd5Hash    ListIndicatorsTypeEnum = "MD5_HASH"
	ListIndicatorsTypeSha1Hash   ListIndicatorsTypeEnum = "SHA1_HASH"
	ListIndicatorsTypeSha256Hash ListIndicatorsTypeEnum = "SHA256_HASH"
	ListIndicatorsTypeIpAddress  ListIndicatorsTypeEnum = "IP_ADDRESS"
	ListIndicatorsTypeUrl        ListIndicatorsTypeEnum = "URL"
)

var mappingListIndicatorsTypeEnum = map[string]ListIndicatorsTypeEnum{
	"DOMAIN_NAME": ListIndicatorsTypeDomainName,
	"FILE_NAME":   ListIndicatorsTypeFileName,
	"MD5_HASH":    ListIndicatorsTypeMd5Hash,
	"SHA1_HASH":   ListIndicatorsTypeSha1Hash,
	"SHA256_HASH": ListIndicatorsTypeSha256Hash,
	"IP_ADDRESS":  ListIndicatorsTypeIpAddress,
	"URL":         ListIndicatorsTypeUrl,
}

var mappingListIndicatorsTypeEnumLowerCase = map[string]ListIndicatorsTypeEnum{
	"domain_name": ListIndicatorsTypeDomainName,
	"file_name":   ListIndicatorsTypeFileName,
	"md5_hash":    ListIndicatorsTypeMd5Hash,
	"sha1_hash":   ListIndicatorsTypeSha1Hash,
	"sha256_hash": ListIndicatorsTypeSha256Hash,
	"ip_address":  ListIndicatorsTypeIpAddress,
	"url":         ListIndicatorsTypeUrl,
}

// GetListIndicatorsTypeEnumValues Enumerates the set of values for ListIndicatorsTypeEnum
func GetListIndicatorsTypeEnumValues() []ListIndicatorsTypeEnum {
	values := make([]ListIndicatorsTypeEnum, 0)
	for _, v := range mappingListIndicatorsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListIndicatorsTypeEnumStringValues Enumerates the set of values in String for ListIndicatorsTypeEnum
func GetListIndicatorsTypeEnumStringValues() []string {
	return []string{
		"DOMAIN_NAME",
		"FILE_NAME",
		"MD5_HASH",
		"SHA1_HASH",
		"SHA256_HASH",
		"IP_ADDRESS",
		"URL",
	}
}

// GetMappingListIndicatorsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIndicatorsTypeEnum(val string) (ListIndicatorsTypeEnum, bool) {
	enum, ok := mappingListIndicatorsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIndicatorsSortOrderEnum Enum with underlying type: string
type ListIndicatorsSortOrderEnum string

// Set of constants representing the allowable values for ListIndicatorsSortOrderEnum
const (
	ListIndicatorsSortOrderAsc  ListIndicatorsSortOrderEnum = "ASC"
	ListIndicatorsSortOrderDesc ListIndicatorsSortOrderEnum = "DESC"
)

var mappingListIndicatorsSortOrderEnum = map[string]ListIndicatorsSortOrderEnum{
	"ASC":  ListIndicatorsSortOrderAsc,
	"DESC": ListIndicatorsSortOrderDesc,
}

var mappingListIndicatorsSortOrderEnumLowerCase = map[string]ListIndicatorsSortOrderEnum{
	"asc":  ListIndicatorsSortOrderAsc,
	"desc": ListIndicatorsSortOrderDesc,
}

// GetListIndicatorsSortOrderEnumValues Enumerates the set of values for ListIndicatorsSortOrderEnum
func GetListIndicatorsSortOrderEnumValues() []ListIndicatorsSortOrderEnum {
	values := make([]ListIndicatorsSortOrderEnum, 0)
	for _, v := range mappingListIndicatorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListIndicatorsSortOrderEnumStringValues Enumerates the set of values in String for ListIndicatorsSortOrderEnum
func GetListIndicatorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListIndicatorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIndicatorsSortOrderEnum(val string) (ListIndicatorsSortOrderEnum, bool) {
	enum, ok := mappingListIndicatorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListIndicatorsSortByEnum Enum with underlying type: string
type ListIndicatorsSortByEnum string

// Set of constants representing the allowable values for ListIndicatorsSortByEnum
const (
	ListIndicatorsSortByConfidence   ListIndicatorsSortByEnum = "confidence"
	ListIndicatorsSortByTimecreated  ListIndicatorsSortByEnum = "timeCreated"
	ListIndicatorsSortByTimeupdated  ListIndicatorsSortByEnum = "timeUpdated"
	ListIndicatorsSortByTimelastseen ListIndicatorsSortByEnum = "timeLastSeen"
)

var mappingListIndicatorsSortByEnum = map[string]ListIndicatorsSortByEnum{
	"confidence":   ListIndicatorsSortByConfidence,
	"timeCreated":  ListIndicatorsSortByTimecreated,
	"timeUpdated":  ListIndicatorsSortByTimeupdated,
	"timeLastSeen": ListIndicatorsSortByTimelastseen,
}

var mappingListIndicatorsSortByEnumLowerCase = map[string]ListIndicatorsSortByEnum{
	"confidence":   ListIndicatorsSortByConfidence,
	"timecreated":  ListIndicatorsSortByTimecreated,
	"timeupdated":  ListIndicatorsSortByTimeupdated,
	"timelastseen": ListIndicatorsSortByTimelastseen,
}

// GetListIndicatorsSortByEnumValues Enumerates the set of values for ListIndicatorsSortByEnum
func GetListIndicatorsSortByEnumValues() []ListIndicatorsSortByEnum {
	values := make([]ListIndicatorsSortByEnum, 0)
	for _, v := range mappingListIndicatorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListIndicatorsSortByEnumStringValues Enumerates the set of values in String for ListIndicatorsSortByEnum
func GetListIndicatorsSortByEnumStringValues() []string {
	return []string{
		"confidence",
		"timeCreated",
		"timeUpdated",
		"timeLastSeen",
	}
}

// GetMappingListIndicatorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListIndicatorsSortByEnum(val string) (ListIndicatorsSortByEnum, bool) {
	enum, ok := mappingListIndicatorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
