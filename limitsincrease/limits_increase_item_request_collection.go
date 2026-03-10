// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Limits Increase API
//
// Use the Limits Increase API to work with limit increase requests.
// For more information, see
// Working with Limit Increase Requests (https://docs.oracle.com/iaas/Content/General/service-limits/requests.htm).
//

package limitsincrease

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LimitsIncreaseItemRequestCollection Collection of property summaries for limit increase items within the indicated limit increase request.
type LimitsIncreaseItemRequestCollection struct {

	// List of limit increase items.
	Items []LimitsIncreaseItemRequestSummary `mandatory:"true" json:"items"`
}

func (m LimitsIncreaseItemRequestCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LimitsIncreaseItemRequestCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
