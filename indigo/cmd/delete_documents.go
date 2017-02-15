package cmd

import (
	"errors"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var deleteDocumentsCmd = &cobra.Command{
	Use:   "documents INDEX_NAME DOCUMENT_IDS",
	Short: "indexes the documents to the Indigo gRPC Server",
	Long:  `The index command indexes the JSON representation of the documents to the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("few arguments")
		}

		indexName := args[0]
		documentIds := args[1]

		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", grpcServerName, grpcServerPort), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		c := proto.NewIndigoClient(conn)
		resp, err := c.DeleteDocuments(context.Background(), &proto.DeleteDocumentsRequest{IndexName: indexName, Ids: documentIds, BatchSize: batchSize})
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", resp.Result)

		return nil
	},
}

func init() {
	deleteDocumentsCmd.Flags().Int32VarP(&batchSize, "batch-size", "b", batchSize, "port number")

	deleteCmd.AddCommand(deleteDocumentsCmd)
}
