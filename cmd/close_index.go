package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/setting"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var closeIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "closes the index to the Indigo gRPC Server",
	Long:  `The close index command closes the index to the Indigo gRPC Server.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if indexName == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index-name").Name)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(viper.GetString("grpc_server"), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.CloseIndex(context.Background(), &proto.CloseIndexRequest{IndexName: indexName})
		if err != nil {
			return err
		}

		switch viper.GetString("output_format") {
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
	},
}

func init() {
	closeIndexCmd.Flags().StringVarP(&indexName, "index-name", "n", setting.DefaultIndexName, "index name")

	closeCmd.AddCommand(closeIndexCmd)
}
