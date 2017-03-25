package cmd

import (
	"github.com/mosuka/indigo/constant"
	"github.com/spf13/cobra"
)

var OpenCmd = &cobra.Command{
	Use:   "open",
	Short: "opens the object to the Indigo gRPC Server",
	Long:  `The open command creates the object to the Indigo gRPC Server.`,
	RunE:  runEOpenCmd,
}

func runEOpenCmd(cmd *cobra.Command, args []string) error {
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
	OpenCmd.PersistentFlags().StringVarP(&gRPCServer, "grpc-server", "g", constant.DefaultGRPCServer, "Indigo gRPC Server to connect to")

	RootCmd.AddCommand(OpenCmd)
}
