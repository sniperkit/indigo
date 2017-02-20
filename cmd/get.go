package cmd

import (
	"errors"
	"github.com/mosuka/indigo/constant"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "gets the object from the Indigo gRPC Server",
	Long:  `The get command gets the object from the Indigo gRPC Server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("few arguments")
		}

		return nil
	},
}

func init() {
	getCmd.PersistentFlags().StringVarP(&gRPCServerName, "grpc-server-name", "n", constant.DefaultGRPCServerName, "Indigo gRPC Sever name")
	getCmd.PersistentFlags().IntVarP(&gRPCServerPort, "grpc-server-port", "p", constant.DefaultGRPCServerPort, "Indigo gRPC Server port number")

	RootCmd.AddCommand(getCmd)
}
