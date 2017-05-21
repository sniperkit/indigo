package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/blevesearch/bleve/mapping"
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/service"
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

func runECreateIndexCmd(cmd *cobra.Command, args []string) error {
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

	createIndexRequest := service.CreateIndexRequest{}
	err := json.Unmarshal(resourceBytes, &createIndexRequest)
	if err != nil {
		return err
	}

	// Index
	if cmd.Flag("index").Changed {
		createIndexRequest.Index = createIndexCmdOpts.index
	}

	// IndexMapping
	if cmd.Flag("index-mapping").Changed {
		indexMapping := &mapping.IndexMappingImpl{}
		err := json.Unmarshal([]byte(createIndexCmdOpts.indexMapping), indexMapping)
		if err != nil {
			return err
		}
		createIndexRequest.IndexMapping = indexMapping
	}

	// IndexType
	if cmd.Flag("index-type").Changed {
		createIndexRequest.IndexType = createIndexCmdOpts.indexType
	}

	// Kvstore
	if cmd.Flag("kvstore").Changed {
		createIndexRequest.Kvstore = createIndexCmdOpts.kvstore
	}

	// Kvconfig
	if cmd.Flag("kvconfig").Changed {
		kvconfig := make(map[string]interface{})
		err := json.Unmarshal([]byte(createIndexCmdOpts.kvconfig), kvconfig)
		if err != nil {
			return err
		}
		createIndexRequest.Kvconfig = kvconfig
	}

	protoCreateIndexRequest, err := createIndexRequest.ProtoMessage()
	if err != nil {
		return err
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
