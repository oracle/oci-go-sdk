// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// Manage Oracle Cloud Infrastructure console announcements.
//

package announcementsservice

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v53/common"
)

// CreateAnnouncementsPreferencesDetails The object used to create announcement email preferences.
type CreateAnnouncementsPreferencesDetails struct {

	// A Boolean value to indicate whether the specified compartment chooses to not to receive informational announcements by email.
	// (Manage preferences for receiving announcements by email by specifying the `preferenceType` attribute instead.)
	IsUnsubscribed *bool `mandatory:"false" json:"isUnsubscribed"`

	// The OCID of the compartment for which you want to manage announcement email preferences. (Specify the tenancy by providing the
	// root compartment OCID.)
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The string representing the user's preference, whether to opt in to only required announcements, to opt in to all announcements, including informational announcements, or to opt out of all announcements.
	PreferenceType BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum `mandatory:"true" json:"preferenceType"`
}

//GetIsUnsubscribed returns IsUnsubscribed
func (m CreateAnnouncementsPreferencesDetails) GetIsUnsubscribed() *bool {
	return m.IsUnsubscribed
}

//GetCompartmentId returns CompartmentId
func (m CreateAnnouncementsPreferencesDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetPreferenceType returns PreferenceType
func (m CreateAnnouncementsPreferencesDetails) GetPreferenceType() BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum {
	return m.PreferenceType
}

func (m CreateAnnouncementsPreferencesDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateAnnouncementsPreferencesDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAnnouncementsPreferencesDetails CreateAnnouncementsPreferencesDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateAnnouncementsPreferencesDetails
	}{
		"CreateAnnouncementsPreferencesDetails",
		(MarshalTypeCreateAnnouncementsPreferencesDetails)(m),
	}

	return json.Marshal(&s)
}
