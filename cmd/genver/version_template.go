package main


const versionTemplate = `
// Copyright (c) 2016, 2017, 2018 Oracle and/or its affiliates. All rights reserved.
// Code generated by go generate; DO NOT EDIT
package common

import (
	"bytes"
	"fmt"
	"sync"
)

const (
	major = "{{.Major}}"
	minor = "{{.Minor}}"
	patch = "{{.Patch}}"
	tag   = "{{.Tag}}"
)

var once sync.Once
var version string

// Returns semantic version
func Version() string {
	once.Do(func() {
		ver := fmt.Sprintf("%s.%s.%s", major, minor, patch)
		verBuilder := bytes.NewBufferString(ver)
		if tag != "" && tag != "-" {
			_, err := verBuilder.WriteString(tag)
			if err == nil {
				verBuilder = bytes.NewBufferString(ver)
			}
		}
		version = verBuilder.String()
	})
	return version
}
`

