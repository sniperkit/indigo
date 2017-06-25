//  Copyright (c) 2017 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/client"
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/util"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"io/ioutil"
	"os"
)

type BulkCommandOptions struct {
	gRPCServer string
	batchSize  int32
	resource   string
}

var bulkCmdOpts BulkCommandOptions

var bulkCmd = &cobra.Command{
	Use:   "bulk",
	Short: "indexes the documents in bulk to the Indigo gRPC Server",
	Long:  `The bulk command indexes the documents in bulk to the Indigo gRPC Server.`,
	RunE:  runEBulkCmd,
}

func runEBulkCmd(cmd *cobra.Command, args []string) error {
	bulkResource := util.BulkResource{}
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
		err := json.Unmarshal(resourceBytes, &bulkResource)
		if err != nil {
			return err
		}
	}

	var b []*proto.BulkRequest_Request
	for _, request := range bulkResource.Requests {
		f, err := util.MarshalAny(request.Document.Fields)
		if err != nil {
			return nil
		}
		d := proto.BulkRequest_Document{
			Id:     request.Document.Id,
			Fields: &f,
		}
		r := proto.BulkRequest_Request{
			Method:   request.Method,
			Document: &d,
		}
		b = append(b, &r)
	}

	protoBulkRequest := &proto.BulkRequest{
		BatchSize: bulkResource.BatchSize,
		Requests:  b,
	}

	if cmd.Flag("batch-size").Changed {
		protoBulkRequest.BatchSize = bulkCmdOpts.batchSize
	}

	client, err := client.NewIndigoGRPCClient(bulkCmdOpts.gRPCServer)
	if err != nil {
		return err
	}
	defer client.Close()

	resp, err := client.Client.Bulk(context.Background(), protoBulkRequest)
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
