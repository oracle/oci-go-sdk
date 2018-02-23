// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
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

//LoadBalancerClientData a client for LoadBalancer
type LoadBalancerClientData struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewLoadBalancerClientWithConfigurationProvider Creates a new default LoadBalancer client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewLoadBalancerClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (LoadBalancerClient, error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return nil, err
	}

	client := LoadBalancerClientData{BaseClient: baseClient}
	client.BasePath = "20170115"
	err = client.setConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	return &client, err
}

// GetBaseClient get the BaseClient object of this client
func (client *LoadBalancerClientData) GetBaseClient() *common.BaseClient {
	return &client.BaseClient
}

// SetRegion overrides the region of this client.
func (client *LoadBalancerClientData) SetRegion(region string) {
	client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "iaas", region)
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *LoadBalancerClientData) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *LoadBalancerClientData) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CreateBackend Adds a backend server to a backend set.
func (client LoadBalancerClientData) CreateBackend(ctx context.Context, request CreateBackendRequest) (response CreateBackendResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/backends", request)
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

// CreateBackendSet Adds a backend set to a load balancer.
func (client LoadBalancerClientData) CreateBackendSet(ctx context.Context, request CreateBackendSetRequest) (response CreateBackendSetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/loadBalancers/{loadBalancerId}/backendSets", request)
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

// CreateCertificate Creates an asynchronous request to add an SSL certificate.
func (client LoadBalancerClientData) CreateCertificate(ctx context.Context, request CreateCertificateRequest) (response CreateCertificateResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/loadBalancers/{loadBalancerId}/certificates", request)
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

// CreateListener Adds a listener to a load balancer.
func (client LoadBalancerClientData) CreateListener(ctx context.Context, request CreateListenerRequest) (response CreateListenerResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/loadBalancers/{loadBalancerId}/listeners", request)
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

// CreateLoadBalancer Creates a new load balancer in the specified compartment. For general information about load balancers,
// see Overview of the Load Balancing Service ({{DOC_SERVER_URL}}/Content/Balance/Concepts/balanceoverview.htm).
// For the purposes of access control, you must provide the OCID of the compartment where you want
// the load balancer to reside. Notice that the load balancer doesn't have to be in the same compartment as the VCN
// or backend set. If you're not sure which compartment to use, put the load balancer in the same compartment as the VCN.
// For information about access control and compartments, see
// Overview of the IAM Service ({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm).
// You must specify a display name for the load balancer. It does not have to be unique, and you can change it.
// For information about Availability Domains, see
// Regions and Availability Domains ({{DOC_SERVER_URL}}/Content/General/Concepts/regions.htm).
// To get a list of Availability Domains, use the `ListAvailabilityDomains` operation
// in the Identity and Access Management Service API.
// All Oracle Cloud Infrastructure resources, including load balancers, get an Oracle-assigned,
// unique ID called an Oracle Cloud Identifier (OCID). When you create a resource, you can find its OCID
// in the response. You can also retrieve a resource's OCID by using a List API operation on that resource type,
// or by viewing the resource in the Console. Fore more information, see
// Resource Identifiers ({{DOC_SERVER_URL}}/Content/General/Concepts/identifiers.htm).
// After you send your request, the new object's state will temporarily be PROVISIONING. Before using the
// object, first make sure its state has changed to RUNNING.
// When you create a load balancer, the system assigns an IP address.
// To get the IP address, use the GetLoadBalancer operation.
func (client LoadBalancerClientData) CreateLoadBalancer(ctx context.Context, request CreateLoadBalancerRequest) (response CreateLoadBalancerResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/loadBalancers", request)
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

// CreatePathRouteSet Adds a path route set on a load balancer.
func (client LoadBalancerClientData) CreatePathRouteSet(ctx context.Context, request CreatePathRouteSetRequest) (response CreatePathRouteSetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/loadBalancers/{loadBalancerId}/pathRouteSets", request)
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

// DeleteBackend Removes a backend server from a given load balancer and backend set.
func (client LoadBalancerClientData) DeleteBackend(ctx context.Context, request DeleteBackendRequest) (response DeleteBackendResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/backends/{backendName}", request)
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

// DeleteBackendSet Deletes the specified backend set. Note that deleting a backend set removes its backend servers from the load balancer.
// Before you can delete a backend set, you must remove it from any active listeners.
func (client LoadBalancerClientData) DeleteBackendSet(ctx context.Context, request DeleteBackendSetRequest) (response DeleteBackendSetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}", request)
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

// DeleteCertificate Deletes an SSL certificate from a load balancer.
func (client LoadBalancerClientData) DeleteCertificate(ctx context.Context, request DeleteCertificateRequest) (response DeleteCertificateResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/loadBalancers/{loadBalancerId}/certificates/{certificateName}", request)
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

// DeleteListener Deletes a listener from a load balancer.
func (client LoadBalancerClientData) DeleteListener(ctx context.Context, request DeleteListenerRequest) (response DeleteListenerResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/loadBalancers/{loadBalancerId}/listeners/{listenerName}", request)
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

// DeleteLoadBalancer Stops a load balancer and removes it from service.
func (client LoadBalancerClientData) DeleteLoadBalancer(ctx context.Context, request DeleteLoadBalancerRequest) (response DeleteLoadBalancerResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/loadBalancers/{loadBalancerId}", request)
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

// DeletePathRouteSet Deletes a Path Route Set from a Load Balancer.
func (client LoadBalancerClientData) DeletePathRouteSet(ctx context.Context, request DeletePathRouteSetRequest) (response DeletePathRouteSetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/loadBalancers/{loadBalancerId}/pathRouteSets/{pathRouteSetName}", request)
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

// GetBackend Gets the specified backend server's configuration information.
func (client LoadBalancerClientData) GetBackend(ctx context.Context, request GetBackendRequest) (response GetBackendResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/backends/{backendName}", request)
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

// GetBackendHealth Gets the current health status of the specified backend server.
func (client LoadBalancerClientData) GetBackendHealth(ctx context.Context, request GetBackendHealthRequest) (response GetBackendHealthResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/backends/{backendName}/health", request)
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

// GetBackendSet Gets the specified backend set's configuration information.
func (client LoadBalancerClientData) GetBackendSet(ctx context.Context, request GetBackendSetRequest) (response GetBackendSetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}", request)
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

// GetBackendSetHealth Gets the health status for the specified backend set.
func (client LoadBalancerClientData) GetBackendSetHealth(ctx context.Context, request GetBackendSetHealthRequest) (response GetBackendSetHealthResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/health", request)
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

// GetHealthChecker Gets the health check policy information for a given load balancer and backend set.
func (client LoadBalancerClientData) GetHealthChecker(ctx context.Context, request GetHealthCheckerRequest) (response GetHealthCheckerResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/healthChecker", request)
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

// GetLoadBalancer Gets the specified load balancer's configuration information.
func (client LoadBalancerClientData) GetLoadBalancer(ctx context.Context, request GetLoadBalancerRequest) (response GetLoadBalancerResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancers/{loadBalancerId}", request)
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

// GetLoadBalancerHealth Gets the health status for the specified load balancer.
func (client LoadBalancerClientData) GetLoadBalancerHealth(ctx context.Context, request GetLoadBalancerHealthRequest) (response GetLoadBalancerHealthResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancers/{loadBalancerId}/health", request)
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

// GetPathRouteSet Gets the specified path route set's configuration information.
func (client LoadBalancerClientData) GetPathRouteSet(ctx context.Context, request GetPathRouteSetRequest) (response GetPathRouteSetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancers/{loadBalancerId}/pathRouteSets/{pathRouteSetName}", request)
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

// GetWorkRequest Gets the details of a work request.
func (client LoadBalancerClientData) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancerWorkRequests/{workRequestId}", request)
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

// ListBackendSets Lists all backend sets associated with a given load balancer.
func (client LoadBalancerClientData) ListBackendSets(ctx context.Context, request ListBackendSetsRequest) (response ListBackendSetsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancers/{loadBalancerId}/backendSets", request)
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

// ListBackends Lists the backend servers for a given load balancer and backend set.
func (client LoadBalancerClientData) ListBackends(ctx context.Context, request ListBackendsRequest) (response ListBackendsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/backends", request)
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

// ListCertificates Lists all SSL certificates associated with a given load balancer.
func (client LoadBalancerClientData) ListCertificates(ctx context.Context, request ListCertificatesRequest) (response ListCertificatesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancers/{loadBalancerId}/certificates", request)
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

// ListLoadBalancerHealths Lists the summary health statuses for all load balancers in the specified compartment.
func (client LoadBalancerClientData) ListLoadBalancerHealths(ctx context.Context, request ListLoadBalancerHealthsRequest) (response ListLoadBalancerHealthsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancerHealths", request)
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

// ListLoadBalancers Lists all load balancers in the specified compartment.
func (client LoadBalancerClientData) ListLoadBalancers(ctx context.Context, request ListLoadBalancersRequest) (response ListLoadBalancersResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancers", request)
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

// ListPathRouteSets Lists all Path Route sets associated with a given load balancer.
func (client LoadBalancerClientData) ListPathRouteSets(ctx context.Context, request ListPathRouteSetsRequest) (response ListPathRouteSetsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancers/{loadBalancerId}/pathRouteSets", request)
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

// ListPolicies Lists the available load balancer policies.
func (client LoadBalancerClientData) ListPolicies(ctx context.Context, request ListPoliciesRequest) (response ListPoliciesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancerPolicies", request)
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

// ListProtocols Lists all supported traffic protocols.
func (client LoadBalancerClientData) ListProtocols(ctx context.Context, request ListProtocolsRequest) (response ListProtocolsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancerProtocols", request)
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

// ListShapes Lists the valid load balancer shapes.
func (client LoadBalancerClientData) ListShapes(ctx context.Context, request ListShapesRequest) (response ListShapesResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancerShapes", request)
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

// ListWorkRequests Lists the work requests for a given load balancer.
func (client LoadBalancerClientData) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/loadBalancers/{loadBalancerId}/workRequests", request)
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

// UpdateBackend Updates the configuration of a backend server within the specified backend set.
func (client LoadBalancerClientData) UpdateBackend(ctx context.Context, request UpdateBackendRequest) (response UpdateBackendResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/backends/{backendName}", request)
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

// UpdateBackendSet Updates a backend set.
func (client LoadBalancerClientData) UpdateBackendSet(ctx context.Context, request UpdateBackendSetRequest) (response UpdateBackendSetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}", request)
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

// UpdateHealthChecker Updates the health check policy for a given load balancer and backend set.
func (client LoadBalancerClientData) UpdateHealthChecker(ctx context.Context, request UpdateHealthCheckerRequest) (response UpdateHealthCheckerResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/healthChecker", request)
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

// UpdateListener Updates a listener for a given load balancer.
func (client LoadBalancerClientData) UpdateListener(ctx context.Context, request UpdateListenerRequest) (response UpdateListenerResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/loadBalancers/{loadBalancerId}/listeners/{listenerName}", request)
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

// UpdateLoadBalancer Updates a load balancer's configuration.
func (client LoadBalancerClientData) UpdateLoadBalancer(ctx context.Context, request UpdateLoadBalancerRequest) (response UpdateLoadBalancerResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/loadBalancers/{loadBalancerId}", request)
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

// UpdatePathRouteSet Updates and overwrite a path route set on a load balancer.
func (client LoadBalancerClientData) UpdatePathRouteSet(ctx context.Context, request UpdatePathRouteSetRequest) (response UpdatePathRouteSetResponse, err error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/loadBalancers/{loadBalancerId}/pathRouteSets/{pathRouteSetName}", request)
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

// LoadBalancerClient is client interface for LoadBalancer
type LoadBalancerClient interface {
	GetBaseClient() *common.BaseClient
	ConfigurationProvider() *common.ConfigurationProvider
	CreateBackend(ctx context.Context, request CreateBackendRequest) (response CreateBackendResponse, err error)
	CreateBackendSet(ctx context.Context, request CreateBackendSetRequest) (response CreateBackendSetResponse, err error)
	CreateCertificate(ctx context.Context, request CreateCertificateRequest) (response CreateCertificateResponse, err error)
	CreateListener(ctx context.Context, request CreateListenerRequest) (response CreateListenerResponse, err error)
	CreateLoadBalancer(ctx context.Context, request CreateLoadBalancerRequest) (response CreateLoadBalancerResponse, err error)
	CreatePathRouteSet(ctx context.Context, request CreatePathRouteSetRequest) (response CreatePathRouteSetResponse, err error)
	DeleteBackend(ctx context.Context, request DeleteBackendRequest) (response DeleteBackendResponse, err error)
	DeleteBackendSet(ctx context.Context, request DeleteBackendSetRequest) (response DeleteBackendSetResponse, err error)
	DeleteCertificate(ctx context.Context, request DeleteCertificateRequest) (response DeleteCertificateResponse, err error)
	DeleteListener(ctx context.Context, request DeleteListenerRequest) (response DeleteListenerResponse, err error)
	DeleteLoadBalancer(ctx context.Context, request DeleteLoadBalancerRequest) (response DeleteLoadBalancerResponse, err error)
	DeletePathRouteSet(ctx context.Context, request DeletePathRouteSetRequest) (response DeletePathRouteSetResponse, err error)
	GetBackend(ctx context.Context, request GetBackendRequest) (response GetBackendResponse, err error)
	GetBackendHealth(ctx context.Context, request GetBackendHealthRequest) (response GetBackendHealthResponse, err error)
	GetBackendSet(ctx context.Context, request GetBackendSetRequest) (response GetBackendSetResponse, err error)
	GetBackendSetHealth(ctx context.Context, request GetBackendSetHealthRequest) (response GetBackendSetHealthResponse, err error)
	GetHealthChecker(ctx context.Context, request GetHealthCheckerRequest) (response GetHealthCheckerResponse, err error)
	GetLoadBalancer(ctx context.Context, request GetLoadBalancerRequest) (response GetLoadBalancerResponse, err error)
	GetLoadBalancerHealth(ctx context.Context, request GetLoadBalancerHealthRequest) (response GetLoadBalancerHealthResponse, err error)
	GetPathRouteSet(ctx context.Context, request GetPathRouteSetRequest) (response GetPathRouteSetResponse, err error)
	GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error)
	ListBackendSets(ctx context.Context, request ListBackendSetsRequest) (response ListBackendSetsResponse, err error)
	ListBackends(ctx context.Context, request ListBackendsRequest) (response ListBackendsResponse, err error)
	ListCertificates(ctx context.Context, request ListCertificatesRequest) (response ListCertificatesResponse, err error)
	ListLoadBalancerHealths(ctx context.Context, request ListLoadBalancerHealthsRequest) (response ListLoadBalancerHealthsResponse, err error)
	ListLoadBalancers(ctx context.Context, request ListLoadBalancersRequest) (response ListLoadBalancersResponse, err error)
	ListPathRouteSets(ctx context.Context, request ListPathRouteSetsRequest) (response ListPathRouteSetsResponse, err error)
	ListPolicies(ctx context.Context, request ListPoliciesRequest) (response ListPoliciesResponse, err error)
	ListProtocols(ctx context.Context, request ListProtocolsRequest) (response ListProtocolsResponse, err error)
	ListShapes(ctx context.Context, request ListShapesRequest) (response ListShapesResponse, err error)
	ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error)
	UpdateBackend(ctx context.Context, request UpdateBackendRequest) (response UpdateBackendResponse, err error)
	UpdateBackendSet(ctx context.Context, request UpdateBackendSetRequest) (response UpdateBackendSetResponse, err error)
	UpdateHealthChecker(ctx context.Context, request UpdateHealthCheckerRequest) (response UpdateHealthCheckerResponse, err error)
	UpdateListener(ctx context.Context, request UpdateListenerRequest) (response UpdateListenerResponse, err error)
	UpdateLoadBalancer(ctx context.Context, request UpdateLoadBalancerRequest) (response UpdateLoadBalancerResponse, err error)
	UpdatePathRouteSet(ctx context.Context, request UpdatePathRouteSetRequest) (response UpdatePathRouteSetResponse, err error)
}
