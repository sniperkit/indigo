package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var GetIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "gets the index information from the Indigo gRPC Server",
	Long:  `The get index command gets the index information from the Indigo gRPC Server.`,
	RunE:  runEGetIndexCmd,
}

type GetIndexResponse struct {
	DocumentCount uint64                    `json:"document_count"`
	IndexStats    map[string]interface{}    `json:"index_stats"`
	IndexMapping  *mapping.IndexMappingImpl `json:"index_mapping"`
}

func runEGetIndexCmd(cmd *cobra.Command, args []string) error {
	index := cmd.Flag("index").Value.String()
	if index == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
	}

	getIndexRequest := &proto.GetIndexRequest{
		Index: index,
	}

	conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.GetIndex(context.Background(), getIndexRequest)
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

	r := GetIndexResponse{
		DocumentCount: resp.DocumentCount,
		IndexStats:    indexStats,
		IndexMapping:  indexMapping,
	}

	switch outputFormat {
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
	GetIndexCmd.Flags().String("index", "", "index name")

	GetCmd.AddCommand(GetIndexCmd)
}
