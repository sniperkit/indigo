package grpc

import (
	"github.com/mosuka/indigo/test"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestIndigoGRPCServer(t *testing.T) {
	dir, _ := os.Getwd()

	port := 0
	path, _ := ioutil.TempDir("/tmp", "indigo")
	indexMappingPath := dir + "/../../example/index_mapping.json"
	indexType := "upside_down"
	kvstore := "boltdb"
	kvconfigPath := dir + "/../../example/kvconfig.json"

	indexMapping, err := test.LoadIndexMapping(indexMappingPath)
	if err != nil {
		t.Errorf("could not load IndexMapping %v", indexMappingPath)
	}
	kvconfig, err := test.LoadKvconfig(kvconfigPath)
	if err != nil {
		t.Errorf("could not load kvconfig %v", kvconfigPath)
	}
	kvconfig["path"] = path + "/store"

	server := NewIndigoGRPCServer(port, path, indexMapping, indexType, kvstore, kvconfig)

	if server == nil {
		t.Fatalf("unexpected error.  expected not nil, actual %v", server)
	}

	err = server.Start(true)
	if err != nil {
		t.Fatalf("unexpected error. %v", err)
	}

	time.Sleep(10 * time.Second)

	err = server.Stop(false)
	if err != nil {
		t.Fatalf("unexpected error. %v", err)
	}

	os.RemoveAll(path)
}
