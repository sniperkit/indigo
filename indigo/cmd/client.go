package cmd

import (
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client for the Indigo gRPC Server",
	Long:  `The client command for the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	clientCmd.PersistentFlags().StringVarP(&grpcServerName, "server-name", "n", grpcServerName, "sever name")
	clientCmd.PersistentFlags().IntVarP(&grpcServerPort, "server-port", "p", grpcServerPort, "port number")

	RootCmd.AddCommand(clientCmd)
}
