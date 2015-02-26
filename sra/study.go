package sra

import (
	"encoding/json"
	"encoding/xml"
	srastudy "github.com/indraniel/go-sra-schemas-1.5/SRA.study.xsd_go"
)

type SraStudyXML struct {
	FileName string
	XML      SraStudy
}

type SraStudy struct {
	XMLName xml.Name `xml:"STUDY_SET"`
	srastudy.TStudySetType
}

func (ss SraStudy) String() string {
	json, err := json.MarshalIndent(ss, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func (ss SraStudy) XMLString() string {
	xml, err := xml.MarshalIndent(ss, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (ss SraStudy) GetAccessions() []string {
	studies := ss.TStudySetType.Studies

	accessions := make([]string, 0)
	for _, v := range studies {
		accessions = append(accessions, v.Accession.String())
	}

	return accessions
}
