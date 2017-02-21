package cmd

import (
	"fmt"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
)

var bulkCmd = &cobra.Command{
	Use:   "bulk",
	Short: "indexes the documents in bulk to the Indigo gRPC Server",
	Long:  `The bulk command indexes the documents in bulk to the Indigo gRPC Server.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if indexName == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index-name").Name)
		}

		if bulkRequestFile == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("bulk-request-file").Name)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		bulkRequest := make([]byte, 0)
		file, err := os.Open(bulkRequestFile)
		if err != nil {
			return err
		}
		defer file.Close()

		bulkRequest, err = ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.Bulk(context.Background(), &proto.BulkRequest{IndexName: indexName, BulkRequest: bulkRequest, BatchSize: batchSize})
		if err != nil {
			return err
		}

		fmt.Printf("%d documents put in bulk\n", resp.PutCount)
		fmt.Printf("%d error documents occurred in bulk\n", resp.PutErrorCount)
		fmt.Printf("%d documents deleted in bulk\n", resp.DeleteCount)

		return nil
	},
}

func init() {
	bulkCmd.Flags().StringVarP(&gRPCServer, "grpc-server", "g", constant.DefaultGRPCServer, "Indigo gRPC Sever")
	bulkCmd.Flags().StringVarP(&indexName, "index-name", "i", constant.DefaultIndexName, "index name")
	bulkCmd.Flags().StringVarP(&bulkRequestFile, "bulk-request-file", "b", constant.DefaultBulkRequestFile, "bulk request file")
	bulkCmd.Flags().Int32VarP(&batchSize, "batch-size", "B", constant.DefaultBatchSize, "batch size")

	RootCmd.AddCommand(bulkCmd)
}
