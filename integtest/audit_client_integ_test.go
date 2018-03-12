// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package integtest

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/oracle/oci-go-sdk/audit"
	"github.com/oracle/oci-go-sdk/common"
)

// GetConfiguration test
func TestAuditClient_GetConfiguration(t *testing.T) {
	c, clerr := audit.NewAuditClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	req := audit.GetConfigurationRequest{
		CompartmentId: common.String(getTenancyID()),
	}

	resp, err := c.GetConfiguration(context.Background(), req)
	failIfError(t, err)
	assert.NotEmpty(t, resp.RetentionPeriodDays)
}

// UpdateConfiguration test
func TestAuditClient_UpdateConfiguration(t *testing.T) {
	c, clerr := audit.NewAuditClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	req := audit.GetConfigurationRequest{
		CompartmentId: common.String(getTenancyID()),
	}

	resp, err := c.GetConfiguration(context.Background(), req)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.RetentionPeriodDays)

	retentionPeriodDays := *resp.RetentionPeriodDays

	updateReq := audit.UpdateConfigurationRequest{
		CompartmentId: common.String(getTenancyID()),
	}

	// update the config
	updateReq.RetentionPeriodDays = common.Int(retentionPeriodDays + 1)
	_, err = c.UpdateConfiguration(context.Background(), updateReq)
	assert.NoError(t, err)

	// check update succeed
	// ideally, we want to check the update succeed or not, but the update takes time to persist
	// and the response return a 202 http code indicates the update is processing
	// sleep here can help but will make the test un-reliable, so just check the error here
	/*resp, err = c.GetConfiguration(context.Background(), req)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.RetentionPeriodDays)
	assert.Equal(t, retentionPeriodDays+1, *resp.RetentionPeriodDays)*/

	// restore the config
	updateReq.RetentionPeriodDays = common.Int(retentionPeriodDays)
	_, err = c.UpdateConfiguration(context.Background(), updateReq)
	assert.NoError(t, err)
}

// ListEvents test
func TestAuditClient_ListEvents(t *testing.T) {
	c, clerr := audit.NewAuditClientWithConfigurationProvider(common.DefaultConfigProvider())
	failIfError(t, clerr)

	// list events for last 5 hour
	req := audit.ListEventsRequest{
		CompartmentId: common.String(getCompartmentID()),
		StartTime:     &common.SDKTime{time.Now().Add(time.Hour * -5)},
		EndTime:       &common.SDKTime{time.Now()},
	}

	resp, err := c.ListEvents(context.Background(), req)
	failIfError(t, err)
	assert.NotEmpty(t, resp)
}
