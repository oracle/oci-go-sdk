// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Resource Details about the ticket item object.
type Resource struct {
	Item Item `mandatory:"false" json:"item"`

	// The list of available Oracle Cloud Infrastructure regions.
	Region RegionEnum `mandatory:"false" json:"region,omitempty"`

	// The list of available Oracle Cloud Infrastructure availability domains.
	AvailabilityDomain AvailabilityDomainEnum `mandatory:"false" json:"availabilityDomain,omitempty"`
}

func (m Resource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Resource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRegionEnum(string(m.Region)); !ok && m.Region != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Region: %s. Supported values are: %s.", m.Region, strings.Join(GetRegionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAvailabilityDomainEnum(string(m.AvailabilityDomain)); !ok && m.AvailabilityDomain != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailabilityDomain: %s. Supported values are: %s.", m.AvailabilityDomain, strings.Join(GetAvailabilityDomainEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Resource) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Item               item                   `json:"item"`
		Region             RegionEnum             `json:"region"`
		AvailabilityDomain AvailabilityDomainEnum `json:"availabilityDomain"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Item.UnmarshalPolymorphicJSON(model.Item.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Item = nn.(Item)
	} else {
		m.Item = nil
	}

	m.Region = model.Region

	m.AvailabilityDomain = model.AvailabilityDomain

	return
}
