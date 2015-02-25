package sra

import (
	"encoding/json"
	"encoding/xml"
	srarun "github.com/indraniel/go-sra-schemas-1.5/SRA.run.xsd_go"
)

type SraRunXML struct {
	FileName string
	XML      SraRun
}

type SraRun struct {
	XMLName xml.Name `xml:"RUN_SET"`
	srarun.TxsdRunSet
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

func (sr SraRun) GetAccessions() []string {
	return []string{}
}
