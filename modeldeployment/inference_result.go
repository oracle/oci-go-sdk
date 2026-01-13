// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Model Deployment Data Plane API
//
// Model deployments are a managed resource in the OCI Data Science service to use to deploy machine learning models as HTTP endpoints in OCI. Deploying machine learning models as web applications (HTTP API endpoints) serving predictions in real time is the most common way that models are productionized. HTTP endpoints are flexible and can serve requests for model predictions.
// For more information, see Model Deployments (https://docs.oracle.com/en-us/iaas/data-science/using/model-dep-about.htm)
//

package modeldeployment

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InferenceResult A model used in x-related-resource for grouping actions with no returned body.
type InferenceResult struct {

	// The predict result returned by model
	Data []byte `mandatory:"true" json:"data"`
}

func (m InferenceResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InferenceResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
