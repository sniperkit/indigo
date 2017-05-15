package cmd

import (
	"encoding/json"
	"fmt"
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

type PutDocumentResource struct {
	Index  string      `json:"index,omitempty"`
	Id     string      `json:"id,omitempty"`
	Fields interface{} `json:"fields,omitempty"`
}

func runEPutDocumentCmd(cmd *cobra.Command, args []string) error {
	var resourceBytes []byte = nil
	if cmd.Flag("resource").Changed {
		if cmd.Flag("resource").Value.String() == "-" {
			resourceBytes, _ = ioutil.ReadAll(os.Stdin)
		} else {
			file, err := os.Open(cmd.Flag("resource").Value.String())
			if err != nil {
				return err
			}
			defer file.Close()

			resourceBytes, err = ioutil.ReadAll(file)
			if err != nil {
				return err
			}
		}
	}

	putDocumentResource := PutDocumentResource{}
	err := json.Unmarshal(resourceBytes, &putDocumentResource)
	if err != nil {
		return err
	}

	fieldsBytes, err := json.Marshal(putDocumentResource.Fields)
	if err != nil {
		return err
	}

	putDocumentRequest := &proto.PutDocumentRequest{
		Index:  putDocumentResource.Index,
		Id:     putDocumentResource.Id,
		Fields: fieldsBytes,
	}

	if cmd.Flag("index").Changed {
		putDocumentRequest.Index = cmd.Flag("index").Value.String()
	}

	if cmd.Flag("id").Changed {
		putDocumentRequest.Id = cmd.Flag("id").Value.String()
	}

	if cmd.Flag("fields").Changed {
		fieldsBytes := []byte(cmd.Flag("fields").Value.String())
		putDocumentRequest.Fields = fieldsBytes
	}

	conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.PutDocument(context.Background(), putDocumentRequest)
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
	PutDocumentCmd.Flags().String("resource", "", "resource file")
	PutDocumentCmd.Flags().String("index", "", "index name")
	PutDocumentCmd.Flags().String("id", "", "document id")
	PutDocumentCmd.Flags().String("fields", "", "document fields")

	PutCmd.AddCommand(PutDocumentCmd)
}
