package cmd

import (
	"errors"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var getMappingCmd = &cobra.Command{
	Use:   "mapping INDEX_NAME",
	Short: "gets the index mapping from the Indigo gRPC Server",
	Long:  `The get mapping command gets the index mapping from the Indigo gRPC Server.`,
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
		resp, err := client.GetMapping(context.Background(), &proto.GetMappingRequest{IndexName: indexName})
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", resp.IndexMapping)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getMappingCmd)
}
