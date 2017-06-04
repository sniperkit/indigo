package cmd

const (
	DefaultOutputFormat     string = "json"
	DefaultVersionFlag      bool   = false
	DefaultResource         string = ""
	DefaultServer           string = "localhost:1289"
	DefaultBatchSize        int32  = 1000
	DefaultId               string = ""
	DefaultDocFields        string = ""
	DefaultQuery            string = ""
	DefaultSize             int    = 10
	DefaultFrom             int    = 0
	DefaultExplain          bool   = false
	DefaultFacets           string = ""
	DefaultHighlight        string = ""
	DefaultHighlightStyle   string = "html"
	DefaultIncludeLocations bool   = false
)

var (
	DefaultFields          []string = []string{}
	DefaultSorts           []string = []string{}
	DefaultHighlightFields []string = []string{}
)
