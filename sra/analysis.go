package sra

import (
	"encoding/json"
	"encoding/xml"
	sraanalysis "github.com/indraniel/go-sra-schemas-1.5/SRA.analysis.xsd_go"
)

type SraAnalysisXML struct {
	FileName string
	XML      SraAnalysis
}

type SraAnalysis struct {
	XMLName xml.Name `xml:"ANALYSIS_SET"`
	sraanalysis.TAnalysisSetType
}

func (sa SraAnalysis) String() string {
	json, err := json.MarshalIndent(sa, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func (sa SraAnalysis) XMLString() string {
	xml, err := xml.MarshalIndent(sa, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (sa SraAnalysis) GetAccessions() []string {
	analyses := sa.TAnalysisSetType.Analysises

	accessions := make([]string, 0)
	for _, v := range analyses {
		accessions = append(accessions, v.Accession.String())
	}

	return accessions
}
