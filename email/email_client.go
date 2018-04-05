// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Email Delivery Service API
//
// API spec for managing OCI Email Delivery services.
//

package email

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//EmailClient a client for Email
type EmailClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewEmailClientWithConfigurationProvider Creates a new default Email client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewEmailClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client EmailClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = EmailClient{BaseClient: baseClient}
	client.BasePath = "20170907"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *EmailClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "email", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *EmailClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.config = &configProvider
	client.SetRegion(region)
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *EmailClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateSender Creates a sender for a tenancy in a given compartment.
func (client EmailClient) CreateSender(ctx context.Context, request CreateSenderRequest) (response CreateSenderResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/senders", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// CreateSuppression Adds recipient email addresses to the suppression list for a tenancy.
func (client EmailClient) CreateSuppression(ctx context.Context, request CreateSuppressionRequest) (response CreateSuppressionResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/suppressions", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// DeleteSender Deletes an approved sender for a tenancy in a given compartment for a
// provided `senderId`.
func (client EmailClient) DeleteSender(ctx context.Context, request DeleteSenderRequest) (response DeleteSenderResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/senders/{senderId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// DeleteSuppression Removes a suppressed recipient email address from the suppression list
// for a tenancy in a given compartment for a provided `suppressionId`.
func (client EmailClient) DeleteSuppression(ctx context.Context, request DeleteSuppressionRequest) (response DeleteSuppressionResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/suppressions/{suppressionId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// GetSender Gets an approved sender for a given `senderId`.
func (client EmailClient) GetSender(ctx context.Context, request GetSenderRequest) (response GetSenderResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/senders/{senderId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// GetSuppression Gets the details of a suppressed recipient email address for a given
// `suppressionId`. Each suppression is given a unique OCID.
func (client EmailClient) GetSuppression(ctx context.Context, request GetSuppressionRequest) (response GetSuppressionResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/suppressions/{suppressionId}", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// ListSenders Gets a collection of approved sender email addresses and sender IDs.
func (client EmailClient) ListSenders(ctx context.Context, request ListSendersRequest) (response ListSendersResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/senders", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}

// ListSuppressions Gets a list of suppressed recipient email addresses for a user. The
// `compartmentId` for suppressions must be a tenancy OCID. The returned list
// is sorted by creation time in descending order.
func (client EmailClient) ListSuppressions(ctx context.Context, request ListSuppressionsRequest) (response ListSuppressionsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/suppressions", request)
	if err != nil {
		return
	}

	httpResponse, err := client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return
}
