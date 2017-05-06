package defaultvalue

const (
	DefaultVersionFlag       bool   = false
	DefaultConfig            string = ""
	DefaultLogFormat         string = "text"
	DefaultLogOutput         string = ""
	DefaultLogLevel          string = "info"
	DefaultGRPCPort          int    = 1289
	DefaultDataDir           string = "/var/indigo/data"
	DefaultOpenExistingIndex bool   = false
	DefaultRESTPort          int    = 2289
	DefaultGRPCServer        string = "localhost:1289"
	DefaultBaseURI           string = "/api"

	DefaultOutputFormat      string = "json"
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
	DefaultQuery             string = ""
	DefaultFrom              int    = 0
	DefaultSize              int    = 10
	DefaultExplain           bool   = false
	DefaultFacets            string = ""
	DefaultHighlight         string = ""
	DefaultHighlightStyle    string = "html"
	DefaultIncludeLocations  bool   = false
)

var (
	DefaultFields          []string = []string{}
	DefaultSorts           []string = []string{}
	DefaultHighlightFields []string = []string{}
)
