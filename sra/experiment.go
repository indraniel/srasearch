package sra

import (
	"encoding/json"
	"encoding/xml"
	sraexp "github.com/indraniel/go-sra-schemas-1.5/SRA.experiment.xsd_go"
)

type SraExperimentXML struct {
	FileName string
	XML      SraExp
}

type SraExp struct {
	XMLName xml.Name `xml:"EXPERIMENT_SET"`
	sraexp.TExperimentSetType
}

func (sr SraExp) String() string {
	json, err := json.MarshalIndent(sr, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func (sr SraExp) XMLString() string {
	xml, err := xml.MarshalIndent(sr, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}
