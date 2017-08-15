package common

type RequestOCI struct {

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations (e.g., if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// may be rejected).
	OpcRetryToken string
}

type PaginatedRequestOCI struct {

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page string

	// The maximum number of items to return in a paginated "List" call.
	Limit int32
}

type ResponseOCI struct {

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID string
}

type PaginatedResponseOCI struct {
	ResponseOCI

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage string
}
