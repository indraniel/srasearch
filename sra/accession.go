package sra

import (
	"fmt"
	"time"
)

type AccessionRecord struct {
	Status     string
	Updated    time.Time
	Published  time.Time
	Received   time.Time
	Visibility string
}

func (ar AccessionRecord) String() string {
	status := fmt.Sprintf("Status     : %s\n", ar.Status)

	updated := fmt.Sprintf("Updated    : %s\n",
		ar.Updated.Format("2006-01-02T15:04:05Z"))

	published := fmt.Sprintf("Published  : %s\n",
		ar.Published.Format("2006-01-02T15:04:05Z"))

	received := fmt.Sprintf("Received   : %s\n",
		ar.Received.Format("2006-01-02T15:04:05Z"))

	visibility := fmt.Sprintf("Visibility : %s\n", ar.Visibility)

	return status + updated + published + received + visibility
}
