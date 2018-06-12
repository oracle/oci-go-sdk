// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.

// Clusters Service API
//
// Test for Container Engine for Kubernetes API

package integtest

import (
	"context"
	"testing"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/containerengine"
	"github.com/stretchr/testify/assert"
)

// ListClusters test -- simple test to make sure it can reach to container engine service API
func TestContainerEngineClient_ListClusters(t *testing.T) {
	c, clerr := containerengine.NewContainerEngineClientWithConfigurationProvider(configurationProvider())
	failIfError(t, clerr)

	// list events for last 5 hour
	req := containerengine.ListClustersRequest{
		CompartmentId: common.String(getCompartmentID()),
	}

	resp, err := c.ListClusters(context.Background(), req)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
}
