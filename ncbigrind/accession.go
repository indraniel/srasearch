package ncbigrind

import (
	"github.com/indraniel/srasearch/utils"

	"archive/tar"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

type AccessionRecord struct {
	Status     string
	Updated    time.Time
	Published  time.Time
	Received   time.Time
	Visibility string
	Experiment string
	Sample     string
	Study      string
	MD5        string
	BioSample  string
	BioProject string
	Alias      string
	Issues     string
}

func (ar AccessionRecord) String() string {
	status := fmt.Sprintf("Status     : %s\n", ar.Status)

	updated := fmt.Sprintf("Updated    : %s\n",
		ar.Updated.Format("2006-01-02T15:04:05Z"))

	published := fmt.Sprintf("Published  : %s\n",
		ar.Published.Format("2006-01-02T15:04:05Z"))

	received := fmt.Sprintf("Received   : %s\n",
		ar.Received.Format("2006-01-02T15:04:05Z"))

	visibility := fmt.Sprintf("Visibility : %s\n", ar.Visibility)
	experiment := fmt.Sprintf("Experiment : %s\n", ar.Experiment)
	sample := fmt.Sprintf("Sample     : %s\n", ar.Sample)
	study := fmt.Sprintf("Study      : %s\n", ar.Study)
	md5 := fmt.Sprintf("MD5        : %s\n", ar.MD5)
	biosample := fmt.Sprintf("BioSample  : %s\n", ar.BioSample)
	bioproject := fmt.Sprintf("BioProject : %s\n", ar.BioProject)
	alias := fmt.Sprintf("Alias      : %s\n", ar.Alias)
	issues := fmt.Sprintf("Issues     : %s\n", ar.Issues)

	return alias + status + issues + updated + published + received +
		visibility + study + bioproject + sample + biosample +
		experiment + md5
}

func CollectAccessionStats(tarfile string) (
	*map[string]*AccessionRecord,
	[]string) {

	data := getAccessionFileContents(tarfile)

	reader := csv.NewReader(data)
	reader.FieldsPerRecord = 21
	reader.Comma = '\t'

	db := make(map[string]*AccessionRecord)
	accessions := make([]string, 0)

	skip_header := true

	i := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		i++

		if skip_header {
			skip_header = false
			continue
		}

		accession := record[0]
		updatedTime := utils.ParseTime(record[3])
		publishedTime := utils.ParseTime(record[4])
		receivedTime := utils.ParseTime(record[5])
		r := &AccessionRecord{
			Status:     record[2],
			Updated:    updatedTime,
			Published:  publishedTime,
			Received:   receivedTime,
			Visibility: record[8],
			Alias:      record[9],
			Experiment: record[10],
			Sample:     record[11],
			Study:      record[12],
			MD5:        record[16],
			BioSample:  record[17],
			BioProject: record[18],
			Issues:     record[20],
		}
		db[accession] = r
		accessions = append(accessions, accession)
	}

	log.Println("Processed", i, "accession records")
	return &db, accessions
}

func getAccessionFileContents(tarfile string) *bytes.Buffer {
	gzreader := utils.OpenGZFile(tarfile)
	defer gzreader.Close()

	tarReader := tar.NewReader(gzreader.Gzf)

	buf := new(bytes.Buffer)

	i := 0
	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		name := header.Name

		switch header.Typeflag {
		case tar.TypeDir:
			continue
		case tar.TypeReg:
			//fmt.Println("(", i, ")", "Name:", name)
			res := strings.Index(name, "/SRA_Accessions_")
			if res != -1 {
				io.Copy(buf, tarReader)
				break
			}
		default:
			fmt.Printf("%s: %c %s %s\n",
				"Yikes! Unable to figure out type",
				header.Typeflag,
				"in file",
				name,
			)
		}

		i++
	}

	return buf
}
