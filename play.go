package main

import (
	"github.com/indraniel/srasearch/sra"
	"encoding/json"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/document"
	"log"
	"os"
	"time"
)

func main() {
	index, err := bleve.Open("srasearch0402.idx")
	if err != nil {
		log.Fatalln("[1]", err)
	}

	//query := bleve.NewQueryStringQuery("SRA003617") // works

	query := bleve.NewQueryStringQuery("SubmissionId:SRA003617") // works

	//	date := "2015-03-15"
	//	query := bleve.NewDateRangeQuery(&date, nil)

	//query := bleve.NewMatchQuery(`SubmissionId:"SRA003617"`)     // works
	//query := bleve.NewMatchQuery("SRA003617") // works
	//query := bleve.NewMatchQuery("WXS") // works
	//query := bleve.NewPhraseQuery([]string{"SRA003617"}, "SubmissionId") // not work
	//query := bleve.NewTermQuery("SRA003617") // not work
	//query := bleve.NewTermQuery("SubmissionId:SRA003617") // not work
	search := bleve.NewSearchRequestOptions(query, 100, 0, false)
	//search := bleve.NewSearchRequest(query)

	search.AddFacet("Types", bleve.NewFacetRequest("Type", 7))

	//search.Highlight = bleve.NewHighlightWithStyle("ansi")
	search.Highlight = bleve.NewHighlightWithStyle("html")
	search.Highlight.AddField("XML.Alias")
	search.Highlight.AddField("XML.Description")
	search.Highlight.AddField("XML.SubmissionId")
	search.Highlight.AddField("SubmissionId")
	search.Highlight.AddField("Published")
	search.Highlight.AddField("Type")
	//	search.Fields = []string{"SubmissionId", "Published", "XML.Alias", "XML.Description"}
	searchResults, err := index.Search(search)
	if err != nil {
		log.Fatalln("[2]", err)
	}
	//	jsonStr, _ := json.MarshalIndent(searchResults.Hits, "", "    ")
	jsonStr, _ := json.MarshalIndent(searchResults, "", "    ")
	fmt.Printf("%s\n", jsonStr)
	fmt.Printf("\n\n-----------\n\n")
	fmt.Println(searchResults)
	fmt.Printf("\n\n-----------\n\n")
	ids := make([]string, 0)
	for _, val := range searchResults.Hits {
		id := val.ID
		doc, _ := index.Document(id)

		rv := struct {
			ID     string                 `json:"id"`
			Fields map[string]interface{} `json:"fields"`
		}{
			ID:     id,
			Fields: map[string]interface{}{},
		}
		for _, field := range doc.Fields {
			var newval interface{}
			switch field := field.(type) {
			case *document.TextField:
				newval = string(field.Value())
			case *document.NumericField:
				n, err := field.Number()
				if err == nil {
					newval = n
				}
			case *document.DateTimeField:
				d, err := field.DateTime()
				if err == nil {
					newval = d.Format(time.RFC3339Nano)
				}
			}
			existing, existed := rv.Fields[field.Name()]
			if existed {
				switch existing := existing.(type) {
				case []interface{}:
					rv.Fields[field.Name()] = append(existing, newval)
				case interface{}:
					arr := make([]interface{}, 2)
					arr[0] = existing
					arr[1] = newval
					rv.Fields[field.Name()] = arr
				}
			} else {
				rv.Fields[field.Name()] = newval
			}
		}

		j2, _ := json.MarshalIndent(rv, "", "    ")
		fmt.Printf("%s\n", j2)
		fmt.Println("\n\n")
		raw, err := index.GetInternal([]byte(id))
		if err != nil {
			log.Fatal("Trouble getting internal doc:", err)
		}
		fmt.Printf("%s\n", raw)
		var si sra.SraItem
		err = json.Unmarshal(raw, &si)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		fmt.Println(si.XML)
		fmt.Println("\n\n")
		fmt.Println(si.XML.XMLString())
		os.Exit(0)
	}
	fmt.Println(ids)
}
