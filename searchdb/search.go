package searchdb

import (
	"encoding/json"
	"fmt"
	"time"

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
	// setup query string
	q := make([]bleve.Query, 0)
	if qryString == "*" {
		// Special Search Term:
		//    Just search for everything successfully processed by
		//    the SRA in a given time range
		tsQuery := bleve.NewDateRangeQuery(&start, &end)
		q = append(q, tsQuery)
	} else {
		inputQuery := bleve.NewQueryStringQuery(qryString)
		tsQuery := bleve.NewDateRangeQuery(&start, &end)
		q = append(q, inputQuery, tsQuery)
	}
	query := bleve.NewConjunctionQuery(q)

	from := (page - 1) * querySize

	// setup search options
	search := bleve.NewSearchRequestOptions(query, querySize, from, false)
	search.Fields = []string{"*"} // search and display on all the fields

	// facet on 'submission', 'study', 'sample', 'experiment', 'run', 'analysis'
	search.AddFacet("Types", bleve.NewFacetRequest("Type", 6))

	// run search query with a 1 minute timeout
	var timeout time.Duration = 1 /* 1 minute */
	var searchResults *bleve.SearchResult = nil
	var err error
	errch := make(chan error)

	go func() {
		results, e := index.Search(search)
		searchResults = results
		errch <- e
	}()

	select {
	case e := <-errch:
		err = e
	case <-time.After(time.Minute * timeout):
		err = fmt.Errorf("Search Timeout (%d mins): %s", timeout,
			"Please refine your query, or narrow your time range",
		)
	}

	if err == nil && (len(searchResults.Hits) == 0 && page != 1) {
		e := fmt.Errorf(
			"Page %d is is out of bounds on search request",
			page,
		)
		return nil, e
	}
	return searchResults, err
}
