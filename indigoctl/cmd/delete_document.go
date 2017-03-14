package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var DeleteDocumentCmd = &cobra.Command{
	Use:   "document",
	Short: "deletes the document from the Indigo gRPC Server",
	Long:  `The delete document command deletes the document from the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if index == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
		}

		if docID == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("doc-id").Name)
		}

		conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.DeleteDocument(context.Background(), &proto.DeleteDocumentRequest{IndexName: index, Id: docID})
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
	},
}

func init() {
	DeleteDocumentCmd.Flags().StringVarP(&index, "index", "i", constant.DefaultIndex, "index name")
	DeleteDocumentCmd.Flags().StringVarP(&docID, "doc-id", "d", constant.DefaultDocID, "document id")

	DeleteCmd.AddCommand(DeleteDocumentCmd)
}
