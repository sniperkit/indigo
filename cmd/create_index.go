package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
)

var createIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "creates the index to the Indigo gRPC Server",
	Long:  `The create index command creates the index to the Indigo gRPC Server.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if indexName == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index-name").Name)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var indexMapping []byte = nil
		var kvConfig []byte = nil

		if indexMappingFile != "" {
			file, err := os.Open(indexMappingFile)
			if err != nil {
				return err
			}
			defer file.Close()

			indexMapping, err = ioutil.ReadAll(file)
			if err != nil {
				return err
			}
		}

		if kvConfigFile != "" {
			file, err := os.Open(kvConfigFile)
			if err != nil {
				return err
			}
			defer file.Close()

			kvConfig, err = ioutil.ReadAll(file)
			if err != nil {
				return err
			}
		}

		conn, err := grpc.Dial(indigoSettings.GetString("grpc_server"), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.CreateIndex(context.Background(), &proto.CreateIndexRequest{IndexName: indexName, IndexMapping: indexMapping, IndexType: indexType, Kvstore: kvStore, Kvconfig: kvConfig})
		if err != nil {
			return err
		}

		switch indigoSettings.GetString("output_format") {
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
	createIndexCmd.Flags().StringVarP(&indexName, "index-name", "n", constant.DefaultIndexName, "index name")
	createIndexCmd.Flags().StringVarP(&indexMappingFile, "index-mapping", "m", constant.DefaultIndexMappingFile, "index mapping file")
	createIndexCmd.Flags().StringVarP(&indexType, "index-type", "t", constant.DefaultIndexType, "index type")
	createIndexCmd.Flags().StringVarP(&kvStore, "kvstore", "s", constant.DefaultKVStore, "kvstore")
	createIndexCmd.Flags().StringVarP(&kvConfigFile, "kvconfig", "c", constant.DefaultKVConfigFile, "kvconfig file")

	createCmd.AddCommand(createIndexCmd)
}
