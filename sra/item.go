package sra

// What do the different SRA accessions represent?
// http://www.ncbi.nlm.nih.gov/books/NBK56913/#search.what_do_the_different_sra_accessi

import (
	"encoding/xml"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type SRASetter interface {
	GetItems() []Itemer
	GetAccessions() []string
}

type Itemer interface {
	String() string
	XMLString() string
	GetAccession() string
}

type SraItem struct {
	Id           string
	SubmissionId string
	XMLFileName  string
	Type         string
	Status       string
	Updated      time.Time
	Published    time.Time
	Received     time.Time
	Visibility   string
	Data         Itemer
}

func (si *SraItem) setId() {
	accession := si.Data.GetAccession()
	si.Id = accession
}

func (si *SraItem) AddAttrFromAccessionRecords(
	db *map[string]*AccessionRecord,
) {
	if data, ok := (*db)[si.Id]; ok {
		si.Status = data.Status
		si.Updated = data.Updated
		si.Published = data.Published
		si.Received = data.Received
		si.Visibility = data.Visibility
	}
}

func NewSraItemsFromXML(filename string, contents []byte) []*SraItem {
	basename := filepath.Base(filename)
	id, sraType, _ := parseXMLFileName(basename)
	set := parseXMLContents(sraType, contents)

	sraItems := make([]*SraItem, 0)

	for _, item := range set.GetItems() {
		si := &SraItem{
			SubmissionId: id,
			XMLFileName:  basename,
			Type:         sraType,
			Data:         item,
		}
		si.setId()
		sraItems = append(sraItems, si)
	}
	return sraItems
}

func parseXMLFileName(filename string) (string, string, string) {
	items := strings.Split(filename, ".")
	submissionAccession, sraType, extension := items[0], items[1], items[2]
	return submissionAccession, sraType, extension
}

func parseXMLContents(sraType string, contents []byte) SRASetter {
	var data SRASetter
	switch sraType {
	case "analysis":
		var analyses SraAnalysisSet
		xml.Unmarshal(contents, &analyses)
		data = analyses
	case "experiment":
		var exps SraExpSet
		xml.Unmarshal(contents, &exps)
		data = exps
	case "run":
		var runs SraRunSet
		xml.Unmarshal(contents, &runs)
		data = runs
	case "sample":
		var samples SraSampleSet
		xml.Unmarshal(contents, &samples)
		data = samples
	case "study":
		var studies SraStudySet
		xml.Unmarshal(contents, &studies)
		data = studies
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
