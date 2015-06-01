package sra

import (
	"encoding/json"
	"encoding/xml"
	srasample "github.com/indraniel/go-sra-schemas-1.5/SRA.sample.xsd_go"
)

type SraSampleSet struct {
	XMLName xml.Name `xml:"SAMPLE_SET"`
	srasample.TSampleSetType
}

type SraSample struct {
	XMLName xml.Name `xml:"SAMPLE"`
	srasample.TSampleType
}

func (ss SraSampleSet) String() string {
	json, err := json.MarshalIndent(ss, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func (ss SraSampleSet) XMLString() string {
	xml, err := xml.MarshalIndent(ss, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (ss SraSampleSet) GetItems() []Itemer {
	tsamples := ss.TSampleSetType.Samples

	samples := make([]Itemer, 0)
	for _, v := range tsamples {
		t := SraSample{TSampleType: *v}
		samples = append(samples, Itemer(t))
	}
	return samples
}

func (ss SraSampleSet) GetAccessions() []string {
	samples := ss.TSampleSetType.Samples

	accessions := make([]string, 0)
	for _, v := range samples {
		accessions = append(accessions, v.Accession.String())
	}

	return accessions
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

func (ss SraSample) GetAccession() string {
	return ss.Accession.String()
}

/* MGI specific */
func (ss SraSample) IMPType() string {
	return "sra-sample"
}
