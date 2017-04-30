package cmd

import (
	"github.com/mosuka/indigo/constant"
)

var (
	versionFlag  bool   = constant.DefaultVersionFlag
	outputFormat string = constant.DefaultOutputFormat

	gRPCServer string = constant.DefaultGRPCServer

	index string = constant.DefaultIndex

	bulkRequest string = constant.DefaultBulkRequest
	batchSize   int32  = constant.DefaultBatchSize

	indexMapping  string = constant.DefaultIndexMapping
	indexType     string = constant.DefaultIndexType
	kvStore       string = constant.DefaultKVStore
	kvConfig      string = constant.DefaultKVConfigFile
	runtimeConfig string = constant.DefaultRuntimeConfig

	docID     string = constant.DefaultDocID
	docFields string = constant.DefaultDocFields

	searchRequest string   = constant.DefaultSearchRequestFile
	query         string   = constant.DefaultQuery
	size          int      = constant.DefaultSize
	from          int      = constant.DefaultFrom
	explain       bool     = constant.DefaultExplain
	fields        []string = constant.DefaultFields
)
