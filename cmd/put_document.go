package cmd

import (
	"fmt"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
)

var putDocumentCmd = &cobra.Command{
	Use:   "document",
	Short: "puts the document to the Indigo gRPC Server",
	Long:  `The index document command puts the document to the Indigo gRPC Server.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if indexName == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index-name").Name)
		}

		if documentID == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("document-id").Name)
		}

		if documentFile == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("document-file").Name)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		document := make([]byte, 0)
		file, err := os.Open(documentFile)
		if err != nil {
			return err
		}
		defer file.Close()

		document, err = ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.PutDocument(context.Background(), &proto.PutDocumentRequest{IndexName: indexName, DocumentID: documentID, Document: document})
		if err != nil {
			return err
		}

		fmt.Printf("%d document put\n", resp.PutCount)
		fmt.Printf("%d error document occurred\n", resp.PutErrorCount)

		return nil
	},
}

func init() {
	putDocumentCmd.Flags().StringVarP(&indexName, "index-name", "i", constant.DefaultIndexName, "index name")
	putDocumentCmd.Flags().StringVarP(&documentID, "document-id", "I", constant.DefaultDocumentID, "document id")
	putDocumentCmd.Flags().StringVarP(&documentFile, "document-file", "d", constant.DefaultDocumentFile, "document file")

	putCmd.AddCommand(putDocumentCmd)
}
