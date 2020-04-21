// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package transfer

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/example/helpers"
)

func TestUploadStreamMultiparts(t *testing.T) {
	type testStruct struct {
		uploadID                 string
		failedPartNumbers        []int
		expectedNumOfUploadParts int
		expectedNumOfCommitParts int
	}

	testDataSet := []testStruct{
		{
			uploadID:                 "id1",
			failedPartNumbers:        []int{},
			expectedNumOfUploadParts: 10,
			expectedNumOfCommitParts: 10,
		},
		{
			uploadID:                 "id2",
			failedPartNumbers:        []int{1, 2},
			expectedNumOfUploadParts: 8,
			expectedNumOfCommitParts: 0,
		},
	}

	streamUpload := streamUpload{}
	var fileSize, partSize int64
	fileSize = 100
	partSize = 10

	ctx := context.Background()
	for _, testData := range testDataSet {
		fake := fake{upLoadID: testData.uploadID, failedPartNumbers: testData.failedPartNumbers, numberOfCommitedParts: common.Int(0), numberOfUploadedParts: common.Int(0)}
		streamUpload.multipartUploader = &fake
		filePath, _ := helpers.WriteTempFileOfSize(fileSize)

		file, err := os.Open(filePath)
		defer file.Close()
		assert.NoError(t, err)

		request := UploadStreamRequest{
			UploadRequest: UploadRequest{PartSize: common.Int64(partSize)},
			StreamReader:  file,
		}

		request.initDefaultValues()
		streamUpload.UploadStream(ctx, request)
		assert.NotEmpty(t, streamUpload.manifest.parts)

		// all parts have been committed
		totalParts := int(fileSize / partSize)
		failedParts := len(fake.failedPartNumbers)
		assert.Equal(t, testData.expectedNumOfUploadParts, totalParts-failedParts)
		assert.NoError(t, err)

		if failedParts != 0 {
			// not parts should be commit as there are failed parts
			assert.Equal(t, 0, *fake.numberOfCommitedParts)
		}

		// finally, all parts should be committed
		assert.Equal(t, testData.expectedNumOfCommitParts, *fake.numberOfCommitedParts)
	}
}
