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

var indexDocumentCmd = &cobra.Command{
	Use:   "document INDEX_NAME ID DOCUMENT",
	Short: "indexes the document to the Indigo gRPC Server",
	Long:  `The index document command indexes the document to the Indigo gRPC Server.`,
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
		resp, err := client.IndexDocument(context.Background(), &proto.IndexDocumentRequest{Name: indexName, Id: id, Document: document})
		if err != nil {
			return err
		}

		fmt.Printf("%d document indexed\n", resp.Count)

		return nil
	},
}

func init() {
	indexCmd.AddCommand(indexDocumentCmd)
}
