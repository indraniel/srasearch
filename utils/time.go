package utils

import (
	"log"
	"time"
)

func ParseTime(ts string) time.Time {
	if ts == "-" {
		return time.Time{}
	}

	t, err := time.Parse("2006-01-02T15:04:05Z", ts)
	if err != nil {
		log.Fatalf("Trouble parsing timestamp: '%s' : %s\n", ts, err)
	}
	return t
}
