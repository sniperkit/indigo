package constant

const (
	DefaultVersionFlag bool = false

	DefaultLogOutputFile string = ""
	DefaultLogLevel      string = "info"
	DefaultLogFormat     string = "text"

	DefaultGRPCServerPort int    = 1289
	DefaultDataDir        string = "./data"

	DefaultRESTServerPort int    = 2289
	DefaultGRPCServer     string = "localhost:1289"
	DefaultBaseURI        string = "/api"

	DefaultIndexName string = ""

	DefaultIndexMappingFile string = ""
	DefaultIndexType        string = "upside_down"
	DefaultIndexStore       string = "boltdb"
	DefaultKVConfigFile     string = ""

	DefaultDocumentID   string = ""
	DefaultDocumentFile string = ""

	DefaultBulkRequestFile string = ""
	DefaultBatchSize       int32  = 1000

	DefaultSearchRequestFile string = ""
)
