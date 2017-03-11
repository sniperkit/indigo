package cmd

import (
	"github.com/mosuka/indigo/constant"
)

var (
	versionFlag bool = constant.DefaultVersionFlag

	//configFile        string = constant.DefaultConfigFile
	//outputFormat      string = constant.DefaultOutputFormat
	//logOutputFile     string = constant.DefaultLogOutputFile
	//logLevel          string = constant.DefaultLogLevel
	//gRPCPort          int    = constant.DefaultGRPCPort
	//dataDir           string = constant.DefaultDataDir
	//openExistingIndex bool   = constant.DefaultOpenExistingIndex
	//batchSize int32 = constant.DefaultBatchSize

	//restPort   int    = constant.DefaultRESTPort
	//baseURI    string = constant.DefaultBaseURI
	//gRPCServer string = constant.DefaultGRPCServer

	indexName         string = constant.DefaultIndexName
	indexMappingFile  string = constant.DefaultIndexMappingFile
	indexType         string = constant.DefaultIndexType
	kvStore           string = constant.DefaultKVStore
	kvConfigFile      string = constant.DefaultKVConfigFile
	runtimeConfigFile string = constant.DefaultRuntimeConfigFile
	documentID        string = constant.DefaultDocumentID
	documentFile      string = constant.DefaultDocumentFile
	bulkRequestFile   string = constant.DefaultBulkRequestFile
	searchRequestFile string = constant.DefaultSearchRequestFile
)
