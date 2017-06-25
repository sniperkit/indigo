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
	"github.com/blevesearch/bleve/mapping"
	"github.com/mosuka/indigo/client"
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/util"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

type GetIndexCommandOptions struct {
	includeIndexMapping bool
	includeIndexType    bool
	includeKvstore      bool
	includeKvconfig     bool
}

var getIndexCmdOpts GetIndexCommandOptions

var getIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "gets the index mapping from the Indigo Server",
	Long:  `The get index command gets the index information from the Indigo Server.`,
	RunE:  runEGetIndexCmd,
}

func runEGetIndexCmd(cmd *cobra.Command, args []string) error {
	protoGetIndexRequest := &proto.GetIndexRequest{
		IncludeIndexMapping: getIndexCmdOpts.includeIndexMapping,
		IncludeIndexType:    getIndexCmdOpts.includeIndexType,
		IncludeKvstore:      getIndexCmdOpts.includeKvstore,
		IncludeKvconfig:     getIndexCmdOpts.includeKvconfig,
	}

	icw, err := client.NewIndigoClientWrapper(getCmdOpts.gRPCServer)
	if err != nil {
		return err
	}
	defer icw.Conn.Close()

	resp, err := icw.Client.GetIndex(context.Background(), protoGetIndexRequest)
	if err != nil {
		return err
	}

	r := util.GetIndexResponse{
		Path:      resp.Path,
		IndexType: resp.IndexType,
		Kvstore:   resp.Kvstore,
	}

	if resp.IndexMapping != nil {
		indexMapping, err := util.UnmarshalAny(resp.IndexMapping)
		if err != nil {
			return err
		}
		r.IndexMapping = indexMapping.(*mapping.IndexMappingImpl)
	}

	if resp.Kvconfig != nil {
		kvconfig, err := util.UnmarshalAny(resp.Kvconfig)
		if err != nil {
			return err
		}
		r.Kvconfig = kvconfig
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
		fmt.Printf("%s\n", resp.String())
	}

	return nil
}

func init() {
	getIndexCmd.Flags().BoolVar(&getIndexCmdOpts.includeIndexMapping, "include-index-mapping", DefaultIncludeIndexMapping, "include index mapping")
	getIndexCmd.Flags().BoolVar(&getIndexCmdOpts.includeIndexType, "include-index-type", DefaultIncludeIndexType, "include index type")
	getIndexCmd.Flags().BoolVar(&getIndexCmdOpts.includeKvstore, "include-kvstore", DefaultIncludeKvstore, "include kvstore")
	getIndexCmd.Flags().BoolVar(&getIndexCmdOpts.includeKvconfig, "include-kvconfig", DefaultIncludeKvconfig, "include kvconfig")

	getCmd.AddCommand(getIndexCmd)
}
