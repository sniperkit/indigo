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

var OpenIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "opens the index to the Indigo gRPC Server",
	Long:  `The open index command opens the index to the Indigo gRPC Server.`,
	RunE:  runEOpenIndexCmd,
}

func runEOpenIndexCmd(cmd *cobra.Command, args []string) error {
	if index == "" {
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

	conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.OpenIndex(context.Background(), &proto.OpenIndexRequest{Index: index, RuntimeConfig: rc})
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
}

func init() {
	OpenIndexCmd.Flags().StringVar(&index, "index", constant.DefaultIndex, "index name")
	OpenIndexCmd.Flags().StringVar(&runtimeConfig, "runtime-config", constant.DefaultRuntimeConfig, "runtime config")

	OpenCmd.AddCommand(OpenIndexCmd)
}
