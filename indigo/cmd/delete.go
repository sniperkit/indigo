package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletes the object to the Indigo gRPC Server",
	Long:  `The delete command deletes the object to the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("few arguments")
		}

		return nil
	},
}

func init() {
	deleteCmd.PersistentFlags().StringVarP(&gRPCServerName, "grpc-server-name", "n", gRPCServerName, "Indigo gRPC Sever name")
	deleteCmd.PersistentFlags().IntVarP(&gRPCServerPort, "grpc-server-port", "p", gRPCServerPort, "Indigo gRPC Server port number")

	RootCmd.AddCommand(deleteCmd)
}
