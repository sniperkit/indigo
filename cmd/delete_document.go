package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var deleteDocumentCmd = &cobra.Command{
	Use:   "document",
	Short: "deletes the document from the Indigo gRPC Server",
	Long:  `The delete document command deletes the document from the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if viper.GetString("index") == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
		}

		if docID == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("doc-id").Name)
		}

		conn, err := grpc.Dial(viper.GetString("grpc_server"), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.DeleteDocument(context.Background(), &proto.DeleteDocumentRequest{IndexName: viper.GetString("index"), Id: docID})
		if err != nil {
			return err
		}

		switch viper.GetString("output_format") {
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
	deleteDocumentCmd.Flags().StringP("index", "i", constant.DefaultIndex, "index name")
	viper.BindPFlag("index", deleteDocumentCmd.Flags().Lookup("index"))

	deleteDocumentCmd.Flags().StringVarP(&docID, "doc-id", "d", constant.DefaultDocID, "document id")

	deleteCmd.AddCommand(deleteDocumentCmd)
}
