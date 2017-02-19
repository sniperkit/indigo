package cmd

import (
	"errors"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var deleteIndexCmd = &cobra.Command{
	Use:   "index NAME",
	Short: "deletes the index from the Indigo gRPC Server",
	Long:  `The delete index command deletes the index from the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("few arguments")
		}

		indexName := args[0]

		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", gRPCServerName, gRPCServerPort), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.DeleteIndex(context.Background(), &proto.DeleteIndexRequest{Name: indexName})
		if err != nil {
			return err
		}

		fmt.Printf("%s deleted\n", resp.Name)

		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteIndexCmd)
}