// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
//
// Helper methods for OCI GOSDK Samples
//

package helpers

import (
	"sync"
)

// SafeResource tracks created resources to be cleaned up
var safeResource = struct {
	mux sync.RWMutex
	m   map[string]string
}{m: make(map[string]string)}

// AddResource add new created resource into a local map
func AddResource(resourceID string, resourceName string) {
	safeResource.mux.Lock()
	safeResource.m[resourceName] = resourceID
	safeResource.mux.Unlock()
}

// RemoveResource removes a resource after resource clean up succeed
func RemoveResource(resourceID string) {
	safeResource.mux.Lock()
	delete(safeResource.m, resourceID)
	safeResource.mux.Unlock()
}

// GetResources return a copy of all resources need to be cleaned up
func GetResources() map[string]string {
	return safeResource.m
}
