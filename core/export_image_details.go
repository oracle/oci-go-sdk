// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oci-go-sdk/common"
	"encoding/json"
)

// ExportImageDetails. The destination details for the image export.
// Set `destinationType` to `objectStorageTuple`
// and use ExportImageViaObjectStorageTupleDetails
// when specifying the namespace, bucket name, and object name.
// Set `destinationType` to `objectStorageUri` and
// use ExportImageViaObjectStorageUriDetails
// when specifying the Object Storage URL.
type ExportImageDetails interface {
}

type exportimagedetails struct {
	DestinationType string `json:"destinationType"`
}

func (m *exportimagedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	err := json.Unmarshal(data, m)
	if err != nil {
		return nil, err
	}

	switch m.DestinationType {
	case "objectStorageUri":
		mm := ExportImageViaObjectStorageUriDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "objectStorageTuple":
		mm := ExportImageViaObjectStorageTupleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

func (model exportimagedetails) String() string {
	return common.PointerString(model)
}
