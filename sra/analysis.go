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

func (sr SraAnalysis) String() string {
	json, err := json.MarshalIndent(sr, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func (sr SraAnalysis) XMLString() string {
	xml, err := xml.MarshalIndent(sr, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (sr SraAnalysis) GetAccessions() []string {
	return []string{}
}
