package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "searches the documents from the Indigo gRPC Server",
	Long:  `The search command searches the documents from the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("few arguments")
		}

		return nil
	},
}

func init() {
	searchCmd.PersistentFlags().StringVarP(&gRPCServerName, "grpc-server-name", "n", gRPCServerName, "Indigo gRPC Sever name")
	searchCmd.PersistentFlags().IntVarP(&gRPCServerPort, "grpc-server-port", "p", gRPCServerPort, "Indigo gRPC Server port number")

	RootCmd.AddCommand(searchCmd)
}
