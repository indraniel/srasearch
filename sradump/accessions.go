package sradump

import (
	"archive/tar"
	"github.com/indraniel/srasearch/sra"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

func CollectAccessionStats(tarfile string) *map[string]*sra.AccessionRecord {
	data := getAccessionFileContents(tarfile)

	reader := csv.NewReader(data)
	reader.FieldsPerRecord = 21
	reader.Comma = '\t'

	db := make(map[string]*sra.AccessionRecord)
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
		updatedTime := parseTime(record[3])
		publishedTime := parseTime(record[4])
		receivedTime := parseTime(record[5])
		r := &sra.AccessionRecord{
			Status:     record[2],
			Updated:    updatedTime,
			Published:  publishedTime,
			Received:   receivedTime,
			Visibility: record[8],
		}
		db[accession] = r
	}

	log.Println("Processed", i, "accession records")
	return &db
}

func parseTime(ts string) time.Time {
	if ts == "-" {
		return time.Time{}
	}

	t, err := time.Parse("2006-01-02T15:04:05Z", ts)
	if err != nil {
		log.Fatalf("Trouble parsing timestamp: '%s' : %s\n", ts, err)
	}
	return t
}

func getAccessionFileContents(tarfile string) *bytes.Buffer {
	f, gzf := openGZFile(tarfile)
	defer closeGZFile(f, gzf)

	tarReader := tar.NewReader(gzf)

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