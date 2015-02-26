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

func (se SraExp) String() string {
	json, err := json.MarshalIndent(se, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func (se SraExp) XMLString() string {
	xml, err := xml.MarshalIndent(se, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (se SraExp) GetAccessions() []string {
	exps := se.TExperimentSetType.Experiments

	accessions := make([]string, 0)
	for _, v := range exps {
		accessions = append(accessions, v.Accession.String())
	}

	return accessions
}
