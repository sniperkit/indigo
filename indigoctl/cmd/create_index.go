package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/blevesearch/bleve/mapping"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
)

type CreateIndexCommandOptions struct {
	index        string
	resource     string
	indexMapping string
	indexType    string
	kvstore      string
	kvconfig     string
}

var createIndexCmdOpts CreateIndexCommandOptions

var createIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "creates the index to the Indigo gRPC Server",
	Long:  `The create index command creates the index to the Indigo gRPC Server.`,
	RunE:  runECreateIndexCmd,
}

type CreateIndexResource struct {
	IndexMapping *mapping.IndexMappingImpl `json:"index_mapping,omitempty"`
	IndexType    string                    `json:"index_type,omitempty"`
	Kvstore      string                    `json:"kvstore,omitempty"`
	Kvconfig     map[string]interface{}    `json:"kvconfig,omitempty"`
}

func runECreateIndexCmd(cmd *cobra.Command, args []string) error {
	if createIndexCmdOpts.index == "" {
		return fmt.Errorf("required flag: --%s", cmd.Flag("index").Name)
	}

	var resourceBytes []byte = nil
	if cmd.Flag("resource").Changed {
		if createIndexCmdOpts.resource == "-" {
			resourceBytes, _ = ioutil.ReadAll(os.Stdin)
		} else {
			file, err := os.Open(createIndexCmdOpts.resource)
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

	protoCreateIndexRequest := &proto.CreateIndexRequest{
		Index:        createIndexCmdOpts.index,
		IndexMapping: indexMappingBytes,
		IndexType:    createIndexResource.IndexType,
		Kvstore:      createIndexResource.Kvstore,
		Kvconfig:     kvconfigBytes,
	}

	if cmd.Flag("index-mapping").Changed {
		protoCreateIndexRequest.IndexMapping = []byte(createIndexCmdOpts.indexMapping)
	}

	if cmd.Flag("index-type").Changed {
		protoCreateIndexRequest.IndexType = createIndexCmdOpts.indexType
	}

	if cmd.Flag("kvstore").Changed {
		protoCreateIndexRequest.Kvstore = createIndexCmdOpts.kvstore
	}

	if cmd.Flag("kvconfig").Changed {
		protoCreateIndexRequest.Kvconfig = []byte(createIndexCmdOpts.kvconfig)
	}

	conn, err := grpc.Dial(createCmdOpts.gRPCServer, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewIndigoClient(conn)
	resp, err := client.CreateIndex(context.Background(), protoCreateIndexRequest)
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
	createIndexCmd.Flags().StringVar(&createIndexCmdOpts.index, "index", DefaultIndex, "index name")
	createIndexCmd.Flags().StringVar(&createIndexCmdOpts.resource, "resource", DefaultResource, "resource file")
	createIndexCmd.Flags().StringVar(&createIndexCmdOpts.indexMapping, "index-mapping", DefaultIndexMapping, "index mapping")
	createIndexCmd.Flags().StringVar(&createIndexCmdOpts.indexType, "index-type", DefaultIndexType, "index type")
	createIndexCmd.Flags().StringVar(&createIndexCmdOpts.kvstore, "kvstore", DefaultKvstore, "kvstore")
	createIndexCmd.Flags().StringVar(&createIndexCmdOpts.kvconfig, "kvconfig", DefaultKvconfig, "kvconfig")

	createCmd.AddCommand(createIndexCmd)
}
