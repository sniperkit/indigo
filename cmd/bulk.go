package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
)

var bulkCmd = &cobra.Command{
	Use:   "bulk",
	Short: "indexes the documents in bulk to the Indigo gRPC Server",
	Long:  `The bulk command indexes the documents in bulk to the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if viper.GetString("index") == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
		}

		if bulkRequest == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("bulk-request").Name)
		}

		br := make([]byte, 0)
		file, err := os.Open(bulkRequest)
		if err != nil {
			return err
		}
		defer file.Close()

		br, err = ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		conn, err := grpc.Dial(viper.GetString("grpc_server"), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.Bulk(context.Background(), &proto.BulkRequest{IndexName: viper.GetString("index"), BulkRequest: br, BatchSize: int32(viper.GetInt("batch_size"))})
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
	bulkCmd.Flags().StringP("grpc-server", "g", constant.DefaultGRPCServer, "Indigo gRPC Sever")
	viper.BindPFlag("grpc_server", bulkCmd.Flags().Lookup("grpc-server"))

	bulkCmd.Flags().Int32P("batch-size", "s", constant.DefaultBatchSize, "batch size")
	viper.BindPFlag("batch_size", bulkCmd.Flags().Lookup("batch-size"))

	bulkCmd.Flags().StringP("index", "i", constant.DefaultIndex, "index name")
	viper.BindPFlag("index", bulkCmd.Flags().Lookup("index"))

	bulkCmd.Flags().StringVarP(&bulkRequest, "bulk-request", "b", constant.DefaultBulkRequest, "bulk request")

	RootCmd.AddCommand(bulkCmd)
}
