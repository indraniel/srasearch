package utils

import (
	"log"
	"os"
)

func CheckFileExists(file string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		log.Fatalf(
			"Could not find '%s' on file system: %s",
			file, err,
		)
	}
}
