// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstancePrincipalDelegationTokenConfigurationProvider_ErrorInput(t *testing.T) {
	delegationToken := ""
	configurationProvider, err := InstancePrincipalDelegationTokenConfigurationProvider(&delegationToken)
	assert.Nil(t, configurationProvider)
	assert.NotNil(t, err)

	configurationProvider, err = InstancePrincipalDelegationTokenConfigurationProvider(nil)
	assert.Nil(t, configurationProvider)
	assert.NotNil(t, err)
}
