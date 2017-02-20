package cmd

import (
	"github.com/mosuka/indigo/constant"
)

var (
	versionFlag bool = constant.DefaultVersionFlag

	logOutputFile string = constant.DefaultLogOutputFile
	logLevel      string = constant.DefaultLogLevel
	logFormat     string = constant.DefaultLogFormat

	gRPCServerName string = constant.DefaultGRPCServerName
	gRPCServerPort int    = constant.DefaultGRPCServerPort
	restServerPort int    = constant.DefaultRESTServerPort

	dataDir string = constant.DefaultDataDir

	baseURI string = constant.DefaultBaseURI

	indexName        string = constant.DefaultIndexName
	indexMappingFile string = constant.DefaultIndexMappingFile
	indexType        string = constant.DefaultIndexType
	indexStore       string = constant.DefaultIndexStore

	batchSize int32 = constant.DefaultBatchSize
)
