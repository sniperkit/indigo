package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type GetDocumentCommandOptions struct {
	id string
}

var getDocumentCmdOpts GetDocumentCommandOptions

var getDocumentCmd = &cobra.Command{
	Use:   "document",
	Short: "gets the document from the Indigo gRPC Server",
	Long:  `The get document command gets the document from the Indigo gRPC Server.`,
	RunE:  runEGetDocumentCmd,
}

type GetDocumentResponse struct {
	Id     string      `json:"id"`
	Fields interface{} `json:"fields"`
}

func runEGetDocumentCmd(cmd *cobra.Command, args []string) error {
	if getDocumentCmdOpts.id == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("id").Name)
	}

	protoGetDocumentRequest := &proto.GetDocumentRequest{
		Id: getDocumentCmdOpts.id,
	}

	conn, err := grpc.Dial(getCmdOpts.gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.GetDocument(context.Background(), protoGetDocumentRequest)
	if err != nil {
		return err
	}

	var fields interface{} = nil
	if err := json.Unmarshal(resp.Fields, &fields); err != nil {
		return err
	}

	r := GetDocumentResponse{
		Id:     resp.Id,
		Fields: fields,
	}

	switch rootCmdOpts.outputFormat {
	case "text":
		fmt.Printf("%s\n", r)
	case "json":
		output, err := json.MarshalIndent(r, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", output)
	default:
		fmt.Printf("%s\n", r)
	}

	return nil
}

func init() {
	getDocumentCmd.Flags().StringVar(&getDocumentCmdOpts.id, "id", DefaultId, "document id")

	getCmd.AddCommand(getDocumentCmd)
}
