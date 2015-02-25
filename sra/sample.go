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

func (sr SraSample) String() string {
	json, err := json.MarshalIndent(sr, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func (sr SraSample) XMLString() string {
	xml, err := xml.MarshalIndent(sr, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}
