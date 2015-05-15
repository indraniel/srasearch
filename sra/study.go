package sra

import (
	"encoding/json"
	"encoding/xml"
	srastudy "github.com/indraniel/go-sra-schemas-1.5/SRA.study.xsd_go"
)

type SraStudySet struct {
	XMLName xml.Name `xml:"STUDY_SET"`
	srastudy.TStudySetType
}

type SraStudy struct {
	XMLName xml.Name `xml:"STUDY"`
	srastudy.TStudyType
}

func (ss SraStudySet) String() string {
	json, err := json.MarshalIndent(ss, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func (ss SraStudySet) XMLString() string {
	xml, err := xml.MarshalIndent(ss, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (ss SraStudySet) GetItems() []Itemer {
	tstudies := ss.TStudySetType.Studies

	studies := make([]Itemer, 0)
	for _, v := range tstudies {
		t := SraStudy{TStudyType: *v}
		studies = append(studies, Itemer(t))
	}
	return studies
}

func (ss SraStudySet) GetAccessions() []string {
	studies := ss.TStudySetType.Studies

	accessions := make([]string, 0)
	for _, v := range studies {
		accessions = append(accessions, v.Accession.String())
	}

	return accessions
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

func (ss SraStudy) GetAccession() string {
	return ss.Accession.String()
}

/* MGI specific */
func (ss SraStudy) IMPType() string {
	return "sra-study"
}
