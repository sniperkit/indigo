package cmd

import "github.com/spf13/viper"

var configDir string = "."
var config *viper.Viper = viper.New()

var logFile string = "./indigo_rest.log"
var logLevel string = "info"
var logFormat string = "text"

var serverName string = "localhost"
var serverPort int = 20000

var gRPCServerName string = "localhost"
var gRPCServerPort int = 10000
