package merge

import (
	"archive/tar"
	"github.com/indraniel/srasearch/sra"
	"github.com/indraniel/srasearch/sradump"
	"github.com/indraniel/srasearch/utils"
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func RunMerge(metadata, uploads, dumpfile, output string) {
	log.Println("Collecting Accession Stats")
	accessionDB, accession_order := CollectAccessionStats(metadata)

	log.Println("Collecting Uploads Stats")
	uploadsDB := sradump.CollectUploadStats(uploads)

	log.Println("Building Data Structure From Prior Dump")
	dumpDB := CollectDumpStats(dumpfile)

	log.Println("Building Incremental Data Structure from tar file")
	incrementalDB := ProcessTarXMLs(metadata, accessionDB, uploadsDB)

	tmpdir, tmpfile := makeTmpFile()
	defer os.Remove(tmpfile)
	defer os.Remove(tmpdir)
	log.Println("Tmp Dump File is:", tmpfile)
	log.Println("Merging Data Structures into Tmp Dump File:", tmpfile)
	merge(accession_order, accessionDB, dumpDB, incrementalDB, uploadsDB, tmpfile)

	log.Println("Compressing Dump File")
	err := sradump.CompressDumpFile(tmpfile, output)
	if err != nil {
		log.Print("Trouble making gzip file:", err)
		return
	}
	log.Println("All Done!")
}

func makeTmpFile() (tmpdir, tmpfile string) {
	tmpdir, err := ioutil.TempDir(os.TempDir(), "sra-dump")
	if err != nil {
		log.Fatal("Trouble making temp dir:", err)
	}

	f, err := ioutil.TempFile(tmpdir, "sra-tmp-dump")
	if err != nil {
		log.Fatal("Trouble making temp file:", err)
	}
	defer f.Close()

	tmpfile = f.Name()
	return
}

func CollectAccessionStats(tarfile string) (
	*map[string]*sra.AccessionRecord,
	[]string) {

	data := getAccessionFileContents(tarfile)

	reader := csv.NewReader(data)
	reader.FieldsPerRecord = 21
	reader.Comma = '\t'

	db := make(map[string]*sra.AccessionRecord)
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
		r := &sra.AccessionRecord{
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

func CollectDumpStats(dumpfile string) *map[string]*sra.SraItem {
	f, gzf := utils.OpenGZFile(dumpfile)
	defer utils.CloseGZFile(f, gzf)

	db := make(map[string]*sra.SraItem)

	count := 1
	reader := bufio.NewReader(gzf)
	line, err := reader.ReadString('\n')

	for err == nil {
		elems := strings.SplitN(line, ",", 2)
		accession, jsonData := elems[0], elems[1]
		raw := strings.TrimRight(jsonData, "\n")
		si := new(sra.SraItem)
		if e := json.Unmarshal([]byte(raw), si); e != nil {
			log.Fatalf(
				"Trouble json parsing accession record: %s\n",
				raw,
			)
		}
		db[accession] = si
		line, err = reader.ReadString('\n')
		count++
	}

	if err != io.EOF {
		log.Fatalln(
			"[err] reading line ", count, "in", f.Name(), ":", err,
		)
	}

	return &db
}

func ProcessTarXMLs(
	tarfile string,
	accessionDB *map[string]*sra.AccessionRecord,
	uploadsDB *map[string][]string,
) *map[string]*sra.SraItem {
	f, gzf := utils.OpenGZFile(tarfile)
	defer utils.CloseGZFile(f, gzf)

	tarReader := tar.NewReader(gzf)

	db := make(map[string]*sra.SraItem)

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
			//			fmt.Println("(", i, ")", "Name:", name)
			if isXML(name) {
				buf := new(bytes.Buffer)
				io.Copy(buf, tarReader)
				sraItems := processXML(accessionDB, uploadsDB, name, buf)
				for _, si := range sraItems {
					db[si.Id] = si
				}
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
		//		if i == 100 {
		//			os.Exit(0)
		//		}
	}

	return &db
}

func isXML(filename string) bool {
	if path.Ext(filename) == ".xml" {
		return true
	}
	return false
}

func processXML(
	accessionDB *map[string]*sra.AccessionRecord,
	uploadsDB *map[string][]string,
	name string,
	contents *bytes.Buffer,
) []*sra.SraItem {
	sraItems := sra.NewSraItemsFromXML(name, contents.Bytes())
	for _, si := range sraItems {
		si.AddAttrFromAccessionRecords(accessionDB)
		si.AddAttrFromUploadRecords(uploadsDB)
	}

	return sraItems
}

func merge(
	accessions []string,
	accessionDB *map[string]*sra.AccessionRecord,
	dumpDB *map[string]*sra.SraItem,
	incrementalDB *map[string]*sra.SraItem,
	uploadsDB *map[string][]string,
	outFile string,
) {

	outPtr, err := os.Create(outFile)
	if err != nil {
		log.Fatal("Trouble opening %s for writing : %s\n", outFile, err)
	}
	defer outPtr.Close()

	for i, accession := range accessions {
		// the "hot" stuff should be in the incremental file
		if sraItem, ok := (*incrementalDB)[accession]; ok {
			recordSraItem(sraItem, outPtr)
			continue
		}

		// the "usual" stuff should be in the prior dump
		if sraItem, ok := (*dumpDB)[accession]; ok {
			recordSraItem(sraItem, outPtr)
			continue
		}

		// problematic cases -- specially handle and note
		if accessionRecord, ok := (*accessionDB)[accession]; ok {
			fmt.Printf(
				"--> [%d] Got a NCBI 'unprocessed' record: %s (%s)\n",
				i, accession, accessionRecord.Status,
			)
			sraItem := new(sra.SraItem)
			sraItem.Id = accession
			sraItem.AddAttrFromAccessionRecords(accessionDB)
			sraItem.AddAttrFromUploadRecords(uploadsDB)
			recordSraItem(sraItem, outPtr)
			continue
		}

		// this shouldn't be happening...but you never know...
		log.Fatalln("[err] Don't know how to merge accession: ", accession, "!")
	}
}

func recordSraItem(si *sra.SraItem, outPtr *os.File) {
	json, err := json.Marshal(si)
	if err != nil {
		log.Fatal("Trouble encoding '%s' into json: \n%+v\n",
			si, err)
	}

	line := strings.Join([]string{si.Id, string(json)}, ",")

	_, err = outPtr.WriteString(line)
	checkWrite(outPtr, err)
	_, err = outPtr.Write([]byte("\n"))
	checkWrite(outPtr, err)
}

func checkWrite(f *os.File, e error) {
	if e != nil {
		log.Fatal("Trouble writing to '%s' : %s\n", f.Name(), e)
	}
}

func getAccessionFileContents(tarfile string) *bytes.Buffer {
	f, gzf := utils.OpenGZFile(tarfile)
	defer utils.CloseGZFile(f, gzf)

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
