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

type PutDocumentCommandOptions struct {
	id       string
	fields   string
	resource string
}

var putDocumentCmdOpts PutDocumentCommandOptions

var putDocumentCmd = &cobra.Command{
	Use:   "document",
	Short: "puts the document to the Indigo gRPC Server",
	Long:  `The index document command puts the document to the Indigo gRPC Server.`,
	RunE:  runEPutDocumentCmd,
}

func runEPutDocumentCmd(cmd *cobra.Command, args []string) error {
	putDocumentResource := util.PutDocumentResource{}
	if cmd.Flag("resource").Changed {
		var resourceBytes []byte = nil
		if putDocumentCmdOpts.resource == "-" {
			resourceBytes, _ = ioutil.ReadAll(os.Stdin)
		} else {
			file, err := os.Open(putDocumentCmdOpts.resource)
			if err != nil {
				return err
			}
			defer file.Close()

			resourceBytes, err = ioutil.ReadAll(file)
			if err != nil {
				return err
			}
		}
		err := json.Unmarshal(resourceBytes, &putDocumentResource)
		if err != nil {
			return err
		}
	}

	fieldsAny, err := util.MarshalAny(putDocumentResource.Fields)
	if err != nil {
		return err
	}

	protoPutDocumentRequest := &proto.PutDocumentRequest{
		Id:     putDocumentResource.Id,
		Fields: &fieldsAny,
	}

	if cmd.Flag("id").Changed {
		protoPutDocumentRequest.Id = putDocumentCmdOpts.id
	}

	if cmd.Flag("fields").Changed {
		var fields map[string]interface{}
		err := json.Unmarshal([]byte(putDocumentCmdOpts.fields), &fields)
		if err != nil {
			return err
		}
		fieldsAny, err := util.MarshalAny(fields)
		if err != nil {
			return err
		}
		protoPutDocumentRequest.Fields = &fieldsAny
	}

	icw, err := client.NewIndigoClientWrapper(putCmdOpts.gRPCServer)
	if err != nil {
		return err
	}
	defer icw.Conn.Close()

	resp, err := icw.Client.PutDocument(context.Background(), protoPutDocumentRequest)
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
	putDocumentCmd.Flags().StringVar(&putDocumentCmdOpts.id, "id", DefaultId, "document id")
	putDocumentCmd.Flags().StringVar(&putDocumentCmdOpts.resource, "resource", DefaultResource, "resource file")
	putDocumentCmd.Flags().StringVar(&putDocumentCmdOpts.fields, "fields", DefaultDocFields, "document fields")

	putCmd.AddCommand(putDocumentCmd)
}
