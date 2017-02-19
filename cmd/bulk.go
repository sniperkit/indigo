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

var bulkCmd = &cobra.Command{
	Use:   "bulk INDEX_NAME BULK_REQUEST",
	Short: "indexes the documents in bulk to the Indigo gRPC Server",
	Long:  `The bulk command indexes the documents in bulk to the Indigo gRPC Server.`,
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
		resp, err := client.Bulk(context.Background(), &proto.BulkRequest{Name: indexName, BulkRequest: documents, BatchSize: batchSize})
		if err != nil {
			return err
		}

		fmt.Printf("%d documents put in bulk\n", resp.PutCount)
		fmt.Printf("%d error documents occurred in bulk\n", resp.PutErrorCount)
		fmt.Printf("%d documents deleted in bulk\n", resp.DeleteCount)

		return nil
	},
}

func init() {
	bulkCmd.Flags().StringVarP(&gRPCServerName, "grpc-server-name", "n", gRPCServerName, "Indigo gRPC Sever name")
	bulkCmd.Flags().IntVarP(&gRPCServerPort, "grpc-server-port", "p", gRPCServerPort, "Indigo gRPC Server port number")
	bulkCmd.Flags().Int32VarP(&batchSize, "batch-size", "b", batchSize, "port number")

	RootCmd.AddCommand(bulkCmd)
}
