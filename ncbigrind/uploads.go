package ncbigrind

import (
	"github.com/indraniel/srasearch/utils"

	"encoding/csv"
	"io"
	"log"
)

func CollectUploadStats(uploads string) *map[string][]string {
	gzreader := utils.OpenGZFile(uploads)
	defer gzreader.Close()

	reader := csv.NewReader(gzreader.Gzf)
	reader.FieldsPerRecord = 20
	reader.LazyQuotes = true
	reader.Comma = '\t'

	db := make(map[string][]string)
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

		accession := record[13]
		fileName := record[3]
		uploadName := record[6]

		if accession == "-" {
			continue
		}

		names, exists := db[accession]
		if !exists {
			names = make([]string, 0)
		}

		if fileName != "=" || fileName != "-" {
			exist_flag := 0
			for _, v := range names {
				if v == fileName {
					exist_flag = 1
					break
				}
			}

			if exist_flag == 0 {
				names = append(names, fileName)
			}
		}

		if uploadName != "=" || uploadName != "-" {
			exist_flag := 0
			for _, v := range names {
				if v == uploadName {
					exist_flag = 1
					break
				}
			}

			if exist_flag == 0 {
				names = append(names, uploadName)
			}
		}

		db[accession] = names
	}

	log.Println("Processed", i, "upload records")
	return &db
}
