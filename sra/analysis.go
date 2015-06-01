package sra

import (
	"encoding/json"
	"encoding/xml"
	sraanalysis "github.com/indraniel/go-sra-schemas-1.5/SRA.analysis.xsd_go"
)

type SraAnalysisSet struct {
	XMLName xml.Name `xml:"ANALYSIS_SET"`
	sraanalysis.TAnalysisSetType
}

type SraAnalysis struct {
	XMLName xml.Name `xml:"ANALYSIS"`
	sraanalysis.TAnalysisType
}

func (sa SraAnalysisSet) String() string {
	json, err := json.MarshalIndent(sa, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func (sa SraAnalysisSet) XMLString() string {
	xml, err := xml.MarshalIndent(sa, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (sa SraAnalysisSet) GetItems() []Itemer {
	tanalyses := sa.TAnalysisSetType.Analysises

	analyses := make([]Itemer, 0)
	for _, v := range tanalyses {
		t := SraAnalysis{TAnalysisType: *v}
		analyses = append(analyses, Itemer(t))
	}
	return analyses
}

func (sa SraAnalysisSet) GetAccessions() []string {
	analyses := sa.TAnalysisSetType.Analysises

	accessions := make([]string, 0)
	for _, v := range analyses {
		accessions = append(accessions, v.Accession.String())
	}

	return accessions
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

func (sa SraAnalysis) GetAccession() string {
	return sa.Accession.String()
}

/* MGI specific */
func (sa SraAnalysis) IMPType() string {
	return "sra-analysis"
}
