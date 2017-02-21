package cmd

import (
	"fmt"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "searches the documents from the Indigo gRPC Server",
	Long:  `The search command searches the documents from the Indigo gRPC Server.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if indexName == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index-name").Name)
		}

		if searchRequestFile == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("search-request-file").Name)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
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

		conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)

		resp, err := client.Search(context.Background(), &proto.SearchRequest{IndexName: indexName, SearchRequest: searchRequest})
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", resp.SearchResult)

		return nil
	},
}

func init() {
	searchCmd.Flags().StringVarP(&gRPCServer, "grpc-server", "g", constant.DefaultGRPCServer, "Indigo gRPC Sever")
	searchCmd.Flags().StringVarP(&indexName, "index-name", "i", constant.DefaultIndexName, "index name")
	searchCmd.Flags().StringVarP(&searchRequestFile, "search-request-file", "r", constant.DefaultSearchRequestFile, "search request file")

	RootCmd.AddCommand(searchCmd)
}
