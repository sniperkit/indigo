package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/setting"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var getDocumentCmd = &cobra.Command{
	Use:   "document",
	Short: "gets the document from the Indigo gRPC Server",
	Long:  `The get document command gets the document from the Indigo gRPC Server.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if indexName == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index-name").Name)
		}

		if documentID == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("id").Name)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(viper.GetString("grpc_server"), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.GetDocument(context.Background(), &proto.GetDocumentRequest{IndexName: indexName, Id: documentID})
		if err != nil {
			return err
		}

		fields := make(map[string]interface{})
		if err := json.Unmarshal(resp.Fields, &fields); err != nil {
			return err
		}

		r := struct {
			ID     string                 `json:"id"`
			Fields map[string]interface{} `json:"fields"`
		}{
			ID:     documentID,
			Fields: fields,
		}

		switch viper.GetString("output_format") {
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
	getDocumentCmd.Flags().StringVarP(&indexName, "index-name", "n", setting.DefaultIndexName, "index name")
	getDocumentCmd.Flags().StringVarP(&documentID, "id", "i", setting.DefaultDocumentID, "document id")

	getCmd.AddCommand(getDocumentCmd)
}
