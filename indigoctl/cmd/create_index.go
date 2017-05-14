package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/blevesearch/bleve/mapping"
	"github.com/mosuka/indigo/defaultvalue"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
)

var CreateIndexCmd = &cobra.Command{
	Use:   "index RESOURCE_FILENAME",
	Short: "creates the index to the Indigo gRPC Server",
	Long:  `The create index command creates the index to the Indigo gRPC Server.`,
	RunE:  runECreateIndexCmd,
}

type CreateIndexResource struct {
	Index        string                    `json:"index,omitempty"`
	IndexMapping *mapping.IndexMappingImpl `json:"index_mapping,omitempty"`
	IndexType    string                    `json:"index_type,omitempty"`
	Kvstore      string                    `json:"kvstore,omitempty"`
	Kvconfig     map[string]interface{}    `json:"kvconfig,omitempty"`
}

func runECreateIndexCmd(cmd *cobra.Command, args []string) error {
	var resourceBytes []byte = nil
	if terminal.IsTerminal(0) {
		if len(args) > 0 {
			resourceBytes = []byte(args[0])
			file, err := os.Open(args[0])
			if err != nil {
				return err
			}
			defer file.Close()

			resourceBytes, err = ioutil.ReadAll(file)
			if err != nil {
				return err
			}
		}
	} else {
		resourceBytes, _ = ioutil.ReadAll(os.Stdin)
	}

	createIndexResource := CreateIndexResource{}
	err := json.Unmarshal(resourceBytes, &createIndexResource)
	if err != nil {
		return err
	}

	indexMappingBytes, err := json.Marshal(createIndexResource.IndexMapping)
	if err != nil {
		return err
	}

	kvconfigBytes, err := json.Marshal(createIndexResource.Kvconfig)
	if err != nil {
		return err
	}

	createIndexRequest := &proto.CreateIndexRequest{
		Index:        createIndexResource.Index,
		IndexMapping: indexMappingBytes,
		IndexType:    createIndexResource.IndexType,
		Kvstore:      createIndexResource.Kvstore,
		Kvconfig:     kvconfigBytes,
	}

	if cmd.Flag("index").Changed {
		createIndexRequest.Index = cmd.Flag("index").Value.String()
	}

	if cmd.Flag("index-mapping").Changed {
		indexMappingBytes := []byte(cmd.Flag("index-mapping").Value.String())
		createIndexRequest.IndexMapping = indexMappingBytes
	}

	if cmd.Flag("index-type").Changed {
		createIndexRequest.IndexType = cmd.Flag("index-type").Value.String()
	}

	if cmd.Flag("kvstore").Changed {
		createIndexRequest.Kvstore = cmd.Flag("kvstore").Value.String()
	}

	if cmd.Flag("kvconfig").Changed {
		kvconfigBytes := []byte(cmd.Flag("kvconfig").Value.String())
		createIndexRequest.Kvconfig = kvconfigBytes
	}

	conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.CreateIndex(context.Background(), createIndexRequest)
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
	CreateIndexCmd.Flags().String("index", defaultvalue.DefaultIndex, "index name")
	CreateIndexCmd.Flags().String("index-mapping", defaultvalue.DefaultIndexMapping, "index mapping")
	CreateIndexCmd.Flags().String("index-type", defaultvalue.DefaultIndexType, "index type")
	CreateIndexCmd.Flags().String("kvstore", defaultvalue.DefaultKVStore, "kvstore")
	CreateIndexCmd.Flags().String("kvconfig", defaultvalue.DefaultKVConfigFile, "kvconfig")

	CreateCmd.AddCommand(CreateIndexCmd)
}
