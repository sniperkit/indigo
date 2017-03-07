package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/constant"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "searches the documents from the Indigo gRPC Server",
	Long:  `The search command searches the documents from the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if indexName == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index-name").Name)
		}

		if searchRequestFile == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("search-request").Name)
		}

		searchRequest := make([]byte, 0)
		file, err := os.Open(searchRequestFile)
		if err != nil {
			return err
		}
		defer file.Close()

		searchRequest, err = ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		conn, err := grpc.Dial(viper.GetString("grpc_server"), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)

		resp, err := client.Search(context.Background(), &proto.SearchRequest{IndexName: indexName, SearchRequest: searchRequest})
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

		switch viper.GetString("output_format") {
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
	searchCmd.Flags().StringP("grpc-server", "g", constant.DefaultGRPCServer, "Indigo gRPC Sever")
	searchCmd.Flags().StringVarP(&indexName, "index-name", "n", constant.DefaultIndexName, "index name")
	searchCmd.Flags().StringVarP(&searchRequestFile, "search-request", "s", constant.DefaultSearchRequestFile, "search request file")

	viper.BindPFlag("grpc_server", searchCmd.Flags().Lookup("grpc-server"))

	RootCmd.AddCommand(searchCmd)
}
