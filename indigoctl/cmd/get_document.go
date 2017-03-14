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

var GetDocumentCmd = &cobra.Command{
	Use:   "document",
	Short: "gets the document from the Indigo gRPC Server",
	Long:  `The get document command gets the document from the Indigo gRPC Server.`,
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
		resp, err := client.GetDocument(context.Background(), &proto.GetDocumentRequest{IndexName: index, Id: docID})
		if err != nil {
			return err
		}

		fields := make(map[string]interface{})
		if err := json.Unmarshal(resp.Fields, &fields); err != nil {
			return err
		}

		r := struct {
			ID     string                 `json:"id"`
			Fields map[string]interface{} `json:"docFields"`
		}{
			ID:     docID,
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
	},
}

func init() {
	GetDocumentCmd.Flags().StringVarP(&index, "index", "i", constant.DefaultIndex, "index name")
	GetDocumentCmd.Flags().StringVarP(&docID, "doc-id", "d", constant.DefaultDocID, "document id")

	GetCmd.AddCommand(GetDocumentCmd)
}
