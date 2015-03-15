package makeindex

import (
	//"github.com/indraniel/srasearch/sra"
	"github.com/indraniel/srasearch/utils"
	"bufio"
	//"encoding/json"
	"github.com/blevesearch/bleve"
	"io"
	"log"
	"strings"
	"time"
)

func CreateSearchIndex(input, output string) {
	batchSize := 100
	f, gzf := utils.OpenGZFile(input)
	defer utils.CloseGZFile(f, gzf)

	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(output, mapping)
	if err != nil {
		log.Fatal("Trouble making a bleve index!")
	}

	reader := bufio.NewReader(gzf)
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
		log.Fatalln("[err] reading line ", count, "in", f.Name(), ":", err)
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
