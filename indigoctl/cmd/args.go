package cmd

import (
	"github.com/mosuka/indigo/constant"
)

var (
	versionFlag  bool   = constant.DefaultVersionFlag
	outputFormat string = constant.DefaultOutputFormat

	gRPCServer string = constant.DefaultGRPCServer

	//configFile        string = constant.DefaultConfig
	//logOutputFile     string = constant.DefaultLogOutput
	//logLevel          string = constant.DefaultLogLevel
	//gRPCPort          int    = constant.DefaultGRPCPort
	//dataDir           string = constant.DefaultDataDir
	//openExistingIndex bool   = constant.DefaultOpenExistingIndex

	index string = constant.DefaultIndex

	bulkRequest string = constant.DefaultBulkRequest
	batchSize   int32  = constant.DefaultBatchSize

	//restPort   int    = constant.DefaultRESTPort
	//baseURI    string = constant.DefaultBaseURI

	indexMapping  string = constant.DefaultIndexMapping
	indexType     string = constant.DefaultIndexType
	kvStore       string = constant.DefaultKVStore
	kvConfig      string = constant.DefaultKVConfigFile
	runtimeConfig string = constant.DefaultRuntimeConfig

	docID     string = constant.DefaultDocID
	docFields string = constant.DefaultDocFields

	searchRequest string = constant.DefaultSearchRequestFile
)
