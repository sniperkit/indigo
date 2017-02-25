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

var openIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "opens the index to the Indigo gRPC Server",
	Long:  `The open index command opens the index to the Indigo gRPC Server.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if indexName == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("name").Name)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var runtimeConfig []byte = nil

		if runtimeConfigFile != "" {
			file, err := os.Open(runtimeConfigFile)
			if err != nil {
				return err
			}
			defer file.Close()

			runtimeConfig, err = ioutil.ReadAll(file)
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
		resp, err := client.OpenIndex(context.Background(), &proto.OpenIndexRequest{IndexName: indexName, RuntimeConfig: runtimeConfig})
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
	openIndexCmd.Flags().StringVarP(&indexName, "name", "n", constant.DefaultIndexName, "index name")
	openIndexCmd.Flags().StringVarP(&runtimeConfigFile, "runtime-config", "r", constant.DefaultRuntimeConfigFile, "runtime config file")

	openCmd.AddCommand(openIndexCmd)
}
