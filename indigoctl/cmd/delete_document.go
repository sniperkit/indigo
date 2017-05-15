package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var DeleteDocumentCmd = &cobra.Command{
	Use:   "document",
	Short: "deletes the document from the Indigo gRPC Server",
	Long:  `The delete document command deletes the document from the Indigo gRPC Server.`,
	RunE:  runEDeleteDocumentCmd,
}

func runEDeleteDocumentCmd(cmd *cobra.Command, args []string) error {
	index := cmd.Flag("index").Value.String()
	if index == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
	}

	id := cmd.Flag("id").Value.String()
	if id == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("id").Name)
	}

	deleteDocumentRequest := &proto.DeleteDocumentRequest{
		Index: index,
		Id:    id,
	}

	conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.DeleteDocument(context.Background(), deleteDocumentRequest)
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
	DeleteDocumentCmd.Flags().String("index", "", "index name")
	DeleteDocumentCmd.Flags().String("id", "", "document id")

	DeleteCmd.AddCommand(DeleteDocumentCmd)
}
