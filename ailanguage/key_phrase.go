// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Infrastructure Artificial Intelligence Services API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately using our proven,
//     pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
//     This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"github.com/oracle/oci-go-sdk/v40/common"
)

// KeyPhrase Key phrase and score for the given text.
type KeyPhrase struct {

	// Key phrase of the the given text.
	Text *string `mandatory:"true" json:"text"`

	// Score of the given key phrase.
	Score *float64 `mandatory:"true" json:"score"`
}

func (m KeyPhrase) String() string {
	return common.PointerString(m)
}
