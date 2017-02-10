package cmd

import "github.com/spf13/viper"

var configDir string = "."
var config *viper.Viper = viper.New()

var logFile string = "./bleve-server.log"
var logLevel string = "info"
var logFormat string = "text"

var serverName string = "localhost"
var serverPort int = 10000

var indexDir string = "./index"
var indexType string = "upside_down"
var indexStore string = "boltdb"
var indexMapping string = "./mapping.json"
