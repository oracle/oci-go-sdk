// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MarketplacePublisherService API
//
// Use the Marketplace Publisher API to manage the publishing of applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplacepublisher

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateHelmChartImageDetails Helmchart image details required to create an helmchart artifact.
type CreateHelmChartImageDetails struct {

	// The source registry OCID of the container image.
	SourceRegistryId *string `mandatory:"true" json:"sourceRegistryId"`

	// The source registry url of the helmchart image.
	SourceRegistryUrl *string `mandatory:"true" json:"sourceRegistryUrl"`

	// The Supported Versions of Kubernetes
	SupportedKubernetesVersions []string `mandatory:"false" json:"supportedKubernetesVersions"`
}

func (m CreateHelmChartImageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateHelmChartImageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
