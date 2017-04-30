package cmd

import (
	"github.com/mosuka/indigo/constant"
	"github.com/spf13/cobra"
)

var CloseCmd = &cobra.Command{
	Use:   "close",
	Short: "closes the object to the Indigo gRPC Server",
	Long:  `The open command creates the object to the Indigo gRPC Server.`,
	RunE:  runECloseCmd,
}

func runECloseCmd(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return cmd.Help()
	}

	_, _, err := cmd.Find(args)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	CloseCmd.PersistentFlags().StringVar(&gRPCServer, "grpc-server", constant.DefaultGRPCServer, "Indigo gRPC Server to connect to")

	RootCmd.AddCommand(CloseCmd)
}
