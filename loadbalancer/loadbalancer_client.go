// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//LoadBalancerClient a client for LoadBalancer
type LoadBalancerClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewLoadBalancerClientWithConfigurationProvider Creates a new default LoadBalancer client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewLoadBalancerClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client LoadBalancerClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	client = LoadBalancerClient{BaseClient: baseClient}
	client.BasePath = "20170115"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *LoadBalancerClient) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "iaas", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *LoadBalancerClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *LoadBalancerClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateBackend Adds a backend server to a backend set.
func (client LoadBalancerClient) CreateBackend(ctx context.Context, request CreateBackendRequest) (response CreateBackendResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createBackend, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateBackendResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) createBackend(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/backends")
	if err != nil {
		return nil, err
	}

	var response CreateBackendResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateBackendSet Adds a backend set to a load balancer.
func (client LoadBalancerClient) CreateBackendSet(ctx context.Context, request CreateBackendSetRequest) (response CreateBackendSetResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createBackendSet, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateBackendSetResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) createBackendSet(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/loadBalancers/{loadBalancerId}/backendSets")
	if err != nil {
		return nil, err
	}

	var response CreateBackendSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateCertificate Creates an asynchronous request to add an SSL certificate.
func (client LoadBalancerClient) CreateCertificate(ctx context.Context, request CreateCertificateRequest) (response CreateCertificateResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createCertificate, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateCertificateResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) createCertificate(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/loadBalancers/{loadBalancerId}/certificates")
	if err != nil {
		return nil, err
	}

	var response CreateCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateListener Adds a listener to a load balancer.
func (client LoadBalancerClient) CreateListener(ctx context.Context, request CreateListenerRequest) (response CreateListenerResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createListener, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateListenerResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) createListener(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/loadBalancers/{loadBalancerId}/listeners")
	if err != nil {
		return nil, err
	}

	var response CreateListenerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// CreateLoadBalancer Creates a new load balancer in the specified compartment. For general information about load balancers,
// see [Overview of the Load Balancing Service]({{DOC_SERVER_URL}}/Content/Balance/Concepts/balanceoverview.htm).
// For the purposes of access control, you must provide the OCID of the compartment where you want
// the load balancer to reside. Notice that the load balancer doesn't have to be in the same compartment as the VCN
// or backend set. If you're not sure which compartment to use, put the load balancer in the same compartment as the VCN.
// For information about access control and compartments, see
// [Overview of the IAM Service]({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm).
// You must specify a display name for the load balancer. It does not have to be unique, and you can change it.
// For information about Availability Domains, see
// [Regions and Availability Domains]({{DOC_SERVER_URL}}/Content/General/Concepts/regions.htm).
// To get a list of Availability Domains, use the `ListAvailabilityDomains` operation
// in the Identity and Access Management Service API.
// All Oracle Cloud Infrastructure resources, including load balancers, get an Oracle-assigned,
// unique ID called an Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID
// in the response. You can also retrieve a resource's OCID by using a List API operation on that resource type,
// or by viewing the resource in the Console. Fore more information, see
// [Resource Identifiers]({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// After you send your request, the new object's state will temporarily be PROVISIONING. Before using the
// object, first make sure its state has changed to RUNNING.
// When you create a load balancer, the system assigns an IP address.
// To get the IP address, use the GetLoadBalancer operation.
func (client LoadBalancerClient) CreateLoadBalancer(ctx context.Context, request CreateLoadBalancerRequest) (response CreateLoadBalancerResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.createLoadBalancer, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(CreateLoadBalancerResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) createLoadBalancer(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPost, "/loadBalancers")
	if err != nil {
		return nil, err
	}

	var response CreateLoadBalancerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteBackend Removes a backend server from a given load balancer and backend set.
func (client LoadBalancerClient) DeleteBackend(ctx context.Context, request DeleteBackendRequest) (response DeleteBackendResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteBackend, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteBackendResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) deleteBackend(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/backends/{backendName}")
	if err != nil {
		return nil, err
	}

	var response DeleteBackendResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteBackendSet Deletes the specified backend set. Note that deleting a backend set removes its backend servers from the load balancer.
// Before you can delete a backend set, you must remove it from any active listeners.
func (client LoadBalancerClient) DeleteBackendSet(ctx context.Context, request DeleteBackendSetRequest) (response DeleteBackendSetResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteBackendSet, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteBackendSetResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) deleteBackendSet(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}")
	if err != nil {
		return nil, err
	}

	var response DeleteBackendSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteCertificate Deletes an SSL certificate from a load balancer.
func (client LoadBalancerClient) DeleteCertificate(ctx context.Context, request DeleteCertificateRequest) (response DeleteCertificateResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteCertificate, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteCertificateResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) deleteCertificate(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/loadBalancers/{loadBalancerId}/certificates/{certificateName}")
	if err != nil {
		return nil, err
	}

	var response DeleteCertificateResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteListener Deletes a listener from a load balancer.
func (client LoadBalancerClient) DeleteListener(ctx context.Context, request DeleteListenerRequest) (response DeleteListenerResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteListener, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteListenerResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) deleteListener(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/loadBalancers/{loadBalancerId}/listeners/{listenerName}")
	if err != nil {
		return nil, err
	}

	var response DeleteListenerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// DeleteLoadBalancer Stops a load balancer and removes it from service.
func (client LoadBalancerClient) DeleteLoadBalancer(ctx context.Context, request DeleteLoadBalancerRequest) (response DeleteLoadBalancerResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.deleteLoadBalancer, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(DeleteLoadBalancerResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) deleteLoadBalancer(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodDelete, "/loadBalancers/{loadBalancerId}")
	if err != nil {
		return nil, err
	}

	var response DeleteLoadBalancerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetBackend Gets the specified backend server's configuration information.
func (client LoadBalancerClient) GetBackend(ctx context.Context, request GetBackendRequest) (response GetBackendResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getBackend, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetBackendResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) getBackend(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/backends/{backendName}")
	if err != nil {
		return nil, err
	}

	var response GetBackendResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetBackendHealth Gets the current health status of the specified backend server.
func (client LoadBalancerClient) GetBackendHealth(ctx context.Context, request GetBackendHealthRequest) (response GetBackendHealthResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getBackendHealth, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetBackendHealthResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) getBackendHealth(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/backends/{backendName}/health")
	if err != nil {
		return nil, err
	}

	var response GetBackendHealthResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetBackendSet Gets the specified backend set's configuration information.
func (client LoadBalancerClient) GetBackendSet(ctx context.Context, request GetBackendSetRequest) (response GetBackendSetResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getBackendSet, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetBackendSetResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) getBackendSet(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}")
	if err != nil {
		return nil, err
	}

	var response GetBackendSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetBackendSetHealth Gets the health status for the specified backend set.
func (client LoadBalancerClient) GetBackendSetHealth(ctx context.Context, request GetBackendSetHealthRequest) (response GetBackendSetHealthResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getBackendSetHealth, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetBackendSetHealthResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) getBackendSetHealth(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/health")
	if err != nil {
		return nil, err
	}

	var response GetBackendSetHealthResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetHealthChecker Gets the health check policy information for a given load balancer and backend set.
func (client LoadBalancerClient) GetHealthChecker(ctx context.Context, request GetHealthCheckerRequest) (response GetHealthCheckerResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getHealthChecker, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetHealthCheckerResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) getHealthChecker(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/healthChecker")
	if err != nil {
		return nil, err
	}

	var response GetHealthCheckerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetLoadBalancer Gets the specified load balancer's configuration information.
func (client LoadBalancerClient) GetLoadBalancer(ctx context.Context, request GetLoadBalancerRequest) (response GetLoadBalancerResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getLoadBalancer, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetLoadBalancerResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) getLoadBalancer(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/loadBalancers/{loadBalancerId}")
	if err != nil {
		return nil, err
	}

	var response GetLoadBalancerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetLoadBalancerHealth Gets the health status for the specified load balancer.
func (client LoadBalancerClient) GetLoadBalancerHealth(ctx context.Context, request GetLoadBalancerHealthRequest) (response GetLoadBalancerHealthResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getLoadBalancerHealth, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetLoadBalancerHealthResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) getLoadBalancerHealth(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/loadBalancers/{loadBalancerId}/health")
	if err != nil {
		return nil, err
	}

	var response GetLoadBalancerHealthResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// GetWorkRequest Gets the details of a work request.
func (client LoadBalancerClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.getWorkRequest, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(GetWorkRequestResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) getWorkRequest(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/loadBalancerWorkRequests/{workRequestId}")
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListBackendSets Lists all backend sets associated with a given load balancer.
func (client LoadBalancerClient) ListBackendSets(ctx context.Context, request ListBackendSetsRequest) (response ListBackendSetsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listBackendSets, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListBackendSetsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) listBackendSets(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/loadBalancers/{loadBalancerId}/backendSets")
	if err != nil {
		return nil, err
	}

	var response ListBackendSetsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListBackends Lists the backend servers for a given load balancer and backend set.
func (client LoadBalancerClient) ListBackends(ctx context.Context, request ListBackendsRequest) (response ListBackendsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listBackends, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListBackendsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) listBackends(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/backends")
	if err != nil {
		return nil, err
	}

	var response ListBackendsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListCertificates Lists all SSL certificates associated with a given load balancer.
func (client LoadBalancerClient) ListCertificates(ctx context.Context, request ListCertificatesRequest) (response ListCertificatesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listCertificates, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListCertificatesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) listCertificates(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/loadBalancers/{loadBalancerId}/certificates")
	if err != nil {
		return nil, err
	}

	var response ListCertificatesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListLoadBalancerHealths Lists the summary health statuses for all load balancers in the specified compartment.
func (client LoadBalancerClient) ListLoadBalancerHealths(ctx context.Context, request ListLoadBalancerHealthsRequest) (response ListLoadBalancerHealthsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listLoadBalancerHealths, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListLoadBalancerHealthsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) listLoadBalancerHealths(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/loadBalancerHealths")
	if err != nil {
		return nil, err
	}

	var response ListLoadBalancerHealthsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListLoadBalancers Lists all load balancers in the specified compartment.
func (client LoadBalancerClient) ListLoadBalancers(ctx context.Context, request ListLoadBalancersRequest) (response ListLoadBalancersResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listLoadBalancers, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListLoadBalancersResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) listLoadBalancers(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/loadBalancers")
	if err != nil {
		return nil, err
	}

	var response ListLoadBalancersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListPolicies Lists the available load balancer policies.
func (client LoadBalancerClient) ListPolicies(ctx context.Context, request ListPoliciesRequest) (response ListPoliciesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listPolicies, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListPoliciesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) listPolicies(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/loadBalancerPolicies")
	if err != nil {
		return nil, err
	}

	var response ListPoliciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListProtocols Lists all supported traffic protocols.
func (client LoadBalancerClient) ListProtocols(ctx context.Context, request ListProtocolsRequest) (response ListProtocolsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listProtocols, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListProtocolsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) listProtocols(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/loadBalancerProtocols")
	if err != nil {
		return nil, err
	}

	var response ListProtocolsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListShapes Lists the valid load balancer shapes.
func (client LoadBalancerClient) ListShapes(ctx context.Context, request ListShapesRequest) (response ListShapesResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listShapes, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListShapesResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) listShapes(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/loadBalancerShapes")
	if err != nil {
		return nil, err
	}

	var response ListShapesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// ListWorkRequests Lists the work requests for a given load balancer.
func (client LoadBalancerClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.listWorkRequests, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(ListWorkRequestsResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) listWorkRequests(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodGet, "/loadBalancers/{loadBalancerId}/workRequests")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateBackend Updates the configuration of a backend server within the specified backend set.
func (client LoadBalancerClient) UpdateBackend(ctx context.Context, request UpdateBackendRequest) (response UpdateBackendResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateBackend, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateBackendResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) updateBackend(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/backends/{backendName}")
	if err != nil {
		return nil, err
	}

	var response UpdateBackendResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateBackendSet Updates a backend set.
func (client LoadBalancerClient) UpdateBackendSet(ctx context.Context, request UpdateBackendSetRequest) (response UpdateBackendSetResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateBackendSet, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateBackendSetResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) updateBackendSet(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}")
	if err != nil {
		return nil, err
	}

	var response UpdateBackendSetResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateHealthChecker Updates the health check policy for a given load balancer and backend set.
func (client LoadBalancerClient) UpdateHealthChecker(ctx context.Context, request UpdateHealthCheckerRequest) (response UpdateHealthCheckerResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateHealthChecker, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateHealthCheckerResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) updateHealthChecker(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/healthChecker")
	if err != nil {
		return nil, err
	}

	var response UpdateHealthCheckerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateListener Updates a listener for a given load balancer.
func (client LoadBalancerClient) UpdateListener(ctx context.Context, request UpdateListenerRequest) (response UpdateListenerResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateListener, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateListenerResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) updateListener(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/loadBalancers/{loadBalancerId}/listeners/{listenerName}")
	if err != nil {
		return nil, err
	}

	var response UpdateListenerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}

// UpdateLoadBalancer Updates a load balancer's configuration.
func (client LoadBalancerClient) UpdateLoadBalancer(ctx context.Context, request UpdateLoadBalancerRequest) (response UpdateLoadBalancerResponse, err error) {
	ociResponse, e := client.Retry(ctx, request, client.updateLoadBalancer, request.GetRetryPolicy())
	if e != nil {
		err = e
		return
	}
	if convertedResponse, convertedResponseOk := ociResponse.(UpdateLoadBalancerResponse); convertedResponseOk {
		response = convertedResponse
	}
	return
}

// lower camel case request => not exported
func (client LoadBalancerClient) updateLoadBalancer(ctx context.Context, request common.OciRequest) (common.OciResponse, error) {
	httpRequest, err := request.GetHttpRequest(http.MethodPut, "/loadBalancers/{loadBalancerId}")
	if err != nil {
		return nil, err
	}

	var response UpdateLoadBalancerResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, nil
}
