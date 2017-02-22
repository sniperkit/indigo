package cmd

import (
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

		conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.CreateIndex(context.Background(), &proto.CreateIndexRequest{IndexName: indexName, IndexMapping: indexMapping, IndexType: indexType, IndexStore: indexStore, KvConfig: kvConfig})
		if err != nil {
			return err
		}

		fmt.Printf("%s created\n", resp.IndexName)

		return nil
	},
}

func init() {
	createIndexCmd.Flags().StringVarP(&indexName, "index-name", "i", constant.DefaultIndexName, "index name")
	createIndexCmd.Flags().StringVarP(&indexMappingFile, "index-mapping", "m", constant.DefaultIndexMappingFile, "index mapping file")
	createIndexCmd.Flags().StringVarP(&indexType, "index-type", "t", constant.DefaultIndexType, "index type")
	createIndexCmd.Flags().StringVarP(&indexStore, "index-store", "s", constant.DefaultIndexStore, "index store")
	createIndexCmd.Flags().StringVarP(&kvConfigFile, "kv-config-file", "k", constant.DefaultKVConfigFile, "KV config file")

	createCmd.AddCommand(createIndexCmd)
}
