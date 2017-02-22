package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var closeIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "closes the index to the Indigo gRPC Server",
	Long:  `The close index command closes the index to the Indigo gRPC Server.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if indexName == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("name").Name)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.CloseIndex(context.Background(), &proto.CloseIndexRequest{IndexName: indexName})
		if err != nil {
			return err
		}

		switch outputFormat {
		case "text":
			fmt.Printf("IndexName: %s\n", resp.IndexName)
		case "json":
			output, err := json.Marshal(resp)
			if err != nil {
				return err
			}

			fmt.Printf("%s\n", output)
		default:
			fmt.Printf("IndexName: %s\n", resp.IndexName)
		}

		return nil
	},
}

func init() {
	closeIndexCmd.Flags().StringVarP(&indexName, "name", "n", constant.DefaultIndexName, "index name")

	closeCmd.AddCommand(closeIndexCmd)
}
