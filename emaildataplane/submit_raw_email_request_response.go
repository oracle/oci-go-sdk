// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package emaildataplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// SubmitRawEmailRequest wrapper for the SubmitRawEmail operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/emaildataplane/SubmitRawEmail.go.html to see an example of how to use SubmitRawEmailRequest.
type SubmitRawEmailRequest struct {

	// The media type of the body.
	ContentType SubmitRawEmailContentTypeEnum `mandatory:"true" contributesTo:"header" name:"content-type"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the approved sender resource.
	CompartmentId *string `mandatory:"true" contributesTo:"header" name:"compartment-id"`

	// The envelope and the header from email address, that is sending the email. Email address must be an approved sender.
	Sender *string `mandatory:"true" contributesTo:"header" name:"sender"`

	// The destination for the email, all recipients including to, cc and bcc addresses.
	Recipients []string `contributesTo:"header" name:"recipients" collectionFormat:"csv"`

	// This should be formatted in valid MIME format. Message can include attachments. MIME libraries should be used to convert the content into the appropriate format.
	RawMessage io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// The request ID for tracing from the system
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The content length of the body.
	ContentLength *int64 `mandatory:"false" contributesTo:"header" name:"content-length"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SubmitRawEmailRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SubmitRawEmailRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
	if err == nil && binaryRequestBody.Seekable() {
		common.UpdateRequestBinaryBody(&httpRequest, binaryRequestBody)
	}
	return httpRequest, err
}

// BinaryRequestBody implements the OCIRequest interface
func (request SubmitRawEmailRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {
	rsc := common.NewOCIReadSeekCloser(request.RawMessage)
	if rsc.Seekable() {
		return rsc, true
	}
	return nil, true

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SubmitRawEmailRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SubmitRawEmailRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSubmitRawEmailContentTypeEnum(string(request.ContentType)); !ok && request.ContentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ContentType: %s. Supported values are: %s.", request.ContentType, strings.Join(GetSubmitRawEmailContentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SubmitRawEmailResponse wrapper for the SubmitRawEmail operation
type SubmitRawEmailResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The EmailRawSubmittedResponse instance
	EmailRawSubmittedResponse `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response SubmitRawEmailResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SubmitRawEmailResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SubmitRawEmailContentTypeEnum Enum with underlying type: string
type SubmitRawEmailContentTypeEnum string

// Set of constants representing the allowable values for SubmitRawEmailContentTypeEnum
const (
	SubmitRawEmailContentTypeRfc822 SubmitRawEmailContentTypeEnum = "message/rfc822"
	SubmitRawEmailContentTypeGlobal SubmitRawEmailContentTypeEnum = "message/global"
)

var mappingSubmitRawEmailContentTypeEnum = map[string]SubmitRawEmailContentTypeEnum{
	"message/rfc822": SubmitRawEmailContentTypeRfc822,
	"message/global": SubmitRawEmailContentTypeGlobal,
}

var mappingSubmitRawEmailContentTypeEnumLowerCase = map[string]SubmitRawEmailContentTypeEnum{
	"message/rfc822": SubmitRawEmailContentTypeRfc822,
	"message/global": SubmitRawEmailContentTypeGlobal,
}

// GetSubmitRawEmailContentTypeEnumValues Enumerates the set of values for SubmitRawEmailContentTypeEnum
func GetSubmitRawEmailContentTypeEnumValues() []SubmitRawEmailContentTypeEnum {
	values := make([]SubmitRawEmailContentTypeEnum, 0)
	for _, v := range mappingSubmitRawEmailContentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSubmitRawEmailContentTypeEnumStringValues Enumerates the set of values in String for SubmitRawEmailContentTypeEnum
func GetSubmitRawEmailContentTypeEnumStringValues() []string {
	return []string{
		"message/rfc822",
		"message/global",
	}
}

// GetMappingSubmitRawEmailContentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubmitRawEmailContentTypeEnum(val string) (SubmitRawEmailContentTypeEnum, bool) {
	enum, ok := mappingSubmitRawEmailContentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
