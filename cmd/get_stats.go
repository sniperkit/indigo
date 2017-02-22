package cmd

import (
	"fmt"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var getStatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "gets the index stats from the Indigo gRPC Server",
	Long:  `The get stats command gets the index stats from the Indigo gRPC Server.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if indexName == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index-name").Name)
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
		resp, err := client.GetStats(context.Background(), &proto.GetStatsRequest{IndexName: indexName})
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", resp.IndexStats)

		return nil
	},
}

func init() {
	getStatsCmd.Flags().StringVarP(&indexName, "index-name", "n", constant.DefaultIndexName, "index name")

	getCmd.AddCommand(getStatsCmd)
}
