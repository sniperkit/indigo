package cmd

import (
	"github.com/mosuka/indigo/setting"
)

var (
	versionFlag bool = setting.DefaultVersionFlag

	configFile        string = setting.DefaultConfigFile
	outputFormat      string = setting.DefaultOutputFormat
	logOutputFile     string = setting.DefaultLogOutputFile
	logLevel          string = setting.DefaultLogLevel
	gRPCPort          int    = setting.DefaultGRPCPort
	dataDir           string = setting.DefaultDataDir
	openExistingIndex bool   = setting.DefaultOpenExistingIndex
	restPort          int    = setting.DefaultRESTPort
	baseURI           string = setting.DefaultBaseURI
	gRPCServer        string = setting.DefaultGRPCServer

	indexName         string = setting.DefaultIndexName
	indexMappingFile  string = setting.DefaultIndexMappingFile
	indexType         string = setting.DefaultIndexType
	kvStore           string = setting.DefaultKVStore
	kvConfigFile      string = setting.DefaultKVConfigFile
	runtimeConfigFile string = setting.DefaultRuntimeConfigFile
	documentID        string = setting.DefaultDocumentID
	documentFile      string = setting.DefaultDocumentFile
	bulkRequestFile   string = setting.DefaultBulkRequestFile
	batchSize         int32  = setting.DefaultBatchSize
	searchRequestFile string = setting.DefaultSearchRequestFile
)
