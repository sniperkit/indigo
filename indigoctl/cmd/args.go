package cmd

import (
	"github.com/mosuka/indigo/defaultvalue"
)

var (
	versionFlag   bool     = defaultvalue.DefaultVersionFlag
	outputFormat  string   = defaultvalue.DefaultOutputFormat
	gRPCServer    string   = defaultvalue.DefaultGRPCServer
	index         string   = defaultvalue.DefaultIndex
	bulkRequest   string   = defaultvalue.DefaultBulkRequest
	batchSize     int32    = defaultvalue.DefaultBatchSize
	indexMapping  string   = defaultvalue.DefaultIndexMapping
	indexType     string   = defaultvalue.DefaultIndexType
	kvStore       string   = defaultvalue.DefaultKVStore
	kvConfig      string   = defaultvalue.DefaultKVConfigFile
	runtimeConfig string   = defaultvalue.DefaultRuntimeConfig
	docID         string   = defaultvalue.DefaultDocID
	docFields     string   = defaultvalue.DefaultDocFields
	searchRequest string   = defaultvalue.DefaultSearchRequestFile
	query         string   = defaultvalue.DefaultQuery
	size          int      = defaultvalue.DefaultSize
	from          int      = defaultvalue.DefaultFrom
	explain       bool     = defaultvalue.DefaultExplain
	fields        []string = defaultvalue.DefaultFields
)
