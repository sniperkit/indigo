package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type GetIndexResponse struct {
	Path         string      `json:"path"`
	IndexMapping interface{} `json:"index_mapping"`
	IndexType    string      `json:"index_type"`
	Kvstore      string      `json:"kvstore"`
	Kvconfig     interface{} `json:"kvconfig"`
}

var getIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "gets the index mapping from the Indigo Server",
	Long:  `The get index command gets the index information from the Indigo Server.`,
	RunE:  runEGetIndexCmd,
}

func runEGetIndexCmd(cmd *cobra.Command, args []string) error {
	protoGetIndexRequest := &proto.GetIndexRequest{}

	conn, err := grpc.Dial(getCmdOpts.gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.GetIndex(context.Background(), protoGetIndexRequest)
	if err != nil {
		return err
	}

	indexMapping, err := proto.UnmarshalAny(resp.IndexMapping)
	if err != nil {
		return err
	}

	kvconfig, err := proto.UnmarshalAny(resp.Kvconfig)
	if err != nil {
		return err
	}

	r := GetIndexResponse{
		Path:         resp.Path,
		IndexMapping: indexMapping,
		IndexType:    resp.IndexType,
		Kvstore:      resp.Kvstore,
		Kvconfig:     kvconfig,
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
	getCmd.AddCommand(getIndexCmd)
}
