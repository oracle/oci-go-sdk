// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package certificates

import (
	"github.com/oracle/oci-go-sdk/v54/common"
	"net/http"
)

// GetCertificateBundleRequest wrapper for the GetCertificateBundle operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificates/GetCertificateBundle.go.html to see an example of how to use GetCertificateBundleRequest.
type GetCertificateBundleRequest struct {

	// The OCID of the certificate.
	CertificateId *string `mandatory:"true" contributesTo:"path" name:"certificateId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The version number of the certificate. The default value is 0, which means that this query parameter is ignored.
	VersionNumber *int64 `mandatory:"false" contributesTo:"query" name:"versionNumber"`

	// The name of the certificate. (This might be referred to as the name of the certificate version, as every certificate consists of at least one version.) Names are unique across versions of a given certificate.
	CertificateVersionName *string `mandatory:"false" contributesTo:"query" name:"certificateVersionName"`

	// The rotation state of the certificate version.
	Stage GetCertificateBundleStageEnum `mandatory:"false" contributesTo:"query" name:"stage" omitEmpty:"true"`

	// The type of certificate bundle. By default, the private key fields are not returned. When querying for certificate bundles, to return results with certificate contents, the private key in PEM format, and the private key passphrase, specify the value of this parameter as `CERTIFICATE_CONTENT_WITH_PRIVATE_KEY`.
	CertificateBundleType GetCertificateBundleCertificateBundleTypeEnum `mandatory:"false" contributesTo:"query" name:"certificateBundleType" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetCertificateBundleRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetCertificateBundleRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetCertificateBundleRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetCertificateBundleRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetCertificateBundleResponse wrapper for the GetCertificateBundle operation
type GetCertificateBundleResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The CertificateBundle instance
	CertificateBundle `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetCertificateBundleResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetCertificateBundleResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetCertificateBundleStageEnum Enum with underlying type: string
type GetCertificateBundleStageEnum string

// Set of constants representing the allowable values for GetCertificateBundleStageEnum
const (
	GetCertificateBundleStageCurrent    GetCertificateBundleStageEnum = "CURRENT"
	GetCertificateBundleStagePending    GetCertificateBundleStageEnum = "PENDING"
	GetCertificateBundleStageLatest     GetCertificateBundleStageEnum = "LATEST"
	GetCertificateBundleStagePrevious   GetCertificateBundleStageEnum = "PREVIOUS"
	GetCertificateBundleStageDeprecated GetCertificateBundleStageEnum = "DEPRECATED"
)

var mappingGetCertificateBundleStage = map[string]GetCertificateBundleStageEnum{
	"CURRENT":    GetCertificateBundleStageCurrent,
	"PENDING":    GetCertificateBundleStagePending,
	"LATEST":     GetCertificateBundleStageLatest,
	"PREVIOUS":   GetCertificateBundleStagePrevious,
	"DEPRECATED": GetCertificateBundleStageDeprecated,
}

// GetGetCertificateBundleStageEnumValues Enumerates the set of values for GetCertificateBundleStageEnum
func GetGetCertificateBundleStageEnumValues() []GetCertificateBundleStageEnum {
	values := make([]GetCertificateBundleStageEnum, 0)
	for _, v := range mappingGetCertificateBundleStage {
		values = append(values, v)
	}
	return values
}

// GetCertificateBundleCertificateBundleTypeEnum Enum with underlying type: string
type GetCertificateBundleCertificateBundleTypeEnum string

// Set of constants representing the allowable values for GetCertificateBundleCertificateBundleTypeEnum
const (
	GetCertificateBundleCertificateBundleTypePublicOnly     GetCertificateBundleCertificateBundleTypeEnum = "CERTIFICATE_CONTENT_PUBLIC_ONLY"
	GetCertificateBundleCertificateBundleTypeWithPrivateKey GetCertificateBundleCertificateBundleTypeEnum = "CERTIFICATE_CONTENT_WITH_PRIVATE_KEY"
)

var mappingGetCertificateBundleCertificateBundleType = map[string]GetCertificateBundleCertificateBundleTypeEnum{
	"CERTIFICATE_CONTENT_PUBLIC_ONLY":      GetCertificateBundleCertificateBundleTypePublicOnly,
	"CERTIFICATE_CONTENT_WITH_PRIVATE_KEY": GetCertificateBundleCertificateBundleTypeWithPrivateKey,
}

// GetGetCertificateBundleCertificateBundleTypeEnumValues Enumerates the set of values for GetCertificateBundleCertificateBundleTypeEnum
func GetGetCertificateBundleCertificateBundleTypeEnumValues() []GetCertificateBundleCertificateBundleTypeEnum {
	values := make([]GetCertificateBundleCertificateBundleTypeEnum, 0)
	for _, v := range mappingGetCertificateBundleCertificateBundleType {
		values = append(values, v)
	}
	return values
}
