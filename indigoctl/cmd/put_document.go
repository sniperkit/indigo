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

type PutDocumentCommandOptions struct {
	index    string
	id       string
	fields   string
	resource string
}

var putDocumentCmdOpts PutDocumentCommandOptions

type PutDocumentResource struct {
	Fields interface{} `json:"fields,omitempty"`
}

var putDocumentCmd = &cobra.Command{
	Use:   "document",
	Short: "puts the document to the Indigo gRPC Server",
	Long:  `The index document command puts the document to the Indigo gRPC Server.`,
	RunE:  runEPutDocumentCmd,
}

func runEPutDocumentCmd(cmd *cobra.Command, args []string) error {
	if putDocumentCmdOpts.index == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
	}

	if putDocumentCmdOpts.id == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("id").Name)
	}

	var resourceBytes []byte = nil
	if cmd.Flag("resource").Changed {
		if putDocumentCmdOpts.resource == "-" {
			resourceBytes, _ = ioutil.ReadAll(os.Stdin)
		} else {
			file, err := os.Open(putDocumentCmdOpts.resource)
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

	protoPutDocumentRequest := &proto.PutDocumentRequest{
		Index:  putDocumentCmdOpts.index,
		Id:     putDocumentCmdOpts.id,
		Fields: fieldsBytes,
	}

	if cmd.Flag("fields").Changed {
		protoPutDocumentRequest.Fields = []byte(putDocumentCmdOpts.fields)
	}

	conn, err := grpc.Dial(putCmdOpts.gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.PutDocument(context.Background(), protoPutDocumentRequest)
	if err != nil {
		return err
	}

	switch rootCmdOpts.outputFormat {
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
	putDocumentCmd.Flags().StringVar(&putDocumentCmdOpts.index, "index", DefaultIndex, "index name")
	putDocumentCmd.Flags().StringVar(&putDocumentCmdOpts.id, "id", DefaultId, "document id")
	putDocumentCmd.Flags().StringVar(&putDocumentCmdOpts.resource, "resource", DefaultResource, "resource file")
	putDocumentCmd.Flags().StringVar(&putDocumentCmdOpts.fields, "fields", DefaultDocFields, "document fields")

	putCmd.AddCommand(putDocumentCmd)
}
