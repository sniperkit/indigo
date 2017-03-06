package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/setting"
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
			return fmt.Errorf("required flag: --%s", cmd.Flag("fields").Name)
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

		conn, err := grpc.Dial(IndigoSettings.GetString("grpc_server"), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.PutDocument(context.Background(), &proto.PutDocumentRequest{IndexName: indexName, Id: documentID, Fields: document})
		if err != nil {
			return err
		}

		switch IndigoSettings.GetString("output_format") {
		case "text":
			fmt.Printf("%s\n", resp.String())
		case "json":
			output, err := json.MarshalIndent(resp, "", "  ")
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", output)
		default:
			fmt.Printf("%s\n", resp.String())
		}

		return nil
	},
}

func init() {
	putDocumentCmd.Flags().StringVarP(&indexName, "index-name", "n", setting.DefaultIndexName, "index name")
	putDocumentCmd.Flags().StringVarP(&documentID, "document-id", "i", setting.DefaultDocumentID, "document id")
	putDocumentCmd.Flags().StringVarP(&documentFile, "fields", "F", setting.DefaultDocumentFile, "fields file")

	putCmd.AddCommand(putDocumentCmd)
}
