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

var deleteBulkCmd = &cobra.Command{
	Use:   "bulk INDEX_NAME DOCUMENT_IDS",
	Short: "deletes the documents in bulk from the Indigo gRPC Server",
	Long:  `The delete documents command deletes the documents in bulk from the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("few arguments")
		}

		indexName := args[0]
		buf := new(bytes.Buffer)
		buf.ReadFrom(strings.NewReader(args[1]))
		ids := buf.Bytes()

		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", gRPCServerName, gRPCServerPort), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.DeleteBulk(context.Background(), &proto.DeleteBulkRequest{Name: indexName, Ids: ids, BatchSize: batchSize})
		if err != nil {
			return err
		}

		fmt.Printf("%d documents deleted in bulk\n", resp.Count)

		return nil
	},
}

func init() {
	deleteBulkCmd.Flags().Int32VarP(&batchSize, "batch-size", "b", batchSize, "port number")

	deleteCmd.AddCommand(deleteBulkCmd)
}
