// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Audit API
//
// API for the Audit Service. You can use this API for queries, but not bulk-export operations.
//

package integtest

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/audit"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testRegionForAudit = common.REGION_PHX
)

func TestAuditClient_GetConfiguration(t *testing.T) {
	t.Skip("Not implemented")
	c := audit.NewAuditClientForRegion(testRegionForAudit)
	request := audit.GetConfigurationRequest{}
	r, err := c.GetConfiguration(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestAuditClient_GetWorkRequest(t *testing.T) {
	t.Skip("Not implemented")
	c := audit.NewAuditClientForRegion(testRegionForAudit)
	request := audit.GetWorkRequestRequest{}
	r, err := c.GetWorkRequest(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestAuditClient_ListEvents(t *testing.T) {
	t.Skip("Not implemented")
	c := audit.NewAuditClientForRegion(testRegionForAudit)
	request := audit.ListEventsRequest{}
	r, err := c.ListEvents(context.Background(), request)
	assert.NotEmpty(t, r, fmt.Sprint(r))
	assert.NoError(t, err)
	return
}

func TestAuditClient_UpdateConfiguration(t *testing.T) {
	t.Skip("Not implemented")
	c := audit.NewAuditClientForRegion(testRegionForAudit)
	request := audit.UpdateConfigurationRequest{}
	err := c.UpdateConfiguration(context.Background(), request)
	assert.NoError(t, err)
	return
}
