package test

import (
	"encoding/json"
	"github.com/blevesearch/bleve/mapping"
	"io/ioutil"
	"os"
)

func LoadIndexMapping(path string) (*mapping.IndexMappingImpl, error) {
	indexMapping := mapping.NewIndexMapping()

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	resourceBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resourceBytes, indexMapping)
	if err != nil {
		return nil, err
	}

	return indexMapping, nil
}

func LoadKvconfig(path string) (map[string]interface{}, error) {
	kvconfig := make(map[string]interface{})

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	resourceBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resourceBytes, &kvconfig)
	if err != nil {
		return nil, err
	}

	return kvconfig, nil
}
