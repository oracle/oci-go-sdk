// Package common Copyright (c) 2016, 2017, 2018 Oracle and/or its affiliates. All rights reserved.
package common

import (
	"fmt"
	"strings"
)

//Region type for regions
type Region string

const (
	//RegionSEA region SEA
	RegionSEA Region = "sea"
	//RegionPHX region PHX
	RegionPHX Region = "us-phoenix-1"
	//RegionIAD region IAD
	RegionIAD Region = "us-ashburn-1"
	//RegionFRA region FRA
	RegionFRA Region = "eu-frankfurt-1"
)

//StringToRegion convert a string to Region type
func StringToRegion(stringRegion string) (r Region, err error) {
	switch strings.ToLower(stringRegion) {
	case "sea":
		r = RegionSEA
	case "phx", "us-phoenix-1":
		r = RegionPHX
	case "iad", "us-ashburn-1":
		r = RegionIAD
	case "fra", "eu-frankfurt-1":
		r = RegionFRA
	default:
		err = fmt.Errorf("region named: %s, not valid", stringRegion)
	}
	return
}
