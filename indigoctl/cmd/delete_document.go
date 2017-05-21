package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type DeleteDocumentCommandOptions struct {
	index string
	id    string
}

var deleteDocumentCmdOpts DeleteDocumentCommandOptions

var deleteDocumentCmd = &cobra.Command{
	Use:   "document",
	Short: "deletes the document from the Indigo gRPC Server",
	Long:  `The delete document command deletes the document from the Indigo gRPC Server.`,
	RunE:  runEDeleteDocumentCmd,
}

func runEDeleteDocumentCmd(cmd *cobra.Command, args []string) error {
	if deleteDocumentCmdOpts.index == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
	}

	if deleteDocumentCmdOpts.id == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("id").Name)
	}

	protoDeleteDocumentRequest := &proto.DeleteDocumentRequest{
		Index: deleteDocumentCmdOpts.index,
		Id:    deleteDocumentCmdOpts.id,
	}

	conn, err := grpc.Dial(deleteCmdOpts.gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.DeleteDocument(context.Background(), protoDeleteDocumentRequest)
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
	deleteDocumentCmd.Flags().StringVar(&deleteDocumentCmdOpts.index, "index", DefaultIndex, "index name")
	deleteDocumentCmd.Flags().StringVar(&deleteDocumentCmdOpts.id, "id", DefaultId, "document id")

	deleteCmd.AddCommand(deleteDocumentCmd)
}
