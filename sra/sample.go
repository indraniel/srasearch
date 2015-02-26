package sra

import (
	"encoding/json"
	"encoding/xml"
	srasample "github.com/indraniel/go-sra-schemas-1.5/SRA.sample.xsd_go"
)

type SraSample struct {
	XMLName xml.Name `xml:"SAMPLE_SET"`
	srasample.TSampleSetType
}

func (ss SraSample) String() string {
	json, err := json.MarshalIndent(ss, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func (ss SraSample) XMLString() string {
	xml, err := xml.MarshalIndent(ss, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (ss SraSample) GetAccessions() []string {
	samples := ss.TSampleSetType.Samples

	accessions := make([]string, 0)
	for _, v := range samples {
		accessions = append(accessions, v.Accession.String())
	}

	return accessions
}
