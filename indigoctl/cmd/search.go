package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/mosuka/indigo/defaultvalue"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
)

var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "searches the documents from the Indigo gRPC Server",
	Long:  `The search command searches the documents from the Indigo gRPC Server.`,
	RunE:  runESearchCmd,
}

func runESearchCmd(cmd *cobra.Command, args []string) error {
	if index == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
	}
	if query == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("query").Name)
	}

	sr := make([]byte, 0)

	if searchRequest != "" {
		file, err := os.Open(searchRequest)
		if err != nil {
			return err
		}
		defer file.Close()

		sr, err = ioutil.ReadAll(file)
		if err != nil {
			return err
		}
	}

	searchRequest := bleve.NewSearchRequest(nil)
	if len(sr) > 0 {
		err := searchRequest.UnmarshalJSON(sr)
		if err != nil {
			return err
		}
	}

	if cmd.Flag("query").Changed {
		searchRequest.Query = bleve.NewQueryStringQuery(query)
	}
	if cmd.Flag("size").Changed {
		searchRequest.Size = size
	}
	if cmd.Flag("from").Changed {
		searchRequest.From = from
	}
	if cmd.Flag("explain").Changed {
		searchRequest.Explain = explain
	}
	if cmd.Flag("field").Changed {
		searchRequest.Fields = fields
	}
	if cmd.Flag("sort").Changed {
		searchRequest.SortBy(sorts)
	}
	if cmd.Flag("facets").Changed {
		facetRequest := bleve.FacetsRequest{}
		err := json.Unmarshal([]byte(facets), &facetRequest)
		if err != nil {
			return err
		}
		searchRequest.Facets = facetRequest
	}
	if cmd.Flag("highlight").Changed {
		highlightRequest := bleve.NewHighlight()
		err := json.Unmarshal([]byte(highlight), highlightRequest)
		if err != nil {
			return err
		}
		searchRequest.Highlight = highlightRequest
	}

	sr, err := json.Marshal(searchRequest)
	if err != nil {
		return err
	}

	conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)

	resp, err := client.Search(context.Background(), &proto.SearchRequest{Index: index, SearchRequest: sr})
	if err != nil {
		return err
	}

	searchResult := make(map[string]interface{})
	if err := json.Unmarshal(resp.SearchResult, &searchResult); err != nil {
		return err
	}

	r := struct {
		SearchResult map[string]interface{} `json:"search_result"`
	}{
		SearchResult: searchResult,
	}

	switch outputFormat {
	case "text":
		fmt.Printf("%s\n", resp.String())
	case "json":
		output, err := json.MarshalIndent(r, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", output)
	default:
		fmt.Printf("%s\n", resp.String())
	}

	return nil
}

func init() {
	SearchCmd.Flags().StringVar(&gRPCServer, "grpc-server", defaultvalue.DefaultGRPCServer, "Indigo gRPC Server to connect to")
	SearchCmd.Flags().StringVar(&index, "index", defaultvalue.DefaultIndex, "index name")
	SearchCmd.Flags().StringVar(&searchRequest, "search-request", defaultvalue.DefaultSearchRequestFile, "search request file")
	SearchCmd.Flags().StringVar(&query, "query", defaultvalue.DefaultQuery, "query string")
	SearchCmd.Flags().IntVar(&size, "size", defaultvalue.DefaultSize, "number of hits to return")
	SearchCmd.Flags().IntVar(&from, "from", defaultvalue.DefaultFrom, "starting from index of the hits to return")
	SearchCmd.Flags().BoolVar(&explain, "explain", defaultvalue.DefaultExplain, "contain an explanation of how scoring of the hits was computed")
	SearchCmd.Flags().StringSliceVar(&fields, "field", defaultvalue.DefaultFields, "specify a set of fields to return")
	SearchCmd.Flags().StringSliceVar(&sorts, "sort", defaultvalue.DefaultSorts, "sorting to perform")
	SearchCmd.Flags().StringVar(&facets, "facets", defaultvalue.DefaultFacets, "faceting to perform")
	SearchCmd.Flags().StringVar(&highlight, "highlight", defaultvalue.DefaultHighlight, "highlighting to perform")

	RootCmd.AddCommand(SearchCmd)
}
