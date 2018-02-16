// Package common Copyright (c) 2016, 2017, 2018 Oracle and/or its affiliates. All rights reserved.
package common

import (
	"log"
	"os"
	"sync"
)

// Simple logging proxy to distinguish control for logging messages
// Debug logging is turned on/off by the presence of the environment variable "OCI_GO_SDK_DEBUG"
var isDebugLogEnabled bool
var checkDebug sync.Once

func getOutputForEnv() {
	checkDebug.Do(func() {
		isDebugLogEnabled = *new(bool)
		_, isDebugLogEnabled = os.LookupEnv("OCI_GO_SDK_DEBUG")
	})

	return
}

// Debugf logs v with the provided format if debug mode is set
func Debugf(format string, v ...interface{}) {
	getOutputForEnv()
	if isDebugLogEnabled {
		log.Printf(format, v...)
	}
}

// Debug  logs v if debug mode is set
func Debug(v ...interface{}) {
	getOutputForEnv()
	if isDebugLogEnabled {
		log.Print(v...)
	}
}

// Debugln logs v appending a new line if debug mode is set
func Debugln(v ...interface{}) {
	getOutputForEnv()
	if isDebugLogEnabled {
		log.Println(v...)
	}
}

// IfDebug executes closure if debug is enabled
func IfDebug(fn func()) {
	if isDebugLogEnabled {
		fn()
	}
}

// Logln logs v appending a new line at the end
func Logln(v ...interface{}) {
	log.Println(v...)
}

// Logf logs v with the provided format
func Logf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
