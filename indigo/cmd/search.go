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
	searchCmd.PersistentFlags().StringVarP(&grpcServerName, "grpc-server-name", "n", grpcServerName, "Indigo gRPC Sever name")
	searchCmd.PersistentFlags().IntVarP(&grpcServerPort, "grpc-server-port", "p", grpcServerPort, "Indigo gRPC Server port number")

	RootCmd.AddCommand(searchCmd)
}
