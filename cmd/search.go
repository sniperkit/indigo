package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strings"
)

var searchCmd = &cobra.Command{
	Use:   "search INDEX_NAME SEARCH_REQUEST",
	Short: "searches the documents from the Indigo gRPC Server",
	Long:  `The search command searches the documents from the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("few arguments")
		}

		indexName := args[0]
		buf := new(bytes.Buffer)
		buf.ReadFrom(strings.NewReader(args[1]))
		searchRequest := buf.Bytes()

		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", gRPCServerName, gRPCServerPort), grpc.WithInsecure())
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
	searchCmd.Flags().StringVarP(&gRPCServerName, "grpc-server-name", "n", constant.DefaultGRPCServerName, "Indigo gRPC Sever name")
	searchCmd.Flags().IntVarP(&gRPCServerPort, "grpc-server-port", "p", constant.DefaultGRPCServerPort, "Indigo gRPC Server port number")

	RootCmd.AddCommand(searchCmd)
}
