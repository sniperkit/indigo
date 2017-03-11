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
	"io/ioutil"
	"os"
)

var putDocumentCmd = &cobra.Command{
	Use:   "document",
	Short: "puts the document to the Indigo gRPC Server",
	Long:  `The index document command puts the document to the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if viper.GetString("index") == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
		}

		if docID == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("doc-id").Name)
		}

		if docFields == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("doc-fields").Name)
		}

		document := make([]byte, 0)
		file, err := os.Open(docFields)
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
		resp, err := client.PutDocument(context.Background(), &proto.PutDocumentRequest{IndexName: viper.GetString("index"), Id: docID, Fields: document})
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
	putDocumentCmd.Flags().StringP("index", "i", constant.DefaultIndex, "index name")
	viper.BindPFlag("index", putDocumentCmd.Flags().Lookup("index"))

	putDocumentCmd.Flags().StringVarP(&docID, "doc-id", "d", constant.DefaultDocID, "document id")

	putDocumentCmd.Flags().StringVarP(&docFields, "doc-fields", "F", constant.DefaultDocFields, "document fields")

	putCmd.AddCommand(putDocumentCmd)
}
