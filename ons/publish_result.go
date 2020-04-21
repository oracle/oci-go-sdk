// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Notifications API
//
// Use the Notifications API to broadcast messages to distributed components by topic, using a publish-subscribe pattern.
// For information about managing topics, subscriptions, and messages, see Notifications Overview (https://docs.cloud.oracle.com/iaas/Content/Notification/Concepts/notificationoverview.htm).
//

package ons

import (
	"github.com/oracle/oci-go-sdk/common"
)

// PublishResult The response to a PublishMessage call.
type PublishResult struct {

	// The UUID of the message.
	MessageId *string `mandatory:"true" json:"messageId"`

	// The time that the service received the message.
	TimeStamp *common.SDKTime `mandatory:"false" json:"timeStamp"`
}

func (m PublishResult) String() string {
	return common.PointerString(m)
}
