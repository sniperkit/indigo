package grpc

import (
	"encoding/json"
	"github.com/blevesearch/bleve/mapping"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"testing"
)

func loadIndexMapping(path string) *mapping.IndexMappingImpl {
	indexMapping := mapping.NewIndexMapping()

	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	resourceBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(resourceBytes, indexMapping)
	if err != nil {
		return nil
	}

	return indexMapping
}

func loadKvconfig(path string) map[string]interface{} {
	kvconfig := make(map[string]interface{})

	file, err := os.Open(viper.GetString("kvconfig"))
	if err != nil {
		return nil
	}
	defer file.Close()

	resourceBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(resourceBytes, &kvconfig)
	if err != nil {
		return nil
	}

	return kvconfig
}

func TestNewIndigoGRPCServer(t *testing.T) {
	port := 12890
	path := "/tmp/hoge"
	indexMapping := loadIndexMapping("../../example/index_mapping.json")
	indexType := "upside_down"
	kvstore := "boltdb"
	kvconfig := loadKvconfig("../../example/kvconfig.json")

	server := NewIndigoGRPCServer(port, path, indexMapping, indexType, kvstore, kvconfig)

	if server == nil {
		t.Fatalf("unexpected error.  expected not nil, actual %v", server)
	}
}
