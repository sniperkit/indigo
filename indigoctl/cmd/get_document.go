package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var GetDocumentCmd = &cobra.Command{
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
	index := cmd.Flag("index").Value.String()
	if index == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
	}

	id := cmd.Flag("id").Value.String()
	if id == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("id").Name)
	}

	getDocumentRequest := &proto.GetDocumentRequest{
		Index: index,
		Id:    id,
	}

	conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.GetDocument(context.Background(), getDocumentRequest)
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

	switch outputFormat {
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
	GetDocumentCmd.Flags().String("index", "", "index name")
	GetDocumentCmd.Flags().String("id", "", "document id")

	GetCmd.AddCommand(GetDocumentCmd)
}
