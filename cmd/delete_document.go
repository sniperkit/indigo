package cmd

import (
	"errors"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var deleteDocumentCmd = &cobra.Command{
	Use:   "document INDEX_NAME DOCUMENT_ID",
	Short: "deletes the document from the Indigo gRPC Server",
	Long:  `The delete document command deletes the document from the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("few arguments")
		}

		indexName := args[0]
		id := args[1]

		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", gRPCServerName, gRPCServerPort), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.DeleteDocument(context.Background(), &proto.DeleteDocumentRequest{IndexName: indexName, DocumentID: id})
		if err != nil {
			return err
		}

		fmt.Printf("%d document deleted\n", resp.DeleteCount)

		return nil
	},
}

func init() {
	deleteCmd.AddCommand(deleteDocumentCmd)
}
