package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var ListIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "closes the index to the Indigo gRPC Server",
	Long:  `The close index command closes the index to the Indigo gRPC Server.`,
	RunE:  runEListIndexCmd,
}

func runEListIndexCmd(cmd *cobra.Command, args []string) error {
	listIndexRequest := &proto.ListIndexRequest{}

	conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.ListIndex(context.Background(), listIndexRequest)
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
	ListCmd.AddCommand(ListIndexCmd)
}
