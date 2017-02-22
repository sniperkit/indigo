package cmd

import (
	"github.com/mosuka/indigo/constant"
)

var (
	versionFlag bool = constant.DefaultVersionFlag

	logOutputFile string = constant.DefaultLogOutputFile
	logLevel      string = constant.DefaultLogLevel
	logFormat     string = constant.DefaultLogFormat

	gRPCServerPort int    = constant.DefaultGRPCServerPort
	dataDir        string = constant.DefaultDataDir

	restServerPort int    = constant.DefaultRESTServerPort
	gRPCServer     string = constant.DefaultGRPCServer
	baseURI        string = constant.DefaultBaseURI

	indexName string = constant.DefaultIndexName

	indexMappingFile string = constant.DefaultIndexMappingFile
	indexType        string = constant.DefaultIndexType
	indexStore       string = constant.DefaultIndexStore
	kvConfigFile     string = constant.DefaultKVConfigFile

	documentID   string = constant.DefaultDocumentID
	documentFile string = constant.DefaultDocumentFile

	bulkRequestFile string = constant.DefaultBulkRequestFile
	batchSize       int32  = constant.DefaultBatchSize

	searchRequestFile string = constant.DefaultSearchRequestFile
)
