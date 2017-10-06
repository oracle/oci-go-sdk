package common

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Simple logging proxy to distinguish control for logging messages
// Debug logging is turned on/off by the presence of the environment variable "OCI_GO_SDK_DEBUG"
var debugLog = log.New(os.Stderr, "DEBUG ", log.Ldate|log.Ltime|log.Lshortfile)
var mainLog = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
var isDebugLogEnabled *bool = nil

func getOutputForEnv() (writer io.Writer) {
	if isDebugLogEnabled == nil {
		isDebugLogEnabled = new(bool)
		_, *isDebugLogEnabled = os.LookupEnv("OCI_GO_SDK_DEBUG")
	}

	writer = ioutil.Discard
	if *isDebugLogEnabled {
		writer = os.Stderr
	}
	return
}

// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...interface{}) {
	debugLog.SetOutput(getOutputForEnv())
	debugLog.Output(3, fmt.Sprintf(format, v...))
}

// Arguments are handled in the manner of fmt.Print.
func Debug(v ...interface{}) {
	debugLog.SetOutput(getOutputForEnv())
	debugLog.Output(3, fmt.Sprint(v...))
}

// Arguments are handled in the manner of fmt.Println.
func Debugln(v ...interface{}) {
	debugLog.SetOutput(getOutputForEnv())
	debugLog.Output(3, fmt.Sprintln(v...))
}

func Logln(v ...interface{}) {
	mainLog.Output(3, fmt.Sprintln(v...))
}
