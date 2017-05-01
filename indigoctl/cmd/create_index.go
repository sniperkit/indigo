package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mosuka/indigo/defaultvalue"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
)

var CreateIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "creates the index to the Indigo gRPC Server",
	Long:  `The create index command creates the index to the Indigo gRPC Server.`,
	RunE:  runECreateIndexCmd,
}

func runECreateIndexCmd(cmd *cobra.Command, args []string) error {
	if index == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
	}

	var im []byte = nil
	var kvc []byte = nil

	if indexMapping != "" {
		file, err := os.Open(indexMapping)
		if err != nil {
			return err
		}
		defer file.Close()

		im, err = ioutil.ReadAll(file)
		if err != nil {
			return err
		}
	}

	if kvConfig != "" {
		file, err := os.Open(kvConfig)
		if err != nil {
			return err
		}
		defer file.Close()

		kvc, err = ioutil.ReadAll(file)
		if err != nil {
			return err
		}
	}

	conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.CreateIndex(context.Background(), &proto.CreateIndexRequest{Index: index, IndexMapping: im, IndexType: indexType, Kvstore: kvStore, Kvconfig: kvc})
	if err != nil {
		return err
	}

	switch outputFormat {
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
	CreateIndexCmd.Flags().StringVar(&index, "index", defaultvalue.DefaultIndex, "index name")
	CreateIndexCmd.Flags().StringVar(&indexMapping, "index-mapping", defaultvalue.DefaultIndexMapping, "index mapping")
	CreateIndexCmd.Flags().StringVar(&indexType, "index-type", defaultvalue.DefaultIndexType, "index type")
	CreateIndexCmd.Flags().StringVar(&kvStore, "kvstore", defaultvalue.DefaultKVStore, "kvstore")
	CreateIndexCmd.Flags().StringVar(&kvConfig, "kvconfig", defaultvalue.DefaultKVConfigFile, "kvconfig")

	CreateCmd.AddCommand(CreateIndexCmd)
}
