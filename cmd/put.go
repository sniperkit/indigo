package cmd

import (
	"errors"
	"github.com/mosuka/indigo/constant"
	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "puts the object from the Indigo gRPC Server",
	Long:  `The put command puts the object from the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("few arguments")
		}

		return nil
	},
}

func init() {
	putCmd.PersistentFlags().StringVarP(&gRPCServerName, "grpc-server-name", "n", constant.DefaultGRPCServerName, "Indigo gRPC Sever name")
	putCmd.PersistentFlags().IntVarP(&gRPCServerPort, "grpc-server-port", "p", constant.DefaultGRPCServerPort, "Indigo gRPC Server port number")

	RootCmd.AddCommand(putCmd)
}
