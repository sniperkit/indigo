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
	"github.com/mosuka/indigo/util"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
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
	// create request
	var bulkRequest *util.BulkRequest
	var err error
	if cmd.Flag("resource").Changed {
		if bulkCmdOpts.resource == "-" {
			bulkRequest, err = util.NewBulkRequest(os.Stdin)
		} else {
			file, err := os.Open(bulkCmdOpts.resource)
			if err != nil {
				return err
			}
			defer file.Close()

			bulkRequest, err = util.NewBulkRequest(file)
			if err != nil {
				return err
			}
		}
	}

	// overwrite request
	if cmd.Flag("batch-size").Changed {
		bulkRequest.BatchSize = bulkCmdOpts.batchSize
	}

	// create proto message
	req, err := bulkRequest.MarshalProto()
	if err != nil {
		return err
	}

	// create client
	icw, err := client.NewIndigoClientWrapper(bulkCmdOpts.gRPCServer)
	if err != nil {
		return err
	}
	defer icw.Conn.Close()

	// request
	resp, err := icw.Client.Bulk(context.Background(), req)
	if err != nil {
		return err
	}

	// create response
	bulkResponse, err := util.NewBulkResponse(resp)
	if err != nil {
		return err
	}

	// output request
	switch rootCmdOpts.outputFormat {
	case "text":
		fmt.Printf("%v\n", bulkResponse)
	case "json":
		output, err := json.MarshalIndent(bulkResponse, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", output)
	default:
		fmt.Printf("%v\n", bulkResponse)
	}

	return nil
}

func init() {
	bulkCmd.Flags().StringVar(&bulkCmdOpts.gRPCServer, "grpc-server", DefaultServer, "Indigo gRPC Server to connect to")
	bulkCmd.Flags().Int32Var(&bulkCmdOpts.batchSize, "batch-size", DefaultBatchSize, "batch size of bulk request")
	bulkCmd.Flags().StringVar(&bulkCmdOpts.resource, "resource", DefaultResource, "resource file")

	RootCmd.AddCommand(bulkCmd)
}
