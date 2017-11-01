// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.
package common

import (
	"fmt"
	"strings"
)

type Region int

const (
	region_invalid Region = 1 << iota
	REGION_SEA
	REGION_PHX
	REGION_IAD
	REGION_FRA
)

func RegionToString(r Region) (s string, err error) {
	switch r {
	case REGION_SEA:
		s = "sea"
	case REGION_PHX:
		s = "us-phoenix-1"
	case REGION_IAD:
		s = "us-ashburn-1"
	case REGION_FRA:
		s = "eu-frankfurt-1"
	default:
		err = fmt.Errorf("Region with value: %d, was not found", r)
	}
	return
}

func StringToRegion(stringRegion string) (r Region, err error) {
	switch strings.ToLower(stringRegion) {
	case "sea":
		r = REGION_SEA
	case "us-phoenix":
		r = REGION_PHX
	case "us-ashburn-1":
		r = REGION_IAD
	case "eu-frankfur-1":
		r = REGION_FRA
	default:
		err = fmt.Errorf("Region named: %s, not valid", stringRegion)
	}
	return
}
