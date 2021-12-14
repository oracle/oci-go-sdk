// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v54/common"
)

// SentimentDocumentResult The document response for sentiment detect call.
type SentimentDocumentResult struct {

	// Document Unique Identifier.
	Key *string `mandatory:"true" json:"key"`

	// List of detected aspects sentiment.
	Aspects []SentimentAspect `mandatory:"true" json:"aspects"`

	// Language code as per ISO 639-1 (https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) standard.
	LanguageCode *string `mandatory:"true" json:"languageCode"`

	// Document level sentiment.
	DocumentSentiment *string `mandatory:"false" json:"documentSentiment"`

	// Scores for each sentiment.
	// Example: {"positive": 1.0, "negative": 0.0}
	DocumentScores map[string]float64 `mandatory:"false" json:"documentScores"`

	// List of detected sentences sentiment.
	Sentences []SentimentSentence `mandatory:"false" json:"sentences"`
}

func (m SentimentDocumentResult) String() string {
	return common.PointerString(m)
}
