package initdump

import (
	"github.com/indraniel/srasearch/ncbiparser"
	"github.com/indraniel/srasearch/sra"
	"github.com/indraniel/srasearch/utils"

	"fmt"
	"log"
	"os"
)

func Main(metadata, uploads, output string) {
	log.Println("Collecting Accession Stats")
	accessionDB, accession_order := ncbiparser.CollectAccessionStats(metadata)

	log.Println("Collecting Uploads Stats")
	uploadsDB := ncbiparser.CollectUploadStats(uploads)

	log.Println("Processing XMLs in metadata/tar File")
	tarDB := ncbiparser.ProcessTarXMLs(metadata, accessionDB, uploadsDB)

	tmpdir, tmpfile := utils.MakeTmpFile()
	defer os.Remove(tmpfile)
	defer os.Remove(tmpdir)
	log.Println("Tmp Dump File is:", tmpfile)

	makeDumpFile(accession_order, accessionDB, tarDB, uploadsDB, tmpfile)

	log.Println("Compressing Dump File")
	err := utils.CompressFile(tmpfile, output)
	if err != nil {
		log.Print("Trouble making gzip file:", err)
		return
	}
	log.Println("All Done!")
}

func makeDumpFile(
	accessions []string,
	accessionDB *map[string]*sra.AccessionRecord,
	tarDB *map[string]*sra.SraItem,
	uploadsDB *map[string][]string,
	outFile string,
) {
	outPtr, err := os.Create(outFile)
	if err != nil {
		log.Fatal("Trouble opening %s for writing : %s\n", outFile, err)
	}
	defer outPtr.Close()

	for i, accession := range accessions {
		// the majority of stuff should be in the tar file XMLs
		if sraItem, ok := (*tarDB)[accession]; ok {
			sraItem.Record(outPtr)
			continue
		}

		// problematic cases -- specially handle and note
		if accessionRecord, ok := (*accessionDB)[accession]; ok {
			fmt.Printf(
				"--> [%d] Got a NCBI record with no XML : %s (%s)\n",
				i, accession, accessionRecord.Status,
			)
			sraItem := new(sra.SraItem)
			sraItem.Id = accession
			sraItem.AddAttrFromAccessionRecords(accessionDB)
			sraItem.AddAttrFromUploadRecords(uploadsDB)
			sraItem.Record(outPtr)
			continue
		}

		// this shouldn't be happening...but you never know...
		log.Fatalln("[err] Don't know how to dump accession: ", accession, "!")
	}
}
