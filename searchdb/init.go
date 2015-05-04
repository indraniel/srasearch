package searchdb

import (
	"errors"

	"github.com/blevesearch/bleve"
)

var IndexPath string
var index bleve.Index

func Init(idxPath string) error {
	IndexPath = idxPath
	idx, err := bleve.Open(IndexPath)
	if err != nil {
		return err
	}
	index = idx
	return nil
}

func EnsureDB() error {
	var err error
	if IndexPath == "" || index == nil {
		err = errors.New("The Search DB hasn't been initialized!")
	}
	return err
}
