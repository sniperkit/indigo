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
	"github.com/buger/jsonparser"
	"github.com/mosuka/indigo/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"io/ioutil"
	"os"
)

type BulkCommandOptions struct {
	gRPCServer string
	batchSize  int32
	request    string
}

var bulkCmdOpts BulkCommandOptions

var bulkCmd = &cobra.Command{
	Use:   "bulk",
	Short: "indexes the documents in bulk to the Indigo gRPC Server",
	Long:  `The bulk command indexes the documents in bulk to the Indigo gRPC Server.`,
	RunE:  runEBulkCmd,
}

func runEBulkCmd(cmd *cobra.Command, args []string) error {
	// read request
	var data []byte
	var err error
	if cmd.Flag("request").Changed {
		if bulkCmdOpts.request == "-" {
			data, err = ioutil.ReadAll(os.Stdin)
		} else {
			file, err := os.Open(bulkCmdOpts.request)
			if err != nil {
				return err
			}
			defer file.Close()
			data, err = ioutil.ReadAll(file)
			if err != nil {
				return err
			}
		}
	}

	// get batch_size
	batchSize, err := jsonparser.GetInt(data, "batch_size")
	if err != nil {
		return err
	}

	// get requests
	requestsBytes, _, _, err := jsonparser.Get(data, "requests")
	if err != nil {
		return err
	}
	var requests []map[string]interface{}
	err = json.Unmarshal(requestsBytes, &requests)
	if err != nil {
		return err
	}

	// overwrite batch size
	if cmd.Flag("batch-size").Changed {
		batchSize = int64(bulkCmdOpts.batchSize)
	}

	// create client
	icw, err := client.NewIndigoClientWrapper(bulkCmdOpts.gRPCServer)
	if err != nil {
		return err
	}
	defer icw.Conn.Close()

	// request
	resp, err := icw.Bulk(context.Background(), requests, int32(batchSize))
	if err != nil {
		return err
	}

	// output request
	switch rootCmdOpts.outputFormat {
	case "text":
		fmt.Printf("%v\n", resp)
	case "json":
		output, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", output)
	default:
		fmt.Printf("%v\n", resp)
	}

	return nil
}

func init() {
	bulkCmd.Flags().StringVar(&bulkCmdOpts.gRPCServer, "grpc-server", DefaultServer, "Indigo gRPC Server to connect to")
	bulkCmd.Flags().Int32Var(&bulkCmdOpts.batchSize, "batch-size", DefaultBatchSize, "batch size of bulk request")
	bulkCmd.Flags().StringVar(&bulkCmdOpts.request, "request", DefaultResource, "request file")

	RootCmd.AddCommand(bulkCmd)
}
