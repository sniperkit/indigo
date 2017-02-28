package cmd

import (
	"encoding/json"
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
			return fmt.Errorf("required flag: --%s", cmd.Flag("bulk-request").Name)
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

		r := struct {
			PutCount      int32 `json:"put_count"`
			PutErrorCount int32 `json:"put_error_count"`
			DeleteCount   int32 `json:"delete_count"`
		}{
			PutCount:      resp.PutCount,
			PutErrorCount: resp.PutErrorCount,
			DeleteCount:   resp.DeleteCount,
		}

		switch outputFormat {
		case "text":
			fmt.Printf("%s\n", resp.String())
		case "json":
			output, err := json.MarshalIndent(r, "", "  ")
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
	bulkCmd.Flags().StringVarP(&gRPCServer, "grpc-server", "g", constant.DefaultGRPCServer, "Indigo gRPC Sever")
	bulkCmd.Flags().StringVarP(&indexName, "index-name", "n", constant.DefaultIndexName, "index name")
	bulkCmd.Flags().StringVarP(&bulkRequestFile, "bulk-request", "b", constant.DefaultBulkRequestFile, "bulk request file")
	bulkCmd.Flags().Int32VarP(&batchSize, "batch-size", "s", constant.DefaultBatchSize, "batch size")

	RootCmd.AddCommand(bulkCmd)
}
