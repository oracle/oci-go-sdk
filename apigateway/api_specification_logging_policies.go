// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ApiSpecificationLoggingPolicies Policies controlling the pushing of logs to OCI Public Logging.
type ApiSpecificationLoggingPolicies struct {
	AccessLog *AccessLogPolicy `mandatory:"false" json:"accessLog"`

	ExecutionLog *ExecutionLogPolicy `mandatory:"false" json:"executionLog"`
}

func (m ApiSpecificationLoggingPolicies) String() string {
	return common.PointerString(m)
}
