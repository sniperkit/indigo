package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type DeleteIndexCommandOptions struct {
	index string
}

var deleteIndexCmdOpts DeleteIndexCommandOptions

var deleteIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "deletes the index from the Indigo gRPC Server",
	Long:  `The delete index command deletes the index from the Indigo gRPC Server.`,
	RunE:  runEDeleteIndexCmd,
}

func runEDeleteIndexCmd(cmd *cobra.Command, args []string) error {
	if deleteIndexCmdOpts.index == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
	}

	protoDeleteIndexRequest := &proto.DeleteIndexRequest{
		Index: deleteIndexCmdOpts.index,
	}

	conn, err := grpc.Dial(deleteCmdOpts.gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.DeleteIndex(context.Background(), protoDeleteIndexRequest)
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
	deleteIndexCmd.Flags().StringVar(&deleteIndexCmdOpts.index, "index", DefaultIndex, "index name")

	deleteCmd.AddCommand(deleteIndexCmd)
}
