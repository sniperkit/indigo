package constant

const (
	DefaultVersionFlag       bool   = false
	DefaultConfigFile        string = ""
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
