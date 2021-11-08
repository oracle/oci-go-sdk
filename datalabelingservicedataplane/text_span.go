// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DlsDataPlane API
//
// A description of the DlsDataPlane API.
//

package datalabelingservicedataplane

import (
	"github.com/oracle/oci-go-sdk/v51/common"
)

// TextSpan A wrapper class for offset and length, which together represent a span of text in a text document.
type TextSpan struct {

	// Offset of the selected text within the entire text.
	Offset *float32 `mandatory:"false" json:"offset"`

	// Length of the selected text.
	Length *float32 `mandatory:"false" json:"length"`
}

func (m TextSpan) String() string {
	return common.PointerString(m)
}
