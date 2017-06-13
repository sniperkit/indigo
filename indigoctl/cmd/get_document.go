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
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/resource"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type GetDocumentCommandOptions struct {
	id string
}

var getDocumentCmdOpts GetDocumentCommandOptions

var getDocumentCmd = &cobra.Command{
	Use:   "document",
	Short: "gets the document from the Indigo gRPC Server",
	Long:  `The get document command gets the document from the Indigo gRPC Server.`,
	RunE:  runEGetDocumentCmd,
}

func runEGetDocumentCmd(cmd *cobra.Command, args []string) error {
	if getDocumentCmdOpts.id == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("id").Name)
	}

	protoGetDocumentRequest := &proto.GetDocumentRequest{
		Id: getDocumentCmdOpts.id,
	}

	conn, err := grpc.Dial(getCmdOpts.gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.GetDocument(context.Background(), protoGetDocumentRequest)
	if err != nil {
		return err
	}

	fields, err := proto.UnmarshalAny(resp.Fields)
	if err != nil {
		return err
	}

	r := resource.GetDocumentResponse{
		Id:     resp.Id,
		Fields: fields.(*map[string]interface{}),
	}

	switch rootCmdOpts.outputFormat {
	case "text":
		fmt.Printf("%s\n", r)
	case "json":
		output, err := json.MarshalIndent(r, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", output)
	default:
		fmt.Printf("%s\n", r)
	}

	return nil
}

func init() {
	getDocumentCmd.Flags().StringVar(&getDocumentCmdOpts.id, "id", DefaultId, "document id")

	getCmd.AddCommand(getDocumentCmd)
}
