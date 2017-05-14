package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
)

var OpenIndexCmd = &cobra.Command{
	Use:   "index RESOURCE_FILENAME",
	Short: "opens the index to the Indigo gRPC Server",
	Long:  `The open index command opens the index to the Indigo gRPC Server.`,
	RunE:  runEOpenIndexCmd,
}

type OpenIndexResource struct {
	Index         string                 `json:"index,omitempty"`
	RuntimeConfig map[string]interface{} `json:"runtime_config,omitempty"`
}

func runEOpenIndexCmd(cmd *cobra.Command, args []string) error {
	var resourceBytes []byte = nil
	if cmd.Flag("resource").Changed {
		if cmd.Flag("resource").Value.String() == "-" {
			resourceBytes, _ = ioutil.ReadAll(os.Stdin)
		} else {
			file, err := os.Open(cmd.Flag("resource").Value.String())
			if err != nil {
				return err
			}
			defer file.Close()

			resourceBytes, err = ioutil.ReadAll(file)
			if err != nil {
				return err
			}
		}
	}

	openIndexResource := OpenIndexResource{}
	err := json.Unmarshal(resourceBytes, &openIndexResource)
	if err != nil {
		return err
	}

	runtimeConfigBytes, err := json.Marshal(openIndexResource.RuntimeConfig)
	if err != nil {
		return err
	}

	openIndexRequest := &proto.OpenIndexRequest{
		Index:         openIndexResource.Index,
		RuntimeConfig: runtimeConfigBytes,
	}

	if cmd.Flag("index").Changed {
		openIndexRequest.Index = cmd.Flag("index").Value.String()
	}

	if cmd.Flag("runtime-config").Changed {
		runtimeConfigBytes := []byte(cmd.Flag("runtime-config").Value.String())
		openIndexRequest.RuntimeConfig = runtimeConfigBytes
	}

	conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.OpenIndex(context.Background(), openIndexRequest)
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
	OpenIndexCmd.Flags().String("resource", "", "resource file")
	OpenIndexCmd.Flags().String("index", "", "index name")
	OpenIndexCmd.Flags().String("runtime-config", "", "runtime config")

	OpenCmd.AddCommand(OpenIndexCmd)
}
