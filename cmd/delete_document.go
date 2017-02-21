package cmd

import (
	"fmt"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var deleteDocumentCmd = &cobra.Command{
	Use:   "document",
	Short: "deletes the document from the Indigo gRPC Server",
	Long:  `The delete document command deletes the document from the Indigo gRPC Server.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if indexName == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index-name").Name)
		}

		if documentID == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("document-id").Name)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.DeleteDocument(context.Background(), &proto.DeleteDocumentRequest{IndexName: indexName, DocumentID: documentID})
		if err != nil {
			return err
		}

		fmt.Printf("%d document deleted\n", resp.DeleteCount)

		return nil
	},
}

func init() {
	deleteDocumentCmd.Flags().StringVarP(&indexName, "index-name", "i", constant.DefaultIndexName, "index name")
	deleteDocumentCmd.Flags().StringVarP(&documentID, "document-id", "I", constant.DefaultDocumentID, "document id")

	deleteCmd.AddCommand(deleteDocumentCmd)
}
