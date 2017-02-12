package cmd

import (
	"github.com/spf13/viper"
)

var versionFlag bool = false

var configDir string = "."
var config *viper.Viper = viper.New()

var grpcServerName string = "localhost"
var grpcServerPort int = 10000
var grpcLogFile string = "./indigo_grpc.log"
var grpcLogLevel string = "info"
var grpcLogFormat string = "text"

var indexDir string = "./index"
var indexType string = "upside_down"
var indexStore string = "boltdb"
var indexMapping string = "./mapping.json"

var restServerName string = "localhost"
var restServerPort int = 10000
var restServerURIPath string = "/api"
var restLogFile string = "./indigo_grpc.log"
var restLogLevel string = "info"
var restLogFormat string = "text"

var batchSize int32 = 1000
var deleteFlag bool = false
