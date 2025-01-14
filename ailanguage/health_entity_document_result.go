// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HealthEntityDocumentResult The document response for health entity detect call.
type HealthEntityDocumentResult struct {

	// Document unique identifier defined by the user.
	Key *string `mandatory:"true" json:"key"`

	// List of detected entities.
	Entities []HealthEntity `mandatory:"true" json:"entities"`

	// Map of resolved entities by entity type
	ResolvedEntities map[string]ResolvedEntities `mandatory:"true" json:"resolvedEntities"`

	// Language code of the document. Please refer to respective model API documentation (https://docs.cloud.oracle.com/iaas/language/using/overview.htm) for supported languages.
	LanguageCode *string `mandatory:"true" json:"languageCode"`

	// List of succeeded document response.
	Relations []RelationEntity `mandatory:"false" json:"relations"`
}

func (m HealthEntityDocumentResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HealthEntityDocumentResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
