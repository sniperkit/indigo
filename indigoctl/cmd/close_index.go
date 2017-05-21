package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type CloseIndexCommandOptions struct {
	index string
}

var closeIndexCmdOpts CloseIndexCommandOptions

var closeIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "closes the index to the Indigo gRPC Server",
	Long:  `The close index command closes the index to the Indigo gRPC Server.`,
	RunE:  runECloseIndexCmd,
}

func runECloseIndexCmd(cmd *cobra.Command, args []string) error {
	if closeIndexCmdOpts.index == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
	}

	protoCloseIndexRequest := &proto.CloseIndexRequest{
		Index: closeIndexCmdOpts.index,
	}

	conn, err := grpc.Dial(closeCmdOpts.gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.CloseIndex(context.Background(), protoCloseIndexRequest)
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
	closeIndexCmd.Flags().StringVar(&closeIndexCmdOpts.index, "index", DefaultIndex, "index name")

	closeCmd.AddCommand(closeIndexCmd)
}
