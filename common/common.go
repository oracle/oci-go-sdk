// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.
package common

import (
	"fmt"
	"strings"
)

type Region string

const (
	REGION_SEA Region = "sea"
	REGION_PHX Region = "us-phoenix-1"
	REGION_IAD Region = "us-ashburn-1"
	REGION_FRA Region = "eu-frankfurt-1"
)

func StringToRegion(stringRegion string) (r Region, err error) {
	switch strings.ToLower(stringRegion) {
	case "sea":
		r = REGION_SEA
	case "phx", "us-phoenix-1":
		r = REGION_PHX
	case "iad", "us-ashburn-1":
		r = REGION_IAD
	case "fra", "eu-frankfurt-1":
		r = REGION_FRA
	default:
		err = fmt.Errorf("Region named: %s, not valid", stringRegion)
	}
	return
}
