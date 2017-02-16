package cmd

import (
	"errors"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var searchDocumentsCmd = &cobra.Command{
	Use:   "documents SEARCH_REQUEST",
	Short: "searches the Indigo gRPC Serve with the search request",
	Long:  `The search documents command searches the documents from  the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("few arguments")
		}

		searchRequest := args[0]

		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", grpcServerName, grpcServerPort), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		c := proto.NewIndigoClient(conn)

		resp, err := c.SearchDocuments(context.Background(), &proto.SearchDocumentsRequest{SearchRequest: searchRequest})
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", resp.SearchResult)

		return nil
	},
}

func init() {
	searchCmd.AddCommand(searchDocumentsCmd)
}
