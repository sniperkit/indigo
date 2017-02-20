package cmd

import (
	"errors"
	"github.com/mosuka/indigo/constant"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "creates the object to the Indigo gRPC Server",
	Long:  `The create command creates the object to the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("few arguments")
		}

		return nil
	},
}

func init() {
	createCmd.PersistentFlags().StringVarP(&gRPCServerName, "grpc-server-name", "n", constant.DefaultGRPCServerName, "Indigo gRPC Sever name")
	createCmd.PersistentFlags().IntVarP(&gRPCServerPort, "grpc-server-port", "p", constant.DefaultGRPCServerPort, "Indigo gRPC Server port number")

	RootCmd.AddCommand(createCmd)
}
