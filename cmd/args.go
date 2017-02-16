package cmd

var versionFlag bool = false

var serverPort int = 1289

var logOutputFile string = ""
var logLevel string = "info"
var logFormat string = "text"

var dataDir string = "./data"

var baseURI string = "/api"

var gRPCServerName string = "localhost"
var gRPCServerPort int = 1289

var indexType string = "upside_down"
var indexStore string = "boltdb"
var kvConfig map[string]interface{} = nil

var batchSize int32 = 1000
