package sra

import (
	"encoding/json"
	"encoding/xml"
	srarun "github.com/indraniel/go-sra-schemas-1.5/SRA.run.xsd_go"
)

type SraRunSet struct {
	XMLName xml.Name `xml:"RUN_SET"`
	srarun.TxsdRunSet
}

type SraRun struct {
	XMLName xml.Name `xml:"RUN"`
	srarun.TRunType
}

func (sr SraRunSet) String() string {
	json, err := json.MarshalIndent(sr, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func (sr SraRunSet) XMLString() string {
	xml, err := xml.MarshalIndent(sr, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (sr SraRunSet) GetItems() []Itemer {
	truns := sr.TxsdRunSet.Runs

	runs := make([]Itemer, 0)
	for _, v := range truns {
		t := SraRun{TRunType: *v}
		runs = append(runs, Itemer(t))
	}
	return runs
}

func (sr SraRunSet) GetAccessions() []string {
	runs := sr.TxsdRunSet.Runs

	accessions := make([]string, 0)
	for _, v := range runs {
		accessions = append(accessions, v.Accession.String())
	}

	return accessions
}

func (sr SraRun) String() string {
	json, err := json.MarshalIndent(sr, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func (sr SraRun) XMLString() string {
	xml, err := xml.MarshalIndent(sr, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (sr SraRun) GetAccession() string {
	return sr.Accession.String()
}

/* MGI Specific */
func (sr SraRun) IMPType() string {
	return "sra-run"
}
