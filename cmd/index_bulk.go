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

var indexBulkCmd = &cobra.Command{
	Use:   "bulk INDEX_NAME DOCUMENTS",
	Short: "indexes the documents in bulk to the Indigo gRPC Server",
	Long:  `The index documents command indexes the documents in bulk to the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("few arguments")
		}

		indexName := args[0]
		buf := new(bytes.Buffer)
		buf.ReadFrom(strings.NewReader(args[1]))
		documents := buf.Bytes()

		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", gRPCServerName, gRPCServerPort), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.IndexBulk(context.Background(), &proto.IndexBulkRequest{Name: indexName, Documents: documents, BatchSize: batchSize})
		if err != nil {
			return err
		}

		fmt.Printf("%d documents indexed in bulk\n", resp.Count)

		return nil
	},
}

func init() {
	indexBulkCmd.Flags().Int32VarP(&batchSize, "batch-size", "b", batchSize, "port number")

	indexCmd.AddCommand(indexBulkCmd)
}
