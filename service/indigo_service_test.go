package service

import (
	"github.com/mosuka/indigo/test"
	"io/ioutil"
	"os"
	"testing"
)

func TestIndigoGRPCService(t *testing.T) {
	dir, _ := os.Getwd()

	path, _ := ioutil.TempDir("/tmp", "indigo")
	indexMappingPath := dir + "/../example/index_mapping.json"
	indexType := "upside_down"
	kvstore := "boltdb"
	kvconfigPath := dir + "/../example/kvconfig.json"

	indexMapping, err := test.LoadIndexMapping(indexMappingPath)
	if err != nil {
		t.Errorf("could not load IndexMapping %v", indexMappingPath)
	}
	kvconfig, err := test.LoadKvconfig(kvconfigPath)
	if err != nil {
		t.Errorf("could not load kvconfig %v", kvconfigPath)
	}
	kvconfig["path"] = path + "/store"

	s := NewIndigoGRPCService(path, indexMapping, indexType, kvstore, kvconfig)
	if s == nil {
		t.Fatalf("unexpected error.  expected not nil, actual %v", s)
	}

}
