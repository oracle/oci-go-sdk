package helper

import "log"

// LogIfError log the error
func LogIfError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
