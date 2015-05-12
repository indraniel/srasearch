package searchdb

import (
	"encoding/json"
	"fmt"

	"github.com/blevesearch/bleve"

	"github.com/indraniel/srasearch/sra"
)

func GetSRAItem(id string) (*sra.SraItem, error) {
	err := EnsureDB()
	if err != nil {
		return nil, err
	}

	raw, err := index.GetInternal([]byte(id))
	if err != nil {
		err := fmt.Errorf("Didn't find '%s' in index!", id)
		return nil, err
	}

	if raw == nil {
		err := fmt.Errorf("Didn't find '%s' in the search DB!", id)
		return nil, err
	}

	si := new(sra.SraItem)
	err = json.Unmarshal(raw, si)
	if err != nil {
		e := fmt.Errorf(
			"Error unmarshaling '%s' from search index! : %s",
			id, err,
		)
		return nil, e
	}
	return si, nil
}

func Query(qryString, start, end string, page, querySize int) (*bleve.SearchResult, error) {
	inputQuery := bleve.NewQueryStringQuery(qryString)
	tsQuery := bleve.NewDateRangeQuery(&start, &end)
	query := bleve.NewConjunctionQuery([]bleve.Query{inputQuery, tsQuery})

	from := (page - 1) * querySize

	search := bleve.NewSearchRequestOptions(query, querySize, from, false)
	search.Fields = []string{"*"}
	search.AddFacet("Types", bleve.NewFacetRequest("Type", 7))
	//	search.Highlight = bleve.NewHighlightWithStyle("html")
	//	search.Highlight.AddField("XML.Alias")
	//	search.Highlight.AddField("XML.Description")
	//	search.Highlight.AddField("XML.SubmissionId")
	//	search.Highlight.AddField("SubmissionId")
	//	search.Highlight.AddField("Published")
	//	search.Highlight.AddField("Type")

	searchResults, err := index.Search(search)
	if err == nil && (len(searchResults.Hits) == 0 && page != 1) {
		e := fmt.Errorf(
			"Page %d is is out of bounds on search request",
			page,
		)
		return nil, e
	}
	return searchResults, err
}
