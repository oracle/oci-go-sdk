// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.
package common

type Region int

const (
	SEA Region = 1 << iota
	PHX
	IAD
	FRA
)

var DefaultRegion = IAD

func RegionToString(r Region) string {
	switch r {
	case SEA:
		return "sea"
	case PHX:
		return "phx"
	case IAD:
		return "iad"
	case FRA:
		return "fra"
	default:
		return ""
	}
}

