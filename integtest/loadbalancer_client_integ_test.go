// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package integtest

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/loadbalancer"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testRegionForLoadBalancer = common.REGION_PHX
)

func TestLoadBalancerClient_CreateBackend(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.CreateBackendRequest{}
	err := c.CreateBackend(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_CreateBackendSet(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.CreateBackendSetRequest{}
	err := c.CreateBackendSet(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_CreateCertificate(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.CreateCertificateRequest{}
	err := c.CreateCertificate(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_CreateListener(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.CreateListenerRequest{}
	err := c.CreateListener(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_CreateLoadBalancer(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.CreateLoadBalancerRequest{}
	err := c.CreateLoadBalancer(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_DeleteBackend(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.DeleteBackendRequest{}
	err := c.DeleteBackend(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_DeleteBackendSet(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.DeleteBackendSetRequest{}
	err := c.DeleteBackendSet(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_DeleteCertificate(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.DeleteCertificateRequest{}
	err := c.DeleteCertificate(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_DeleteListener(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.DeleteListenerRequest{}
	err := c.DeleteListener(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_DeleteLoadBalancer(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.DeleteLoadBalancerRequest{}
	err := c.DeleteLoadBalancer(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_GetBackend(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.GetBackendRequest{}
	r, err := c.GetBackend(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_GetBackendHealth(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.GetBackendHealthRequest{}
	r, err := c.GetBackendHealth(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_GetBackendSet(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.GetBackendSetRequest{}
	r, err := c.GetBackendSet(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_GetBackendSetHealth(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.GetBackendSetHealthRequest{}
	r, err := c.GetBackendSetHealth(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_GetHealthChecker(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.GetHealthCheckerRequest{}
	r, err := c.GetHealthChecker(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_GetLoadBalancer(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.GetLoadBalancerRequest{}
	r, err := c.GetLoadBalancer(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_GetLoadBalancerHealth(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.GetLoadBalancerHealthRequest{}
	r, err := c.GetLoadBalancerHealth(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_GetWorkRequest(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.GetWorkRequestRequest{}
	r, err := c.GetWorkRequest(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_ListBackendSets(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.ListBackendSetsRequest{}
	r, err := c.ListBackendSets(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_ListBackends(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.ListBackendsRequest{}
	r, err := c.ListBackends(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_ListCertificates(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.ListCertificatesRequest{}
	r, err := c.ListCertificates(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_ListLoadBalancerHealths(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.ListLoadBalancerHealthsRequest{}
	r, err := c.ListLoadBalancerHealths(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_ListLoadBalancers(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.ListLoadBalancersRequest{}
	r, err := c.ListLoadBalancers(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_ListPolicies(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.ListPoliciesRequest{}
	r, err := c.ListPolicies(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_ListProtocols(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.ListProtocolsRequest{}
	r, err := c.ListProtocols(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_ListShapes(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.ListShapesRequest{}
	r, err := c.ListShapes(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_ListWorkRequests(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.ListWorkRequestsRequest{}
	r, err := c.ListWorkRequests(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_UpdateBackend(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.UpdateBackendRequest{}
	err := c.UpdateBackend(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_UpdateBackendSet(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.UpdateBackendSetRequest{}
	err := c.UpdateBackendSet(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_UpdateHealthChecker(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.UpdateHealthCheckerRequest{}
	err := c.UpdateHealthChecker(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_UpdateListener(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.UpdateListenerRequest{}
	err := c.UpdateListener(context.Background(), request)
	assert.NoError(t, err)
	return
}

func TestLoadBalancerClient_UpdateLoadBalancer(t *testing.T) {
	t.Skip("Not implemented")
	c := loadbalancer.NewLoadBalancerClientForRegion(testRegionForLoadBalancer)
	request := loadbalancer.UpdateLoadBalancerRequest{}
	err := c.UpdateLoadBalancer(context.Background(), request)
	assert.NoError(t, err)
	return
}
