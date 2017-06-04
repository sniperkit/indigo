package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
)

type BulkCommandOptions struct {
	gRPCServer string
	batchSize  int32
	resource   string
}

var bulkCmdOpts BulkCommandOptions

type BulkRequest struct {
	Method string      `json:"method,omitempty"`
	Id     string      `json:"id,omitempty"`
	Fields interface{} `json:"fields,omitempty"`
}

type BulkResource struct {
	BatchSize    int32         `json:"batch_size,omitempty"`
	BulkRequests []BulkRequest `json:"bulk_requests,omitempty"`
}

var bulkCmd = &cobra.Command{
	Use:   "bulk",
	Short: "indexes the documents in bulk to the Indigo gRPC Server",
	Long:  `The bulk command indexes the documents in bulk to the Indigo gRPC Server.`,
	RunE:  runEBulkCmd,
}

func runEBulkCmd(cmd *cobra.Command, args []string) error {
	var resourceBytes []byte = nil
	if cmd.Flag("resource").Changed {
		if bulkCmdOpts.resource == "-" {
			resourceBytes, _ = ioutil.ReadAll(os.Stdin)
		} else {
			file, err := os.Open(bulkCmdOpts.resource)
			if err != nil {
				return err
			}
			defer file.Close()

			resourceBytes, err = ioutil.ReadAll(file)
			if err != nil {
				return err
			}
		}
	}

	bulkResource := BulkResource{}
	err := json.Unmarshal(resourceBytes, &bulkResource)
	if err != nil {
		return err
	}

	bulkRequestsBytes, err := json.Marshal(bulkResource.BulkRequests)
	if err != nil {
		return err
	}

	protoBulkRequest := &proto.BulkRequest{
		BatchSize:    bulkResource.BatchSize,
		BulkRequests: bulkRequestsBytes,
	}

	if cmd.Flag("batch-size").Changed {
		protoBulkRequest.BatchSize = bulkCmdOpts.batchSize
	}

	conn, err := grpc.Dial(bulkCmdOpts.gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.Bulk(context.Background(), protoBulkRequest)
	if err != nil {
		return err
	}

	switch rootCmdOpts.outputFormat {
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
	bulkCmd.Flags().StringVar(&bulkCmdOpts.gRPCServer, "grpc-server", DefaultServer, "Indigo gRPC Server to connect to")
	bulkCmd.Flags().Int32Var(&bulkCmdOpts.batchSize, "batch-size", DefaultBatchSize, "batch size of bulk request")
	bulkCmd.Flags().StringVar(&bulkCmdOpts.resource, "resource", DefaultResource, "resource file")

	RootCmd.AddCommand(bulkCmd)
}
