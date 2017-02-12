package cmd

import (
	"errors"
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var indexCmd = &cobra.Command{
	Use:   "index REQUEST",
	Short: "indexes the documents to the Indigo gRPC Server",
	Long:  `The index command indexes the JSON representation of the documents to the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("must specify DOCUMENTS")
		}

		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", grpcServerName, grpcServerPort), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		c := proto.NewIndigoClient(conn)

		var resp *proto.IndexResponse
		if deleteFlag {
			resp, err = c.Delete(context.Background(), &proto.DeleteRequest{Ids: args[0], BatchSize: batchSize})
		} else {
			resp, err = c.Index(context.Background(), &proto.IndexRequest{Documents: args[0], BatchSize: batchSize})
		}
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", resp.Result)

		return nil
	},
}

func init() {
	indexCmd.Flags().Int32VarP(&batchSize, "batch-size", "b", batchSize, "port number")
	indexCmd.Flags().BoolVarP(&deleteFlag, "delete", "d", deleteFlag, "delete documents")

	clientCmd.AddCommand(indexCmd)
}
