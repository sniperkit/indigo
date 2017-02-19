package cmd

import (
	"errors"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var getDocumentCmd = &cobra.Command{
	Use:   "document INDEX_NAME DOCUMENT_ID",
	Short: "gets the document from the Indigo gRPC Server",
	Long:  `The get document command gets the document from the Indigo gRPC Server.`,
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
		resp, err := client.GetDocument(context.Background(), &proto.GetDocumentRequest{IndexName: indexName, DocumentID: id})
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", resp.Document)

		return nil
	},
}

func init() {
	getCmd.AddCommand(getDocumentCmd)
}
