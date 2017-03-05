package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var getIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "gets the index information from the Indigo gRPC Server",
	Long:  `The get index command gets the index information from the Indigo gRPC Server.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if indexName == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index-name").Name)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(indigoSettings.GetString("grpc_server"), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.GetIndex(context.Background(), &proto.GetIndexRequest{IndexName: indexName})
		if err != nil {
			return err
		}

		indexStats := make(map[string]interface{})
		if err := json.Unmarshal(resp.IndexStats, &indexStats); err != nil {
			return err
		}

		indexMapping := bleve.NewIndexMapping()
		if err := json.Unmarshal(resp.IndexMapping, &indexMapping); err != nil {
			return err
		}

		r := struct {
			DocumentCount uint64                    `json:"document_count"`
			IndexStats    map[string]interface{}    `json:"index_stats"`
			IndexMapping  *mapping.IndexMappingImpl `json:"index_mapping"`
		}{
			DocumentCount: resp.DocumentCount,
			IndexStats:    indexStats,
			IndexMapping:  indexMapping,
		}

		switch indigoSettings.GetString("output_format") {
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
	},
}

func init() {
	getIndexCmd.Flags().StringVarP(&indexName, "index-name", "n", constant.DefaultIndexName, "index name")

	getCmd.AddCommand(getIndexCmd)
}
