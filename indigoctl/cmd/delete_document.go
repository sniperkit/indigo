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
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

type DeleteDocumentCommandOptions struct {
	id string
}

var deleteDocumentCmdOpts DeleteDocumentCommandOptions

var deleteDocumentCmd = &cobra.Command{
	Use:   "document",
	Short: "deletes the document from the Indigo gRPC Server",
	Long:  `The delete document command deletes the document from the Indigo gRPC Server.`,
	RunE:  runEDeleteDocumentCmd,
}

func runEDeleteDocumentCmd(cmd *cobra.Command, args []string) error {
	if deleteDocumentCmdOpts.id == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("id").Name)
	}

	protoDeleteDocumentRequest := &proto.DeleteDocumentRequest{
		Id: deleteDocumentCmdOpts.id,
	}

	icw, err := client.NewIndigoClientWrapper(deleteCmdOpts.gRPCServer)
	if err != nil {
		return err
	}
	defer icw.Conn.Close()

	resp, err := icw.Client.DeleteDocument(context.Background(), protoDeleteDocumentRequest)
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
	deleteDocumentCmd.Flags().StringVar(&deleteDocumentCmdOpts.id, "id", DefaultId, "document id")

	deleteCmd.AddCommand(deleteDocumentCmd)
}
