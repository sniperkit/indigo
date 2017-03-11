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

var openIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "opens the index to the Indigo gRPC Server",
	Long:  `The open index command opens the index to the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if viper.GetString("index") == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
		}

		var rc []byte = nil

		if runtimeConfig != "" {
			file, err := os.Open(runtimeConfig)
			if err != nil {
				return err
			}
			defer file.Close()

			rc, err = ioutil.ReadAll(file)
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
		resp, err := client.OpenIndex(context.Background(), &proto.OpenIndexRequest{IndexName: viper.GetString("index"), RuntimeConfig: rc})
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
	openIndexCmd.Flags().StringP("index", "i", constant.DefaultIndex, "index name")
	viper.BindPFlag("index", openIndexCmd.Flags().Lookup("index"))

	openIndexCmd.Flags().StringVarP(&runtimeConfig, "runtime-config", "r", constant.DefaultRuntimeConfig, "runtime config")

	openCmd.AddCommand(openIndexCmd)
}
