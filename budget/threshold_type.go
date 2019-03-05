// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// BudgetsControlPlane API
//
// Use the BudgetsControlPlane API to manage budgets and budget alerts.
//

package budget

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ThresholdType The type of threshold. Valid values are PERCENTAGE or ABSOLUTE.
type ThresholdType struct {
}

func (m ThresholdType) String() string {
	return common.PointerString(m)
}
