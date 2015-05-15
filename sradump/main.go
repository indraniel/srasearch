package sradump

import (
	"io/ioutil"
	"log"
	"os"
)

func RunSraDump(metadata, uploads, output string) {
	log.Println("Collecting Accession Stats")
	dbAccessions := CollectAccessionStats(metadata)

	log.Println("Collecting Uploads Stats")
	dbUploads := CollectUploadStats(uploads)

	log.Println("Processing XMLs / Creating Dump File")

	tmpdir, tmpfile := makeTmpFile()
	defer os.Remove(tmpfile)
	defer os.Remove(tmpdir)
	log.Println("Tmp Dump File is:", tmpfile)
	ProcessTarXMLs(metadata, dbAccessions, dbUploads, tmpfile)

	log.Println("Compressing Dump File")
	err := CompressDumpFile(tmpfile, output)
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
