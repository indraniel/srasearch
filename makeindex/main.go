package makeindex

import (
	//"github.com/indraniel/srasearch/sra"
	"github.com/indraniel/srasearch/utils"

	"github.com/blevesearch/bleve"

	"bufio"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

func CreateSearchIndex(input, output string) {
	createBaseDir(output)
	if _, err := os.Stat(output); err == nil {
		log.Fatalf("Index: '%s' %s %s",
			output,
			"already exists!",
			"Please move or delete it and try again.",
		)
	}

	log.Println("Starting to make a brand new index:", output)

	batchSize := 100

	gzreader := utils.OpenGZFile(input)
	defer gzreader.Close()

	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(output, mapping)
	if err != nil {
		log.Fatalln("Trouble making a bleve index! :", err)
	}

	reader := bufio.NewReader(gzreader.Gzf)
	line, err := reader.ReadString('\n')

	count := 0
	batch := index.NewBatch()
	startTime := time.Now()
	batchCount := 0
	for err == nil {
		elems := strings.SplitN(line, ",", 2)
		docId, json := elems[0], elems[1]
		json = strings.TrimRight(json, "\n")

		batch.Index(docId, []byte(json))
		batch.SetInternal([]byte(docId), []byte(json))
		batchCount++

		if batchCount >= batchSize {
			//log.Println("Submitting a batch to the index")
			err := index.Batch(batch)
			//log.Println("[DONE] Submitting a batch to the index")
			if err != nil {
				log.Fatal("Couldn't batch!")
				//				return err
			}
			batch = index.NewBatch()
			batchCount = 0
		}
		count++
		if count%1000 == 0 {
			indexDuration := time.Since(startTime)
			indexDurationSeconds := float64(indexDuration) / float64(time.Second)
			timePerDoc := float64(indexDuration) / float64(count)
			log.Printf("Indexed %d documents, in %.2fs (average %.2fms/doc)\n", count, indexDurationSeconds, timePerDoc/float64(time.Millisecond))
		}
		line, err = reader.ReadString('\n')
	}

	if err != io.EOF {
		log.Fatalln("[err] reading line ", count, "in", gzreader.File.Name(), ":", err)
	}

	// flush the last batch
	if batchCount > 0 {
		err = index.Batch(batch)
		if err != nil {
			log.Fatal(err)
		}
	}

	indexDuration := time.Since(startTime)
	indexDurationSeconds := float64(indexDuration) / float64(time.Second)
	timePerDoc := float64(indexDuration) / float64(count)
	log.Printf("Indexed %d documents, in %.2fs (average %.2fms/doc)\n", count, indexDurationSeconds, timePerDoc/float64(time.Millisecond))
	log.Println("All Done")
}

func createBaseDir(dirPath string) {
	base := path.Dir(dirPath)
	if _, err := os.Stat(base); os.IsNotExist(err) {
		log.Printf("'%s' : Doesn't exist - creating path", base)
		e := os.MkdirAll(base, 0776)
		if e != nil {
			log.Fatalln(
				"Trouble creating directory:",
				base, ":",
				e,
			)
		}
	}
}

//func buildIndexMapping() (*bleve.IndexMapping, error) {
//}

// v1
//		var si sra.SraItem
//		err := json.Unmarshal([]byte(line), &si)
//		if err != nil {
//			log.Fatal("Trouble unmarshal JSON to SraItem:", err)
//		}
//		//		log.Println("(", i, ")", "Indexing", si.Id)
//
//		err = index.Index(si.Id, si)
//		if err != nil {
//			log.Fatal("Trouble indexing sra item:", err)
//		}
//		i++
//		if i%25000 == 0 {
//			log.Println("Indexed", i, "entries")
//		}
