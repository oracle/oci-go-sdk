package common

import (
	"bytes"
	"fmt"
	"sync"
)

const (
	major = 0
	minor = 0
	patch = 0
	tag   = ""
)

var once sync.Once
var version string

// Returns semantic version
func Version() string {
	once.Do(func() {
		ver := fmt.Sprintf("%d.%d.%d", major, minor, patch)
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
