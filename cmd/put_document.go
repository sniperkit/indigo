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

var putDocumentCmd = &cobra.Command{
	Use:   "document INDEX_NAME DOCUMENT_ID DOCUMENT",
	Short: "puts the document to the Indigo gRPC Server",
	Long:  `The index document command puts the document to the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("few arguments")
		}

		indexName := args[0]
		id := args[1]
		buf := new(bytes.Buffer)
		buf.ReadFrom(strings.NewReader(args[2]))
		document := buf.Bytes()

		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", gRPCServerName, gRPCServerPort), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.PutDocument(context.Background(), &proto.PutDocumentRequest{IndexName: indexName, DocumentID: id, Document: document})
		if err != nil {
			return err
		}

		fmt.Printf("%d document put\n", resp.PutCount)
		fmt.Printf("%d error document occurred\n", resp.PutErrorCount)

		return nil
	},
}

func init() {
	putCmd.AddCommand(putDocumentCmd)
}
