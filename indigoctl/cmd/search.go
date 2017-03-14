package cmd

import (
	"encoding/json"
	"fmt"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		if index == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
		}

		if searchRequest == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("search-request").Name)
		}

		sr := make([]byte, 0)
		file, err := os.Open(searchRequest)
		if err != nil {
			return err
		}
		defer file.Close()

		sr, err = ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)

		resp, err := client.Search(context.Background(), &proto.SearchRequest{IndexName: index, SearchRequest: sr})
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
	},
}

func init() {
	SearchCmd.Flags().StringVarP(&gRPCServer, "grpc-server", "g", constant.DefaultGRPCServer, "Indigo gRPC Sever")
	SearchCmd.Flags().StringVarP(&index, "index", "i", constant.DefaultIndex, "index name")
	SearchCmd.Flags().StringVarP(&searchRequest, "search-request", "s", constant.DefaultSearchRequestFile, "search request file")

	RootCmd.AddCommand(SearchCmd)
}
