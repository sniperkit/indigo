package cmd

import (
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "client for the Indigo gRPC Server",
	Long:  `The client command for the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	searchCmd.PersistentFlags().StringVarP(&grpcServerName, "grpc-server-name", "n", grpcServerName, "Indigo gRPC Sever name")
	searchCmd.PersistentFlags().IntVarP(&grpcServerPort, "grpc-server-port", "p", grpcServerPort, "Indigo gRPC Server port number")

	RootCmd.AddCommand(searchCmd)
}
