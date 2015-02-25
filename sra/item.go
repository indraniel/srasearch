package sra

import (
	"encoding/xml"
	"log"
	"path/filepath"
	"strings"
)

type Schemer interface {
	String() string
}

type SraItem struct {
	Id   string
	Type string
	Data Schemer
}

func NewSraItemFromXML(filename string, contents []byte) *SraItem {
	basename := filepath.Base(filename)
	id, sraType, _ := parseXMLFileName(basename)
	data := parseXMLContents(sraType, contents)
	si := &SraItem{
		Id:   id,
		Type: sraType,
		Data: data,
	}
	return si
}

func parseXMLFileName(filename string) (string, string, string) {
	items := strings.Split(filename, ".")
	id, sraType, extension := items[0], items[1], items[2]
	return id, sraType, extension
}

func parseXMLContents(sraType string, contents []byte) Schemer {
	var data Schemer
	switch sraType {
	case "analysis":
		var analysis SraAnalysis
		xml.Unmarshal(contents, &analysis)
		data = analysis
	case "experiment":
		var exp SraExp
		xml.Unmarshal(contents, &exp)
		data = exp
	case "run":
		var run SraRun
		xml.Unmarshal(contents, &run)
		data = run
	case "sample":
		var sample SraSample
		xml.Unmarshal(contents, &sample)
		data = sample
	case "study":
		var study SraStudy
		xml.Unmarshal(contents, &study)
		data = study
	case "submission":
		var submission SraSubmission
		xml.Unmarshal(contents, &submission)
		data = submission
	default:
		log.Fatalf(
			"Don't know how to parse '%s' XML contents!",
			sraType,
		)
	}
	return data
}
