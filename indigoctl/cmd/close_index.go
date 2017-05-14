package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var CloseIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "closes the index to the Indigo gRPC Server",
	Long:  `The close index command closes the index to the Indigo gRPC Server.`,
	RunE:  runECloseIndexCmd,
}

func runECloseIndexCmd(cmd *cobra.Command, args []string) error {
	closeIndexRequest := &proto.CloseIndexRequest{}

	if cmd.Flag("index").Changed {
		if cmd.Flag("index").Value.String() == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
		}
		closeIndexRequest.Index = cmd.Flag("index").Value.String()
	}

	conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.CloseIndex(context.Background(), closeIndexRequest)
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
	CloseIndexCmd.MarkFlagRequired("index")
	CloseIndexCmd.Flags().String("index", "", "index name")

	CloseCmd.AddCommand(CloseIndexCmd)
}
