package cmd

import (
	"github.com/mosuka/indigo/setting"
)

var (
	versionFlag       bool   = setting.DefaultVersionFlag
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
