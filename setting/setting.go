package setting

import (
	"github.com/spf13/viper"
)

const (
	DefaultVersionFlag       bool   = false
	DefaultOutputFormat      string = "text"
	DefaultLogOutputFile     string = ""
	DefaultLogLevel          string = "info"
	DefaultGRPCPort          int    = 1289
	DefaultDataDir           string = "./data"
	DefaultOpenExistingIndex bool   = false
	DefaultRESTPort          int    = 2289
	DefaultGRPCServer        string = "localhost:1289"
	DefaultBaseURI           string = "/api"
	DefaultIndexName         string = ""
	DefaultIndexMappingFile  string = ""
	DefaultIndexType         string = "upside_down"
	DefaultKVStore           string = "boltdb"
	DefaultKVConfigFile      string = ""
	DefaultRuntimeConfigFile string = ""
	DefaultDocumentID        string = ""
	DefaultDocumentFile      string = ""
	DefaultBulkRequestFile   string = ""
	DefaultBatchSize         int32  = 1000
	DefaultSearchRequestFile string = ""
)

//var IndigoSettings = viper.New()

func init() {
	viper.SetDefault("output_format", DefaultOutputFormat)
	viper.SetDefault("log_output", DefaultLogOutputFile)
	viper.SetDefault("log_level", DefaultLogLevel)
	viper.SetDefault("grpc_port", DefaultGRPCPort)
	viper.SetDefault("data_dir", DefaultDataDir)
	viper.SetDefault("open_existing_index", DefaultOpenExistingIndex)
	viper.SetDefault("rest_port", DefaultRESTPort)
	viper.SetDefault("grpc_server", DefaultGRPCServer)
	viper.SetDefault("base_uri", DefaultBaseURI)

	viper.SetEnvPrefix("indigo")
	viper.BindEnv("output_format")
	viper.BindEnv("log_output")
	viper.BindEnv("log_level")
	viper.BindEnv("grpc_port")
	viper.BindEnv("data_dir")
	viper.BindEnv("open_existing_index")
	viper.BindEnv("rest_port")
	viper.BindEnv("grpc_server")
	viper.BindEnv("base_uri")
}
