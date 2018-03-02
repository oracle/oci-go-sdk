// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package integtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/loadbalancer"
	"github.com/stretchr/testify/assert"
)

func TestLoadBalancerClient_CreateBackend(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.CreateBackendRequest{}
	_, err := c.CreateBackend(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_CreateBackendSet(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.CreateBackendSetRequest{}
	_, err := c.CreateBackendSet(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_CreateCertificate(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.CreateCertificateRequest{}
	_, err := c.CreateCertificate(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_CreateListener(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.CreateListenerRequest{}
	_, err := c.CreateListener(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_CreateLoadBalancer(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.CreateLoadBalancerRequest{}
	_, err := c.CreateLoadBalancer(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_DeleteBackend(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.DeleteBackendRequest{}
	_, err := c.DeleteBackend(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_DeleteBackendSet(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.DeleteBackendSetRequest{}
	_, err := c.DeleteBackendSet(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_DeleteCertificate(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.DeleteCertificateRequest{}
	_, err := c.DeleteCertificate(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_DeleteListener(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.DeleteListenerRequest{}
	_, err := c.DeleteListener(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_DeleteLoadBalancer(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.DeleteLoadBalancerRequest{}
	_, err := c.DeleteLoadBalancer(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_GetBackend(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.GetBackendRequest{}
	r, err := c.GetBackend(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_GetBackendHealth(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.GetBackendHealthRequest{}
	r, err := c.GetBackendHealth(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_GetBackendSet(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.GetBackendSetRequest{}
	r, err := c.GetBackendSet(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_GetBackendSetHealth(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.GetBackendSetHealthRequest{}
	r, err := c.GetBackendSetHealth(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_GetHealthChecker(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.GetHealthCheckerRequest{}
	r, err := c.GetHealthChecker(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_GetLoadBalancer(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.GetLoadBalancerRequest{}
	r, err := c.GetLoadBalancer(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_GetLoadBalancerHealth(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.GetLoadBalancerHealthRequest{}
	r, err := c.GetLoadBalancerHealth(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_GetWorkRequest(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.GetWorkRequestRequest{}
	r, err := c.GetWorkRequest(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_ListBackendSets(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.ListBackendSetsRequest{}
	r, err := c.ListBackendSets(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_ListBackends(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.ListBackendsRequest{}
	r, err := c.ListBackends(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_ListCertificates(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.ListCertificatesRequest{}
	r, err := c.ListCertificates(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_ListLoadBalancerHealths(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.ListLoadBalancerHealthsRequest{}
	r, err := c.ListLoadBalancerHealths(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_ListLoadBalancers(t *testing.T) {
	// make sure the list API will return at least one item
	createOrGetLoadBalancer(t)
	loadbalancers := listActiveLoadBalancers(t)
	assert.NotEmpty(t, loadbalancers, fmt.Sprint(loadbalancers))
	return
}

func TestLoadBalancerClient_ListPolicies(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.ListPoliciesRequest{}
	r, err := c.ListPolicies(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_ListProtocols(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.ListProtocolsRequest{}
	r, err := c.ListProtocols(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_ListShapes(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.ListShapesRequest{}
	r, err := c.ListShapes(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_ListWorkRequests(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.ListWorkRequestsRequest{}
	r, err := c.ListWorkRequests(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_UpdateBackend(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.UpdateBackendRequest{}
	_, err := c.UpdateBackend(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_UpdateBackendSet(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.UpdateBackendSetRequest{}
	_, err := c.UpdateBackendSet(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_UpdateHealthChecker(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.UpdateHealthCheckerRequest{}
	_, err := c.UpdateHealthChecker(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_UpdateListener(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.UpdateListenerRequest{}
	_, err := c.UpdateListener(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_UpdateLoadBalancer(t *testing.T) {
	t.Skip("Not implemented")
	c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)
	request := loadbalancer.UpdateLoadBalancerRequest{}
	_, err := c.UpdateLoadBalancer(context.Background(), request)
	assert.NoError(t, err)
	return
}
