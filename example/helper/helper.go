// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package helper

import "log"

// LogIfError is equivalent to Println() followed by a call to os.Exit(1) if error is not nil
func LogIfError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
