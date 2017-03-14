package constant

const (
	DefaultVersionFlag       bool   = false
	DefaultConfig            string = ""
	DefaultOutputFormat      string = "text"
	DefaultLogOutput         string = ""
	DefaultLogLevel          string = "info"
	DefaultGRPCPort          int    = 1289
	DefaultDataDir           string = "./data"
	DefaultOpenExistingIndex bool   = false
	DefaultRESTPort          int    = 2289
	DefaultGRPCServer        string = "localhost:1289"
	DefaultBaseURI           string = "/api"
	DefaultIndex             string = ""
	DefaultIndexMapping      string = ""
	DefaultIndexType         string = "upside_down"
	DefaultKVStore           string = "boltdb"
	DefaultKVConfigFile      string = ""
	DefaultRuntimeConfig     string = ""
	DefaultDocID             string = ""
	DefaultDocFields         string = ""
	DefaultBulkRequest       string = ""
	DefaultBatchSize         int32  = 1000
	DefaultSearchRequestFile string = ""
)
