package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/mosuka/indigo/constant"
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

	searchRequest.Query = bleve.NewQueryStringQuery(query)
	searchRequest.Size = size
	searchRequest.From = from
	searchRequest.Explain = explain
	searchRequest.Fields = fields

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
	SearchCmd.Flags().StringVar(&gRPCServer, "grpc-server", constant.DefaultGRPCServer, "Indigo gRPC Server to connect to")
	SearchCmd.Flags().StringVar(&index, "index", constant.DefaultIndex, "index name")
	SearchCmd.Flags().StringVar(&searchRequest, "search-request", constant.DefaultSearchRequestFile, "search request file")
	SearchCmd.Flags().StringVar(&query, "query", constant.DefaultQuery, "query string")
	SearchCmd.Flags().IntVar(&size, "size", constant.DefaultSize, "number of hits to return")
	SearchCmd.Flags().IntVar(&from, "from", constant.DefaultFrom, "starting from index of the hits to return")
	SearchCmd.Flags().BoolVar(&explain, "explain", constant.DefaultExplain, "contain an explanation of how scoring of the hits was computed")
	SearchCmd.Flags().StringSliceVar(&fields, "field", constant.DefaultFields, "specify a set of fields to return")

	RootCmd.AddCommand(SearchCmd)
}
