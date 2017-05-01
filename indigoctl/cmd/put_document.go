package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/defaultvalue"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
)

var PutDocumentCmd = &cobra.Command{
	Use:   "document",
	Short: "puts the document to the Indigo gRPC Server",
	Long:  `The index document command puts the document to the Indigo gRPC Server.`,
	RunE:  runEPutDocumentCmd,
}

func runEPutDocumentCmd(cmd *cobra.Command, args []string) error {
	if index == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
	}

	if docID == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("id").Name)
	}

	if docFields == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("fields").Name)
	}

	document := make([]byte, 0)
	file, err := os.Open(docFields)
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
	resp, err := client.PutDocument(context.Background(), &proto.PutDocumentRequest{Index: index, Id: docID, Fields: document})
	if err != nil {
		return err
	}

	switch outputFormat {
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
}

func init() {
	PutDocumentCmd.Flags().StringVar(&index, "index", defaultvalue.DefaultIndex, "index name")
	PutDocumentCmd.Flags().StringVar(&docID, "id", defaultvalue.DefaultDocID, "document id")
	PutDocumentCmd.Flags().StringVar(&docFields, "fields", defaultvalue.DefaultDocFields, "document fields")

	PutCmd.AddCommand(PutDocumentCmd)
}
