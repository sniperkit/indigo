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

var createIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "creates the index to the Indigo gRPC Server",
	Long:  `The create index command creates the index to the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if viper.GetString("index") == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
		}

		var im []byte = nil
		var kvc []byte = nil

		if indexMapping != "" {
			file, err := os.Open(indexMapping)
			if err != nil {
				return err
			}
			defer file.Close()

			im, err = ioutil.ReadAll(file)
			if err != nil {
				return err
			}
		}

		if kvConfig != "" {
			file, err := os.Open(kvConfig)
			if err != nil {
				return err
			}
			defer file.Close()

			kvc, err = ioutil.ReadAll(file)
			if err != nil {
				return err
			}
		}

		conn, err := grpc.Dial(viper.GetString("grpc_server"), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.CreateIndex(context.Background(), &proto.CreateIndexRequest{IndexName: viper.GetString("index"), IndexMapping: im, IndexType: indexType, Kvstore: kvStore, Kvconfig: kvc})
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
	createIndexCmd.Flags().StringP("index", "i", constant.DefaultIndex, "index name")
	viper.BindPFlag("index", createIndexCmd.Flags().Lookup("index"))

	createIndexCmd.Flags().StringVarP(&indexMapping, "index-mapping", "m", constant.DefaultIndexMapping, "index mapping")

	createIndexCmd.Flags().StringVarP(&indexType, "index-type", "t", constant.DefaultIndexType, "index type")

	createIndexCmd.Flags().StringVarP(&kvStore, "kv-store", "s", constant.DefaultKVStore, "kvstore")

	createIndexCmd.Flags().StringVarP(&kvConfig, "kv-config", "k", constant.DefaultKVConfigFile, "kvconfig")

	createCmd.AddCommand(createIndexCmd)
}
