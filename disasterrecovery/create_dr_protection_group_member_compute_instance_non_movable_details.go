// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDrProtectionGroupMemberComputeInstanceNonMovableDetails Create properties for a non-movable compute instance member.
type CreateDrProtectionGroupMemberComputeInstanceNonMovableDetails struct {

	// The OCID of the member.
	// Example: `ocid1.instance.oc1..uniqueID`
	MemberId *string `mandatory:"true" json:"memberId"`

	// A flag indicating whether the non-movable compute instance should be started and stopped during DR operations.
	// *Prechecks cannot be executed on stopped instances that are configured to be started.*
	IsStartStopEnabled *bool `mandatory:"false" json:"isStartStopEnabled"`

	// A list of operations performed on file systems used by the compute instance.
	FileSystemOperations []CreateComputeInstanceNonMovableFileSystemOperationDetails `mandatory:"false" json:"fileSystemOperations"`

	// Deprecated. Use the 'blockVolumeAttachAndMountOperations' attribute instead of this.
	// A list of operations performed on block volumes used by the compute instance.
	BlockVolumeOperations []CreateComputeInstanceNonMovableBlockVolumeOperationDetails `mandatory:"false" json:"blockVolumeOperations"`

	BlockVolumeAttachAndMountOperations *CreateComputeInstanceNonMovableBlockVolumeAttachAndMountOperationsDetails `mandatory:"false" json:"blockVolumeAttachAndMountOperations"`
}

// GetMemberId returns MemberId
func (m CreateDrProtectionGroupMemberComputeInstanceNonMovableDetails) GetMemberId() *string {
	return m.MemberId
}

func (m CreateDrProtectionGroupMemberComputeInstanceNonMovableDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDrProtectionGroupMemberComputeInstanceNonMovableDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDrProtectionGroupMemberComputeInstanceNonMovableDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDrProtectionGroupMemberComputeInstanceNonMovableDetails CreateDrProtectionGroupMemberComputeInstanceNonMovableDetails
	s := struct {
		DiscriminatorParam string `json:"memberType"`
		MarshalTypeCreateDrProtectionGroupMemberComputeInstanceNonMovableDetails
	}{
		"COMPUTE_INSTANCE_NON_MOVABLE",
		(MarshalTypeCreateDrProtectionGroupMemberComputeInstanceNonMovableDetails)(m),
	}

	return json.Marshal(&s)
}
