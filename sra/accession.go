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
	Experiment string
	Sample     string
	Study      string
	MD5        string
	BioSample  string
	BioProject string
	Alias      string
	Issues     string
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
	experiment := fmt.Sprintf("Experiment : %s\n", ar.Experiment)
	sample := fmt.Sprintf("Sample     : %s\n", ar.Sample)
	study := fmt.Sprintf("Study      : %s\n", ar.Study)
	md5 := fmt.Sprintf("MD5        : %s\n", ar.MD5)
	biosample := fmt.Sprintf("BioSample  : %s\n", ar.BioSample)
	bioproject := fmt.Sprintf("BioProject : %s\n", ar.BioProject)
	alias := fmt.Sprintf("Alias      : %s\n", ar.Alias)
	issues := fmt.Sprintf("Issues     : %s\n", ar.Issues)

	return alias + status + issues + updated + published + received +
		visibility + study + bioproject + sample + biosample +
		experiment + md5
}
