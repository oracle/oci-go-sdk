// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
//     pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
//     This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// BatchDetectLanguageTextClassificationDetails The documents details for text classification call.
type BatchDetectLanguageTextClassificationDetails struct {

	// List of Documents for detect text classification.
	Documents []TextClassificationDocument `mandatory:"true" json:"documents"`
}

func (m BatchDetectLanguageTextClassificationDetails) String() string {
	return common.PointerString(m)
}
