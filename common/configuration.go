// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.
package common

// Configuration provides configuration options for client requests
type ClientRequestConfiguration struct {
	Timeout int32
}

type Configuration struct {
	ClientRequestConfiguration
	Region string
	Endpoint string
}

