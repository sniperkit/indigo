package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "client for the Indigo gRPC Server",
	Long:  `The client command for the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	createCmd.PersistentFlags().StringVarP(&grpcServerName, "grpc-server-name", "n", grpcServerName, "Indigo gRPC Sever name")
	createCmd.PersistentFlags().IntVarP(&grpcServerPort, "grpc-server-port", "p", grpcServerPort, "Indigo gRPC Server port number")

	RootCmd.AddCommand(createCmd)
}
