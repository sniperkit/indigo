package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/defaultvalue"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
)

var BulkCmd = &cobra.Command{
	Use:   "bulk",
	Short: "indexes the documents in bulk to the Indigo gRPC Server",
	Long:  `The bulk command indexes the documents in bulk to the Indigo gRPC Server.`,
	RunE:  runEBulkCmd,
}

type BulkRequest struct {
	Method string      `json:"method,omitempty"`
	Id     string      `json:"id,omitempty"`
	Fields interface{} `json:"fields,omitempty"`
}

type BulkResource struct {
	BatchSize    int32         `json:"batch_size,omitempty"`
	BulkRequests []BulkRequest `json:"bulk_requests,omitempty"`
}

func runEBulkCmd(cmd *cobra.Command, args []string) error {
	index := cmd.Flag("index").Value.String()
	if index == "" {
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

	conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.Bulk(context.Background(), &proto.BulkRequest{Index: index, BulkRequests: br, BatchSize: batchSize})
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
	BulkCmd.Flags().StringVar(&gRPCServer, "grpc-server", defaultvalue.DefaultGRPCServer, "Indigo gRPC Server to connect to")
	BulkCmd.Flags().String("index", "", "index name")
	BulkCmd.Flags().StringVar(&bulkRequest, "bulk-request", defaultvalue.DefaultBulkRequest, "bulk request")
	BulkCmd.Flags().Int32Var(&batchSize, "batch-size", defaultvalue.DefaultBatchSize, "batch size of bulk request")

	RootCmd.AddCommand(BulkCmd)
}
