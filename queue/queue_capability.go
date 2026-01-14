// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Queue API
//
// Use the Queue API to produce and consume messages, create queues, and manage related items. For more information, see Queue (https://docs.oracle.com/iaas/Content/queue/overview.htm).
//

package queue

import (
	"github.com/oracle/oci-go-sdk/v65/common"
)

// QueueCapability A capability configuration for a queue.
type QueueCapability struct {

	// The type of capability.
	// Example: CONSUMER_GROUPS
	Type *string `mandatory:"false" json:"type,omitempty"`

	// Whether the primary consumer group is enabled for this capability.
	IsPrimaryConsumerGroupEnabled *bool `mandatory:"false" json:"isPrimaryConsumerGroupEnabled,omitempty"`
}

func (m QueueCapability) String() string {
	return common.PointerString(m)
}
