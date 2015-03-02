package main

import (
	"fmt"
	"github.com/blevesearch/bleve"
	"log"
)

func main() {
	index, err := bleve.Open("srasearch.idx")
	if err != nil {
		log.Fatalln("[1]", err)
	}

	query := bleve.NewQueryStringQuery("SubmissionId:SRA003617") // works
	//query := bleve.NewMatchQuery(`SubmissionId:"SRA003617"`)     // works
	//query := bleve.NewMatchQuery("SRA003617") // works
	//query := bleve.NewMatchQuery("WXS") // works
	//query := bleve.NewPhraseQuery([]string{"SRA003617"}, "SubmissionId") // not work
	//query := bleve.NewTermQuery("SRA003617") // not work
	//query := bleve.NewTermQuery("SubmissionId:SRA003617") // not work
	search := bleve.NewSearchRequestOptions(query, 100, 0, false)
	//search := bleve.NewSearchRequest(query)
	search.Highlight = bleve.NewHighlightWithStyle("ansi")
	search.Highlight.AddField("Data.Alias")
	search.Fields = []string{"SubmissionId", "Published", "Data.Alias", "Data.Description"}
	searchResults, err := index.Search(search)
	if err != nil {
		log.Fatalln("[2]", err)
	}
	fmt.Printf("%v\n", searchResults)
}
