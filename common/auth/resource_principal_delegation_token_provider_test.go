// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package auth

import (
	"testing"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/stretchr/testify/assert"
)

func TestResourcePrincipalDelegationTokenConfigurationProvider_ErrorInput(t *testing.T) {
	delegationToken := ""
	region := common.StringToRegion("us-ashburn-1")
	configurationProvider, err := ResourcePrincipalDelegationTokenConfigurationProvider(&delegationToken)
	assert.Nil(t, configurationProvider)
	assert.NotNil(t, err)

	configurationProvider, err = ResourcePrincipalDelegationTokenConfigurationProvider(nil)
	assert.Nil(t, configurationProvider)
	assert.NotNil(t, err)

	configurationProviderForRegion, err := ResourcePrincipalDelegationTokenConfigurationProviderForRegion(&delegationToken, region)
	assert.Nil(t, configurationProviderForRegion)
	assert.NotNil(t, err)

	configurationProviderForRegionNotFound, err := ResourcePrincipalDelegationTokenConfigurationProviderForRegion(&delegationToken, "")
	assert.Nil(t, configurationProviderForRegionNotFound)
	assert.NotNil(t, err)
}
