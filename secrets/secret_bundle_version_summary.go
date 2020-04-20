// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secrets
//
// API for retrieving secrets from vaults.
//

package secrets

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SecretBundleVersionSummary The properties of the secret bundle. (Secret bundle version summary objects do not include the actual contents of the secret.)
type SecretBundleVersionSummary struct {

	// The OCID of the secret.
	SecretId *string `mandatory:"true" json:"secretId"`

	// The version number of the secret.
	VersionNumber *int64 `mandatory:"true" json:"versionNumber"`

	// The time when the secret bundle was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The version name of the secret bundle, as provided when the secret was created or last rotated.
	VersionName *string `mandatory:"false" json:"versionName"`

	// An optional property indicating when to delete the secret version, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`

	// An optional property indicating when the secret version will expire, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfExpiry *common.SDKTime `mandatory:"false" json:"timeOfExpiry"`

	// A list of possible rotation states for the secret bundle.
	Stages []SecretBundleVersionSummaryStagesEnum `mandatory:"false" json:"stages,omitempty"`
}

func (m SecretBundleVersionSummary) String() string {
	return common.PointerString(m)
}

// SecretBundleVersionSummaryStagesEnum Enum with underlying type: string
type SecretBundleVersionSummaryStagesEnum string

// Set of constants representing the allowable values for SecretBundleVersionSummaryStagesEnum
const (
	SecretBundleVersionSummaryStagesCurrent    SecretBundleVersionSummaryStagesEnum = "CURRENT"
	SecretBundleVersionSummaryStagesPending    SecretBundleVersionSummaryStagesEnum = "PENDING"
	SecretBundleVersionSummaryStagesLatest     SecretBundleVersionSummaryStagesEnum = "LATEST"
	SecretBundleVersionSummaryStagesPrevious   SecretBundleVersionSummaryStagesEnum = "PREVIOUS"
	SecretBundleVersionSummaryStagesDeprecated SecretBundleVersionSummaryStagesEnum = "DEPRECATED"
)

var mappingSecretBundleVersionSummaryStages = map[string]SecretBundleVersionSummaryStagesEnum{
	"CURRENT":    SecretBundleVersionSummaryStagesCurrent,
	"PENDING":    SecretBundleVersionSummaryStagesPending,
	"LATEST":     SecretBundleVersionSummaryStagesLatest,
	"PREVIOUS":   SecretBundleVersionSummaryStagesPrevious,
	"DEPRECATED": SecretBundleVersionSummaryStagesDeprecated,
}

// GetSecretBundleVersionSummaryStagesEnumValues Enumerates the set of values for SecretBundleVersionSummaryStagesEnum
func GetSecretBundleVersionSummaryStagesEnumValues() []SecretBundleVersionSummaryStagesEnum {
	values := make([]SecretBundleVersionSummaryStagesEnum, 0)
	for _, v := range mappingSecretBundleVersionSummaryStages {
		values = append(values, v)
	}
	return values
}
