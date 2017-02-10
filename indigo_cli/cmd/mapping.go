package cmd

import (
	"fmt"
	"github.com/mosuka/indigo/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var mappingCmd = &cobra.Command{
	Use:   "mapping",
	Short: "prints the index mapping used for the Indigo gRPC Server",
	Long:  `The mapping command prints a JSON representation of the index mapping used for the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", serverName, serverPort), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		c := proto.NewIndigoClient(conn)
		resp, err := c.Mapping(context.Background(), &proto.MappingRequest{})
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", resp.Mapping)

		return nil
	},
}

func init() {
	mappingCmd.Flags().StringVarP(&serverName, "grpc-name", "n", serverName, "sever name")
	mappingCmd.Flags().IntVarP(&serverPort, "grpc-port", "p", serverPort, "port number")

	RootCmd.AddCommand(mappingCmd)
}
