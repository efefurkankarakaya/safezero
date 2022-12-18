package utils

import "log"

func ReflectError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// If there's no error, then print.
func PrintSuccess(message string, err error) {
	if err == nil {
		println(message)
	}
}
