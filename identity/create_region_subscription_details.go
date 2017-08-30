// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

type CreateRegionSubscriptionDetails struct {

	// The regions's key.
	// Allowed values are:
	// - `PHX`
	// - `IAD`
	// Example: `PHX`
	RegionKey string `json:"regionKey,omitempty"`
}
