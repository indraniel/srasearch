package sra

import (
	"encoding/json"
	"encoding/xml"
	srasubmission "github.com/indraniel/go-sra-schemas-1.5/SRA.submission.xsd_go"
)

type SraSubmission struct {
	XMLName xml.Name `xml:"SUBMISSION"`
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

func (ss SraSubmission) GetItems() []Itemer {
	return []Itemer{Itemer(ss)}
}

func (ss SraSubmission) GetAccessions() []string {
	accessions := []string{ss.Accession.String()}
	return accessions
}

func (ss SraSubmission) GetAccession() string {
	return ss.Accession.String()
}

func (ss SraSubmission) IMPType() string {
	return "sra-submission"
}
