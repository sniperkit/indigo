package cmd

import (
	"fmt"
	"github.com/mosuka/bleve-server/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var mappingCmd = &cobra.Command{
	Use:   "mapping",
	Short: "prints the index mapping used for Bleve Server",
	Long:  `The mapping command prints a JSON representation of the index mapping used for Bleve Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", serverName, serverPort), grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()

		c := proto.NewBleveClient(conn)
		resp, err := c.Mapping(context.Background(), &proto.MappingRequest{})
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", resp.Mapping)

		return nil
	},
}

func init() {
	mappingCmd.Flags().StringVarP(&serverName, "server-name", "n", serverName, "sever name")
	mappingCmd.Flags().IntVarP(&serverPort, "server-port", "p", serverPort, "port number")

	RootCmd.AddCommand(mappingCmd)
}
