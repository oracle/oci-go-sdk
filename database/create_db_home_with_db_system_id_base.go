// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateDbHomeWithDbSystemIdBase is an interface representing the polymorphic json shape of this model
type CreateDbHomeWithDbSystemIdBase interface {

	// The OCID of the DB System.
	GetDbSystemID() *string

	// The user-provided name of the database home.
	GetDisplayName() *string
}

type createdbhomewithdbsystemidbase struct {
	JsonData    []byte
	DbSystemID  *string `mandatory:"true" json:"dbSystemId,omitempty"`
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`
	Source      string  `json:"source"`
}

//UnmarshalJSON unmarshals json
func (m *createdbhomewithdbsystemidbase) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedbhomewithdbsystemidbase createdbhomewithdbsystemidbase
	s := struct {
		Model Unmarshalercreatedbhomewithdbsystemidbase
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DbSystemID = s.Model.DbSystemID
	m.DisplayName = s.Model.DisplayName
	m.Source = s.Model.Source

	return err
}

//UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdbhomewithdbsystemidbase) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	var err error
	switch m.Source {
	case "DB_BACKUP":
		mm := CreateDbHomeWithDbSystemIdFromBackupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := CreateDbHomeWithDbSystemIdDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

//GetDbSystemID returns DbSystemID
func (m createdbhomewithdbsystemidbase) GetDbSystemID() *string {
	return m.DbSystemID
}

//GetDisplayName returns DisplayName
func (m createdbhomewithdbsystemidbase) GetDisplayName() *string {
	return m.DisplayName
}

func (m createdbhomewithdbsystemidbase) String() string {
	return common.PointerString(m)
}
