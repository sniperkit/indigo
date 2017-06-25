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
	"github.com/blevesearch/bleve"
	"github.com/mosuka/indigo/client"
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/util"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"io/ioutil"
	"os"
)

type SearchCommandOptions struct {
	gRPCServer       string
	resource         string
	query            string
	size             int
	from             int
	explain          bool
	fields           []string
	sorts            []string
	facets           string
	highlight        string
	highlightStyle   string
	highlightFields  []string
	includeLocations bool
}

var searchCmdOpts SearchCommandOptions

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "searches the documents from the Indigo gRPC Server",
	Long:  `The search command searches the documents from the Indigo gRPC Server.`,
	RunE:  runESearchCmd,
}

func runESearchCmd(cmd *cobra.Command, args []string) error {
	var resourceBytes []byte = nil
	if cmd.Flag("resource").Changed {
		if searchCmdOpts.resource == "-" {
			resourceBytes, _ = ioutil.ReadAll(os.Stdin)
		} else {
			file, err := os.Open(searchCmdOpts.resource)
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

	searchRequest := bleve.NewSearchRequest(nil)
	if len(resourceBytes) > 0 {
		err := searchRequest.UnmarshalJSON(resourceBytes)
		if err != nil {
			return err
		}
	}

	if cmd.Flag("query").Changed {
		searchRequest.Query = bleve.NewQueryStringQuery(searchCmdOpts.query)
	}

	if cmd.Flag("size").Changed {
		searchRequest.Size = searchCmdOpts.size
	}

	if cmd.Flag("from").Changed {
		searchRequest.From = searchCmdOpts.from
	}

	if cmd.Flag("explain").Changed {
		searchRequest.Explain = searchCmdOpts.explain
	}

	if cmd.Flag("field").Changed {
		searchRequest.Fields = searchCmdOpts.fields
	}

	if cmd.Flag("sort").Changed {
		searchRequest.SortBy(searchCmdOpts.sorts)
	}

	if cmd.Flag("facets").Changed {
		facetRequest := bleve.FacetsRequest{}
		err := json.Unmarshal([]byte(searchCmdOpts.facets), &facetRequest)
		if err != nil {
			return err
		}
		searchRequest.Facets = facetRequest
	}

	if cmd.Flag("highlight").Changed {
		highlightRequest := bleve.NewHighlight()
		err := json.Unmarshal([]byte(searchCmdOpts.highlight), highlightRequest)
		if err != nil {
			return err
		}
		searchRequest.Highlight = highlightRequest
	}

	if cmd.Flag("highlight-style").Changed || cmd.Flag("highlight-field").Changed {
		highlightRequest := bleve.NewHighlightWithStyle(searchCmdOpts.highlightStyle)
		highlightRequest.Fields = searchCmdOpts.highlightFields
		searchRequest.Highlight = highlightRequest
	}

	if cmd.Flag("include-locations").Changed {
		searchRequest.IncludeLocations = searchCmdOpts.includeLocations
	}

	searchRequestAny, err := util.MarshalAny(searchRequest)
	if err != nil {
		return err
	}

	protoPutDocumentRequest := &proto.SearchRequest{
		SearchRequest: &searchRequestAny,
	}

	client, err := client.NewIndigoGRPCClient(searchCmdOpts.gRPCServer)
	if err != nil {
		return err
	}
	defer client.Close()

	resp, err := client.Client.Search(context.Background(), protoPutDocumentRequest)
	if err != nil {
		return err
	}

	searchResult, err := util.UnmarshalAny(resp.SearchResult)

	r := util.SearchResponse{
		SearchResult: searchResult.(*bleve.SearchResult),
	}

	switch rootCmdOpts.outputFormat {
	case "text":
		fmt.Printf("%s\n", resp.String())
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
	searchCmd.Flags().StringVar(&searchCmdOpts.gRPCServer, "grpc-server", DefaultServer, "Indigo gRPC Server to connect to")
	searchCmd.Flags().StringVar(&searchCmdOpts.resource, "resource", DefaultResource, "resource file")
	searchCmd.Flags().StringVar(&searchCmdOpts.query, "query", DefaultQuery, "query string")
	searchCmd.Flags().IntVar(&searchCmdOpts.size, "size", DefaultSize, "number of hits to return")
	searchCmd.Flags().IntVar(&searchCmdOpts.from, "from", DefaultFrom, "starting from index of the hits to return")
	searchCmd.Flags().BoolVar(&searchCmdOpts.explain, "explain", DefaultExplain, "contain an explanation of how scoring of the hits was computed")
	searchCmd.Flags().StringSliceVar(&searchCmdOpts.fields, "field", DefaultFields, "specify a set of fields to return")
	searchCmd.Flags().StringSliceVar(&searchCmdOpts.sorts, "sort", DefaultSorts, "sorting to perform")
	searchCmd.Flags().StringVar(&searchCmdOpts.facets, "facets", DefaultFacets, "faceting to perform")
	searchCmd.Flags().StringVar(&searchCmdOpts.highlight, "highlight", DefaultHighlight, "highlighting to perform")
	searchCmd.Flags().StringVar(&searchCmdOpts.highlightStyle, "highlight-style", DefaultHighlightStyle, "highlighting style")
	searchCmd.Flags().StringSliceVar(&searchCmdOpts.highlightFields, "highlight-field", DefaultHighlightFields, "specify a set of fields to highlight")
	searchCmd.Flags().BoolVar(&searchCmdOpts.includeLocations, "include-locations", DefaultIncludeLocations, "include terms locations")

	RootCmd.AddCommand(searchCmd)
}
