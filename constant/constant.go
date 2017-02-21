package constant

const (
	DefaultVersionFlag bool = false

	DefaultLogOutputFile string = ""
	DefaultLogLevel      string = "info"
	DefaultLogFormat     string = "text"

	//DefaultGRPCServerName string = "localhost"
	DefaultGRPCServerPort int = 1289
	DefaultRESTServerPort int = 2289

	DefaultGRPCServer string = "localhost:1289"

	DefaultDataDir string = "./data"

	DefaultBaseURI string = "/api"

	DefaultIndexName        string = ""
	DefaultIndexMappingFile string = ""
	DefaultIndexType        string = "upside_down"
	DefaultIndexStore       string = "boltdb"

	DefaultDocumentID   string = ""
	DefaultDocumentFile string = ""

	DefaultBulkRequestFile string = ""
	DefaultBatchSize       int32  = 1000

	DefaultSearchRequestFile string = ""
)
