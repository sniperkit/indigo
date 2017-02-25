package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/mosuka/indigo/constant"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var getMappingCmd = &cobra.Command{
	Use:   "mapping",
	Short: "gets the index mapping from the Indigo gRPC Server",
	Long:  `The get mapping command gets the index mapping from the Indigo gRPC Server.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if indexName == "" {
			return fmt.Errorf("required flag: --%s", cmd.Flag("index-name").Name)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		client := proto.NewIndigoClient(conn)
		resp, err := client.GetMapping(context.Background(), &proto.GetMappingRequest{IndexName: indexName})
		if err != nil {
			return err
		}

		switch outputFormat {
		case "text":
			fmt.Printf("%s\n", resp.String())
		case "json":
			result := make(map[string]interface{})

			indexMapping := bleve.NewIndexMapping()
			if err := json.Unmarshal(resp.IndexMapping, &indexMapping); err != nil {
				return err
			}
			result["indexMapping"] = indexMapping

			output, err := json.MarshalIndent(result, "", "  ")
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
	getMappingCmd.Flags().StringVarP(&indexName, "index-name", "n", constant.DefaultIndexName, "index name")

	getCmd.AddCommand(getMappingCmd)
}
