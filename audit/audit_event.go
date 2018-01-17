// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Audit API
//
// API for the Audit Service. You can use this API for queries, but not bulk-export operations.
//

package audit

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AuditEvent The representation of AuditEvent
type AuditEvent struct {

	// The OCID of the tenant.
	TenantId *string `mandatory:"false" json:"tenantId,omitempty"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId,omitempty"`

	// The GUID of the event.
	EventId *string `mandatory:"false" json:"eventId,omitempty"`

	// The source of the event.
	EventSource *string `mandatory:"false" json:"eventSource,omitempty"`

	// The type of the event.
	EventType *string `mandatory:"false" json:"eventType,omitempty"`

	// The time the event occurred, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.
	EventTime *common.SDKTime `mandatory:"false" json:"eventTime,omitempty"`

	// The OCID of the user whose action triggered the event.
	PrincipalId *string `mandatory:"false" json:"principalId,omitempty"`

	// The credential ID of the user. This value is extracted from the HTTP 'Authorization' request header. It consists of the tenantId, userId, and user fingerprint, all delimited by a slash (/).
	CredentialId *string `mandatory:"false" json:"credentialId,omitempty"`

	// The HTTP method of the request.
	RequestAction *string `mandatory:"false" json:"requestAction,omitempty"`

	// The opc-request-id of the request.
	RequestId *string `mandatory:"false" json:"requestId,omitempty"`

	// The user agent of the client that made the request.
	RequestAgent *string `mandatory:"false" json:"requestAgent,omitempty"`

	// The HTTP header fields and values in the request.
	RequestHeaders map[string][]string `mandatory:"false" json:"requestHeaders,omitempty"`

	// The IP address of the source of the request.
	RequestOrigin *string `mandatory:"false" json:"requestOrigin,omitempty"`

	// The query parameter fields and values for the request.
	RequestParameters map[string][]string `mandatory:"false" json:"requestParameters,omitempty"`

	// The resource targeted by the request.
	RequestResource *string `mandatory:"false" json:"requestResource,omitempty"`

	// The headers of the response.
	ResponseHeaders map[string][]string `mandatory:"false" json:"responseHeaders,omitempty"`

	// The status code of the response.
	ResponseStatus *string `mandatory:"false" json:"responseStatus,omitempty"`

	// The time of the response to the audited request, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.
	ResponseTime *common.SDKTime `mandatory:"false" json:"responseTime,omitempty"`

	// Metadata of interest from the response payload. For example, the OCID of a resource.
	ResponsePayload map[string]interface{} `mandatory:"false" json:"responsePayload,omitempty"`
}

func (m AuditEvent) String() string {
	return common.PointerString(m)
}
