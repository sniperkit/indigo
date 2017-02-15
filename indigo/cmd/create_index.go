package cmd

import (
	"errors"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var createIndexCmd = &cobra.Command{
	Use:   "index INDEX_NAME INDEX_MAPPING",
	Short: "creates the index to the Indigo gRPC Server",
	Long:  `The create index command creates the index to the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("few arguments")
		}

		indexName := args[0]
		indexMapping := args[1]

		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", grpcServerName, grpcServerPort), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		c := proto.NewIndigoClient(conn)
		resp, err := c.CreateIndex(context.Background(), &proto.CreateIndexRequest{IndexName: indexName, IndexMapping: indexMapping, IndexType: indexType, IndexStore: indexStore})
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", resp.Result)

		return nil
	},
}

func init() {
	createIndexCmd.Flags().StringVarP(&indexType, "index-type", "t", indexType, "index type")
	createIndexCmd.Flags().StringVarP(&indexStore, "index-store", "s", indexStore, "index store")

	createCmd.AddCommand(createIndexCmd)
}
