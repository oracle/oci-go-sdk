// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package transfer

import (
	"errors"
	"io"

	"github.com/oracle/oci-go-sdk/common"
)

// UploadStreamRequest defines the input parameters for UploadFile method
type UploadStreamRequest struct {
	UploadRequest

	// The reader of input stream
	StreamReader io.Reader
}

var errorInvalidStream = errors.New("uploadStream is required")

const defaultStreamPartSize = 10 * 1024 * 1024 // 10MB

func (request UploadStreamRequest) validate() error {
	err := request.UploadRequest.validate()

	if err != nil {
		return err
	}

	if request.StreamReader == nil {
		return errorInvalidStream
	}

	return nil
}

func (request *UploadStreamRequest) initDefaultValues() error {
	if request.PartSize == nil {
		request.PartSize = common.Int64(defaultStreamPartSize)
	}

	return request.UploadRequest.initDefaultValues()
}
