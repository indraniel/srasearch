package sra

import (
	"encoding/json"
	"encoding/xml"
	srasubmission "github.com/indraniel/go-sra-schemas-1.5/SRA.submission.xsd_go"
)

type SraSubmissionXML struct {
	FileName string
	XML      SraSubmission
}

type SraSubmission struct {
	XMLName xml.Name `xml:"RUN_SET"`
	srasubmission.TSubmissionType
}

func (sr SraSubmission) String() string {
	json, err := json.MarshalIndent(sr, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func (sr SraSubmission) XMLString() string {
	xml, err := xml.MarshalIndent(sr, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (sr SraSubmission) GetAccessions() []string {
	return []string{}
}
