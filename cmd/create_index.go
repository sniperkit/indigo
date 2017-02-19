package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strings"
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
		buf := new(bytes.Buffer)
		buf.ReadFrom(strings.NewReader(args[1]))
		indexMapping := buf.Bytes()

		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", gRPCServerName, gRPCServerPort), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.CreateIndex(context.Background(), &proto.CreateIndexRequest{Name: indexName, Mapping: indexMapping, Type: indexType, Store: indexStore})
		if err != nil {
			return err
		}

		fmt.Printf("%s created\n", resp.Name)

		return nil
	},
}

func init() {
	createIndexCmd.Flags().StringVarP(&indexType, "index-type", "t", indexType, "index type")
	createIndexCmd.Flags().StringVarP(&indexStore, "index-store", "s", indexStore, "index store")

	createCmd.AddCommand(createIndexCmd)
}
