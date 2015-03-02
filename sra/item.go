package sra

// What do the different SRA accessions represent?
// http://www.ncbi.nlm.nih.gov/books/NBK56913/#search.what_do_the_different_sra_accessi

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
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

func (si *SraItem) UnmarshalJSON(data []byte) error {
	var aux struct {
		Id           string
		SubmissionId string
		XMLFileName  string
		Type         string
		Status       string
		Visibility   string
		Updated      string
		Published    string
		Received     string
		Data         json.RawMessage
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&aux); err != nil {
		return fmt.Errorf("decode sra item: %v", err)
	}

	si.Id = aux.Id
	si.SubmissionId = aux.SubmissionId
	si.XMLFileName = aux.XMLFileName
	si.Type = aux.Type
	si.Status = aux.Status
	si.Visibility = aux.Visibility

	t, err := time.Parse("2006-01-02T15:04:05Z", aux.Updated)
	if err != nil {
		return fmt.Errorf(
			"%s: '%s' : %s\n",
			"JSON Unmarshal error: Trouble parsing timestamp",
			aux.Updated,
			err,
		)
	}
	si.Updated = t

	t, err = time.Parse("2006-01-02T15:04:05Z", aux.Published)
	if err != nil {
		return fmt.Errorf(
			"%s: '%s' : %s\n",
			"JSON Unmarshal error: Trouble parsing timestamp",
			aux.Published,
			err,
		)
	}
	si.Published = t

	t, err = time.Parse("2006-01-02T15:04:05Z", aux.Received)
	if err != nil {
		return fmt.Errorf(
			"%s: '%s' : %s\n",
			"JSON Unmarshal error: Trouble parsing timestamp",
			aux.Received,
			err,
		)
	}
	si.Received = t

	var item Itemer
	switch si.Type {
	case "analysis":
		var analysis SraAnalysis
		err = json.Unmarshal([]byte(aux.Data), &analysis)
		if err != nil {
			return fmt.Errorf(
				"[SraAnalysis] JSON Unmarshal error: %v", err,
			)
		}
		item = analysis
	case "experiment":
		var exp SraExp
		err = json.Unmarshal([]byte(aux.Data), &exp)
		if err != nil {
			return fmt.Errorf(
				"[SraExp] JSON Unmarshal error: %v", err,
			)
		}
		item = exp
	case "run":
		var run SraRun
		err = json.Unmarshal([]byte(aux.Data), &run)
		if err != nil {
			return fmt.Errorf(
				"[SraRun] JSON Unmarshal error: %v", err,
			)
		}
		item = run
	case "sample":
		var sample SraSample
		err = json.Unmarshal([]byte(aux.Data), &sample)
		if err != nil {
			return fmt.Errorf(
				"[SraSample] JSON Unmarshal error: %v", err,
			)
		}
		item = sample
	case "study":
		var study SraStudy
		err = json.Unmarshal([]byte(aux.Data), &study)
		if err != nil {
			return fmt.Errorf(
				"[SraStudy] JSON Unmarshal error: %v", err,
			)
		}
		item = study
	case "submission":
		var submission SraSubmission
		err = json.Unmarshal([]byte(aux.Data), &submission)
		if err != nil {
			return fmt.Errorf(
				"[SraSubmission] JSON Unmarshal error: %v", err,
			)
		}
		item = submission
	default:
		return fmt.Errorf(
			"Don't know how to parse sra item of type '%s'\n",
			si.Type,
		)
	}

	si.Data = item
	return nil
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
