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

	// create request
	var getDocumentRequest *util.GetDocumentRequest
	getDocumentRequest, err := util.NewGetDocumentRequest(getDocumentCmdOpts.id)
	if err != nil {
		return err
	}

	// create proto message
	req, err := getDocumentRequest.MarshalProto()
	if err != nil {
		return err
	}

	// create client
	icw, err := client.NewIndigoClientWrapper(getCmdOpts.gRPCServer)
	if err != nil {
		return err
	}
	defer icw.Conn.Close()

	// request
	resp, err := icw.Client.GetDocument(context.Background(), req)
	if err != nil {
		return err
	}

	// create response
	getDocumentResponse, err := util.NewGetDocumentRespone(resp)
	if err != nil {
		return err
	}

	// output response
	switch rootCmdOpts.outputFormat {
	case "text":
		fmt.Printf("%v\n", getDocumentResponse)
	case "json":
		output, err := json.MarshalIndent(getDocumentResponse, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", output)
	default:
		fmt.Printf("%v\n", getDocumentResponse)
	}

	return nil
}

func init() {
	getDocumentCmd.Flags().StringVar(&getDocumentCmdOpts.id, "id", DefaultId, "document id")

	getCmd.AddCommand(getDocumentCmd)
}
