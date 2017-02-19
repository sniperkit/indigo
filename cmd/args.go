package cmd

import (
	"github.com/mosuka/indigo/constant"
)

var versionFlag bool = false

var logOutputFile string = ""
var logLevel string = "info"
var logFormat string = "text"

var gRPCServerName string = "localhost"
var gRPCServerPort int = 1289
var restServerPort int = 2289

var dataDir string = "./data"

var baseURI string = "/api"

var indexType string = constant.DefaultIndexType
var indexStore string = constant.DefaultIndexStore
var kvConfig map[string]interface{} = nil

var batchSize int32 = constant.DefaultBatchSize
