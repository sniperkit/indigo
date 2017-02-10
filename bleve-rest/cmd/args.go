package cmd

import "github.com/spf13/viper"

var configDir string = "."
var config *viper.Viper = viper.New()

var logFile string = "./bleve-http.log"
var logLevel string = "info"
var logFormat string = "text"

var serverName string = "localhost"
var serverPort int = 20000

var bleveServerName string = "localhost"
var bleveServerPort int = 10000
