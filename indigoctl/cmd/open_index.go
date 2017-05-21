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

type OpenIndexCommandOptions struct {
	index         string
	resource      string
	runtimeConfig string
}

var openIndexCmdOpts OpenIndexCommandOptions

var openIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "opens the index to the Indigo gRPC Server",
	Long:  `The open index command opens the index to the Indigo gRPC Server.`,
	RunE:  runEOpenIndexCmd,
}

type OpenIndexResource struct {
	RuntimeConfig map[string]interface{} `json:"runtime_config,omitempty"`
}

func runEOpenIndexCmd(cmd *cobra.Command, args []string) error {
	if openIndexCmdOpts.index == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
	}

	var resourceBytes []byte = nil
	if cmd.Flag("resource").Changed {
		if openIndexCmdOpts.resource == "-" {
			resourceBytes, _ = ioutil.ReadAll(os.Stdin)
		} else {
			file, err := os.Open(openIndexCmdOpts.resource)
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

	protoOpenIndexRequest := &proto.OpenIndexRequest{
		Index:         openIndexCmdOpts.index,
		RuntimeConfig: runtimeConfigBytes,
	}

	if cmd.Flag("index").Changed {
		protoOpenIndexRequest.Index = openIndexCmdOpts.index
	}

	if cmd.Flag("runtime-config").Changed {
		protoOpenIndexRequest.RuntimeConfig = []byte(openIndexCmdOpts.runtimeConfig)
	}

	conn, err := grpc.Dial(openCmdOpts.gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.OpenIndex(context.Background(), protoOpenIndexRequest)
	if err != nil {
		return err
	}

	switch rootCmdOpts.outputFormat {
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
	openIndexCmd.Flags().StringVar(&openIndexCmdOpts.index, "index", DefaultIndex, "index name")
	openIndexCmd.Flags().StringVar(&openIndexCmdOpts.resource, "resource", DefaultResource, "resource file")
	openIndexCmd.Flags().StringVar(&openIndexCmdOpts.runtimeConfig, "runtime-config", DefaultRuntimeConfig, "runtime config")

	openCmd.AddCommand(openIndexCmd)
}
