// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.
package common

import (
	"fmt"
)

type Region int

const (
	SEA Region = 1 << iota
	PHX
	IAD
	FRA
)

var DefaultRegion = PHX

func RegionToString(r Region) (s string, err error) {
	switch r {
	case SEA:
		s =  "sea"
	case PHX:
		s = "us-phoenix-1"
	case IAD:
		s =  "iad"
	case FRA:
		s =  "fra"
	default:
		err = fmt.Errorf("Region with value: %d, was not found", r)
	}
	return
}

