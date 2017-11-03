// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// CrossConnectStatus. The status of the cross-connect.
type CrossConnectStatus struct {

	// The OCID of the cross-connect.
	CrossConnectID *string `mandatory:"true" json:"crossConnectId,omitempty"`

	// Whether Oracle's side of the interface is up or down.
	InterfaceState CrossConnectStatusInterfaceStateEnum `mandatory:"false" json:"interfaceState,omitempty"`

	// The light level of the cross-connect (in dBm).
	// Example: `14.0`
	LightLevelIndBm *float32 `mandatory:"false" json:"lightLevelIndBm,omitempty"`

	// Status indicator corresponding to the light level.
	//   * **NO_LIGHT:** No measurable light
	//   * **LOW_WARN:** There's measurable light but it's too low
	//   * **HIGH_WARN:** Light level is too high
	//   * **BAD:** There's measurable light but the signal-to-noise ratio is bad
	//   * **GOOD:** Good light level
	LightLevelIndicator CrossConnectStatusLightLevelIndicatorEnum `mandatory:"false" json:"lightLevelIndicator,omitempty"`
}

func (model CrossConnectStatus) String() string {
	return common.PointerString(model)
}

type CrossConnectStatusInterfaceStateEnum string
type CrossConnectStatusInterfaceState struct{}

const (
	CROSS_CONNECT_STATUS_INTERFACE_STATE_UP      CrossConnectStatusInterfaceStateEnum = "UP"
	CROSS_CONNECT_STATUS_INTERFACE_STATE_DOWN    CrossConnectStatusInterfaceStateEnum = "DOWN"
	CROSS_CONNECT_STATUS_INTERFACE_STATE_UNKNOWN CrossConnectStatusInterfaceStateEnum = "UNKNOWN"
)

var mapping_crossconnectstatus_interfaceState = map[string]CrossConnectStatusInterfaceStateEnum{
	"UP":      CROSS_CONNECT_STATUS_INTERFACE_STATE_UP,
	"DOWN":    CROSS_CONNECT_STATUS_INTERFACE_STATE_DOWN,
	"UNKNOWN": CROSS_CONNECT_STATUS_INTERFACE_STATE_UNKNOWN,
}

func (receiver CrossConnectStatusInterfaceState) Values() []CrossConnectStatusInterfaceStateEnum {
	values := make([]CrossConnectStatusInterfaceStateEnum, 0)
	for _, v := range mapping_crossconnectstatus_interfaceState {
		if v != CROSS_CONNECT_STATUS_INTERFACE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

func (receiver CrossConnectStatusInterfaceState) IsValid(toBeChecked string) bool {
	for _, v := range receiver.Values() {
		if CrossConnectStatusInterfaceStateEnum(toBeChecked) == v {
			return true
		}
	}
	return false
}

func (receiver CrossConnectStatusInterfaceState) From(toBeConverted string) CrossConnectStatusInterfaceStateEnum {
	if val, ok := mapping_crossconnectstatus_interfaceState[toBeConverted]; ok {
		return val
	}
	return CROSS_CONNECT_STATUS_INTERFACE_STATE_UNKNOWN
}

type CrossConnectStatusLightLevelIndicatorEnum string
type CrossConnectStatusLightLevelIndicator struct{}

const (
	CROSS_CONNECT_STATUS_LIGHT_LEVEL_INDICATOR_NO_LIGHT  CrossConnectStatusLightLevelIndicatorEnum = "NO_LIGHT"
	CROSS_CONNECT_STATUS_LIGHT_LEVEL_INDICATOR_LOW_WARN  CrossConnectStatusLightLevelIndicatorEnum = "LOW_WARN"
	CROSS_CONNECT_STATUS_LIGHT_LEVEL_INDICATOR_HIGH_WARN CrossConnectStatusLightLevelIndicatorEnum = "HIGH_WARN"
	CROSS_CONNECT_STATUS_LIGHT_LEVEL_INDICATOR_BAD       CrossConnectStatusLightLevelIndicatorEnum = "BAD"
	CROSS_CONNECT_STATUS_LIGHT_LEVEL_INDICATOR_GOOD      CrossConnectStatusLightLevelIndicatorEnum = "GOOD"
	CROSS_CONNECT_STATUS_LIGHT_LEVEL_INDICATOR_UNKNOWN   CrossConnectStatusLightLevelIndicatorEnum = "UNKNOWN"
)

var mapping_crossconnectstatus_lightLevelIndicator = map[string]CrossConnectStatusLightLevelIndicatorEnum{
	"NO_LIGHT":  CROSS_CONNECT_STATUS_LIGHT_LEVEL_INDICATOR_NO_LIGHT,
	"LOW_WARN":  CROSS_CONNECT_STATUS_LIGHT_LEVEL_INDICATOR_LOW_WARN,
	"HIGH_WARN": CROSS_CONNECT_STATUS_LIGHT_LEVEL_INDICATOR_HIGH_WARN,
	"BAD":       CROSS_CONNECT_STATUS_LIGHT_LEVEL_INDICATOR_BAD,
	"GOOD":      CROSS_CONNECT_STATUS_LIGHT_LEVEL_INDICATOR_GOOD,
	"UNKNOWN":   CROSS_CONNECT_STATUS_LIGHT_LEVEL_INDICATOR_UNKNOWN,
}

func (receiver CrossConnectStatusLightLevelIndicator) Values() []CrossConnectStatusLightLevelIndicatorEnum {
	values := make([]CrossConnectStatusLightLevelIndicatorEnum, 0)
	for _, v := range mapping_crossconnectstatus_lightLevelIndicator {
		if v != CROSS_CONNECT_STATUS_LIGHT_LEVEL_INDICATOR_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

func (receiver CrossConnectStatusLightLevelIndicator) IsValid(toBeChecked string) bool {
	for _, v := range receiver.Values() {
		if CrossConnectStatusLightLevelIndicatorEnum(toBeChecked) == v {
			return true
		}
	}
	return false
}

func (receiver CrossConnectStatusLightLevelIndicator) From(toBeConverted string) CrossConnectStatusLightLevelIndicatorEnum {
	if val, ok := mapping_crossconnectstatus_lightLevelIndicator[toBeConverted]; ok {
		return val
	}
	return CROSS_CONNECT_STATUS_LIGHT_LEVEL_INDICATOR_UNKNOWN
}
