package constant

const (
	DefaultVersionFlag bool = false

	DefaultLogOutputFile string = ""
	DefaultLogLevel      string = "info"
	DefaultLogFormat     string = "text"

	DefaultGRPCServerName string = "localhost"
	DefaultGRPCServerPort int    = 1289
	DefaultRESTServerPort int    = 2289

	DefaultDataDir string = "./data"

	DefaultBaseURI string = "/api"

	DefaultIndexName        string = "default"
	DefaultIndexMappingFile string = ""
	DefaultIndexType        string = "upside_down"
	DefaultIndexStore       string = "boltdb"

	DefaultBatchSize int32 = 1000
)
