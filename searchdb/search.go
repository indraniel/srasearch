package searchdb

import (
	"encoding/json"
	"fmt"

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
