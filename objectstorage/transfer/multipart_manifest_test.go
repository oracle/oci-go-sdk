// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package transfer

import (
	"os"
	"path"
	"testing"

	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/stretchr/testify/assert"
)

func TestSplitFileParts(t *testing.T) {
	type splitFilePartsTest struct {
		contentLen           int
		partSize             int64
		expectedPartNum      int
		expectedInitOffset   int64
		expectedLastOffset   int64
		expectedLastPartSize int64
		enableCheckSum       bool
	}

	testDataSet := []splitFilePartsTest{
		{
			contentLen:           100,
			partSize:             10,
			expectedPartNum:      10,
			expectedInitOffset:   10,
			expectedLastOffset:   90,
			expectedLastPartSize: 10,
			enableCheckSum:       false,
		},
		{
			contentLen:           30,
			partSize:             14,
			expectedPartNum:      3,
			expectedInitOffset:   14,
			expectedLastOffset:   28,
			expectedLastPartSize: 2,
			enableCheckSum:       false,
		},
	}

	for _, testData := range testDataSet {
		filePath, fileSize := helpers.WriteTempFileOfSize(int64(testData.contentLen))
		file, err := os.Open(filePath)
		assert.NoError(t, err)

		manifest := &multipartManifest{parts: make(map[string]map[int]uploadPart)}

		// UploadFileMultiparts closes the done channel when it returns; it may do so before
		// receiving all the values from result and errc channel
		done := make(chan struct{})
		partsChannel := manifest.splitFileToParts(done, testData.partSize, &testData.enableCheckSum, file, fileSize)

		// read through channel
		parts := []uploadPart{}
		for part := range partsChannel {
			assert.NoError(t, part.err)
			parts = append(parts, part)
		}

		assert.Equal(t, testData.expectedPartNum, len(parts))
		assert.Equal(t, testData.expectedInitOffset, parts[1].offset)
		assert.Equal(t, testData.partSize, parts[1].size)
		assert.Equal(t, testData.expectedLastOffset, parts[len(parts)-1].offset)
		assert.Equal(t, testData.expectedLastPartSize, parts[len(parts)-1].size)

		file.Close()
		os.Remove(path.Base(filePath))
		close(done)
	}
}
