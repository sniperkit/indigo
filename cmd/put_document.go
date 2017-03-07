package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/constant"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
	"github.com/spf13/viper"
)

var putDocumentCmd = &cobra.Command{
	Use:   "document",
	Short: "puts the document to the Indigo gRPC Server",
	Long:  `The index document command puts the document to the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if indexName == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index-name").Name)
		}

		if documentID == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("id").Name)
		}

		if documentFile == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("fields").Name)
		}

		document := make([]byte, 0)
		file, err := os.Open(documentFile)
		if err != nil {
			return err
		}
		defer file.Close()

		document, err = ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		conn, err := grpc.Dial(viper.GetString("grpc_server"), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.PutDocument(context.Background(), &proto.PutDocumentRequest{IndexName: indexName, Id: documentID, Fields: document})
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
	putDocumentCmd.Flags().StringVarP(&indexName, "index-name", "n", constant.DefaultIndexName, "index name")
	putDocumentCmd.Flags().StringVarP(&documentID, "id", "i", constant.DefaultDocumentID, "document id")
	putDocumentCmd.Flags().StringVarP(&documentFile, "fields", "F", constant.DefaultDocumentFile, "fields file")

	putCmd.AddCommand(putDocumentCmd)
}
