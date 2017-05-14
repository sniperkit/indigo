package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/defaultvalue"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var DeleteIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "deletes the index from the Indigo gRPC Server",
	Long:  `The delete index command deletes the index from the Indigo gRPC Server.`,
	RunE:  runEDeleteIndexCmd,
}

func runEDeleteIndexCmd(cmd *cobra.Command, args []string) error {
	deleteIndexRequest := &proto.DeleteIndexRequest{}

	if cmd.Flag("index").Changed {
		if cmd.Flag("index").Value.String() == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
		}
		deleteIndexRequest.Index = cmd.Flag("index").Value.String()
	}

	conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.DeleteIndex(context.Background(), deleteIndexRequest)
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
	DeleteIndexCmd.Flags().String("index", defaultvalue.DefaultIndex, "index name")

	DeleteCmd.AddCommand(DeleteIndexCmd)
}
