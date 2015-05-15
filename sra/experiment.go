package sra

import (
	"encoding/json"
	"encoding/xml"
	sraexp "github.com/indraniel/go-sra-schemas-1.5/SRA.experiment.xsd_go"
)

type SraExpSet struct {
	XMLName xml.Name `xml:"EXPERIMENT_SET"`
	sraexp.TExperimentSetType
}

type SraExp struct {
	XMLName xml.Name `xml:"EXPERIMENT"`
	sraexp.TExperimentType
}

func (se SraExpSet) String() string {
	json, err := json.MarshalIndent(se, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func (se SraExpSet) XMLString() string {
	xml, err := xml.MarshalIndent(se, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (se SraExpSet) GetItems() []Itemer {
	texperiments := se.TExperimentSetType.Experiments

	experiments := make([]Itemer, 0)
	for _, v := range texperiments {
		t := SraExp{TExperimentType: *v}
		experiments = append(experiments, SraExp(t))
	}
	return experiments
}

func (se SraExpSet) GetAccessions() []string {
	exps := se.TExperimentSetType.Experiments

	accessions := make([]string, 0)
	for _, v := range exps {
		accessions = append(accessions, v.Accession.String())
	}

	return accessions
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

func (se SraExp) GetAccession() string {
	return se.Accession.String()
}

/* MGI specific */
func (se SraExp) IMPType() string {
	return "sra-experiment"
}
