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
	"github.com/oracle/oci-go-sdk/v52/common"
)

// AnnouncementsPreferencesSummary The summary object for announcement email preferences.
type AnnouncementsPreferencesSummary struct {

	// The OCID of the compartment for which the email preferences apply. Because announcements are
	// specific to a tenancy, specify the tenancy by providing the root compartment OCID.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The ID of the preferences.
	Id *string `mandatory:"false" json:"id"`

	// A Boolean value to indicate whether the specified compartment chooses to not to receive informational announcements by email.
	// (Manage preferences for receiving announcements by email by specifying the `preferenceType` attribute instead.)
	IsUnsubscribed *bool `mandatory:"false" json:"isUnsubscribed"`

	// When the preferences were set initially.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// When the preferences were last updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The string representing the user's preference regarding receiving announcements by email.
	PreferenceType BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum `mandatory:"false" json:"preferenceType,omitempty"`
}

//GetCompartmentId returns CompartmentId
func (m AnnouncementsPreferencesSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetId returns Id
func (m AnnouncementsPreferencesSummary) GetId() *string {
	return m.Id
}

//GetIsUnsubscribed returns IsUnsubscribed
func (m AnnouncementsPreferencesSummary) GetIsUnsubscribed() *bool {
	return m.IsUnsubscribed
}

//GetTimeCreated returns TimeCreated
func (m AnnouncementsPreferencesSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m AnnouncementsPreferencesSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetPreferenceType returns PreferenceType
func (m AnnouncementsPreferencesSummary) GetPreferenceType() BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum {
	return m.PreferenceType
}

func (m AnnouncementsPreferencesSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AnnouncementsPreferencesSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAnnouncementsPreferencesSummary AnnouncementsPreferencesSummary
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAnnouncementsPreferencesSummary
	}{
		"AnnouncementsPreferencesSummary",
		(MarshalTypeAnnouncementsPreferencesSummary)(m),
	}

	return json.Marshal(&s)
}
