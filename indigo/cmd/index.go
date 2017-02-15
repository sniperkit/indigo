package cmd

import (
	"github.com/spf13/cobra"
)

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "client for the Indigo gRPC Server",
	Long:  `The client command for the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	indexCmd.PersistentFlags().StringVarP(&grpcServerName, "grpc-server-name", "n", grpcServerName, "Indigo gRPC Sever name")
	indexCmd.PersistentFlags().IntVarP(&grpcServerPort, "grpc-server-port", "p", grpcServerPort, "Indigo gRPC Server port number")

	RootCmd.AddCommand(indexCmd)
}
