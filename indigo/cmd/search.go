package cmd

import (
	"errors"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var searchCmd = &cobra.Command{
	Use:   "search REQUEST",
	Short: "searches the Indigo gRPC Serve with the search request",
	Long:  `The search command searches the Indigo gRPC Server with the JSON representation of the search request.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("must specify QUERY")
		}

		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", grpcServerName, grpcServerPort), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		c := proto.NewIndigoClient(conn)

		resp, err := c.Search(context.Background(), &proto.SearchRequest{Request: args[0]})
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", resp.Result)

		return nil
	},
}

func init() {
	clientCmd.AddCommand(searchCmd)
}
