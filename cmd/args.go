package cmd

import (
	"github.com/mosuka/indigo/constant"
)

var (
	versionFlag bool = constant.DefaultVersionFlag

	//configFile        string = constant.DefaultConfig
	//outputFormat      string = constant.DefaultOutputFormat
	//logOutputFile     string = constant.DefaultLogOutput
	//logLevel          string = constant.DefaultLogLevel
	//gRPCPort          int    = constant.DefaultGRPCPort
	//dataDir           string = constant.DefaultDataDir
	//openExistingIndex bool   = constant.DefaultOpenExistingIndex
	//batchSize int32 = constant.DefaultBatchSize

	//restPort   int    = constant.DefaultRESTPort
	//baseURI    string = constant.DefaultBaseURI
	//gRPCServer string = constant.DefaultGRPCServer

	//index         string = constant.DefaultIndex

	indexMapping  string = constant.DefaultIndexMapping
	indexType     string = constant.DefaultIndexType
	kvStore       string = constant.DefaultKVStore
	kvConfig      string = constant.DefaultKVConfigFile
	runtimeConfig string = constant.DefaultRuntimeConfig

	docID     string = constant.DefaultDocID
	docFields string = constant.DefaultDocFields

	bulkRequest   string = constant.DefaultBulkRequest
	searchRequest string = constant.DefaultSearchRequestFile
)
