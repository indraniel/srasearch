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

func (ss SraSubmission) String() string {
	json, err := json.MarshalIndent(ss, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func (ss SraSubmission) XMLString() string {
	xml, err := xml.MarshalIndent(ss, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (ss SraSubmission) GetAccessions() []string {
	submission := ss.TSubmissionType
	accessions := []string{submission.Accession.String()}
	return accessions
}
