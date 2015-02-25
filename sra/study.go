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

func (sr SraStudy) String() string {
	json, err := json.MarshalIndent(sr, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func (sr SraStudy) XMLString() string {
	xml, err := xml.MarshalIndent(sr, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}
