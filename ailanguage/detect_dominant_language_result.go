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
	"github.com/oracle/oci-go-sdk/v49/common"
)

// DetectDominantLanguageResult Result of language detect call.
type DetectDominantLanguageResult struct {

	// Results are sorted in descending order of the scores. Most likely language is on top.
	// languages: [{"name": "Bosnian","code": "bs","score": 0.6942308391868572},
	//            {"name": "Croatian","code": "hr","score": 0.15768701487872652},
	//            {"name": "Serbo-Croatian","code": "sh","score": 0.1480651612334694}]
	Languages []DetectedLanguage `mandatory:"true" json:"languages"`
}

func (m DetectDominantLanguageResult) String() string {
	return common.PointerString(m)
}
